# badgerd.nl

> **Migration in progress:** the site is being ported from Nuxt 2 to a single Go + [templ](https://templ.guide)
> binary (see [migration_guide.md](migration_guide.md)). Production still deploys the Nuxt app via
> `.github/workflows/cd.yml` until cutover.

## Go + templ site

Requirements: Go 1.26+, [just](https://github.com/casey/just), curl. templ is pinned in `go.mod` (`go tool templ`);
the Tailwind v3 standalone CLI is downloaded (checksum-verified) into `bin/` by `just tools`.

```bash
just build        # templ generate -> Tailwind CSS -> bin/badgerd-web (static assets embedded)
just run          # build and serve on :5000 (PORT=5001 just run for another port)
just test         # go test ./...
just regress http://localhost:5001   # capture a snapshot and diff it against regression/baseline
just image        # podman build; `just up` / `just down` for podman compose (127.0.0.1:${HOST_PORT:-5001})
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

Shell snippets in templ must use `@components.CodeBlock(`...`)` (templ collapses whitespace in element text), and
text starting with `for`/`if`/`switch` or containing a standalone `@` must be wrapped in `{ "..." }`.

## Nuxt site (current production)

### Build Setup

```bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev

# build for production and launch server
$ yarn build
$ yarn start

# generate static project
$ yarn generate
```

For detailed explanation on how things work, check out the [documentation](https://nuxtjs.org).

## Special Directories

You can create the following extra directories, some of which have special behaviors. Only `pages` is required; you can delete them if you don't want to use their functionality.

### `assets`

The assets directory contains your uncompiled assets such as Stylus or Sass files, images, or fonts.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/assets).

### `components`

The components directory contains your Vue.js components. Components make up the different parts of your page and can be reused and imported into your pages, layouts and even other components.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/components).

### `layouts`

Layouts are a great help when you want to change the look and feel of your Nuxt app, whether you want to include a sidebar or have distinct layouts for mobile and desktop.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/layouts).

### `pages`

This directory contains your application views and routes. Nuxt will read all the `*.vue` files inside this directory and setup Vue Router automatically.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/get-started/routing).

### `plugins`

The plugins directory contains JavaScript plugins that you want to run before instantiating the root Vue.js Application. This is the place to add Vue plugins and to inject functions or constants. Every time you need to use `Vue.use()`, you should create a file in `plugins/` and add its path to plugins in `nuxt.config.js`.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/plugins).

### `static`

This directory contains your static files. Each file inside this directory is mapped to `/`.

Example: `/static/robots.txt` is mapped as `/robots.txt`.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/static).

### `store`

This directory contains your Vuex store files. Creating a file in this directory automatically activates Vuex.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/store).
