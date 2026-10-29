// Command web serves badgerd.nl: SSR pages with live stock from governor,
// static assets, SEO endpoints and the /api + /gif reverse proxies.
//
// Configuration (environment):
//
//	PORT               listen port (default 5000)
//	BASE_URL           public origin for canonical/sitemap URLs (default https://badgerd.nl)
//	GOVERNOR_BASE_URL  backend for stock, shipping and checkout (default https://governor.badgerd.nl)
//	GIPHY_BASE_URL     upstream for /gif/* (default https://media.giphy.com)
//	STATIC_DIR         serve assets from this directory instead of the embedded copy (development)
//	TZ                 local time zone for date-based banners (deploy/env/*.env set Europe/Amsterdam)
//	ENVIRONMENT        deployment name reported by /healthz (staging, prod)
//	IMAGE_DIGEST       container image digest reported by /healthz (written by the Quadlet unit's ExecStartPre)
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
	_ "time/tzdata" // distroless has no zoneinfo; TZ (e.g. Europe/Amsterdam) drives the break banner window

	"badgerd.nl/internal/governor"
	"badgerd.nl/internal/server"
	"badgerd.nl/web"
)

// Set at link time (-ldflags "-X main.commit=..."; see Containerfile).
var version, commit, buildTime string

func main() {
	healthcheck := flag.Bool("healthcheck", false, "probe http://127.0.0.1:$PORT/healthz and exit 0 if healthy (for container health checks)")
	flag.Parse()

	port := env("PORT", "5000")
	if *healthcheck {
		os.Exit(probe(port))
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if err := run(logger, port); err != nil {
		logger.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger, port string) error {
	govURL, err := parseURL("GOVERNOR_BASE_URL", "https://governor.badgerd.nl")
	if err != nil {
		return err
	}
	giphyURL, err := parseURL("GIPHY_BASE_URL", "https://media.giphy.com")
	if err != nil {
		return err
	}
	if dir := os.Getenv("STATIC_DIR"); dir != "" {
		web.UseDir(dir)
		logger.Info("serving static assets from disk", "dir", dir)
	}

	gov, err := governor.New(govURL.String())
	if err != nil {
		return err
	}
	handler := server.New(server.Config{
		BaseURL:     env("BASE_URL", "https://badgerd.nl"),
		GovernorURL: govURL,
		GiphyURL:    giphyURL,
		Logger:      logger,
		Info: server.Info{
			Version:     version,
			Commit:      buildCommit(),
			BuildTime:   buildTime,
			Environment: os.Getenv("ENVIRONMENT"),
			ImageDigest: os.Getenv("IMAGE_DIGEST"),
		},
	}, gov)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errc := make(chan error, 1)
	go func() {
		logger.Info("listening", "addr", srv.Addr)
		errc <- srv.ListenAndServe()
	}()

	select {
	case err := <-errc:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
	}

	logger.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}

func probe(port string) int {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://127.0.0.1:" + port + "/healthz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 1
	}
	return 0
}

// buildCommit prefers the link-time commit, falling back to the VCS stamp Go
// embeds when building from a git checkout (`just build`).
func buildCommit() string {
	if commit != "" {
		return commit
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	var rev, dirty string
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			rev = s.Value
		case "vcs.modified":
			if s.Value == "true" {
				dirty = "-dirty"
			}
		}
	}
	if rev == "" {
		return ""
	}
	return rev + dirty
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func parseURL(key, def string) (*url.URL, error) {
	raw := env(key, def)
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("%s: invalid URL %q", key, raw)
	}
	return u, nil
}
