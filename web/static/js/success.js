// Order confirmation page: link the purchased product and show a random gif.
;(() => {
  const gifs = [
    'l0amJzVHIAfl7jMDos',
    'BpGWitbFZflfSUYuZ9',
    'cXblnKXr2BQOaYnTni',
    'IwAZ6dvvvaTtdI8SD5',
    '6R2mLi910HL4VXFwOG',
    'MZocLC5dJprPTcrm65',
    'DhstvI3zZ598Nb1rFf',
    'Hm3rh1nMYe9BR20ThG',
    'XQf2LxzdXBt8yo6NcA',
    'w0MYyUAvYCS64',
    'sgZdwNkOH43wZOEUJ2',
    '3ohhwemK8gvyPkHVzq',
    'xTiTnLWl6ftNuZAe8E',
    'i1hiQy3uVZ0KQ',
    'NOSHrLbvT589ttrt8I',
    'yqtpq8rqqXBh6',
    '7ZvPlxBHwfK1y',
    'ghutdpgRkhkxq',
    'fE32ZJgwTysQo',
  ]
  const img = document.getElementById('celebration-gif')
  if (img) img.src = `/gif/media/${gifs[Math.floor(Math.random() * gifs.length)]}/giphy.gif`

  const codename = new URLSearchParams(window.location.search).get('codename')
  const link = document.getElementById('product-link')
  if (!codename || !link) return
  fetch('/api/stripe/products')
    .then((res) => (res.ok ? res.json() : []))
    .then((products) => {
      const product = products.find((p) => p.codename === codename)
      if (product && /^https?:\/\//.test(product.url)) link.href = product.url
    })
    .catch(() => {})
})()
