// Command logogen regenerates the raster logos served by the Go site from the
// master artwork (see migration_guide.md §3.4):
//
//   - logo-96.png / logo-128.png: square, transparent nav logos rendered from
//     web/source/logo-nobg.svg. Unfilled paths get the nav colour (zinc-50);
//     paths with an explicit fill (the red apostrophe) keep it, mirroring how
//     the old inline SVG inherited `fill-zinc-*` from its parent.
//   - logo-512.png: the JSON-LD brand logo, a ≤512×512 downscale of
//     web/static/logo-nobg.png.
//
// Usage (from the repository root): just logos
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"golang.org/x/image/draw"
)

func main() {
	svgPath := flag.String("svg", "web/source/logo-nobg.svg", "master SVG artwork")
	pngPath := flag.String("png", "web/static/logo-nobg.png", "master raster artwork (for the JSON-LD logo)")
	outDir := flag.String("out", "web/static", "output directory")
	flag.Parse()

	// zinc-50: rendered with opacity-80 on the zinc-700 nav it reads as the old fill-zinc-200.
	const navFill = "#fafafa"
	for _, size := range []int{96, 128} {
		img, err := renderSVG(*svgPath, size, navFill)
		if err != nil {
			log.Fatal(err)
		}
		if err := writePNG(filepath.Join(*outDir, fmt.Sprintf("logo-%d.png", size)), img); err != nil {
			log.Fatal(err)
		}
	}

	img, err := downscalePNG(*pngPath, 512)
	if err != nil {
		log.Fatal(err)
	}
	if err := writePNG(filepath.Join(*outDir, "logo-512.png"), img); err != nil {
		log.Fatal(err)
	}
}

func renderSVG(path string, size int, fill string) (image.Image, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Give the root <svg> a default fill; paths with their own fill keep it.
	src = bytes.Replace(src, []byte("<svg"), []byte(fmt.Sprintf(`<svg fill="%s"`, fill)), 1)
	icon, err := oksvg.ReadIconStream(bytes.NewReader(src), oksvg.IgnoreErrorMode)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	// Fit the (wide) artwork into a square canvas, centred, like
	// preserveAspectRatio="xMidYMid meet" did for the inline SVG.
	vw, vh := icon.ViewBox.W, icon.ViewBox.H
	scale := float64(size) / max(vw, vh)
	w, h := vw*scale, vh*scale
	icon.SetTarget((float64(size)-w)/2, (float64(size)-h)/2, w, h)

	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	scanner := rasterx.NewScannerGV(size, size, img, img.Bounds())
	icon.Draw(rasterx.NewDasher(size, size, scanner), 1)
	return img, nil
}

func downscalePNG(path string, maxSide int) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	src, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}
	b := src.Bounds()
	scale := float64(maxSide) / float64(max(b.Dx(), b.Dy()))
	if scale >= 1 {
		return src, nil
	}
	dst := image.NewNRGBA(image.Rect(0, 0, int(float64(b.Dx())*scale+0.5), int(float64(b.Dy())*scale+0.5)))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, b, draw.Src, nil)
	return dst, nil
}

func writePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	enc := png.Encoder{CompressionLevel: png.BestCompression}
	if err := enc.Encode(f, img); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	log.Printf("wrote %s", path)
	return nil
}
