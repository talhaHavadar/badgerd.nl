package server

import (
	"net/http"
	"net/url"
	"path"
	"strings"
)

// trailingSlash applies the site-wide URL policy (migration_guide.md §9): every
// GET/HEAD for an extension-less path without a trailing slash is 301-redirected
// to the slash form, query string preserved. Files, /api and /gif are untouched.
func trailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if (r.Method == http.MethodGet || r.Method == http.MethodHead) &&
			!strings.HasSuffix(p, "/") &&
			path.Ext(p) == "" &&
			p != "/healthz" &&
			!strings.HasPrefix(p, "/api/") &&
			!strings.HasPrefix(p, "/gif/") &&
			// Leave unclean paths (e.g. "//evil.example") to ServeMux's own cleaning
			// so we never emit a protocol-relative Location.
			path.Clean(p) == p {
			target := url.URL{Path: p + "/", RawQuery: r.URL.RawQuery}
			http.Redirect(w, r, target.String(), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}
