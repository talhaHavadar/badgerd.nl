package server

import (
	"bytes"
	"context"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"sync"

	"github.com/a-h/templ"

	"badgerd.nl/internal/catalog"
	"badgerd.nl/internal/governor"
	"badgerd.nl/internal/seo"
	"badgerd.nl/web"
	"badgerd.nl/web/components"
	"badgerd.nl/web/pages"
)

var legalPages = []struct {
	path string
	page func(components.Page) templ.Component
}{
	{"/terms/", pages.Terms},
	{"/privacy/", pages.Privacy},
	{"/refund/", pages.Refund},
	{"/legal/", pages.LegalNotice},
}

func (s *Server) page(meta seo.Meta, scripts ...string) components.Page {
	return components.Page{Meta: meta, ShowBreak: components.BreakNotice.Active(s.now()), Scripts: scripts}
}

// render buffers the component so a template error becomes a clean 500.
func (s *Server) render(w http.ResponseWriter, r *http.Request, status int, c templ.Component) {
	var buf bytes.Buffer
	if err := c.Render(r.Context(), &buf); err != nil {
		s.log.Error("render failed", "path", r.URL.Path, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusOK, pages.Index(s.page(seo.Index(s.cfg.BaseURL))))
}

func (s *Server) handleStatic(urlPath string, page func(components.Page) templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.render(w, r, http.StatusOK, page(s.page(seo.Page(s.cfg.BaseURL, urlPath))))
	})
}

func (s *Server) handleSuccess(w http.ResponseWriter, r *http.Request) {
	meta := seo.Page(s.cfg.BaseURL, "/success/")
	s.render(w, r, http.StatusOK, pages.Success(s.page(meta, "js/success.js"), r.URL.Query().Get("order_id")))
}

// handleProduct renders a product page with live stock. If governor is
// unavailable the page still renders, with purchasing disabled (out-of-stock
// state) rather than failing.
func (s *Server) handleProduct(p catalog.Product) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stock, countries := s.liveData(r.Context(), p)
		form := p.Form(stock)
		meta := seo.Product(s.cfg.BaseURL, p, form.Availability, s.now())
		c, err := pages.Product(s.page(meta, "js/carousel.js"), p, form, countries)
		if err != nil {
			s.log.Error("product page", "slug", p.Slug, "err", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		s.render(w, r, http.StatusOK, c)
	})
}

func (s *Server) liveData(ctx context.Context, p catalog.Product) (catalog.Stock, []governor.Country) {
	var (
		wg        sync.WaitGroup
		stock     catalog.Stock
		countries []governor.Country
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		gp, found, err := s.stock.Product(ctx, p.Codename)
		switch {
		case err != nil:
			s.log.Warn("governor products unavailable; rendering without stock", "slug", p.Slug, "err", err)
		case !found:
			s.log.Warn("product missing from governor; rendering without stock", "codename", p.Codename)
		default:
			stock = catalog.Stock{Current: gp.CurrentStock, Reserved: gp.ReservedStock, IgnoreStock: gp.IgnoreStock}
		}
	}()
	go func() {
		defer wg.Done()
		cs, err := s.stock.Countries(ctx)
		if err != nil {
			s.log.Warn("governor shipping rates unavailable", "err", err)
			return
		}
		countries = cs
	}()
	wg.Wait()
	return stock, countries
}

// handleFiles serves web/static, or the 404 page.
func (s *Server) handleFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	name := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	fsys := web.Static()
	if fi, err := fs.Stat(fsys, name); name == "" || err != nil || fi.IsDir() {
		s.notFound(w, r)
		return
	}
	if r.URL.Query().Has("v") {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	} else {
		w.Header().Set("Cache-Control", "public, max-age=3600")
	}
	http.ServeFileFS(w, r, fsys, name)
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	meta := seo.Meta{Title: "Page not found | " + seo.DefaultTitle, Description: seo.DefaultDescription}
	s.render(w, r, http.StatusNotFound, pages.NotFound(s.page(meta)))
}
