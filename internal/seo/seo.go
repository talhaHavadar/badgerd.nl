// Package seo models per-page <head> metadata and schema.org JSON-LD.
package seo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"badgerd.nl/internal/catalog"
)

const (
	DefaultTitle       = "Badgerd Technologies B.V."
	DefaultDescription = "Discover our innovative solutions for interesting problems in our life. Buy one of our products to support us!"

	brandName   = "Badgerd"
	brandSlogan = "Creating software for hardware, hardware for humanity"
)

// Meta is the per-page <head> content.
type Meta struct {
	Title       string
	Description string
	Canonical   string
	OG          *OpenGraph
	JSONLD      []any // rendered as one <script type="application/ld+json"> array
}

// OpenGraph tags.
type OpenGraph struct {
	Title       string
	Description string
	Type        string
	Image       string
	URL         string
	Locale      string
}

// JSONLDScript marshals the JSON-LD payload. encoding/json escapes <, > and &,
// so the result is safe to embed in a <script> element.
func (m Meta) JSONLDScript() (string, error) {
	if len(m.JSONLD) == 0 {
		return "", nil
	}
	b, err := json.Marshal(m.JSONLD)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Page returns Meta for a page without OpenGraph data (legal pages, success).
func Page(baseURL, path string) Meta {
	return Meta{Title: DefaultTitle, Description: DefaultDescription, Canonical: baseURL + path}
}

// Index returns the landing-page Meta.
func Index(baseURL string) Meta {
	return Meta{
		Title:       "Sett | Badgerd Technologies",
		Description: DefaultDescription,
		// The live site's canonical for "/" has no trailing slash; og:url does.
		Canonical: baseURL,
		OG: &OpenGraph{
			Title:       "Badgerd Technologies",
			Description: DefaultDescription,
			Type:        "website",
			URL:         baseURL + "/",
			Locale:      "en_US",
		},
		JSONLD: []any{
			WebPage{Context: "http://schema.org", ID: baseURL + "/#webpage", Type: "WebPage", URL: baseURL + "/", Name: "Badgerd Technologies"},
			Breadcrumbs(baseURL, catalog.DefaultBreadcrumbs),
		},
	}
}

// Product returns the Meta for a product page; offer availability follows live stock.
func Product(baseURL string, p catalog.Product, avail catalog.Availability, now time.Time) Meta {
	pageURL := baseURL + p.Path()
	return Meta{
		Title:       p.Title,
		Description: p.Description,
		Canonical:   pageURL,
		OG: &OpenGraph{
			Title:       p.OGTitle,
			Description: p.Description,
			Type:        "website",
			Image:       p.OGImage,
			URL:         pageURL,
			Locale:      "en_US",
		},
		JSONLD: []any{
			ProductLD{
				Context:          "http://schema.org",
				ID:               p.LDID,
				Type:             "Product",
				URL:              pageURL,
				Name:             p.LDName,
				MainEntityOfPage: pageURL,
				Image:            p.LDImage,
				Description:      p.LDDescription,
				Brand:            Brand{Logo: baseURL + "/logo-512.png", Name: brandName, Slogan: brandSlogan},
				AggregateRating:  AggregateRating{Type: "AggregateRating", RatingValue: 5, ReviewCount: 1},
				Review: []Review{{
					Type:          "Review",
					Author:        Person{Type: "Person", Name: p.Review.Author},
					DatePublished: p.Review.Date,
					ReviewBody:    p.Review.Body,
					Name:          "First User",
					ReviewRating:  Rating{Type: "Rating", BestRating: 5, RatingValue: 5, WorstRating: 1},
				}},
				Offers: []Offer{{
					Type:            "Offer",
					PriceCurrency:   "EUR",
					Availability:    avail.SchemaURL(),
					Price:           p.PriceEUR,
					SKU:             p.Codename,
					PriceValidUntil: fmt.Sprintf("%d-12-31", now.Year()),
				}},
			},
			Breadcrumbs(baseURL, p.Breadcrumbs),
		},
	}
}

// Breadcrumbs builds a BreadcrumbList with sequential positions.
func Breadcrumbs(baseURL string, crumbs []catalog.Crumb) BreadcrumbList {
	bl := BreadcrumbList{Context: "https://schema.org/", Type: "BreadcrumbList"}
	for i, c := range crumbs {
		bl.ItemListElement = append(bl.ItemListElement, ListItem{
			Type: "ListItem", Position: i + 1, Name: c.Name, Item: baseURL + c.Path,
		})
	}
	return bl
}

// TrimBase normalises a configured base URL (no trailing slash).
func TrimBase(baseURL string) string { return strings.TrimRight(baseURL, "/") }

// JSON-LD types; field order matches the former nuxt-jsonld output.

type WebPage struct {
	Context string `json:"@context"`
	ID      string `json:"@id"`
	Type    string `json:"@type"`
	URL     string `json:"url"`
	Name    string `json:"name"`
}

type ProductLD struct {
	Context          string          `json:"@context"`
	ID               string          `json:"@id"`
	Type             string          `json:"@type"`
	URL              string          `json:"url"`
	Name             string          `json:"name"`
	MainEntityOfPage string          `json:"mainEntityOfPage"`
	Image            string          `json:"image"`
	Description      string          `json:"description"`
	Brand            Brand           `json:"brand"`
	AggregateRating  AggregateRating `json:"aggregateRating"`
	Review           []Review        `json:"review"`
	Offers           []Offer         `json:"offers"`
}

type Brand struct {
	Logo   string `json:"logo"`
	Name   string `json:"name"`
	Slogan string `json:"slogan"`
}

type AggregateRating struct {
	Type        string `json:"@type"`
	RatingValue int    `json:"ratingValue"`
	ReviewCount int    `json:"reviewCount"`
}

type Review struct {
	Type          string `json:"@type"`
	Author        Person `json:"author"`
	DatePublished string `json:"datePublished"`
	ReviewBody    string `json:"reviewBody"`
	Name          string `json:"name"`
	ReviewRating  Rating `json:"reviewRating"`
}

type Person struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type Rating struct {
	Type        string `json:"@type"`
	BestRating  int    `json:"bestRating"`
	RatingValue int    `json:"ratingValue"`
	WorstRating int    `json:"worstRating"`
}

type Offer struct {
	Type            string  `json:"@type"`
	PriceCurrency   string  `json:"priceCurrency"`
	Availability    string  `json:"availability"`
	Price           float64 `json:"price"`
	SKU             string  `json:"sku"`
	PriceValidUntil string  `json:"priceValidUntil"`
}

type BreadcrumbList struct {
	Context         string     `json:"@context"`
	Type            string     `json:"@type"`
	ItemListElement []ListItem `json:"itemListElement"`
}

type ListItem struct {
	Type     string `json:"@type"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Item     string `json:"item"`
}
