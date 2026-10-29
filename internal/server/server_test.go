package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"badgerd.nl/internal/catalog"
	"badgerd.nl/internal/governor"
)

type fakeStock struct {
	products  map[string]governor.Product
	countries []governor.Country
	err       error
}

func (f *fakeStock) Product(_ context.Context, codename string) (governor.Product, bool, error) {
	if f.err != nil {
		return governor.Product{}, false, f.err
	}
	p, ok := f.products[codename]
	return p, ok, nil
}

func (f *fakeStock) Countries(context.Context) ([]governor.Country, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.countries, nil
}

func mustURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	return u
}

func newTestServer(t *testing.T, stock StockSource, now time.Time) *Server {
	t.Helper()
	return New(Config{
		BaseURL:     "https://badgerd.nl/",
		GovernorURL: mustURL(t, "http://governor.invalid"),
		GiphyURL:    mustURL(t, "http://giphy.invalid"),
		Now:         func() time.Time { return now },
		Logger:      discardLogger(),
	}, stock)
}

func get(t *testing.T, h http.Handler, target string) (*http.Response, string) {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, target, nil))
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	return res, string(body)
}

var afterBreak = time.Date(2026, time.October, 1, 12, 0, 0, 0, time.Local)

func TestPagesRender(t *testing.T) {
	s := newTestServer(t, &fakeStock{}, afterBreak)
	tests := []struct {
		path, title, canonical string
	}{
		{"/", "Sett | Badgerd Technologies", "https://badgerd.nl"},
		{"/sdwire3/", "Buy SDWire3 With USB3 | Quickstart With Badgerd SDWire3", "https://badgerd.nl/sdwire3/"},
		{"/sdwirec/", "Buy SDWire With USB-C | Quickstart With Badgerd SDWireC", "https://badgerd.nl/sdwirec/"},
		{"/nrfnano/", "Buy nRFNANO IoT Development Board | Quickstart With Badgerd nRFNANO (nrf52840)", "https://badgerd.nl/nrfnano/"},
		{"/usbmux/", "Buy Badgerd USBMux | Quickstart With USBMux", "https://badgerd.nl/usbmux/"},
		{"/success/", "Badgerd Technologies B.V.", "https://badgerd.nl/success/"},
		{"/terms/", "Badgerd Technologies B.V.", "https://badgerd.nl/terms/"},
		{"/privacy/", "Badgerd Technologies B.V.", "https://badgerd.nl/privacy/"},
		{"/refund/", "Badgerd Technologies B.V.", "https://badgerd.nl/refund/"},
		{"/legal/", "Badgerd Technologies B.V.", "https://badgerd.nl/legal/"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			res, body := get(t, s, tt.path)
			if res.StatusCode != http.StatusOK {
				t.Fatalf("status %d", res.StatusCode)
			}
			if ct := res.Header.Get("Content-Type"); ct != "text/html; charset=utf-8" {
				t.Errorf("content-type %q", ct)
			}
			if !strings.Contains(body, "<title>"+htmlEscape(tt.title)+"</title>") {
				t.Errorf("missing title %q", tt.title)
			}
			if !strings.Contains(body, `<link rel="canonical" href="`+tt.canonical+`">`) {
				t.Errorf("missing canonical %q", tt.canonical)
			}
			if strings.Contains(body, "We are on a break!") {
				t.Error("break banner shown outside its window")
			}
		})
	}
}

func htmlEscape(s string) string { return strings.ReplaceAll(s, "&", "&amp;") }

func TestProductJSONLD(t *testing.T) {
	s := newTestServer(t, &fakeStock{products: map[string]governor.Product{
		"badgerd_usbmux": {Codename: "badgerd_usbmux", CurrentStock: 3},
	}}, afterBreak)
	for _, p := range catalog.Products {
		_, body := get(t, s, p.Path())
		if n := strings.Count(body, `<script type="application/ld+json">`); n != 1 {
			t.Fatalf("%s: %d JSON-LD scripts", p.Slug, n)
		}
		for _, want := range []string{`"@type":"Product"`, `"sku":"` + p.Codename + `"`, `"@type":"BreadcrumbList"`, `"priceValidUntil":"2026-12-31"`} {
			if !strings.Contains(body, want) {
				t.Errorf("%s: JSON-LD missing %s", p.Slug, want)
			}
		}
	}
	_, body := get(t, s, "/usbmux/")
	if !strings.Contains(body, `"availability":"https://schema.org/InStock"`) {
		t.Error("usbmux offer availability should follow live stock")
	}
}

func TestProductStockStates(t *testing.T) {
	countries := []governor.Country{{Code: "NL", Name: "Netherlands"}}
	tests := []struct {
		name    string
		product governor.Product
		want    []string
		reject  []string
	}{
		{
			name:    "out of stock",
			product: governor.Product{CurrentStock: 0},
			want:    []string{"Out of Stock", `name="quantity" disabled`, `<button type="submit" disabled`},
			reject:  []string{"In Stock", "Last 0 items", "Looking for a bulk order"},
		},
		{
			name:    "low stock",
			product: governor.Product{CurrentStock: 5, ReservedStock: 2},
			want: []string{"In Stock", `text-orange-600">Last 5 items`, `<option value="5">5</option>`,
				"2 items reserved, might be yours, try again in 30 minutes", ">Buy</button>", `<option value="NL">Netherlands</option>`,
				`<div class="font-semibold">€180.00 + Shipping</div>`},
			reject: []string{`<option value="6">`, `name="quantity" disabled`, "incl. VAT"},
		},
		{
			name:    "pre-order",
			product: governor.Product{IgnoreStock: true},
			want:    []string{"Open for Pre-Order", ">Pre Order</button>", `<option value="20">20</option>`},
			reject:  []string{"PREORDER30", "discount", "Out of Stock"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.product.Codename = "badgerd_usbmux"
			s := newTestServer(t, &fakeStock{
				products:  map[string]governor.Product{"badgerd_usbmux": tt.product},
				countries: countries,
			}, afterBreak)
			_, body := get(t, s, "/usbmux/")
			for _, w := range tt.want {
				if !strings.Contains(body, w) {
					t.Errorf("missing %q", w)
				}
			}
			for _, r := range tt.reject {
				if strings.Contains(body, r) {
					t.Errorf("unexpected %q", r)
				}
			}
		})
	}
}

func TestProductDegradesWhenGovernorDown(t *testing.T) {
	s := newTestServer(t, &fakeStock{err: errors.New("connection refused")}, afterBreak)
	res, body := get(t, s, "/sdwirec/")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status %d, want 200", res.StatusCode)
	}
	if !strings.Contains(body, "Out of Stock") || !strings.Contains(body, `<button type="submit" disabled`) {
		t.Error("purchase should be disabled when governor is unavailable")
	}
}

func TestBreakBanner(t *testing.T) {
	s := newTestServer(t, &fakeStock{}, time.Date(2026, time.August, 31, 23, 0, 0, 0, time.Local))
	if _, body := get(t, s, "/"); !strings.Contains(body, "We are on a break!") {
		t.Error("break banner missing inside its window")
	}
}

func TestSuccessEscapesOrderID(t *testing.T) {
	s := newTestServer(t, &fakeStock{}, afterBreak)
	_, body := get(t, s, "/success/?order_id=%3Cscript%3E42&codename=badgerd_usbmux")
	if !strings.Contains(body, "Order# &lt;script&gt;42") {
		t.Error("order id not rendered escaped")
	}
}

func TestTrailingSlashPolicy(t *testing.T) {
	s := newTestServer(t, &fakeStock{}, afterBreak)
	tests := []struct {
		target, location string
		status           int
	}{
		{"/usbmux", "/usbmux/", http.StatusMovedPermanently},
		{"/terms", "/terms/", http.StatusMovedPermanently},
		{"/terms?utm=x", "/terms/?utm=x", http.StatusMovedPermanently},
		{"/favicon.ico", "", http.StatusOK},
		{"/robots.txt", "", http.StatusOK},
		{"/healthz", "", http.StatusOK},
	}
	for _, tt := range tests {
		res, _ := get(t, s, tt.target)
		if res.StatusCode != tt.status || res.Header.Get("Location") != tt.location {
			t.Errorf("%s: got %d %q, want %d %q", tt.target, res.StatusCode, res.Header.Get("Location"), tt.status, tt.location)
		}
	}
	res, _ := get(t, s, "//evil.example")
	if loc := res.Header.Get("Location"); strings.HasPrefix(loc, "//") {
		t.Errorf("protocol-relative redirect %q", loc)
	}
}

func TestStaticAndNotFound(t *testing.T) {
	s := newTestServer(t, &fakeStock{}, afterBreak)
	for path, ct := range map[string]string{
		"/nrfnano/nrfnano_bootloader-0.7.0_s140_6.1.1.hex": "application/octet-stream",
		"/site.webmanifest": "application/manifest+json",
		"/logo-96.png":      "image/png",
	} {
		res, _ := get(t, s, path)
		if res.StatusCode != http.StatusOK || res.Header.Get("Content-Type") != ct {
			t.Errorf("%s: %d %q, want 200 %q", path, res.StatusCode, res.Header.Get("Content-Type"), ct)
		}
	}
	if res, _ := get(t, s, "/missing.png"); res.StatusCode != http.StatusNotFound {
		t.Errorf("missing file: status %d", res.StatusCode)
	}
	if res, _ := get(t, s, "/nope/"); res.StatusCode != http.StatusNotFound {
		t.Errorf("unknown page: status %d", res.StatusCode)
	}
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/usbmux/", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST page: status %d", rec.Code)
	}
}

func TestRobotsAndSitemap(t *testing.T) {
	s := newTestServer(t, &fakeStock{}, afterBreak)
	_, robots := get(t, s, "/robots.txt")
	if robots != "User-agent: *\nDisallow: \nSitemap: https://badgerd.nl/sitemap.xml" {
		t.Errorf("robots.txt = %q", robots)
	}
	res, sitemap := get(t, s, "/sitemap.xml")
	if !strings.HasPrefix(res.Header.Get("Content-Type"), "application/xml") {
		t.Errorf("sitemap content-type %q", res.Header.Get("Content-Type"))
	}
	locs := regexp.MustCompile(`<loc>([^<]+)</loc>`).FindAllStringSubmatch(sitemap, -1)
	var got []string
	for _, m := range locs {
		got = append(got, m[1])
	}
	want := []string{"legal", "nrfnano", "privacy", "refund", "sdwire3", "sdwirec", "success", "terms", "usbmux", ""}
	if len(got) != len(want) {
		t.Fatalf("sitemap locs = %v", got)
	}
	for i, w := range want {
		exp := "https://badgerd.nl/" + w
		if w != "" {
			exp += "/"
		}
		if got[i] != exp {
			t.Errorf("loc[%d] = %q, want %q", i, got[i], exp)
		}
	}
}

func TestHealthz(t *testing.T) {
	now := time.Date(2026, time.October, 1, 12, 0, 0, 0, time.UTC)
	s := New(Config{
		BaseURL:     "https://badgerd.nl",
		GovernorURL: mustURL(t, "http://governor.invalid"),
		GiphyURL:    mustURL(t, "http://giphy.invalid"),
		Now:         func() time.Time { return now },
		Logger:      discardLogger(),
		Info:        Info{Commit: "abc123", Environment: "staging", ImageDigest: "sha256:feed"},
	}, &fakeStock{})
	now = now.Add(90 * time.Second)

	res, body := get(t, s, "/healthz")
	if res.StatusCode != http.StatusOK || res.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("got %d %q", res.StatusCode, res.Header.Get("Content-Type"))
	}
	var got map[string]any
	if err := json.Unmarshal([]byte(body), &got); err != nil {
		t.Fatalf("%v: %s", err, body)
	}
	want := map[string]any{
		"status":         "ok",
		"commit":         "abc123",
		"environment":    "staging",
		"image_digest":   "sha256:feed",
		"started_at":     "2026-10-01T12:00:00Z",
		"uptime_seconds": float64(90),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
