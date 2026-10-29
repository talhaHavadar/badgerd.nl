// Package server wires the badgerd.nl HTTP surface: SSR pages, static
// assets, SEO endpoints and the same-origin /api and /gif reverse proxies.
package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"mime"
	"net/http"
	"net/url"
	"time"

	"badgerd.nl/internal/catalog"
	"badgerd.nl/internal/governor"
	"badgerd.nl/internal/seo"
)

// Config configures the server.
type Config struct {
	BaseURL     string   // public origin for canonical/sitemap URLs, e.g. https://badgerd.nl
	GovernorURL *url.URL // upstream for /api/*
	GiphyURL    *url.URL // upstream for /gif/*
	Logger      *slog.Logger
	Now         func() time.Time // defaults to time.Now
	Info        Info             // reported by /healthz
}

// Info describes the running instance; empty fields are omitted from /healthz.
type Info struct {
	Version     string `json:"version,omitempty"`
	Commit      string `json:"commit,omitempty"`
	BuildTime   string `json:"build_time,omitempty"`
	Environment string `json:"environment,omitempty"`
	ImageDigest string `json:"image_digest,omitempty"`
}

// StockSource provides live product data (implemented by *governor.Client).
type StockSource interface {
	Product(ctx context.Context, codename string) (governor.Product, bool, error)
	Countries(ctx context.Context) ([]governor.Country, error)
}

// Server is the badgerd.nl HTTP handler.
type Server struct {
	cfg     Config
	stock   StockSource
	log     *slog.Logger
	now     func() time.Time
	started time.Time
	handler http.Handler
}

func init() {
	// The runtime image has no /etc/mime.types; match the content types the Nuxt site served.
	for ext, typ := range map[string]string{
		".ico":         "image/x-icon",
		".webmanifest": "application/manifest+json",
		".hex":         "application/octet-stream",
		".uf2":         "application/octet-stream",
		".txt":         "text/plain; charset=utf-8",
	} {
		_ = mime.AddExtensionType(ext, typ)
	}
}

// New returns the site handler.
func New(cfg Config, stock StockSource) *Server {
	cfg.BaseURL = seo.TrimBase(cfg.BaseURL)
	s := &Server{cfg: cfg, stock: stock, log: cfg.Logger, now: cfg.Now}
	if s.log == nil {
		s.log = slog.Default()
	}
	if s.now == nil {
		s.now = time.Now
	}
	s.started = s.now()
	s.handler = s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.handler.ServeHTTP(w, r) }

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", s.handleIndex)
	for _, p := range catalog.Products {
		mux.Handle("GET "+p.Path()+"{$}", s.handleProduct(p))
	}
	for _, lp := range legalPages {
		mux.Handle("GET "+lp.path+"{$}", s.handleStatic(lp.path, lp.page))
	}
	mux.HandleFunc("GET /success/{$}", s.handleSuccess)

	mux.HandleFunc("GET /sitemap.xml", s.handleSitemap)
	mux.HandleFunc("GET /robots.txt", s.handleRobots)
	mux.HandleFunc("GET /healthz", s.handleHealthz)

	mux.Handle("/api/", newProxy(s.cfg.GovernorURL, "/api", true, s.log))
	mux.Handle("GET /gif/", newProxy(s.cfg.GiphyURL, "/gif", false, s.log))

	// Static assets, falling back to the 404 page. Registered without a method
	// so it doesn't conflict with the method-less /api/ pattern.
	mux.HandleFunc("/", s.handleFiles)

	return s.logRequests(trailingSlash(mux))
}

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	now := s.now()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
		Info
		StartedAt     string `json:"started_at"`
		UptimeSeconds int64  `json:"uptime_seconds"`
	}{
		Status:        "ok",
		Info:          s.cfg.Info,
		StartedAt:     s.started.UTC().Format(time.RFC3339),
		UptimeSeconds: int64(now.Sub(s.started).Seconds()),
	})
}

func (s *Server) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		s.log.Info("request", "method", r.Method, "path", r.URL.Path, "status", rec.status, "duration", time.Since(start))
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// Unwrap lets http.ResponseController reach the underlying writer (flushing for proxies).
func (r *statusRecorder) Unwrap() http.ResponseWriter { return r.ResponseWriter }
