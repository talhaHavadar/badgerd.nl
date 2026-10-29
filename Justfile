# badgerd.nl — Go + templ site (see migration_guide.md).
# Requires: go (1.26+), just, curl. `just tools` downloads the pinned Tailwind CLI.

# templ is pinned in go.mod (`tool` directive) and run with `go tool templ`.
tailwind_version := "v3.4.17"
tailwind_sha256  := "7d24f7fa191d2193b78cd5f5a42a6093e14409521908529f42d80b11fde1f1d4" # tailwindcss-linux-x64
tailwind         := "./bin/tailwindcss"
templ            := "go tool templ"
bin              := "bin/badgerd-web"

default: build

# download the pinned Tailwind standalone CLI (no Node needed)
tools:
    mkdir -p bin
    test -x {{tailwind}} || curl -fsSL -o {{tailwind}} https://github.com/tailwindlabs/tailwindcss/releases/download/{{tailwind_version}}/tailwindcss-linux-x64
    echo "{{tailwind_sha256}}  {{tailwind}}" | sha256sum -c --quiet
    chmod +x {{tailwind}}

# regenerate templ Go code from *.templ
generate:
    {{templ}} generate -path web

# format *.templ files
fmt:
    {{templ}} fmt web
    gofmt -w cmd internal web

# compile + minify Tailwind CSS
css: tools
    {{tailwind}} -c web/tailwind.config.js -i web/css/input.css -o web/static/css/app.css --minify

# full build: templ -> css -> static Go binary (assets embedded)
build: generate css
    CGO_ENABLED=0 go build -trimpath -o {{bin}} ./cmd/web

# build then run locally on :5000 (override with PORT=...)
run: build
    PORT=${PORT:-5000} {{bin}}

test: generate
    go test ./...

vet: generate
    go vet ./...

# regenerate logo PNGs from the master artwork (tools/logogen)
logos:
    cd tools/logogen && go run . -svg ../../web/source/logo-nobg.svg -png ../../web/static/logo-nobg.png -out ../../web/static

# build the OCI image with podman (see Containerfile)
image tag="badgerd-web:latest":
    podman build -t {{tag}} -f Containerfile .

# print the systemd services generated from deploy/quadlet (validates the units)
quadlet-check:
    QUADLET_UNIT_DIRS={{justfile_directory()}}/deploy/quadlet /usr/libexec/podman/quadlet -dryrun -user

# on the server (as the deploy user): install deploy/ (Quadlet units, env files, auto-update timer) and start or
# rolling-restart the replicas. Idempotent; the repo is the source of truth. `just deploy --dry-run` only prints.
deploy *args:
    #!/usr/bin/env bash
    set -euo pipefail

    DRY_RUN=false
    case "{{args}}" in
      "") ;;
      --dry-run) DRY_RUN=true ;;
      *) echo "usage: just deploy [--dry-run]" >&2; exit 2 ;;
    esac

    cd "{{justfile_directory()}}"

    IMAGE=ghcr.io/talhahavadar/badgerd.nl
    CONFIG="${XDG_CONFIG_HOME:-$HOME/.config}"
    QUADLET_DIR="$CONFIG/containers/systemd"
    ENV_DIR="$CONFIG/badgerd"
    TIMER_DIR="$CONFIG/systemd/user/podman-auto-update.timer.d"
    # Start order: B replicas first, so on a host still running the old podman-compose container on :5000, prod-b comes
    # up on :5002 before anything touches :5000.
    UNITS=(badgerd-staging-b badgerd-staging-a badgerd-prod-b badgerd-prod-a)

    log() { printf '==> %s\n' "$*"; }
    warn() { printf 'WARNING: %s\n' "$*" >&2; }
    die() { printf 'ERROR: %s\n' "$*" >&2; exit 1; }
    run() {
      if $DRY_RUN; then printf '    [dry-run] %s\n' "$*"; else "$@"; fi
    }

    env_of() { local u=${1#badgerd-}; echo "${u%-*}"; }
    port_of() { sed -n 's/^PublishPort=127\.0\.0\.1:\([0-9]*\):5000$/\1/p' "deploy/quadlet/$1.container"; }

    # ---- preflight ----
    [ "$(id -u)" -ne 0 ] || die "run as the deploy user, not root (rootless podman + systemd --user)"
    [ -d deploy/quadlet ] || die "deploy/quadlet not found; run from a badgerd.nl checkout"
    for cmd in podman systemctl curl ss; do
      command -v "$cmd" >/dev/null || die "$cmd is not installed"
    done
    systemctl --user show-environment >/dev/null 2>&1 \
      || die "no systemd user session; log in via SSH as $(id -un) (or set XDG_RUNTIME_DIR=/run/user/$(id -u))"
    [ "$(loginctl show-user "$(id -un)" -p Linger --value 2>/dev/null)" = yes ] \
      || die "lingering is disabled, so the containers won't start at boot or survive logout: sudo loginctl enable-linger $(id -un)"

    # ---- install files ----
    CHANGED=()   # environments (prod/staging) whose unit or env file changed
    install_file() { # src dest
      if [ -f "$2" ] && cmp -s "$1" "$2"; then return 1; fi
      log "install $1 -> $2"
      run install -D -m 0644 "$1" "$2"
    }
    mark_changed() { [[ " ${CHANGED[*]} " == *" $1 "* ]] || CHANGED+=("$1"); }

    for f in deploy/quadlet/*.container; do
      u=$(basename "$f" .container)
      if install_file "$f" "$QUADLET_DIR/$u.container"; then mark_changed "$(env_of "$u")"; fi
    done
    for f in deploy/env/*.env; do
      if install_file "$f" "$ENV_DIR/$(basename "$f")"; then mark_changed "$(basename "$f" .env)"; fi
    done
    timer_changed=false
    if install_file deploy/systemd/podman-auto-update.timer.d/override.conf "$TIMER_DIR/override.conf"; then
      timer_changed=true
    fi

    log "systemctl --user daemon-reload"
    run systemctl --user daemon-reload
    if ! $DRY_RUN; then
      for u in "${UNITS[@]}"; do
        systemctl --user cat "$u.service" >/dev/null 2>&1 \
          || die "quadlet did not generate $u.service; check: /usr/libexec/podman/quadlet -dryrun -user"
      done
    fi

    # ---- pull images (fails early with a clear message instead of a unit start timeout) ----
    declare -A HAVE_IMAGE
    for env in staging prod; do
      log "pull $IMAGE:$env"
      if $DRY_RUN || podman pull -q "$IMAGE:$env" >/dev/null; then
        HAVE_IMAGE[$env]=true
      else
        HAVE_IMAGE[$env]=false
        hint="merge to main (publish.yml)"
        [ "$env" = prod ] && hint="run the 'Promote to production' workflow"
        warn "$IMAGE:$env is not available yet; $hint, then re-run. Skipping $env replicas."
      fi
    done

    # ---- start / restart replicas, one at a time ----
    # Each start/restart blocks until the unit's ExecStartPost sees /healthz, so the other replica keeps serving.
    BLOCKED=()
    for u in "${UNITS[@]}"; do
      env=$(env_of "$u")
      port=$(port_of "$u")
      ${HAVE_IMAGE[$env]} || continue
      if systemctl --user is-active --quiet "$u.service"; then
        if [[ " ${CHANGED[*]} " == *" $env "* ]]; then
          log "restart $u (config changed)"
          run systemctl --user restart "$u.service"
        else
          log "$u is running and up to date"
        fi
      elif [ -n "$(ss -Hltn "sport = :$port")" ]; then
        warn "127.0.0.1:$port is held by something else (the old podman-compose badgerd-web?); not starting $u"
        BLOCKED+=("$u:$port")
      else
        log "start $u (127.0.0.1:$port)"
        run systemctl --user start "$u.service"
      fi
    done

    # ---- auto-update timer ----
    if ! systemctl --user is-enabled --quiet podman-auto-update.timer || ! systemctl --user is-active --quiet podman-auto-update.timer; then
      log "enable podman-auto-update.timer"
      run systemctl --user enable --now podman-auto-update.timer
    elif $timer_changed; then
      log "restart podman-auto-update.timer (schedule changed)"
      run systemctl --user restart podman-auto-update.timer
    fi

    # ---- report ----
    if ! $DRY_RUN; then
      echo
      for u in "${UNITS[@]}"; do
        port=$(port_of "$u")
        state=$(systemctl --user is-active "$u.service" || true)
        s=-
        if [ "$state" = active ]; then
          if curl -fsS -o /dev/null --max-time 2 "http://127.0.0.1:$port/healthz" 2>/dev/null; then s=healthy; else s=UNHEALTHY; fi
        fi
        printf '    %-18s 127.0.0.1:%-5s %-9s %s\n' "$u" "$port" "$state" "$s"
      done
      systemctl --user list-timers --no-pager podman-auto-update.timer | sed -n '1,2p'
    fi

    for site in badgerd.nl staging.badgerd.nl; do
      f=/etc/nginx/sites-enabled/$site
      [ -r "$f" ] || continue
      up=badgerd_prod; [ "$site" = staging.badgerd.nl ] && up=badgerd_staging
      grep -q "proxy_pass http://$up;" "$f" \
        || warn "$f doesn't proxy to the '$up' upstream yet (only one replica gets traffic); see README 'nginx upstreams'"
    done

    if [ ${#BLOCKED[@]} -gt 0 ]; then
      cat >&2 <<EOF

    Cutover from podman-compose still pending (${BLOCKED[*]}). With the nginx upstream for both
    replicas in place and reloaded (sudo nginx -t && sudo systemctl reload nginx):

        pod="\$(podman inspect badgerd-web --format '{{{{.PodName}}')"
        podman rm -f badgerd-web
        [ -n "\$pod" ] && podman pod rm -f "\$pod"
        just deploy
    EOF
    fi
