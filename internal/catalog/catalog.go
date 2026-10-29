// Package catalog holds the static product metadata (copy, pricing, SEO,
// purchase rules) that used to be hard-coded in each pages/<product>.vue.
// Only volatile stock data comes from governor at request time.
package catalog

// Image is a carousel slide.
type Image struct {
	Src string
	Alt string
}

// Review feeds the JSON-LD Product.review block.
type Review struct {
	Author string
	Date   string
	Body   string
}

// Crumb is a BreadcrumbList entry.
type Crumb struct {
	Name string
	Path string
}

// Product is the static description of a product page.
type Product struct {
	Slug     string // URL path segment: /<slug>/
	Codename string // governor/Stripe codename, posted to checkout
	Name     string // display name in the purchase form

	PriceEUR     float64 // JSON-LD offer price
	PriceLabel   string  // e.g. "€180.00"
	PriceNote    string  // optional small print after the price, e.g. "(Free Shipping)"
	PlusShipping bool    // append " + Shipping" after the price

	// Purchase rules.
	PreOrderLabel string // availability text when governor sets ignore_stock
	FiftyPlus     bool   // show "50+ items" instead of the exact count above 50
	MinQty        int    // smallest quantity offered once stock allows bulk orders
	MaxQty        int    // largest quantity offered
	ContactPrompt string // bulk/B2B prompt under the form
	ContactEmail  string

	Images []Image

	// SEO.
	Title         string
	Description   string
	OGTitle       string
	OGImage       string
	LDID          string
	LDName        string
	LDImage       string
	LDDescription string
	Review        Review
	Breadcrumbs   []Crumb
}

const (
	defaultContactPrompt = "Looking for a bulk order or B2B?"
	defaultContactEmail  = "info+quote@badgerd.nl"
	defaultPreOrderLabel = "Open for Pre-Order"
)

// DefaultBreadcrumbs is the breadcrumb trail shared by the landing page and most product pages.
var DefaultBreadcrumbs = []Crumb{
	{"Store", "/"},
	{"SDWireC", "/sdwirec/"},
	{"USBMux", "/usbmux/"},
	{"nRFNANO", "/nrfnano/"},
}

// Products lists every product page, in navigation order.
var Products = []Product{
	{
		Slug:          "sdwire3",
		Codename:      "badgerd_sdwire3",
		Name:          "SDWire3",
		PriceEUR:      100,
		PriceLabel:    "€100.00",
		PriceNote:     "(Free Shipping)",
		PreOrderLabel: defaultPreOrderLabel,
		FiftyPlus:     true,
		MinQty:        1,
		MaxQty:        20,
		ContactPrompt: defaultContactPrompt,
		ContactEmail:  defaultContactEmail,
		Images: []Image{
			{"/sdwire3/sdwire3-1.jpg", "SDWire3 Top View"},
			{"/sdwire3/sdwire3-2.jpg", "SDWire3 Bottom View"},
			{"/sdwire3/sdwire3-3.jpg", "SDWire3 Being Used in Remote"},
		},
		Title:         "Buy SDWire3 With USB3 | Quickstart With Badgerd SDWire3",
		Description:   "New generation of SDWireC with SDR104 speeds and faster host interaction with USB3",
		OGTitle:       "Badgerd SDWire3",
		OGImage:       "https://www.badgerd.nl/sdwire3/sdwire3-1.jpg",
		LDID:          "https://badgerd.nl/#SDWire3",
		LDName:        "Badgerd SDWire3 | SDWire | Next Gen.",
		LDImage:       "https://www.badgerd.nl/sdwire3/sdwire3-1.jpg",
		LDDescription: "USB-C Variant of standard SDWire product, now we are calling it as SDWire3",
		Review: Review{
			Author: "Talha",
			Date:   "2025-03-01",
			Body:   "It does what you expect from an SDWire device. And having it with USB-C is another plus for me.",
		},
		Breadcrumbs: []Crumb{
			{"Store", "/"},
			{"SDWire3", "/sdwire3/"},
			{"SDWireC", "/sdwirec/"},
			{"USBMux", "/usbmux/"},
			{"nRFNANO", "/nrfnano/"},
		},
	},
	{
		Slug:          "nrfnano",
		Codename:      "badgerd_nrfnano",
		Name:          "nRFNANO",
		PriceEUR:      19,
		PriceLabel:    "€19.00",
		PlusShipping:  true,
		PreOrderLabel: defaultPreOrderLabel,
		FiftyPlus:     true,
		MinQty:        1,
		MaxQty:        20,
		ContactPrompt: defaultContactPrompt,
		ContactEmail:  defaultContactEmail,
		Images: []Image{
			{"/nrfnano/nrfnano_0.jpg", "nRFNANO Top View"},
			{"/nrfnano/nrfnano_1.jpg", "nRFNANO Top View Plugged In"},
			{"/nrfnano/diagram.png", "nRFNANO Pinout Diagram"},
		},
		Title:         "Buy nRFNANO IoT Development Board | Quickstart With Badgerd nRFNANO (nrf52840)",
		Description:   "The Badgerd nRFNANO is an Arduino ProMicro form factor nrf52840 development board. nRFNANO comes with fully compatible pinout with nice!nano. In addition to that you will have an exposed pad to add extra safety for your charging circuit which allows you to connect a temperature sensor.",
		OGTitle:       "Badgerd nRFNANO",
		OGImage:       "https://www.badgerd.nl/nrfnano/nrfnano_0.jpg",
		LDID:          "https://badgerd.nl/#nrfnano",
		LDName:        "Badgerd nRFNANO | nRFNANO",
		LDImage:       "https://www.badgerd.nl/nrfnano/nrfnano_0.jpg",
		LDDescription: "The Badgerd nRFNANO is an Arduino ProMicro form factor nrf52840 development board.",
		Review: Review{
			Author: "Talha",
			Date:   "2024-01-19",
			Body:   "nRFNANO is super easy to learn embedded development for iot by providing circuitpython as a default firmware",
		},
		Breadcrumbs: DefaultBreadcrumbs,
	},
	{
		Slug:          "sdwirec",
		Codename:      "badgerd_sdwirec",
		Name:          "SDWireC",
		PriceEUR:      85,
		PriceLabel:    "€85.00",
		PlusShipping:  true,
		PreOrderLabel: "Only Available for Bulk Orders",
		FiftyPlus:     true,
		MinQty:        5,
		MaxQty:        100,
		ContactPrompt: defaultContactPrompt,
		ContactEmail:  defaultContactEmail,
		Images: []Image{
			{"/sdwirec/sdwirec-1.jpg", "SDWireC Top View"},
			{"/sdwirec/sdwirec-2.jpg", "SDWireC Bottom View"},
			{"/sdwirec/sdwirec-3.jpg", "SDWireC Being Used in Remote"},
		},
		Title:         "Buy SDWire With USB-C | Quickstart With Badgerd SDWireC",
		Description:   "Our way of implementing the infamous sdwire board by Tizen. Explore our variant of sdwire with an USB-C now it is SDWireC.",
		OGTitle:       "Badgerd SDWireC",
		OGImage:       "https://www.badgerd.nl/sdwirec/badgerd-sdwirec-cad.png",
		LDID:          "https://badgerd.nl/#sdwirec",
		LDName:        "Badgerd SDWireC | SDWire",
		LDImage:       "https://www.badgerd.nl/sdwirec/sdwirec-5.jpg",
		LDDescription: "USB-C Variant of standard SDWire product, now we are calling it as SDWireC",
		Review: Review{
			Author: "Talha",
			Date:   "2023-01-19",
			Body:   "It does what you expect from an SDWire device. And having it with USB-C is another plus for me.",
		},
		Breadcrumbs: DefaultBreadcrumbs,
	},
	{
		Slug:          "usbmux",
		Codename:      "badgerd_usbmux",
		Name:          "USBMux",
		PriceEUR:      180,
		PriceLabel:    "€180.00",
		PlusShipping:  true,
		PreOrderLabel: defaultPreOrderLabel,
		MinQty:        1,
		MaxQty:        20,
		ContactPrompt: "Looking for a bulk order or B2B(VAT-free)?",
		ContactEmail:  "info+quoteusbmux@badgerd.nl",
		Images: []Image{
			{"/usbmux/usbmux_2.jpg", "USBMux Cad View"},
			{"/usbmux/usbmux_1.jpg", "USBMux Cad View"},
			{"/usbmux/usbmux_3.jpg", "USBMux Cad View"},
			{"/usbmux/usbmux_4.jpg", "USBMux Cad View"},
			{"/usbmux/usbmux_5.jpg", "USBMux Cad View"},
			{"/usbmux/usbmux_6.jpg", "USBMux Cad View"},
		},
		Title:         "Buy Badgerd USBMux | Quickstart With USBMux",
		Description:   "USBMux is a tool that helps you to switch connections between 3 USB ports. You can control the connection of the ports from a computer.",
		OGTitle:       "Badgerd USBMux",
		OGImage:       "https://www.badgerd.nl/usbmux/usbmux-1.jpg",
		LDID:          "https://badgerd.nl/#usbmux",
		LDName:        "Badgerd USBMux | USBMux",
		LDImage:       "https://www.badgerd.nl/usbmux/usbmux-1.jpg",
		LDDescription: "USBMux is a USB switch that you can control the connection of 3 different USB Ports",
		Review: Review{
			Author: "Talha",
			Date:   "2023-05-19",
			Body:   "It does what you expect from a USBMux/Switch.",
		},
		Breadcrumbs: DefaultBreadcrumbs,
	},
}

// BySlug returns the product served at /<slug>/.
func BySlug(slug string) (Product, bool) {
	for _, p := range Products {
		if p.Slug == slug {
			return p, true
		}
	}
	return Product{}, false
}

// Path is the canonical (trailing-slash) URL path of the product page.
func (p Product) Path() string { return "/" + p.Slug + "/" }
