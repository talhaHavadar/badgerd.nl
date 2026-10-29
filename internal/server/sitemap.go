package server

import (
	"encoding/xml"
	"net/http"
	"sort"

	"badgerd.nl/internal/catalog"
)

// sitemapPaths lists every HTML route (trailing-slash form), sorted with "/" last
// like the former @nuxtjs/sitemap output.
func sitemapPaths() []string {
	paths := []string{"/success/"}
	for _, p := range catalog.Products {
		paths = append(paths, p.Path())
	}
	for _, lp := range legalPages {
		paths = append(paths, lp.path)
	}
	sort.Strings(paths)
	return append(paths, "/")
}

type urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []struct {
		Loc string `xml:"loc"`
	} `xml:"url"`
}

// handleSitemap serves sitemap.xml. URLs use BaseURL (https), an intentional
// change from the live site's http:// locs (migration_guide.md §9).
func (s *Server) handleSitemap(w http.ResponseWriter, r *http.Request) {
	us := urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for _, p := range sitemapPaths() {
		us.URLs = append(us.URLs, struct {
			Loc string `xml:"loc"`
		}{s.cfg.BaseURL + p})
	}
	out, err := xml.Marshal(us)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Write([]byte(xml.Header[:len(xml.Header)-1]))
	w.Write(out)
}

// handleRobots serves robots.txt, byte-identical to the live site.
func (s *Server) handleRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("User-agent: *\nDisallow: \nSitemap: " + s.cfg.BaseURL + "/sitemap.xml"))
}
