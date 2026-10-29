// Package governor is a typed client for the governor.badgerd.nl backend,
// which owns live Stripe stock data and shipping rates.
package governor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// Product is the volatile per-product data served by GET /stripe/products.
type Product struct {
	Codename      string `json:"codename"`
	ID            string `json:"id"`
	PriceID       string `json:"price_id"`
	CurrentStock  int    `json:"current_stock"`
	ReservedStock int    `json:"reserved_stock"`
	IgnoreStock   bool   `json:"ignore_stock"`
	URL           string `json:"url"`
}

// Country is a shippable destination from GET /shippingrates.
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Client fetches and briefly caches governor data. It is safe for concurrent use.
type Client struct {
	base *url.URL
	http *http.Client
	ttl  time.Duration
	now  func() time.Time

	mu        sync.Mutex
	products  cacheEntry[[]Product]
	countries cacheEntry[[]Country]
}

type cacheEntry[T any] struct {
	value   T
	expires time.Time
	ok      bool
}

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient overrides the HTTP client (and therefore the request timeout).
func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.http = h } }

// WithCacheTTL sets how long successful responses are reused. Zero disables caching.
func WithCacheTTL(d time.Duration) Option { return func(c *Client) { c.ttl = d } }

// WithClock overrides the time source (for tests).
func WithClock(now func() time.Time) Option { return func(c *Client) { c.now = now } }

// DefaultTimeout bounds every governor request so a slow backend can't stall page renders.
const DefaultTimeout = 3 * time.Second

// DefaultCacheTTL keeps stock "live" while absorbing bursts of page views.
const DefaultCacheTTL = 20 * time.Second

// New returns a Client for the governor at baseURL (e.g. https://governor.badgerd.nl).
func New(baseURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(strings.TrimRight(baseURL, "/"))
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("governor: invalid base URL %q", baseURL)
	}
	c := &Client{
		base: u,
		http: &http.Client{Timeout: DefaultTimeout},
		ttl:  DefaultCacheTTL,
		now:  time.Now,
	}
	for _, o := range opts {
		o(c)
	}
	return c, nil
}

// Products returns all products known to governor.
func (c *Client) Products(ctx context.Context) ([]Product, error) {
	return cached(c, &c.products, func() ([]Product, error) {
		var ps []Product
		if err := c.getJSON(ctx, "/stripe/products", &ps); err != nil {
			return nil, err
		}
		return ps, nil
	})
}

// Product returns the product with the given codename. found is false when
// governor answered but doesn't list the codename.
func (c *Client) Product(ctx context.Context, codename string) (p Product, found bool, err error) {
	ps, err := c.Products(ctx)
	if err != nil {
		return Product{}, false, err
	}
	for _, p := range ps {
		if p.Codename == codename {
			return p, true, nil
		}
	}
	return Product{}, false, nil
}

// Countries returns the shippable countries sorted by name.
func (c *Client) Countries(ctx context.Context) ([]Country, error) {
	return cached(c, &c.countries, func() ([]Country, error) {
		var raw json.RawMessage
		if err := c.getJSON(ctx, "/shippingrates", &raw); err != nil {
			return nil, err
		}
		return ParseShippingRates(raw)
	})
}

// ParseShippingRates accepts both response shapes governor has used:
//
//	{"countries":[{"code":"AR","name":"Argentina",...}, ...]}
//	{"AR":{"name":"Argentina",...}, ...}
//
// and returns the countries sorted by name.
func ParseShippingRates(data []byte) ([]Country, error) {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, fmt.Errorf("governor: decode shipping rates: %w", err)
	}

	var countries []Country
	if list, ok := obj["countries"]; ok {
		if err := json.Unmarshal(list, &countries); err != nil {
			return nil, fmt.Errorf("governor: decode shipping rates countries: %w", err)
		}
	} else {
		for code, v := range obj {
			var entry struct {
				Name string `json:"name"`
			}
			if err := json.Unmarshal(v, &entry); err != nil {
				return nil, fmt.Errorf("governor: decode shipping rate %q: %w", code, err)
			}
			countries = append(countries, Country{Code: code, Name: entry.Name})
		}
	}

	sort.SliceStable(countries, func(i, j int) bool {
		if countries[i].Name != countries[j].Name {
			return countries[i].Name < countries[j].Name
		}
		return countries[i].Code < countries[j].Code
	})
	return countries, nil
}

// cached serves a fresh cache entry, otherwise fetches. On fetch failure a
// stale entry is served (stale stock beats an unpurchasable page); with no
// entry at all the error is returned.
func cached[T any](c *Client, e *cacheEntry[T], fetch func() (T, error)) (T, error) {
	c.mu.Lock()
	if e.ok && c.now().Before(e.expires) {
		v := e.value
		c.mu.Unlock()
		return v, nil
	}
	c.mu.Unlock()

	v, err := fetch()

	c.mu.Lock()
	defer c.mu.Unlock()
	if err != nil {
		if e.ok {
			return e.value, nil
		}
		var zero T
		return zero, err
	}
	if c.ttl > 0 {
		*e = cacheEntry[T]{value: v, expires: c.now().Add(c.ttl), ok: true}
	}
	return v, nil
}

func (c *Client) getJSON(ctx context.Context, path string, dst any) error {
	u := c.base.JoinPath(path)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("governor: GET %s: %w", path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<16))
		return fmt.Errorf("governor: GET %s: %w", path, &StatusError{Code: resp.StatusCode})
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4<<20)).Decode(dst); err != nil {
		return fmt.Errorf("governor: GET %s: decode: %w", path, err)
	}
	return nil
}

// StatusError reports a non-200 governor response.
type StatusError struct{ Code int }

func (e *StatusError) Error() string { return fmt.Sprintf("unexpected status %d", e.Code) }
