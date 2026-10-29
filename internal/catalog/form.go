package catalog

import "strconv"

// Stock is the volatile, governor-owned part of a product.
type Stock struct {
	Current     int
	Reserved    int
	IgnoreStock bool // governor accepts orders regardless of stock (pre-order / bulk-only)
}

// Availability is the purchase state shown next to the price.
type Availability int

const (
	OutOfStock Availability = iota
	InStock
	PreOrder
)

// SchemaURL is the schema.org ItemAvailability for JSON-LD offers.
func (a Availability) SchemaURL() string {
	switch a {
	case InStock:
		return "https://schema.org/InStock"
	case PreOrder:
		return "https://schema.org/PreOrder"
	default:
		return "https://schema.org/OutOfStock"
	}
}

// CountTone colours the stock-count badge.
type CountTone int

const (
	ToneNone CountTone = iota
	ToneLow            // orange: fewer than 10 left
	ToneOK             // green
)

// Form is everything the purchase form needs to render.
type Form struct {
	Availability Availability
	AvailLabel   string

	CountLabel string // "Last 5 items", "50+ items", or "" (no badge)
	CountTone  CountTone

	Quantities  []int // options for the quantity <select>; empty when nothing can be ordered
	DefaultQty  int
	Disabled    bool // inputs and submit disabled
	ButtonLabel string

	ShowContact bool // bulk/B2B prompt
	Reserved    int  // >0: "N items reserved" notice
}

// smallStockLimit: below this, customers may order up to the whole remaining stock.
const smallStockLimit = 20

// Form derives the purchase-form state from live stock. Rules mirror the
// former .vue templates (see migration_guide.md §6), with these deliberate fixes:
//   - no "Last 0 items" badge when stock is 0;
//   - stock of exactly 20 (nrfnano/usbmux) or 20–95 (sdwirec) no longer yields an
//     empty quantity list; quantities run MinQty..min(stock, MaxQty) instead;
//   - the submit label is "Pre Order" only for pre-orders, "Buy" otherwise.
func (p Product) Form(s Stock) Form {
	cur := max(s.Current, 0)
	f := Form{Reserved: max(s.Reserved, 0)}

	switch {
	case s.IgnoreStock:
		f.Availability, f.AvailLabel = PreOrder, p.PreOrderLabel
	case cur > 0:
		f.Availability, f.AvailLabel = InStock, "In Stock"
	default:
		f.Availability, f.AvailLabel = OutOfStock, "Out of Stock"
	}

	switch {
	case cur == 0:
	case cur < 10:
		f.CountLabel, f.CountTone = "Last "+strconv.Itoa(cur)+" items", ToneLow
	case cur > 50 && p.FiftyPlus:
		f.CountLabel, f.CountTone = "50+ items", ToneOK
	default:
		f.CountLabel, f.CountTone = "Last "+strconv.Itoa(cur)+" items", ToneOK
	}

	switch {
	case s.IgnoreStock:
		f.Quantities = qtyRange(p.MinQty, p.MaxQty)
	case cur > 0 && cur < smallStockLimit:
		f.Quantities = qtyRange(1, cur)
	case cur >= smallStockLimit:
		f.Quantities = qtyRange(p.MinQty, min(cur, p.MaxQty))
	}
	if len(f.Quantities) > 0 {
		f.DefaultQty = f.Quantities[0]
	}

	f.Disabled = f.Availability == OutOfStock
	f.ButtonLabel = "Buy"
	if f.Availability == PreOrder {
		f.ButtonLabel = "Pre Order"
	}
	f.ShowContact = cur > 0 || s.IgnoreStock
	return f
}

func qtyRange(lo, hi int) []int {
	if lo < 1 {
		lo = 1
	}
	if hi < lo {
		return nil
	}
	q := make([]int, 0, hi-lo+1)
	for n := lo; n <= hi; n++ {
		q = append(q, n)
	}
	return q
}
