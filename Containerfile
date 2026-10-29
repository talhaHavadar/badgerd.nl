# badgerd.nl Go + templ site (see migration_guide.md §8.1).
# Build: podman build -t badgerd-web -f Containerfile .   (or: just image)

# ---- build stage ----
FROM docker.io/library/golang:1.26 AS build
ARG TARGETARCH=amd64
ARG TAILWIND_VERSION=v4.3.3
ARG TAILWIND_SHA256_AMD64=dc61b3ac6b8c9ca874c0cc4c57b2409791a64c5540404ca5f5367360babc313a
ARG TAILWIND_SHA256_ARM64=55fd0b241214eff3de1e8ee4f22796662f2d2e7a49bcfca7477cfd0bac398195
WORKDIR /src

RUN set -eu; \
    case "$TARGETARCH" in \
      amd64) arch=x64; sum="$TAILWIND_SHA256_AMD64" ;; \
      arm64) arch=arm64; sum="$TAILWIND_SHA256_ARM64" ;; \
      *) echo "unsupported arch $TARGETARCH" >&2; exit 1 ;; \
    esac; \
    curl -fsSL -o /usr/local/bin/tailwindcss \
      "https://github.com/tailwindlabs/tailwindcss/releases/download/${TAILWIND_VERSION}/tailwindcss-linux-${arch}"; \
    echo "${sum}  /usr/local/bin/tailwindcss" | sha256sum -c -; \
    chmod +x /usr/local/bin/tailwindcss

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY web ./web
# Reported by /healthz; publish.yml sets them. Declared late so they don't bust the cached layers above.
ARG VERSION=
ARG COMMIT=
ARG BUILD_TIME=
RUN go tool templ generate -path web \
 && tailwindcss -i web/css/input.css -o web/static/css/app.css --minify \
 && go test ./... \
 && CGO_ENABLED=0 GOOS=linux go build -trimpath \
      -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}" \
      -o /out/badgerd-web ./cmd/web

# ---- runtime stage ----
# Static assets are embedded in the binary, so the image is a single file.
# Health checks are defined in deploy/quadlet/*.container (`badgerd-web -healthcheck`).
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/badgerd-web /app/badgerd-web
ENV PORT=5000
EXPOSE 5000
USER nonroot:nonroot
ENTRYPOINT ["/app/badgerd-web"]
