# Migration Guide: Nuxt 2 → Go + templ

> Status: **Phases 0–5 implemented on branch `go-migration`; Phase 6 build/container files done (prod CD unchanged); Phase 7 (staging soak + cutover) pending.** See §15 for the implementation log and deviations from this plan.
> Target stack: **Go (`net/http` stdlib) + [templ](https://templ.guide)** server, Tailwind via the **standalone CLI**, deployed as a single static binary.
> Author aid: generated from a full audit of this repository and the live site `https://badgerd.nl/` on 2026-10-28.

---

## 1. Executive summary

`badgerd.nl` is a small product/marketing site for Badgerd Technologies B.V. It presents 5 product pages and several legal pages, and it integrates with the backend **`governor.badgerd.nl`** for live Stripe product/stock data, shipping rates, and checkout.

Although the content is mostly static, **the site is not statically hosted today** — it runs as a **Nuxt 2 SSR Node process** behind **nginx**, managed by `pm2` (`pm2 restart badgerd-nl`). The Node process does three jobs that any replacement must preserve:

1. **Server-side rendering** of pages, including **live stock/price** injected at request time (needed for SEO and correct in/out-of-stock UI).
2. A **same-origin reverse proxy**: `/api/* → https://governor.badgerd.nl` (strips `/api`) and `/gif/* → https://media.giphy.com`. This lets the checkout `<form>` POST same-origin and avoids CORS (governor does **not** send CORS headers on `/stripe/*`).
3. **SEO surface**: `sitemap.xml`, `robots.txt`, per-page `<title>`/meta/canonical/OpenGraph, and JSON-LD.

The migration replaces the Node/Nuxt runtime with a **single Go binary** that renders templ components, injects live stock server-side, and reverse-proxies `/api` and `/gif` using the standard library (`net/http/httputil.ReverseProxy`). No `node_modules` at runtime; Tailwind is compiled at build time via the standalone CLI binary.

**Why Go + templ (recap):** 1:1 architectural match with the current SSR server; type-checked, component-oriented templates (closest DX to the existing `.vue` components); single static binary; stdlib covers HTTP serving, templating, and proxying. See §3 for the detailed rationale.

---

## 2. Current architecture (audit findings)

### 2.1 Stack

- **Framework:** Nuxt 2.15 / Vue 2.7, `@nuxt/content` (**unused** — only the stock `content/hello.md` placeholder exists; safe to drop), `@nuxtjs/axios` + `@nuxtjs/proxy`, `@nuxtjs/sitemap`, `@nuxtjs/robots`, `nuxt-jsonld`, `@nuxtjs/google-fonts`.
- **Styling:** Tailwind CSS 3 (`@nuxtjs/tailwindcss`) + `tw-elements` (Bootstrap-like components, provides the image **carousel** and ripple buttons).
- **Deploy:** `.github/workflows/cd.yml` → on push to `main`, self-hosted runner runs `yarn && yarn build && pm2 restart badgerd-nl`. Served behind `nginx/1.24.0`.
- **Toolchain quirk:** scripts require `NODE_OPTIONS=--openssl-legacy-provider` (Node ≥17 OpenSSL 3 workaround for old webpack) — a smell the migration removes.

### 2.2 Routes / pages (`pages/*.vue`)

| Route       | File                 | Type               | Dynamic data                   |
| ----------- | -------------------- | ------------------ | ------------------------------ |
| `/`         | `index.vue` (508 ln) | Landing / store    | none (JSON-LD WebPage)         |
| `/sdwire3/` | `sdwire3.vue` (474)  | Product            | stock + shipping (SSR)         |
| `/sdwirec/` | `sdwirec.vue` (568)  | Product            | stock + shipping (SSR)         |
| `/nrfnano/` | `nrfnano.vue` (470)  | Product            | stock + shipping (SSR)         |
| `/usbmux/`  | `usbmux.vue` (609)   | Product            | stock + shipping (SSR)         |
| `/success/` | `success.vue` (76)   | Order confirmation | client-side fetch + random gif |
| `/terms/`   | `terms.vue` (300)    | Legal              | none                           |
| `/privacy/` | `privacy.vue` (169)  | Legal              | none                           |
| `/refund/`  | `refund.vue` (72)    | Legal              | none                           |
| `/legal/`   | `legal.vue` (34)     | Legal              | none                           |

Shared chrome: `layouts/default.vue` (306 ln) — top nav, footer, a seasonal "lightrope" lights banner + "on a break" notice gated by `new Date()`), and the custom lightrope CSS lives in that file's `<style>` block (the only non-Tailwind CSS). `components/Logo.vue` is a **1.5 MB inline SVG** logo (replaced by a small PNG in the migration — see §3.4).

### 2.3 Dynamic behavior to preserve (the hard part)

**Product data model** (from live `GET /api/stripe/products`):

```json
{
  "codename": "badgerd_usbmux",
  "id": "prod_…",
  "price_id": "price_…",
  "current_stock": 0,
  "reserved_stock": 0,
  "ignore_stock": false,
  "url": "https://www.badgerd.nl/usbmux"
}
```

Each product page does, in `asyncData` (SSR):

1. `GET /api/stripe/products` → find product by hard-coded `codename` (`badgerd_sdwire3`, `badgerd_sdwirec`, `badgerd_nrfnano`, `badgerd_usbmux`).
2. `GET /api/shippingrates` → build sorted country list. Response has two shapes handled in code: `{countries:[{code,name,…}]}` **or** a map `{AR:{name,…},…}`. (Live currently returns the `{countries:[…]}` shape.)

**Stock-driven UI states** (must be reproduced exactly):

- `current_stock == 0 && !ignore_stock` → **Out of Stock**; disable quantity/country selects + submit button.
- `!ignore_stock && current_stock > 0` → **In Stock** (green).
- `ignore_stock` → **Open for Pre-Order** (orange), button label becomes **"Pre Order"** / pre-order.
- `0 < current_stock < 10` → "Last N items" (orange); `10 ≤ current_stock` non-zero → "Last N items" (green) (per current template conditions).
- Quantity options: `1..current_stock` when `0 < current_stock < 20`; else `1..20` when `current_stock > 20 || ignore_stock`.
- `reserved_stock > 0` → "N items reserved, might be yours, try again in 30 minutes".

> **Note (promo codes):** the current Nuxt template hard-codes a pre-order discount note on every pre-order product. Discount codes are **not** universal across products, so the migration deliberately **does not** carry any discount-code copy into the shared product template. If a promo is ever needed, model it as an optional per-product field in `internal/catalog` and render it only for products that set it — never as a blanket pre-order message.

**Checkout form** (same-origin POST, proxied to governor):

```html
<form action="/api/stripe/checkout" method="post">
  <input type="hidden" name="codename" value="badgerd_usbmux" />
  <select name="quantity" required>
    …
  </select>
  <select name="country_code" required>
    …(countries)…
  </select>
  <button type="submit">Buy / Pre Order</button>
</form>
```

Governor responds with a redirect to Stripe Checkout. **No checkout logic is reimplemented** — the Go server only forwards the POST.

**`/success/`** (`success.vue`): reads `order_id` + `codename` from query string, client-side `GET /api/stripe/products` to resolve product, shows a random Giphy gif via the `/gif` proxy. This stays client-side JS.

### 2.4 Reverse proxy (from `nuxt.config.js`)

```
/api/*  → https://governor.badgerd.nl   (pathRewrite ^/api → '')
/gif/*  → https://media.giphy.com        (pathRewrite ^/gif → '')
```

Verified on live: `/api/stripe/products`, `/api/shippingrates`, `/gif/...giphy.gif` all 200. Direct `governor.badgerd.nl/stripe/products` works but sends **no `Access-Control-Allow-Origin`** → the proxy (or equivalent CORS) is mandatory.

### 2.5 SEO surface (verified live)

- `GET /robots.txt` → `User-agent: *` / `Disallow: ` (**empty — everything allowed**, not `Disallow: /`) / `Sitemap: https://badgerd.nl/sitemap.xml`, no trailing newline (64 bytes, per the 2026-10-28 baseline).
- `GET /sitemap.xml` → lists all 10 routes with **trailing slash** and — note — **`http://` scheme** (`<loc>http://badgerd.nl/legal/</loc>`). Decide whether to keep `http` or fix to `https` (see §9).
- **Trailing-slash redirects (live behavior, for reference):** product routes 301-redirect to slash (`/usbmux` → **301** → `/usbmux/`), but the legal routes currently serve **200 at both** `/terms` and `/terms/` (no redirect). Verified live 2026-10-28. **Migration policy decision (§9): apply one global rule — every non-slash path 301-redirects to its trailing-slash form** (no per-route exceptions). This standardizes the legal routes and is documented as an intentional change.
- Per-page `<title>`, `meta[name=description]`, canonical, OG tags, and `application/ld+json` blocks (Product + AggregateRating + Offer + BreadcrumbList on product pages; WebPage on index). Example captured for `usbmux` in Appendix B.
- `og:image`/JSON-LD `image` reference `https://www.badgerd.nl/...` absolute URLs.

### 2.6 Assets

- `assets/` (~21 MB, processed by webpack today via `require()` in templates): `assets/css/main.css` (just the 3 `@tailwind` directives), `assets/<product>/*.jpg|png` carousel/CAD images.
- `static/` (~10 MB, served from `/`): favicons, `site.webmanifest`, `logo-nobg.png`, per-product hero images, plus downloadable firmware for nrfnano (`.hex`, `.uf2`).
- In Go, **all images move under one served directory** (e.g. `web/static/`) and templates reference them by absolute URL path (no `require()`).
- **⚠ Logo bloat (high-impact):** `components/Logo.vue` is a **1.5 MB** unoptimized Inkscape SVG **inlined into every page** via the layout — it is the main reason the home HTML is ~1.6 MB on the wire. It renders at only **`h-10 w-10` (40×40 px)** in the nav, recolored `fill-zinc-200` → `hover:fill-zinc-50`. The JSON-LD `brand.logo` also points at `static/logo-nobg.png`, which is a **5000×4296 (1.8 MB)** PNG. Both are wildly oversized for their use. See the decision in §3.4.

---

## 3. Target architecture (Go + templ)

```
                        ┌──────────────────────────────────────────┐
  Internet ── nginx ──▶ │  badgerd-web (single Go binary)           │
                        │                                           │
                        │  net/http mux:                            │
                        │   GET  /                → render Index    │
                        │   GET  /usbmux/ etc.    → render Product  │
                        │        (fetch governor stock at request)  │
                        │   GET  /terms/ …        → render Legal     │
                        │   GET  /success/        → render Success   │
                        │   GET  /sitemap.xml     → generated        │
                        │   GET  /robots.txt      → static/handler   │
                        │   GET  /<static asset>  → http.FileServer  │
                        │   ANY  /api/*  ─────────▶ ReverseProxy ──▶ governor.badgerd.nl
                        │   GET  /gif/*  ─────────▶ ReverseProxy ──▶ media.giphy.com
                        │   307/301 trailing-slash normalization     │
                        └──────────────────────────────────────────┘
```

### 3.1 Why these choices

- **`net/http` stdlib server** (no web framework): the routing need is ~12 fixed routes + a file server + two proxies. `http.ServeMux` (Go 1.22+ pattern matching) is enough. Go 1.26 is installed.
- **templ** for views: typed, compiled `.templ` components → generated Go. Mirrors the current component model (layout + page components + a reusable product-form component). Gives compile-time safety on the stock conditionals that are easy to get wrong.
- **`httputil.ReverseProxy`** for `/api` and `/gif`: direct replacement for `@nuxtjs/proxy`, including path rewrite via `Director`/`Rewrite`.
- **Tailwind standalone CLI** (`tailwindcss` single binary, no Node): compiles `web/css/input.css` → `web/static/css/app.css` at build time. The existing `tailwind.config.js` `content` globs are repointed at `**/*.templ`. **`tw-elements` dependency:** it currently supplies the carousel + ripple. Options (pick in Phase 2): (a) keep a tiny vendored `tw-elements` JS + its Tailwind plugin, loaded only where needed; (b) replace the carousel with a ~40-line vanilla-JS/CSS carousel and drop `tw-elements` entirely (preferred — removes the last Node/Tailwind-plugin coupling). The seasonal "lightrope" CSS is hand-written and moves verbatim into `input.css`.

### 3.2 Proposed project layout

```
.
├── cmd/web/main.go              # wiring: mux, routes, proxies, server
├── internal/
│   ├── server/
│   │   ├── routes.go            # route table + trailing-slash middleware
│   │   ├── proxy.go             # /api and /gif reverse proxies
│   │   ├── handlers.go          # page handlers
│   │   ├── sitemap.go           # sitemap.xml + robots.txt
│   │   └── seo.go               # per-page meta + JSON-LD structs
│   ├── governor/
│   │   └── client.go            # typed client: Products(), ShippingRates()
│   └── catalog/
│       └── catalog.go           # product metadata (codename, price, copy, images, SEO)
├── web/
│   ├── components/              # *.templ: layout, nav, footer, product_form, carousel
│   ├── pages/                   # *.templ: index, product, legal, success
│   ├── css/input.css            # @tailwind + lightrope CSS
│   └── static/                  # all images, favicons, manifest, firmware, built app.css
├── tailwind.config.js           # content globs → web/**/*.templ
├── Justfile                     # templ generate, tailwind build, go build, run, container
├── Containerfile                # OCI image build (multi-stage) for podman/docker
├── compose.yaml                 # podman compose service definition
├── go.mod
└── migration_guide.md
```

### 3.3 Data/SSR strategy

- **Static product metadata** (price, name, descriptions, images, SEO copy, JSON-LD, hard-coded `codename`) lives in Go (`internal/catalog`), **not** fetched. Only **volatile** fields (`current_stock`, `reserved_stock`, `ignore_stock`) come from governor at request time.
- **Request-time fetch** `governor.Products()` + `ShippingRates()` on each product-page GET, with:
  - a short **timeout** (e.g. 3s) and **graceful degradation**: if governor is unreachable, render the page with `current_stock` treated as unknown → show the product but disable purchase (match "Out of Stock" safe default) rather than 500. (Improves on today's behavior where an SSR fetch failure breaks the page.)
  - optional **in-process cache** (e.g. 15–30 s TTL) to cut governor load; must stay short enough that stock looks live.
- **Country list**: fetched from `ShippingRates()`; handle both response shapes (array and map) as the current code does.

### 3.4 Logo / asset decision (resolved)

**Decision: replace the 1.5 MB inline SVG header logo with a small raster PNG; stop inlining it.**

Rationale: the header logo renders at 40×40 px and does not need vector detail at that size; inlining 1.5 MB of path data into every HTML response is the single biggest page-weight problem. A transparent PNG exported at 2×–3× the display size is a few KB and cuts ~1.5 MB from **every** page.

Implementation:
- Export `logo-nobg.svg` to a **transparent PNG** sized for retina: **`logo-96.png` (96×96)** served at `h-10 w-10` (40 px). Optionally add `logo-128.png` and use `srcset` for 3× displays. Place under `web/static/`.
- Reference it with a plain `<img src="/logo-96.png" class="h-10 w-10 …" alt="Badgerd">` in `components/logo.templ` (no inline SVG, no `fill-*` recoloring).
- **Hover state:** PNG can't be recolored via Tailwind `fill-*`. Reproduce the subtle light-gray→white hover with opacity instead: render a light/near-white logo PNG and apply `opacity-80 transition hover:opacity-100` (visually equivalent to the current `fill-zinc-200`→`fill-zinc-50`). If an exact two-tone match is required, ship two PNGs and swap on hover, but opacity is simpler and sufficient.
- **JSON-LD `brand.logo`:** generate a right-sized `logo-512.png` (≤512×512) from the source art and point the structured-data `logo` at it instead of the 5000×4296 / 1.8 MB `logo-nobg.png`.
- **Keep the source SVG** (`web/static/logo-nobg.svg` or in `assets/`) in the repo as the master for regenerating PNGs; just don't ship it inline.
- **Alternative considered (not chosen):** run `svgo` to shrink the inline SVG. Rejected because the art is detailed illustration (svgo won't get it tiny) and the 40 px render makes vector precision pointless; a PNG is smaller and simpler. Revisit only if a crisp, arbitrarily-scalable, CSS-recolorable logo becomes a requirement.

Favicons/manifest already exist as correctly-sized PNGs and carry over unchanged.

---

## 4. Implementation plan (phased)

Each phase ends in a buildable, testable state. Do the work on a branch (`feat/go-templ-migration`); keep Nuxt deployable until §12 regression passes.

### Phase 0 — Baseline capture (do first; feeds §12)

- Snapshot the live site: HTML of all 10 routes, `robots.txt`, `sitemap.xml`, response headers, `/api/stripe/products`, `/api/shippingrates`, a `/gif` sample. Store under `regression/baseline/` (git-ignored or a separate dir). Script in §12.1.
- Record live quirks: `http://` sitemap scheme, trailing-slash 301s, `og:image` on `www.` host, current stock values per product.

### Phase 1 — Scaffolding

- `go mod init badgerd.nl` (or module path of choice). Add `templ` (`go install github.com/a-h/templ/cmd/templ@latest`) and the `templ` runtime dep.
- Download the **Tailwind standalone CLI** binary; add a `css` recipe to the `Justfile`.
- `cmd/web/main.go`: `http.Server` on `$PORT` (default 5000, matching current `start` script), graceful shutdown, `http.FileServer` for `web/static`, a `/healthz`.
- Trailing-slash middleware: **global rule** — every request to a non-slash path 301-redirects to its trailing-slash form (one middleware, no per-route policy). This standardizes the legal routes (which the live site currently serves at both forms); documented as an intentional change in §9.

### Phase 2 — Shared chrome + styling

- Port `layouts/default.vue` → `components/layout.templ` (nav, footer, seasonal banner using Go `time.Now()`), `components/logo.templ` (**small PNG `<img>`, not inline SVG — see §3.4**; export `logo-96.png`/`logo-512.png` from the source art).
- Move lightrope CSS into `web/css/input.css`; wire Tailwind build; repoint `tailwind.config.js` content globs to `web/**/*.templ`.
- Decide `tw-elements` fate (§3.1): implement vanilla carousel component `components/carousel.templ` (+ small JS) or vendor `tw-elements`.
- Port the static legal pages (`terms`, `privacy`, `refund`, `legal`) → `pages/legal_*.templ` (pure markup) and the index landing page (no dynamic data).

### Phase 3 — Governor client + product pages

- `internal/governor/client.go`: typed `Product` struct + `Products(ctx)` and `ShippingRates(ctx)` (handle both shapes), timeout, optional cache.
- `internal/catalog`: encode per-product static metadata + SEO + JSON-LD.
- `components/product_form.templ`: the stock/quantity/country/submit logic — **single reusable component** driving all four products (the current pages duplicate this; consolidate).
- `pages/product.templ` + handler: fetch stock, render. Port the long marketing/quick-start copy per product (mostly static HTML; copy verbatim from the `.vue` files).
- `pages/success.templ` + the small client-side JS for order lookup + random gif.

### Phase 4 — SEO endpoints

- `internal/server/seo.go`: per-route `<title>`, description, canonical, OG tags; JSON-LD (`text/templ` `templ.JSONScript` or marshalled structs rendered in a `<script type="application/ld+json">`).
- `sitemap.go`: generate `sitemap.xml` (match route list + trailing slash; choose scheme per §9) and serve `robots.txt`.

### Phase 5 — Proxies

- `internal/server/proxy.go`: two `httputil.ReverseProxy` instances with path rewrite (`/api`→governor root, `/gif`→giphy root). Preserve method/body (POST checkout), set appropriate `Host`/forwarded headers. Confirm the checkout POST round-trips (test against governor staging if available; otherwise inspect the redirect without completing payment).

### Phase 6 — Build, CI/CD, deploy (see §10–11)

- `Justfile` recipes + a `Containerfile`/`compose.yaml`, and an updated `.github/workflows/cd.yml` that builds the OCI image (or the binary) and restarts the service via **podman compose** on the server.

### Phase 7 — Regression + cutover (see §12–13).

---

## 5. Page-by-page mapping

| Nuxt source                                  | Go/templ target                                                  | Notes                                                                                                                           |
| -------------------------------------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `layouts/default.vue`                        | `components/layout.templ` + `footer.templ` + `nav.templ`         | Seasonal banner via `time.Now()`; move `<style>` lightrope CSS → `input.css`.                                                   |
| `components/Logo.vue`                        | `components/logo.templ`                                          | **Replace 1.5 MB inline SVG with small PNG `<img>`** (§3.4); opacity hover instead of `fill-*`.                                |
| `pages/index.vue`                            | `pages/index.templ` + handler                                    | Drop `@nuxt/content` `$content('hello')` (unused placeholder). Keep WebPage JSON-LD + meta.                                     |
| `pages/{sdwire3,sdwirec,nrfnano,usbmux}.vue` | `pages/product.templ` + `catalog` entries + `product_form.templ` | Consolidate 4 near-duplicate pages into 1 template + data. Preserve each page's unique copy/quick-start, images, JSON-LD, meta. |
| `pages/success.vue`                          | `pages/success.templ` + `success.js`                             | Client-side order lookup + random gif retained.                                                                                 |
| `pages/{terms,privacy,refund,legal}.vue`     | `pages/legal_*.templ`                                            | Static; copy markup verbatim; keep canonical.                                                                                   |
| `nuxt.config.js` head/link                   | `components/head.templ` + per-page SEO                           | Global favicons/manifest/canonical/lang.                                                                                        |
| `@nuxtjs/sitemap`,`robots`                   | `sitemap.go`                                                     | Generated.                                                                                                                      |
| `@nuxtjs/proxy`                              | `proxy.go`                                                       | `/api`, `/gif`.                                                                                                                 |
| `nuxt-jsonld`                                | `seo.go` JSON-LD                                                 | Marshalled structs.                                                                                                             |

---

## 6. Dynamic logic reference (port verbatim into `product_form.templ`)

Pseudocode the template must implement (derived from the `.vue` conditions):

```
inStock      = current_stock > 0 && !ignore_stock
outOfStock   = current_stock == 0 && !ignore_stock
preOrder     = ignore_stock
lowStockWarn = current_stock > 0 && current_stock < 10
qtyMax       = (current_stock in 1..19) ? current_stock : ((current_stock > 20 || ignore_stock) ? 20 : 0)
disabled     = outOfStock
buttonLabel  = preOrder ? "Pre Order" : "Buy"
showReserved = reserved_stock > 0
// NOTE: no blanket pre-order discount note — see §2.3. Promo copy, if ever needed,
// is an optional per-product field in internal/catalog, not a shared-template default.
```

Keep the exact label strings and Tailwind color classes for visual parity (orange/green/red). Preserve hidden `codename`, `quantity`, `country_code` field names and `action="/api/stripe/checkout" method="post"`.

> **Audit correction (implemented in `internal/catalog/form.go`):** the four `.vue` pages do *not* share one formula. Per-product differences are now catalog fields:
>
> | Product | Pre-order (`ignore_stock`) label | ">50" badge | Bulk quantities | B2B prompt / mailto |
> | ------- | -------------------------------- | ----------- | --------------- | ------------------- |
> | sdwire3 | Open for Pre-Order | "50+ items" | 1–20 (`>= 20`) | "…bulk order or B2B?" / `info+quote@` |
> | nrfnano | Open for Pre-Order | "50+ items" | 1–20 (`> 20` — bug) | "…bulk order or B2B?" / `info+quote@` |
> | sdwirec | **Only Available for Bulk Orders** | "50+ items" | **5–100** (`>= 96`) | "…bulk order or B2B?" / `info+quote@` |
> | usbmux  | Open for Pre-Order | exact count | 1–20 (`> 20` — bug) | "…B2B(VAT-free)?" / `info+quoteusbmux@` |
>
> Price lines differ too (sdwire3: "(Free Shipping)" without "+ Shipping"; the others "+ Shipping"). The former "(incl. VAT)" notes were removed on request (§15.3).
> Implemented rule: `1..stock` when `0 < stock < 20`; `MinQty..min(stock, MaxQty)` when `stock >= 20`; `MinQty..MaxQty` when `ignore_stock`. See §15 for the bugs this fixes.

---

## 7. Governor client contract

| Call     | Method/Path (via proxy)            | Shape                                                            |
| -------- | ---------------------------------- | ---------------------------------------------------------------- |
| Products | `GET /api/stripe/products`         | `[]Product` (see §2.3)                                           |
| Shipping | `GET /api/shippingrates`           | `{countries:[{code,name,rate,zone}]}` **or** `{<code>:{name,…}}` |
| Checkout | `POST /api/stripe/checkout` (form) | server forwards; governor → redirect to Stripe                   |

Go struct:

```go
type Product struct {
    Codename     string `json:"codename"`
    ID           string `json:"id"`
    PriceID      string `json:"price_id"`
    CurrentStock int    `json:"current_stock"`
    ReservedStock int   `json:"reserved_stock"`
    IgnoreStock  bool   `json:"ignore_stock"`
    URL          string `json:"url"`
}
```

---

## 8. Build & tooling

**`Justfile`** ([just](https://github.com/casey/just) command runner) recipes:

```just
# Justfile
tailwind := "./bin/tailwindcss"   # standalone CLI binary
bin      := "bin/badgerd-web"

# regenerate templ Go from *.templ
generate:
    templ generate ./web/...

# compile + minify Tailwind CSS
css:
    {{tailwind}} -i web/css/input.css -o web/static/css/app.css --minify

# full build: templ -> css -> static Go binary
build: generate css
    CGO_ENABLED=0 go build -o {{bin}} ./cmd/web

# build then run locally on :5000
run: build
    PORT=5000 {{bin}}

test:
    go test ./...

# build the OCI image with podman (see Containerfile)
image tag="badgerd-web:latest":
    podman build -t {{tag}} .

# bring the app up via podman compose
up: 
    podman compose up -d --build

down:
    podman compose down
```

Pin the `templ` and Tailwind CLI versions (commit the Tailwind binary or install in CI). Remove `NODE_OPTIONS=--openssl-legacy-provider` entirely.

### 8.1 Containerization (podman)

Ship the app as an OCI image so it can run under **podman compose** on the server. Use a **multi-stage** build so the final image is a minimal static binary + `web/static` assets (no Go toolchain, no Node).

**`Containerfile`:**
```dockerfile
# ---- build stage ----
FROM docker.io/library/golang:1.26 AS build
WORKDIR /src
# templ + tailwind standalone CLI
RUN go install github.com/a-h/templ/cmd/templ@latest
ADD https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 /usr/local/bin/tailwindcss
RUN chmod +x /usr/local/bin/tailwindcss
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN templ generate ./web/... \
 && tailwindcss -i web/css/input.css -o web/static/css/app.css --minify \
 && CGO_ENABLED=0 GOOS=linux go build -o /out/badgerd-web ./cmd/web

# ---- runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/badgerd-web /app/badgerd-web
COPY --from=build /src/web/static /app/web/static
EXPOSE 5000
USER nonroot:nonroot
ENTRYPOINT ["/app/badgerd-web"]
```
Notes: embedding `web/static` into the binary with Go `embed` is an alternative that yields a **single-file** image (then the runtime stage copies only the binary). Pick one approach and be consistent. Use a distroless/static base for a tiny, rootless-friendly image (good fit for podman).

**`compose.yaml`** (podman compose):
```yaml
services:
  web:
    build:
      context: .
      dockerfile: Containerfile
    image: badgerd-web:latest
    container_name: badgerd-web
    restart: unless-stopped
    ports:
      - "127.0.0.1:5000:5000"      # nginx on the host proxies to this
    environment:
      PORT: "5000"
      BASE_URL: "https://badgerd.nl"
      GOVERNOR_BASE_URL: "https://governor.badgerd.nl"
      GIPHY_BASE_URL: "https://media.giphy.com"
    healthcheck:
      test: ["CMD", "/app/badgerd-web", "-healthcheck"]   # or wget to /healthz
      interval: 30s
      timeout: 3s
      retries: 3
```
Run on the server with `podman compose up -d --build` (or `just up`). nginx continues to terminate TLS and `proxy_pass` to `127.0.0.1:5000`. For boot persistence, generate a systemd unit from the pod/container (`podman generate systemd` or a Quadlet `.container` file) so it starts on reboot.

---

## 9. SEO parity decisions

- **Sitemap scheme:** live emits `http://`. Recommendation: switch to `https://` (canonical host is https) — a deliberate, documented change rather than bug-for-bug parity. Note it in the cutover PR.
- **Trailing slash (global policy):** **all** routes 301-redirect non-slash → slash via a single middleware. Product routes already behave this way; the legal routes currently 200 at both forms, so standardizing them is an intentional, documented change that preserves the canonical slash URLs already in the sitemap.
- **`www.` vs apex in OG/JSON-LD images:** keep existing absolute `https://www.badgerd.nl/...` image URLs (they resolve) to avoid changing social-preview cache keys, unless consolidating hosts is desired.
- **JSON-LD:** reproduce Product/Offer/AggregateRating/Review/BreadcrumbList for product pages and WebPage for index. Set `Offer.availability` from live stock (`InStock`/`OutOfStock`) to improve on the current hard-coded value (optional enhancement — flag in PR).
- Preserve every `<title>` and `meta[name=description]` string exactly (Appendix A).

---

## 10. Deployment

**Primary path: containerized, run with podman compose** (see §8.1 for `Containerfile`/`compose.yaml`).

- The app runs as a rootless **podman** container published on `127.0.0.1:5000`; **nginx** on the host keeps terminating TLS and `proxy_pass`es to it — no nginx change needed.
- Deploy/update on the server: `podman compose up -d --build` (or `just up`). Roll back with `podman compose down` + redeploy the previous image tag.
- **Boot persistence:** generate a systemd unit for the container with `podman generate systemd --new --name badgerd-web` (or a **Quadlet** `badgerd-web.container` file in `~/.config/containers/systemd/`), and `systemctl --user enable` it so it survives reboots.
- **Image distribution:** either build on the self-hosted runner directly on the server host (simplest — runner and server are the same box today), or push to a registry (GHCR) and `podman pull` on the server.

Config via env (in `compose.yaml`): `PORT` (5000), `BASE_URL` (`https://badgerd.nl`, for canonical/sitemap), `GOVERNOR_BASE_URL` (`https://governor.badgerd.nl`), `GIPHY_BASE_URL` (`https://media.giphy.com`).

**Alternative (no container):** build the static binary in CI and run it under a `systemd` unit (`Restart=always`, `Environment=PORT=5000 …`), replacing the current `pm2` process. Kept only as a fallback; the container path is preferred per project direction.

Updated `.github/workflows/cd.yml` (self-hosted runner) outline — container path:

```yaml
- uses: actions/checkout@v4
- run: podman compose up -d --build        # builds Containerfile, restarts the service
# (optional) prune old images:
- run: podman image prune -f
```

Alternative (binary path): `setup-go` → `go install …/templ` → `just build` → `systemctl restart badgerd-web`.

Config via env: `PORT`, `GOVERNOR_BASE_URL` (default `https://governor.badgerd.nl`), `GIPHY_BASE_URL`, `BASE_URL` (for canonical/sitemap, default `https://badgerd.nl`).

---

## 11. Cutover & rollback

1. Deploy the Go container to a **staging port** (e.g. publish `127.0.0.1:5001`) alongside the running Nuxt app. Point nginx at it via a test server block / preview host.
2. Run the full §12 regression suite against staging.
3. Flip nginx `proxy_pass` from `:5000` (Nuxt) to the Go container; reload nginx.
4. **Rollback:** revert nginx `proxy_pass` to the still-running Nuxt process (`pm2 restart badgerd-nl`), or `podman compose down` the new service. Keep the Nuxt deployment installed for at least one release cycle; tag the previous container image so it can be re-pulled.
5. After a stable period, delete Nuxt files, Node deps, and the `openssl-legacy-provider` scripts.

---

## 12. Regression testing against `https://badgerd.nl/`

Goal: prove the Go site is behaviorally equivalent to the live Nuxt site (minus intentional changes in §9). Combine **automated diffing** with a **manual checklist**.

### 12.1 Automated: capture baseline then diff

Capture from **production** (Phase 0) and from **staging** (Phase 7), then diff normalized output.

Baseline capture script (`regression/capture.sh`):

```bash
#!/usr/bin/env bash
set -euo pipefail
BASE="${1:?usage: capture.sh <base-url> <out-dir>}"
OUT="${2:?}"; mkdir -p "$OUT"
ROUTES=( / /sdwire3/ /sdwirec/ /nrfnano/ /usbmux/ /success/ /terms/ /privacy/ /refund/ /legal/ )
for r in "${ROUTES[@]}"; do
  f="$OUT/$(echo "$r" | sed 's#/#_#g')"
  curl -sS "$BASE$r" -o "$f.html"
  curl -sS -D "$f.headers" -o /dev/null "$BASE$r"
done
curl -sS "$BASE/robots.txt"  -o "$OUT/robots.txt"
curl -sS "$BASE/sitemap.xml" -o "$OUT/sitemap.xml"
curl -sS "$BASE/api/stripe/products" -o "$OUT/products.json"
curl -sS "$BASE/api/shippingrates"   -o "$OUT/shippingrates.json"
# trailing-slash redirect + gif proxy status
for r in /usbmux /sdwirec /nrfnano /sdwire3; do
  echo "$r -> $(curl -sS -o /dev/null -w '%{http_code} %{redirect_url}' "$BASE$r")"
done > "$OUT/redirects.txt"
curl -sS -o /dev/null -w '%{http_code} %{content_type}\n' \
  "$BASE/gif/media/l0amJzVHIAfl7jMDos/giphy.gif" > "$OUT/gif.txt"
```

Run: `./regression/capture.sh https://badgerd.nl regression/baseline` and later `./regression/capture.sh http://localhost:5001 regression/candidate`.

**Normalization before diffing** (stock changes and ordering cause false diffs):

- Strip volatile numbers: stock counts, "Last N items", `reserved_stock` messages, random gif URLs, dates (`date:` header), ETags, `content-length`.
- Pretty-print HTML (e.g. `tidy -q -i` or `xmllint --format`) and JSON (`jq -S`) before `diff`.
- Compare **structure and copy**, not byte-for-byte. A helper `regression/diff.sh` that runs `diff <(normalize baseline) <(normalize candidate)` per route.

**Assertions to automate** (shell/`jq`/`grep` or a small Go test using `net/http/httptest` against the handler):

- [ ] Each route returns `200` and `content-type: text/html; charset=utf-8`.
- [ ] `<title>` for each route equals the baseline title (Appendix A).
- [ ] `meta[name=description]` content matches baseline.
- [ ] `<link rel=canonical>` present and correct per route.
- [ ] Each product page contains exactly one `application/ld+json` Product block with matching `name`, `sku`/`codename`, `offers.price`, and a BreadcrumbList.
- [ ] `/robots.txt` byte-equal to baseline.
- [ ] `/sitemap.xml` lists the same 10 `<loc>`s (modulo the intentional `http`→`https` scheme change).
- [ ] Trailing slash (global policy): **every** non-slash route → `301` → slash form, including the legal routes (`/terms` → `301` → `/terms/`). This is the documented §9 change vs. the live baseline (which 200s the legal routes at both forms).
- [ ] `/api/stripe/products` and `/api/shippingrates` via the Go proxy return the **same JSON** as direct governor (proxy correctness).
- [ ] `/gif/<id>/giphy.gif` → `200 image/gif` through the Go proxy.
- [ ] Stock-state rendering: with a **mocked governor** returning `current_stock:0/ignore_stock:false` → page shows "Out of Stock" + disabled controls; `current_stock:5` → "In Stock" + "Last 5 items" + quantity options 1–5; `ignore_stock:true` → "Open for Pre-Order" + button "Pre Order" + quantity 1–20 (and **no** discount-code copy). (Use `httptest.Server` as fake governor in Go unit tests.)

### 12.2 Go unit/integration tests

- `governor` client: table tests for **both** shippingrates shapes + timeout/degradation path.
- `product_form` logic: table tests over stock permutations asserting rendered state (render templ to a buffer, assert substrings/classes).
- `sitemap`/`robots`: golden-file tests.
- Proxy: `httptest` fake upstream asserting path rewrite (`/api/x`→`/x`) and method/body preservation for POST.
- Handlers: `httptest` requests asserting status, headers, canonical, JSON-LD presence.

### 12.3 Manual QA checklist (staging, real browser)

- [ ] Visual parity on each route at mobile + desktop widths (nav, footer, carousel, fonts/Montserrat, colors).
- [ ] Image **carousel** works (prev/next) on all 4 product pages.
- [ ] Seasonal banner logic (temporarily force the date) shows/hides correctly.
- [ ] Country `<select>` populated and sorted; quantity select bounds correct for each stock state.
- [ ] **Checkout** end-to-end to Stripe (use a real low-value test or Stripe test mode via governor staging): POST reaches governor and redirects to Stripe Checkout; success return to `/success/?order_id=…&codename=…` resolves product + shows a gif.
- [ ] Downloadable firmware links on `/nrfnano/` (`.hex`, `.uf2`) serve with correct content types.
- [ ] 404 behavior for unknown routes is sane.
- [ ] Lighthouse/PageSpeed: SEO ≥ parity, performance ≥ current (expect improvement).
- [ ] `view-source` confirms SSR HTML (not an empty SPA shell) for product pages.

### 12.4 Exit criteria

All §12.1 assertions pass, §12.2 `go test ./...` green, §12.3 checklist complete, and the only intentional baseline diffs are the documented §9 changes.

---

## 13. Risks & mitigations

| Risk                                         | Impact                  | Mitigation                                                                                       |
| -------------------------------------------- | ----------------------- | ------------------------------------------------------------------------------------------------ |
| Governor down during SSR                     | Product pages fail      | Timeout + degrade to "unavailable/disabled purchase" instead of 500; short cache.                |
| CORS: moving stock fetch client-side         | Broken fetch            | Keep the server-side `/api` proxy (chosen design) — no CORS dependency.                          |
| Checkout POST regression                     | Lost sales              | Dedicated proxy test preserving method/body/headers; manual Stripe test-mode E2E before cutover. |
| `tw-elements` carousel/ripple parity         | Visual/functional drift | Replace with vanilla carousel component and verify in QA, or vendor `tw-elements`.               |
| Tailwind class purge misses `.templ`         | Broken styles           | Repoint `content` globs to `web/**/*.templ`; diff computed CSS; visual QA.                       |
| Sitemap `http`→`https` or `www` changes      | SEO wobble              | Document as intentional; submit updated sitemap; 301s preserve URLs.                             |
| Hidden copy differences in long product text | Content drift           | Copy marketing/quick-start text verbatim; HTML structural diff in §12.1.                         |
| Firmware/static asset path changes           | Broken downloads        | Keep same URL paths under `web/static`; assert in QA.                                            |
| Logo PNG loses `fill-*` hover recolor        | Minor visual drift      | Use light PNG + `opacity-80 hover:opacity-100` (§3.4); verify hover in QA; two-PNG swap if exact match needed. |
| Rootless podman can't bind low ports / SELinux volume labels | Container won't serve | Publish to high port `127.0.0.1:5000` (nginx fronts TLS); use `:Z` on any bind mounts; test `podman compose up` on the server before cutover. |

---

## 14. Effort estimate (rough)

- Phase 1–2 (scaffold + chrome + styling + legal/index): ~1–1.5 days.
- Phase 3 (governor client + product template + form logic): ~1.5–2 days.
- Phase 4–5 (SEO + proxies): ~1 day.
- Phase 6–7 (Containerfile/compose + CI/CD + regression + cutover): ~1.5–2 days.
- Total: ~5.5–7 focused days, plus a soak period running both stacks.

---

## Appendix A — Live titles / descriptions (baseline)

Capture exact strings with Phase 0 script; known examples:

- `/usbmux/` → `<title>Buy Badgerd USBMux | Quickstart With USBMux</title>`, description: "USBMux is a tool that helps you to switch connections between 3 USB ports. You can control the connection of the ports from a computer."
- `/` → `<title>Sett | Badgerd Technologies</title>`.
  (Fill in sdwire3/sdwirec/nrfnano/legal pages from capture.)

## Appendix B — Live `usbmux` JSON-LD (Product) shape

```
Product { url, name:"Badgerd USBMux | USBMux", image:"https://www.badgerd.nl/usbmux/usbmux-1.jpg",
  brand{logo,name:"Badgerd",slogan}, aggregateRating{ratingValue:5,reviewCount:1},
  review[{author,datePublished,reviewBody,reviewRating}],
  offers[{priceCurrency:"EUR",availability,price:180,sku:"badgerd_usbmux",priceValidUntil:"2024-12-31"}] }
+ BreadcrumbList[Store, SDWireC, USBMux, nRFNANO]
```

Note `priceValidUntil:2024-12-31` is stale in the current site — refresh during migration.

## Appendix C — Verified live facts (2026-10-28)

- `server: nginx/1.24.0 (Ubuntu)`; home `content-type: text/html; charset=utf-8`.
- `/api/stripe/products` live stock: nrfnano 0, usbmux 0, sdwirec 100, sdwire3 100.
- `/api/shippingrates` returns `{countries:[…]}` shape.
- `/gif/...` → `200 image/gif`, upstream sends `access-control-allow-origin: *`.
- `/usbmux` → `301 /usbmux/` (all 4 product routes redirect); legal routes (`/terms`,`/privacy`,`/refund`,`/legal`) return `200` at both slash and non-slash forms.
- `sitemap.xml` uses `http://` scheme + trailing slashes; `robots.txt` references `https://badgerd.nl/sitemap.xml`.
- governor `/stripe/products` sends **no** `access-control-allow-origin` (proxy required).

---

## 15. Implementation log (2026-10-29, branch `go-migration`)

### 15.1 What exists

| Area | Files | Notes |
| ---- | ----- | ----- |
| Entrypoint | `cmd/web/main.go` | env config (§10), graceful shutdown, `-healthcheck` flag (distroless has no wget), embedded tzdata |
| Server | `internal/server/{server,routes,handlers,proxy,sitemap}.go` | Go 1.22 `ServeMux` patterns; global trailing-slash 301 middleware; 404 page; `/healthz` |
| Governor client | `internal/governor/client.go` | 3 s timeout, 20 s cache, **serves stale data if governor fails**, both shipping-rate shapes |
| Catalog + form rules | `internal/catalog/{catalog,form}.go` | all static product data/SEO; `Product.Form(stock)` drives the purchase form |
| SEO | `internal/seo/seo.go` | meta/OG/canonical, JSON-LD structs in nuxt-jsonld field order |
| Views | `web/components/*.templ`, `web/pages/*.templ` | one shared product template + per-product copy components |
| Assets | `web/static/` (merged `assets/` + `static/`, identical duplicates), `web/web.go` | **embedded** into the binary (`go:embed`); `?v=<sha>` cache-busting for CSS/JS |
| Styles/JS | `web/css/input.css`, `web/tailwind.config.js`, `web/static/js/{carousel,success}.js` | tw-elements dropped (vanilla carousel, ~90 lines) |
| Logos | `tools/logogen` (separate Go module), `web/source/logo-nobg.svg` | `logo-96/128.png` (3.7/5.5 KB) for nav, `logo-512.png` (74 KB) for JSON-LD |
| Tests | `internal/*/*_test.go` | form-state table tests, governor shapes/cache/stale/timeout, handlers, stock states, degradation, redirects, proxies (POST/body/Host/XFF, cookie stripping, 502), robots/sitemap |
| Regression | `regression/capture.sh`, `regression/compare.py` | stdlib-only diff of title/meta/OG/canonical/JSON-LD/visible text + robots/sitemap/proxy/redirects |
| Build/deploy | `Justfile`, `Containerfile`, `compose.yaml`, `.containerignore`, `.github/workflows/go.yml` | image 38.5 MB; `go.yml` runs vet/test/build on push/PR and can deploy staging (`workflow_dispatch`) |

The Nuxt app is untouched and still deployable (root `tailwind.config.js`, `assets/`, `static/`, `cd.yml` unchanged).

### 15.2 Verification results

- `go vet ./...` and `go test ./...` pass (also run inside the container build).
- `just regress http://localhost:5001` against the 2026-10-28 production baseline: **0 unexpected differences** across all 10 routes (titles, descriptions, OG tags, canonicals, JSON-LD modulo §15.3 changes, visible text), `robots.txt` byte-identical, same 10 sitemap locs, `/api` JSON identical through the proxy, `/gif` → `200 image/gif`, every non-slash route → 301.
- Headless Chromium screenshots (1440×900 and 390×844) of `/`, `/usbmux/`, `/sdwirec/`, `/terms/` vs production: ≤0.08 % pixels differ, confined to the logo (raster vs vector) and the intentional button label.
- Carousel (next/prev/wrap/indicators/autoplay) and `/success/` (product link + gif via proxy) exercised in Chromium with no JS errors.
- Home page HTML: **1.6 MB → 17 KB**.

### 15.3 Intentional changes vs. production (flag in the cutover PR)

1. Sitemap uses `https://` (§9); the unused news/image/video/xhtml namespaces were dropped.
2. Legal routes now 301 to the slash form (global policy, §9).
3. No `PREORDER30` discount copy (§2.3).
4. Submit label follows availability (`Pre Order` only when `ignore_stock`); usbmux previously always said "Pre Order".
5. Bug fixes in quantity options: nrfnano/usbmux with stock **exactly 20** and sdwirec with stock **20–95** rendered an empty quantity list (unorderable); they now offer `MinQty..min(stock, MaxQty)`.
6. No "Last 0 items" badge next to "Out of Stock" (shown live on nrfnano).
7. JSON-LD: `availability` follows live stock (InStock/OutOfStock/PreOrder); `priceValidUntil` = Dec 31 of the current year (was stale/invalid, e.g. `2026-06-31`); `brand.logo` → `/logo-512.png`; BreadcrumbList positions renumbered sequentially (sdwire3 had two `position: 2`); **sdwire3 offer price 85 → 100** to match the visible €100.00 (confirm against Stripe).
8. Header logo is a PNG with `opacity-90 hover:opacity-100` (90 % of zinc-50 on zinc-700 ≈ the old zinc-200 fill; §3.4 suggested 80).
9. tw-elements ripple effect on the submit button removed; selects got real `aria-label`s ("Quantity"/"Country" instead of `.form-select-sm`).
10. Governor outages render the page with purchasing disabled (and stale cached stock is preferred when available) instead of an error page; unknown paths get a styled 404.
11. `/success/` renders the order number server-side (escaped); product lookup and gif stay client-side.
12. `og:image` for nrfnano, sdwire3 and usbmux pointed at files that 404 (`nrfnano-0.jpg`, `badgerd-sdwire3.png`, `badgerd-usbmux-og.png`); they now use `nrfnano_0.jpg`, `sdwire3-1.jpg`, `usbmux-1.jpg` (fixed in both `internal/catalog` and the Nuxt pages). `TestProductImagesExist` guards every catalog image reference.
13. The "(incl. VAT)" note next to product prices was removed (product owner request, 2026-10-29) in both stacks; shipping info stays ("+ Shipping", SDWire3 "(Free Shipping)").

### 15.4 Findings / follow-ups

- **Branch is one commit behind `origin/main`** (`672ba64 sdwire3: update instructions for Rasbian`, which production already serves). Its content is ported into `web/pages/sdwire3.templ`; rebase `go-migration` onto `origin/main` before merging (the only overlap is `pages/sdwire3.vue`, which this branch doesn't modify).
- Carousel CSS: tw-elements only translates the *outgoing* slide with `.active.carousel-item-start/end`; dropping the `.active` qualifier makes the incoming slide fly past and snap back. Fixed in `web/css/input.css`; frame-sampled in Chromium to match production's trajectory.
- The landing page's `og:image` is empty (as in production); consider a dedicated share image.
- Tailwind is pinned to the **v3.4.17** standalone CLI: v4 (current "latest", which the §8.1 Containerfile snippet would have downloaded) is CSS-first and changes utility semantics.
- templ pitfalls hit during the port: element text collapses whitespace (shell snippets use `@components.CodeBlock` with Go raw strings); a standalone `@` or a text line starting with `for`/`if`/`switch` must be wrapped in `{ "..." }`; prettier-style split closing tags (`</a\n>`) don't parse.
- `podman build` ignores `HEALTHCHECK` for OCI images; the health check lives in `compose.yaml`.
- `/success/` is still listed in the sitemap for parity; consider removing it (thank-you pages shouldn't be indexed).
- Remaining before cutover (§11–12): deploy to `127.0.0.1:5001` on the server (`go.yml` → *Run workflow* → deploy_staging, or `just up`), add an nginx preview host, re-run `just regress https://<preview-host>`, do the manual checklist in §12.3 (notably a real Stripe checkout round-trip), then switch `cd.yml` to `podman compose up -d --build` with `HOST_PORT=5000`.

