package server

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

// newProxy forwards requests under prefix to target with the prefix stripped,
// replacing @nuxtjs/proxy (`pathRewrite: {'^/api': ”}`, changeOrigin). Method
// and body are preserved, so the checkout form POST reaches governor and its
// redirect to Stripe is passed back untouched.
func newProxy(target *url.URL, prefix string, forwardClient bool, log *slog.Logger) http.Handler {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.ResponseHeaderTimeout = 30 * time.Second

	return &httputil.ReverseProxy{
		Transport: transport,
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.Out.URL.Path = stripPrefix(pr.In.URL.Path, prefix)
			pr.Out.URL.RawPath = ""
			if raw := pr.In.URL.RawPath; raw != "" {
				pr.Out.URL.RawPath = stripPrefix(raw, prefix)
			}
			pr.SetURL(target) // also sets the outbound Host to the upstream (changeOrigin)

			if !forwardClient {
				// Third-party upstream: don't leak cookies or client details.
				pr.Out.Header.Del("Cookie")
				return
			}
			// Keep the X-Forwarded-* chain from nginx, appending our hop.
			pr.Out.Header["X-Forwarded-For"] = pr.In.Header["X-Forwarded-For"]
			pr.SetXForwarded()
			for _, h := range []string{"X-Forwarded-Proto", "X-Forwarded-Host"} {
				if v := pr.In.Header.Get(h); v != "" {
					pr.Out.Header.Set(h, v)
				}
			}
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			log.Error("proxy error", "upstream", target.Host, "path", r.URL.Path, "err", err)
			http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		},
	}
}

func stripPrefix(p, prefix string) string {
	p = strings.TrimPrefix(p, prefix)
	if p == "" {
		return "/"
	}
	return p
}
