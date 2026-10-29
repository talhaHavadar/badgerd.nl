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

# bring the app up via podman compose (127.0.0.1:${HOST_PORT:-5000})
up:
    podman compose up -d --build

down:
    podman compose down
