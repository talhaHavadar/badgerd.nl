package catalog

import (
	"io/fs"
	"strings"
	"testing"

	"badgerd.nl/web"
)

// Every image a product page references (carousel, og:image, JSON-LD) must be a
// file we actually serve; production once shipped three 404ing og:image URLs.
func TestProductImagesExist(t *testing.T) {
	for _, p := range Products {
		refs := []string{p.OGImage, p.LDImage}
		for _, img := range p.Images {
			refs = append(refs, img.Src)
		}
		for _, ref := range refs {
			name := strings.TrimPrefix(ref, "https://www.badgerd.nl")
			name = strings.TrimPrefix(name, "/")
			if _, err := fs.Stat(web.Static(), name); err != nil {
				t.Errorf("%s: %q is not in web/static: %v", p.Slug, ref, err)
			}
		}
	}
}
