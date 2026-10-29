package catalog

import (
	"slices"
	"testing"
)

func mustProduct(t *testing.T, slug string) Product {
	t.Helper()
	p, ok := BySlug(slug)
	if !ok {
		t.Fatalf("no product %q", slug)
	}
	return p
}

func TestForm(t *testing.T) {
	type want struct {
		avail      Availability
		availLabel string
		count      string
		tone       CountTone
		qtyFirst   int
		qtyLast    int
		qtyLen     int
		disabled   bool
		button     string
		contact    bool
	}
	tests := []struct {
		name  string
		slug  string
		stock Stock
		want  want
	}{
		{"out of stock", "usbmux", Stock{Current: 0},
			want{OutOfStock, "Out of Stock", "", ToneNone, 0, 0, 0, true, "Buy", false}},
		{"low stock", "usbmux", Stock{Current: 5},
			want{InStock, "In Stock", "Last 5 items", ToneLow, 1, 5, 5, false, "Buy", true}},
		{"mid stock", "nrfnano", Stock{Current: 15},
			want{InStock, "In Stock", "Last 15 items", ToneOK, 1, 15, 15, false, "Buy", true}},
		{"exactly 20 is orderable", "nrfnano", Stock{Current: 20},
			want{InStock, "In Stock", "Last 20 items", ToneOK, 1, 20, 20, false, "Buy", true}},
		{"fifty plus", "sdwire3", Stock{Current: 100},
			want{InStock, "In Stock", "50+ items", ToneOK, 1, 20, 20, false, "Buy", true}},
		{"usbmux shows exact count above 50", "usbmux", Stock{Current: 60},
			want{InStock, "In Stock", "Last 60 items", ToneOK, 1, 20, 20, false, "Buy", true}},
		{"pre-order", "usbmux", Stock{Current: 0, IgnoreStock: true},
			want{PreOrder, "Open for Pre-Order", "", ToneNone, 1, 20, 20, false, "Pre Order", true}},
		{"sdwirec bulk with high stock", "sdwirec", Stock{Current: 100},
			want{InStock, "In Stock", "50+ items", ToneOK, 5, 100, 96, false, "Buy", true}},
		{"sdwirec mid stock capped at stock", "sdwirec", Stock{Current: 40},
			want{InStock, "In Stock", "Last 40 items", ToneOK, 5, 40, 36, false, "Buy", true}},
		{"sdwirec small stock allows singles", "sdwirec", Stock{Current: 3},
			want{InStock, "In Stock", "Last 3 items", ToneLow, 1, 3, 3, false, "Buy", true}},
		{"sdwirec bulk-only", "sdwirec", Stock{IgnoreStock: true},
			want{PreOrder, "Only Available for Bulk Orders", "", ToneNone, 5, 100, 96, false, "Pre Order", true}},
		{"negative stock treated as zero", "nrfnano", Stock{Current: -2},
			want{OutOfStock, "Out of Stock", "", ToneNone, 0, 0, 0, true, "Buy", false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := mustProduct(t, tt.slug).Form(tt.stock)
			w := tt.want
			if f.Availability != w.avail || f.AvailLabel != w.availLabel {
				t.Errorf("availability = %v %q, want %v %q", f.Availability, f.AvailLabel, w.avail, w.availLabel)
			}
			if f.CountLabel != w.count || f.CountTone != w.tone {
				t.Errorf("count = %q/%v, want %q/%v", f.CountLabel, f.CountTone, w.count, w.tone)
			}
			if len(f.Quantities) != w.qtyLen {
				t.Fatalf("len(quantities) = %d, want %d", len(f.Quantities), w.qtyLen)
			}
			if w.qtyLen > 0 && (f.Quantities[0] != w.qtyFirst || f.Quantities[len(f.Quantities)-1] != w.qtyLast || f.DefaultQty != w.qtyFirst) {
				t.Errorf("quantities %d..%d default %d, want %d..%d", f.Quantities[0], f.Quantities[len(f.Quantities)-1], f.DefaultQty, w.qtyFirst, w.qtyLast)
			}
			if f.Disabled != w.disabled || f.ButtonLabel != w.button || f.ShowContact != w.contact {
				t.Errorf("disabled=%v button=%q contact=%v, want %v %q %v", f.Disabled, f.ButtonLabel, f.ShowContact, w.disabled, w.button, w.contact)
			}
		})
	}
}

func TestFormReserved(t *testing.T) {
	f := mustProduct(t, "usbmux").Form(Stock{Current: 3, Reserved: 2})
	if f.Reserved != 2 {
		t.Fatalf("Reserved = %d, want 2", f.Reserved)
	}
}

func TestCatalogIntegrity(t *testing.T) {
	seen := map[string]bool{}
	for _, p := range Products {
		if seen[p.Slug] || seen[p.Codename] {
			t.Errorf("duplicate slug/codename %q/%q", p.Slug, p.Codename)
		}
		seen[p.Slug], seen[p.Codename] = true, true
		if p.MinQty < 1 || p.MaxQty < p.MinQty || len(p.Images) == 0 || p.Title == "" || p.Description == "" {
			t.Errorf("%s: incomplete catalog entry", p.Slug)
		}
		if !slices.ContainsFunc(p.Breadcrumbs, func(c Crumb) bool { return c.Path == "/" }) {
			t.Errorf("%s: breadcrumbs missing Store", p.Slug)
		}
	}
}
