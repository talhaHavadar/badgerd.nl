package server

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func discardLogger() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }

type seen struct {
	method, path, query, host, body, cookie, xff, xfproto string
}

func upstream(t *testing.T, got *seen) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		*got = seen{r.Method, r.URL.Path, r.URL.RawQuery, r.Host, string(b), r.Header.Get("Cookie"), r.Header.Get("X-Forwarded-For"), r.Header.Get("X-Forwarded-Proto")}
		if r.URL.Path == "/stripe/checkout" {
			http.Redirect(w, r, "https://checkout.stripe.com/c/pay/cs_test_123", http.StatusSeeOther)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
}

func TestAPIProxyForwardsCheckoutPOST(t *testing.T) {
	var got seen
	up := upstream(t, &got)
	defer up.Close()
	upURL, _ := url.Parse(up.URL)

	s := New(Config{BaseURL: "https://badgerd.nl", GovernorURL: upURL, GiphyURL: upURL, Logger: discardLogger(), Now: time.Now}, &fakeStock{})

	form := "codename=badgerd_usbmux&quantity=2&country_code=NL"
	req := httptest.NewRequest(http.MethodPost, "/api/stripe/checkout", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Forwarded-For", "203.0.113.7")
	req.Header.Set("X-Forwarded-Proto", "https")
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)

	if rec.Code != http.StatusSeeOther || rec.Header().Get("Location") != "https://checkout.stripe.com/c/pay/cs_test_123" {
		t.Fatalf("got %d %q, want governor's redirect passed through", rec.Code, rec.Header().Get("Location"))
	}
	if got.method != http.MethodPost || got.path != "/stripe/checkout" || got.body != form {
		t.Errorf("upstream saw %s %s body=%q", got.method, got.path, got.body)
	}
	if got.host != upURL.Host {
		t.Errorf("upstream Host = %q, want %q (changeOrigin)", got.host, upURL.Host)
	}
	if !strings.HasPrefix(got.xff, "203.0.113.7") || got.xfproto != "https" {
		t.Errorf("forwarded headers not preserved: xff=%q proto=%q", got.xff, got.xfproto)
	}
}

func TestAPIProxyRewritesPathAndQuery(t *testing.T) {
	var got seen
	up := upstream(t, &got)
	defer up.Close()
	upURL, _ := url.Parse(up.URL)
	s := New(Config{BaseURL: "https://badgerd.nl", GovernorURL: upURL, GiphyURL: upURL, Logger: discardLogger()}, &fakeStock{})

	res, _ := get(t, s, "/api/stripe/products?x=1")
	if res.StatusCode != http.StatusOK || got.path != "/stripe/products" || got.query != "x=1" {
		t.Errorf("got %d upstream %s?%s", res.StatusCode, got.path, got.query)
	}
}

func TestGifProxyStripsCookies(t *testing.T) {
	var got seen
	up := upstream(t, &got)
	defer up.Close()
	upURL, _ := url.Parse(up.URL)
	s := New(Config{BaseURL: "https://badgerd.nl", GovernorURL: upURL, GiphyURL: upURL, Logger: discardLogger()}, &fakeStock{})

	req := httptest.NewRequest(http.MethodGet, "/gif/media/abc/giphy.gif", nil)
	req.Header.Set("Cookie", "session=secret")
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)
	if got.path != "/media/abc/giphy.gif" || got.cookie != "" {
		t.Errorf("upstream saw path=%q cookie=%q", got.path, got.cookie)
	}
}

func TestProxyUpstreamDown(t *testing.T) {
	dead, _ := url.Parse("http://127.0.0.1:1")
	s := New(Config{BaseURL: "https://badgerd.nl", GovernorURL: dead, GiphyURL: dead, Logger: discardLogger()}, &fakeStock{})
	if res, _ := get(t, s, "/api/stripe/products"); res.StatusCode != http.StatusBadGateway {
		t.Errorf("status %d, want 502", res.StatusCode)
	}
}
