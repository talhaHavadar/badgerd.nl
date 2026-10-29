#!/usr/bin/env python3
"""Compare two snapshots produced by regression/capture.sh (migration_guide.md §12.1).

Usage: python3 regression/compare.py regression/baseline regression/candidate

For every HTML route it compares <title>, meta description, OpenGraph tags,
canonical link, JSON-LD and the normalised visible text. It also checks
robots.txt, sitemap.xml locs, trailing-slash redirects, the gif proxy and that
the /api proxy returns the same JSON.

Volatile values (stock counts, order ids) and the intentional changes listed in
migration_guide.md §9 / §15 are normalised away; anything left is reported and
makes the script exit 1.
"""
import difflib
import json
import os
import re
import sys
from html.parser import HTMLParser

ROUTES = ['index', 'sdwire3', 'sdwirec', 'nrfnano', 'usbmux', 'success', 'terms', 'privacy', 'refund', 'legal']
VOID = {'area', 'base', 'br', 'col', 'embed', 'hr', 'img', 'input', 'link', 'meta', 'source', 'track', 'wbr'}
SKIP = {'script', 'style', 'svg', 'noscript', 'template', 'title'}
# Whitespace between block-level elements is insignificant; inline joins are not.
BLOCK = {'address', 'article', 'aside', 'blockquote', 'body', 'button', 'code', 'dd', 'div', 'dl', 'dt', 'footer',
         'form', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'header', 'hr', 'html', 'img', 'input', 'li', 'main', 'nav', 'ol',
         'option', 'p', 'section', 'select', 'table', 'td', 'th', 'tr', 'ul'}

# Visible-text normalisation for volatile or intentionally changed content.
TEXT_RULES = [
    # Intentional: no "Last 0 items" badge on out-of-stock products.
    (r'Last 0 items', ''),
    (r'Last \d+ items', 'Last N items'),
    (r'50\+ items', 'Last N items'),
    (r'\d+ items reserved, might be yours, try again in 30 minutes', ''),
    # Quantity <option> lists depend on live stock.
    (r'(?:\b\d{1,3}\b\s+){2,}(?=Your Country)', 'QTY '),
    # Intentional: removed blanket pre-order discount copy (guide §2.3).
    (r'Use PREORDER30 at checkout to get 30% discount\.\s*', ''),
    # Intentional: submit label follows availability ("Pre Order" only for pre-orders).
    (r'\b(?:Buy|Pre Order)\s+(?=Looking for|\d|$)', 'SUBMIT '),
    (r'\bPre Order\b', 'SUBMIT'),
    (r'(Your Country.*?)\bBuy\b', r'\1SUBMIT'),
    # Intentional: "incl. VAT" note removed from product price lines.
    (r'\(incl\. VAT and ', '('),
    (r'\(incl\. VAT\*?\)\s*', ''),
]


# Intentional meta changes: og:image URLs that 404'd in production were repointed.
META_CHANGES = {
    'https://www.badgerd.nl/nrfnano/nrfnano-0.jpg': 'https://www.badgerd.nl/nrfnano/nrfnano_0.jpg',
    'https://www.badgerd.nl/sdwire3/badgerd-sdwire3.png': 'https://www.badgerd.nl/sdwire3/sdwire3-1.jpg',
    'https://www.badgerd.nl/usbmux/badgerd-usbmux-og.png': 'https://www.badgerd.nl/usbmux/usbmux-1.jpg',
}


class Extract(HTMLParser):
    def __init__(self):
        super().__init__(convert_charrefs=True)
        self.title = None
        self.meta = {}
        self.canonical = None
        self.jsonld = []
        self.text = []
        self._stack = []  # (tag, hidden, is_flex_container)
        self._in_title = False
        self._in_ld = False
        self._ld_buf = ''

    def _hidden(self):
        return any(h for _, h, _ in self._stack)

    def _in_flex(self):
        return bool(self._stack) and self._stack[-1][2]

    def handle_starttag(self, tag, attrs):
        a = dict(attrs)
        if tag == 'meta':
            key = a.get('name') or a.get('property')
            if key and key != 'viewport':
                self.meta[key] = a.get('content', '')
        elif tag == 'link' and a.get('rel') == 'canonical':
            self.canonical = a.get('href')
        elif tag == 'title':
            self._in_title = True
            self.title = ''
        elif tag == 'script' and a.get('type') == 'application/ld+json':
            self._in_ld = True
            self._ld_buf = ''
        # Flex items are blockified, so whitespace around them doesn't render either.
        if tag in BLOCK or tag == 'br' or self._in_flex():
            self.text.append(' ')
        if tag in VOID:
            return
        hidden = tag in SKIP or 'display:none' in (a.get('style') or '').replace(' ', '')
        flex = bool({'flex', 'inline-flex'} & set((a.get('class') or '').split()))
        self._stack.append((tag, hidden, flex))

    def handle_startendtag(self, tag, attrs):
        self.handle_starttag(tag, attrs)
        if tag not in VOID:
            self.handle_endtag(tag)

    def handle_endtag(self, tag):
        if tag in BLOCK or (len(self._stack) > 1 and self._stack[-2][2] and self._stack[-1][0] == tag):
            self.text.append(' ')
        if tag == 'title':
            self._in_title = False
        if tag == 'script' and self._in_ld:
            self._in_ld = False
            try:
                data = json.loads(self._ld_buf)
                self.jsonld.extend(data if isinstance(data, list) else [data])
            except json.JSONDecodeError:
                self.jsonld.append({'_invalid': self._ld_buf})
        for i in range(len(self._stack) - 1, -1, -1):
            if self._stack[i][0] == tag:
                del self._stack[i:]
                break

    def handle_data(self, data):
        if self._in_title:
            self.title += data
        if self._in_ld:
            self._ld_buf += data
        if not self._hidden():
            self.text.append(data)


def parse(path):
    p = Extract()
    with open(path, encoding='utf-8') as f:
        p.feed(f.read())
    text = re.sub(r'\s+', ' ', ''.join(p.text)).strip()
    for pat, rep in TEXT_RULES:
        text = re.sub(pat, rep, text)
    return p, re.sub(r'\s+', ' ', text).strip()


def norm_jsonld(items):
    """Drop fields that intentionally change (guide §9): availability now follows
    stock, priceValidUntil is refreshed, brand logo is the 512px PNG, breadcrumb
    positions are renumbered sequentially."""
    out = json.loads(json.dumps(items))
    for it in out:
        if it.get('@type') == 'Product':
            it.get('brand', {}).pop('logo', None)
            for o in it.get('offers', []):
                o.pop('availability', None)
                o.pop('priceValidUntil', None)
                o.pop('price', None)
        if it.get('@type') == 'BreadcrumbList':
            for li in it.get('itemListElement', []):
                li.pop('position', None)
    return out


def words(text):
    # One word per line keeps the diff readable for long single-line text.
    return text.split(' ')


def main(base, cand):
    problems = 0

    def report(route, what, detail):
        nonlocal problems
        problems += 1
        print(f'✗ {route}: {what}')
        if detail:
            print('    ' + detail.replace('\n', '\n    '))

    for r in ROUTES:
        bp, bt = parse(os.path.join(base, r + '.html'))
        cp, ct = parse(os.path.join(cand, r + '.html'))
        if bp.title != cp.title:
            report(r, 'title', f'{bp.title!r} != {cp.title!r}')
        bp.meta = {k: META_CHANGES.get(v, v) for k, v in bp.meta.items()}
        if bp.meta != cp.meta:
            keys = sorted(set(bp.meta) | set(cp.meta))
            diff = [f'{k}: {bp.meta.get(k)!r} -> {cp.meta.get(k)!r}' for k in keys if bp.meta.get(k) != cp.meta.get(k)]
            report(r, 'meta', '\n'.join(diff))
        if bp.canonical != cp.canonical:
            report(r, 'canonical', f'{bp.canonical!r} != {cp.canonical!r}')
        bj, cj = norm_jsonld(bp.jsonld), norm_jsonld(cp.jsonld)
        if bj != cj:
            d = difflib.unified_diff(json.dumps(bj, indent=1, sort_keys=True).splitlines(),
                                     json.dumps(cj, indent=1, sort_keys=True).splitlines(), lineterm='', n=1)
            report(r, 'json-ld', '\n'.join(list(d)[2:]))
        if bt != ct:
            d = difflib.unified_diff(words(bt), words(ct), lineterm='', n=4)
            report(r, 'visible text', ' '.join(x if x.startswith(('+', '-', '@')) else x for x in list(d)[2:]))
        if bt == ct:
            print(f'✓ {r}: visible text identical ({len(ct)} chars)')

    def read(name):
        with open(os.path.join(base, name), encoding='utf-8') as f1, open(os.path.join(cand, name), encoding='utf-8') as f2:
            return f1.read(), f2.read()

    b, c = read('robots.txt')
    if b != c:
        report('robots.txt', 'differs', f'{b!r} != {c!r}')
    else:
        print('✓ robots.txt byte-identical')

    b, c = read('sitemap.xml')
    locs = lambda s: sorted(re.sub(r'^https?://', '', l) for l in re.findall(r'<loc>([^<]+)</loc>', s))
    if locs(b) != locs(c):
        report('sitemap.xml', 'locs differ', f'{locs(b)} != {locs(c)}')
    else:
        print(f'✓ sitemap.xml: same {len(locs(c))} locs (scheme normalised)')

    for name in ('products.json', 'shippingrates.json'):
        b, c = read(name)
        if json.loads(b) != json.loads(c):
            report(name, 'proxy JSON differs (stock may have changed between captures)', '')
        else:
            print(f'✓ {name} identical through proxy')

    b, c = read('gif.txt')
    if b.strip() != c.strip():
        report('gif', 'proxy status differs', f'{b.strip()} != {c.strip()}')
    else:
        print(f'✓ gif proxy: {c.strip()}')

    # Global trailing-slash policy (guide §9): everything 301s, including legal routes.
    _, c = read('redirects.txt')
    for line in c.strip().splitlines():
        m = re.match(r'(\S+) -> (\d+) (\S*)', line)
        if not m or m.group(2) != '301' or not m.group(3).endswith(m.group(1) + '/'):
            report('redirects', 'expected 301 to trailing slash', line)
    print('✓ redirects checked')

    print(f'\n{problems} unexpected difference(s)')
    return 1 if problems else 0


if __name__ == '__main__':
    if len(sys.argv) != 3:
        sys.exit(__doc__)
    sys.exit(main(sys.argv[1], sys.argv[2]))
