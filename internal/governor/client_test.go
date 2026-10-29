package governor

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestParseShippingRates(t *testing.T) {
	tests := []struct {
		name string
		body string
		want []Country
	}{
		{
			name: "countries array shape",
			body: `{"countries":[{"code":"NL","rate":"shr_1","name":"Netherlands","zone":"1"},{"code":"AR","rate":"shr_2","name":"Argentina","zone":"10"}]}`,
			want: []Country{{"AR", "Argentina"}, {"NL", "Netherlands"}},
		},
		{
			name: "code map shape",
			body: `{"NL":{"name":"Netherlands","rate":"shr_1"},"AR":{"name":"Argentina"},"BE":{"name":"Belgium"}}`,
			want: []Country{{"AR", "Argentina"}, {"BE", "Belgium"}, {"NL", "Netherlands"}},
		},
		{name: "empty map", body: `{}`, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseShippingRates([]byte(tt.body))
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %v, want %v", got, tt.want)
				}
			}
		})
	}

	if _, err := ParseShippingRates([]byte(`[1,2]`)); err == nil {
		t.Fatal("expected error for non-object body")
	}
}

func TestProductLookupAndCache(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stripe/products" {
			http.NotFound(w, r)
			return
		}
		hits.Add(1)
		w.Write([]byte(`[{"codename":"badgerd_usbmux","current_stock":5,"reserved_stock":1,"ignore_stock":false}]`))
	}))
	defer srv.Close()

	now := time.Unix(0, 0)
	c, err := New(srv.URL, WithCacheTTL(10*time.Second), WithClock(func() time.Time { return now }))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	p, found, err := c.Product(ctx, "badgerd_usbmux")
	if err != nil || !found || p.CurrentStock != 5 || p.ReservedStock != 1 {
		t.Fatalf("Product() = %+v, %v, %v", p, found, err)
	}
	if _, found, _ := c.Product(ctx, "badgerd_unknown"); found {
		t.Fatal("unknown codename reported as found")
	}
	if hits.Load() != 1 {
		t.Fatalf("expected 1 upstream hit within TTL, got %d", hits.Load())
	}

	now = now.Add(11 * time.Second)
	if _, _, err := c.Product(ctx, "badgerd_usbmux"); err != nil {
		t.Fatal(err)
	}
	if hits.Load() != 2 {
		t.Fatalf("expected refetch after TTL, got %d hits", hits.Load())
	}
}

func TestServesStaleOnFailureAndErrorsWithoutCache(t *testing.T) {
	var fail atomic.Bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fail.Load() {
			http.Error(w, "boom", http.StatusBadGateway)
			return
		}
		w.Write([]byte(`{"countries":[{"code":"NL","name":"Netherlands"}]}`))
	}))
	defer srv.Close()

	now := time.Unix(0, 0)
	c, _ := New(srv.URL, WithCacheTTL(time.Second), WithClock(func() time.Time { return now }))
	ctx := context.Background()

	if _, err := c.Countries(ctx); err != nil {
		t.Fatal(err)
	}
	fail.Store(true)
	now = now.Add(time.Minute)
	got, err := c.Countries(ctx)
	if err != nil || len(got) != 1 {
		t.Fatalf("expected stale countries, got %v, %v", got, err)
	}

	fresh, _ := New(srv.URL)
	if _, err := fresh.Countries(ctx); err == nil {
		t.Fatal("expected error with no cached value")
	}
}

func TestTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer srv.Close()

	c, _ := New(srv.URL, WithHTTPClient(&http.Client{Timeout: 20 * time.Millisecond}))
	if _, err := c.Products(context.Background()); err == nil {
		t.Fatal("expected timeout error")
	}
}

func TestNewRejectsInvalidBaseURL(t *testing.T) {
	if _, err := New("governor.badgerd.nl"); err == nil {
		t.Fatal("expected error for URL without scheme")
	}
}
