// Package web owns the site's static assets and templ views.
package web

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"io/fs"
	"os"
	"strings"
	"sync"
)

//go:embed static
var embedded embed.FS

var (
	mu       sync.RWMutex
	static   fs.FS = mustSub(embedded, "static")
	versions       = map[string]string{}
)

// Static returns the filesystem served at "/".
func Static() fs.FS {
	mu.RLock()
	defer mu.RUnlock()
	return static
}

// UseDir serves static assets from dir instead of the embedded copy (local development).
func UseDir(dir string) {
	mu.Lock()
	defer mu.Unlock()
	static = os.DirFS(dir)
	versions = map[string]string{}
}

// AssetURL returns the URL path for a static asset with a content-hash query
// string, so it can be cached aggressively and still update on deploy.
func AssetURL(name string) string {
	name = strings.TrimPrefix(name, "/")
	mu.RLock()
	v, ok := versions[name]
	fsys := static
	mu.RUnlock()
	if !ok {
		if b, err := fs.ReadFile(fsys, name); err == nil {
			sum := sha256.Sum256(b)
			v = hex.EncodeToString(sum[:4])
		}
		mu.Lock()
		versions[name] = v
		mu.Unlock()
	}
	if v == "" {
		return "/" + name
	}
	return "/" + name + "?v=" + v
}

func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
