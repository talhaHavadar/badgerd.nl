package pages

import (
	"fmt"

	"github.com/a-h/templ"

	"badgerd.nl/internal/catalog"
	"badgerd.nl/internal/governor"
	"badgerd.nl/web/components"
)

// productContent is the page-specific copy around the shared product layout.
type productContent struct {
	banner      templ.Component // optional strip above the hero
	minHeight   string          // hero min-height (taller pages compensate for the banner)
	github      string
	links       templ.Component // links next to the GitHub icon
	description templ.Component
	guide       templ.Component // quick-start / FAQ section below the fold
}

var contents = map[string]productContent{
	"sdwire3": {
		minHeight:   "min-h-[calc(100vh-3rem)]",
		github:      "https://github.com/Badger-Embedded/SDWire3",
		links:       sdwire3Links(),
		description: sdwire3Description(),
		guide:       sdwire3Guide(),
	},
	"nrfnano": {
		minHeight:   "min-h-[calc(100vh-3rem)]",
		github:      "https://github.com/Badger-Embedded/badgerd-nrfnano",
		links:       nrfnanoLinks(),
		description: nrfnanoDescription(),
		guide:       nrfnanoGuide(),
	},
	"sdwirec": {
		minHeight:   "min-h-[calc(100vh-3rem)]",
		github:      "https://github.com/Badger-Embedded/badgerd-sdwirec",
		links:       sdwirecLinks(),
		description: sdwirecDescription(),
		guide:       sdwirecGuide(),
	},
	"usbmux": {
		banner:      usbmuxBanner(),
		minHeight:   "min-h-[calc(100vh-10rem)]",
		github:      "https://github.com/Badger-Embedded/muxer",
		links:       usbmuxLinks(),
		description: usbmuxDescription(),
		guide:       usbmuxGuide(),
	},
}

// Product renders a product page. Every catalog product must have content here
// (enforced by tests).
func Product(page components.Page, p catalog.Product, f catalog.Form, countries []governor.Country) (templ.Component, error) {
	c, ok := contents[p.Slug]
	if !ok {
		return nil, fmt.Errorf("pages: no content for product %q", p.Slug)
	}
	return product(page, p, c, f, countries), nil
}
