# badgerd.nl Go + templ site (see migration_guide.md §8.1).
# Build: podman build -t badgerd-web -f Containerfile .   (or: just image)

# ---- build stage ----
FROM docker.io/library/golang:1.26 AS build
ARG TARGETARCH=amd64
ARG TAILWIND_VERSION=v3.4.17
ARG TAILWIND_SHA256_AMD64=7d24f7fa191d2193b78cd5f5a42a6093e14409521908529f42d80b11fde1f1d4
ARG TAILWIND_SHA256_ARM64=69b1378b8133192d7d2feb12a116fa12d035594f58db3eff215879e4ad8cf39b
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
RUN go tool templ generate -path web \
 && tailwindcss -c web/tailwind.config.js -i web/css/input.css -o web/static/css/app.css --minify \
 && go test ./... \
 && CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/badgerd-web ./cmd/web

# ---- runtime stage ----
# Static assets are embedded in the binary, so the image is a single file.
# Health checks are defined in compose.yaml (`badgerd-web -healthcheck`).
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/badgerd-web /app/badgerd-web
ENV PORT=5000
EXPOSE 5000
USER nonroot:nonroot
ENTRYPOINT ["/app/badgerd-web"]
