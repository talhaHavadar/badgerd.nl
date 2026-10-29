# badgerd.nl

Product site for Badgerd Technologies B.V.: a single Go binary rendering [templ](https://templ.guide) views, with live
stock from `governor.badgerd.nl` and same-origin `/api` + `/gif` reverse proxies. It replaced the former Nuxt 2 app;
see [migration_guide.md](migration_guide.md) for the design and the migration log.

Requirements: Go 1.26+, [just](https://github.com/casey/just), curl. templ is pinned in `go.mod` (`go tool templ`);
the Tailwind v3 standalone CLI is downloaded (checksum-verified) into `bin/` by `just tools`.

```bash
just build        # templ generate -> Tailwind CSS -> bin/badgerd-web (static assets embedded)
just run          # build and serve on :5000 (PORT=5001 just run for another port)
just test         # go test ./...
just regress http://localhost:5001   # capture a snapshot and diff it against regression/baseline
just image        # podman build; `just up` / `just down` for podman compose (127.0.0.1:${HOST_PORT:-5000})
just logos        # regenerate logo PNGs from web/source/logo-nobg.svg
```

Runtime configuration (env): `PORT` (5000), `BASE_URL` (`https://badgerd.nl`), `GOVERNOR_BASE_URL`
(`https://governor.badgerd.nl`), `GIPHY_BASE_URL` (`https://media.giphy.com`), `STATIC_DIR` (serve assets from disk
while developing, e.g. `STATIC_DIR=web/static`), `TZ`.

Layout:

- `cmd/web` — entrypoint (config, graceful shutdown, `-healthcheck` probe).
- `internal/server` — routes, trailing-slash policy, page handlers, `/api` + `/gif` reverse proxies, sitemap/robots.
- `internal/governor` — typed governor client (stock, shipping rates) with timeout and short cache.
- `internal/catalog` — static product data and the stock → purchase-form rules.
- `internal/seo` — per-page meta and JSON-LD.
- `web/components`, `web/pages` — templ views; `web/css/input.css` + `web/tailwind.config.js` — styles;
  `web/static` — assets served at `/` (embedded into the binary).

Shell snippets in templ must use ``@components.CodeBlock(`...`)`` (templ collapses whitespace in element text), and
text starting with `for`/`if`/`switch` or containing a standalone `@` must be wrapped in `{ "..." }`.

## Deployment

Pushing to `main` runs `.github/workflows/cd.yml` on the self-hosted runner: it builds the image with
`podman compose up -d --build` (tests run inside the build) and serves it on `127.0.0.1:5000`, where nginx proxies to
it. `.github/workflows/ci.yml` runs vet/test/build on pushes and pull requests. To roll back, revert the offending
commit on `main` (or redeploy an older checkout with `just up`).
