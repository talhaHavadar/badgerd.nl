# badgerd.nl

Product site for Badgerd Technologies B.V.: a single Go binary rendering [templ](https://templ.guide) views, with live
stock from `governor.badgerd.nl` and same-origin `/api` + `/gif` reverse proxies. It replaced the former Nuxt 2 app;
see [migration_guide.md](migration_guide.md) for the design and the migration log.

Requirements: Go 1.26+, [just](https://github.com/casey/just), curl. templ is pinned in `go.mod` (`go tool templ`);
the Tailwind v3 standalone CLI is downloaded (checksum-verified) into `bin/` by `just tools`.

```bash
just build         # templ generate -> Tailwind CSS -> bin/badgerd-web (static assets embedded)
just run           # build and serve on :5000 (PORT=5001 just run for another port)
just test          # go test ./...
just image         # podman build -t badgerd-web:latest
just quadlet-check # validate deploy/quadlet units
just deploy        # deploy quadlet units
just logos         # regenerate logo PNGs from web/source/logo-nobg.svg
```

Runtime configuration (env):

```
# PORT defaults to 5000
# PORT=5000
# serve assets from disk while developing
# STATIC_DIR=web/static
BASE_URL=https://badgerd.nl
GOVERNOR_BASE_URL=https://governor.badgerd.nl
GIPHY_BASE_URL=https://media.giphy.com
TZ=Europe/Amsterdam
```

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

```
merge to main ──publish.yml──▶ ghcr.io/talhahavadar/badgerd.nl:staging ──auto-update──▶ staging.badgerd.nl (:5001 + :5003)
run promote.yml ─────────────▶ ghcr.io/talhahavadar/badgerd.nl:prod    ──auto-update──▶ badgerd.nl         (:5000 + :5002)
```

- `.github/workflows/ci.yml` runs vet/test/build on pushes and pull requests.
- `.github/workflows/publish.yml` builds the image (tests run inside the build), smoke-tests it and pushes `latest`,
  `staging`, `YYYYMMDD` and `sha-<short>` from `main` (`<tag>` from `v*` git tags; pull requests only build). It uses
  the job's short-lived `GITHUB_TOKEN`, passed to `podman login --password-stdin`.
- `.github/workflows/promote.yml` (manual, `production` environment; add required reviewers there to gate it) re-tags
  the current `staging` digest as `prod` plus an immutable `prod-YYYYMMDD-HHMMSS`, without rebuilding.
- On the server each environment runs two replicas as rootless podman [Quadlet](https://docs.podman.io/en/v4.9.3/markdown/podman-systemd.unit.5.html)
  units (`deploy/quadlet/`), configured by `deploy/env/<env>.env`. nginx balances over both replicas.
  `podman-auto-update.timer` (every minute, `deploy/systemd/`) pulls new `staging`/`prod` digests and restarts the
  replicas one at a time. Each unit counts as started only once `/healthz` answers (`ExecStartPost`), so nginx keeps
  serving from the other replica. If the new image never gets healthy, `podman auto-update` rolls that replica
  back to the previous image.

**Rollback:** run _Promote to production_ with `source` set to an earlier `prod-YYYYMMDD-HHMMSS` (listed in earlier
run summaries) or `sha-<short>` tag. The server picks it up within a minute.

Useful on the server: `systemctl --user status 'badgerd-*'`, `journalctl --user -u badgerd-prod-a`,
`podman auto-update --dry-run`, `systemctl --user list-timers podman-auto-update.timer`.

### Server setup: `just deploy`

Run `just deploy` as the deploy user from the root of a checkout on the server (`git pull` first; `just deploy
--dry-run` only prints). It is idempotent and treats the repo as the source of truth:

- installs `deploy/quadlet/*.container` → `~/.config/containers/systemd/`, `deploy/env/*.env` → `~/.config/badgerd/`
  and the timer override → `~/.config/systemd/user/podman-auto-update.timer.d/`, then `daemon-reload`;
- pulls `:staging` and `:prod` (skipping an environment whose tag doesn't exist yet);
- starts replicas that aren't running, and restarts, one at a time, the replicas of an environment whose unit or env
  file changed (each waits for `/healthz`, so the other replica keeps serving);
- enables `podman-auto-update.timer` and prints the replica status;
- warns when an nginx site doesn't use the upstream yet, and skips a replica whose port is still held by
  something else (e.g. the old podman-compose container).

Quadlet units can't be `enable`d; `WantedBy=default.target` starts them at boot, which needs lingering
(`sudo loginctl enable-linger <user>`; `just deploy` checks it). To change an environment's settings (e.g. a staging
`GOVERNOR_BASE_URL`), edit `deploy/env/<env>.env`, commit, and `just deploy` on the server.

### nginx upstreams config

nginx upstreams (`upstream` blocks go at the top of the site file, outside `server {}`). `proxy_next_upstream` retries
a request on the other replica when one refuses connections during a restart:

```nginx
upstream badgerd_prod {
    server 127.0.0.1:5000 max_fails=1 fail_timeout=5s;
    server 127.0.0.1:5002 max_fails=1 fail_timeout=5s;
}

# in the badgerd.nl server { } block (replaces the old proxy_pass to localhost:5000 and its Upgrade/Connection headers)
location / {
    proxy_pass http://badgerd_prod;
    proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-Host $host;
    proxy_connect_timeout 2s;
    proxy_next_upstream error timeout;
    proxy_next_upstream_tries 2;
}
```

Staging is the same with `upstream badgerd_staging { server 127.0.0.1:5001 ...; server 127.0.0.1:5003 ...; }` and
`proxy_pass http://badgerd_staging;`. Keep its `X-Robots-Tag "noindex, nofollow"` header.
