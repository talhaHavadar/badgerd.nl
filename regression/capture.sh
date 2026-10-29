#!/usr/bin/env bash
# Capture a behavioral snapshot of a badgerd.nl deployment for migration regression
# testing (see ../migration_guide.md §12). Run once against production to create a
# baseline, then against the Go/templ staging server, and diff the normalized output.
#
# Usage:
#   ./regression/capture.sh https://badgerd.nl        regression/baseline
#   ./regression/capture.sh http://localhost:5001     regression/candidate
set -euo pipefail

BASE="${1:?usage: capture.sh <base-url> <out-dir>}"
OUT="${2:?usage: capture.sh <base-url> <out-dir>}"
mkdir -p "$OUT"

ROUTES=( / /sdwire3/ /sdwirec/ /nrfnano/ /usbmux/ /success/ /terms/ /privacy/ /refund/ /legal/ )

echo "Capturing ${#ROUTES[@]} HTML routes from $BASE ..."
for r in "${ROUTES[@]}"; do
  name="$(echo "$r" | sed 's#/#_#g; s#^_##; s#_$##')"
  [ -z "$name" ] && name="index"
  curl -fsS "$BASE$r" -o "$OUT/$name.html"
  curl -fsS -D "$OUT/$name.headers" -o /dev/null "$BASE$r"
done

echo "Capturing SEO + API surface ..."
curl -fsS "$BASE/robots.txt"            -o "$OUT/robots.txt"
curl -fsS "$BASE/sitemap.xml"           -o "$OUT/sitemap.xml"
curl -fsS "$BASE/api/stripe/products"   -o "$OUT/products.json"
curl -fsS "$BASE/api/shippingrates"     -o "$OUT/shippingrates.json"

echo "Capturing trailing-slash redirects ..."
: > "$OUT/redirects.txt"
for r in /usbmux /sdwirec /nrfnano /sdwire3 /terms /privacy /refund /legal; do
  code_url="$(curl -sS -o /dev/null -w '%{http_code} %{redirect_url}' "$BASE$r" || true)"
  echo "$r -> $code_url" >> "$OUT/redirects.txt"
done

echo "Capturing gif proxy status ..."
curl -sS -o /dev/null -w '%{http_code} %{content_type}\n' \
  "$BASE/gif/media/l0amJzVHIAfl7jMDos/giphy.gif" > "$OUT/gif.txt" || true

echo "Done. Snapshot written to $OUT"
echo "Tip: normalize (strip stock counts, dates, etags, content-length) before diffing baseline vs candidate."
