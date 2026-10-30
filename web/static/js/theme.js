// Light/dark theme toggle (nav button with data-theme-toggle). The initial `.dark` class on <html> is set before
// first paint by the inline script in components.Layout. A choice is stored only when it differs from the OS
// preference, so switching back to the OS theme means "follow the OS" again.
;(() => {
  const root = document.documentElement
  const media = matchMedia('(prefers-color-scheme: dark)')
  const stored = () => {
    try {
      return localStorage.getItem('theme')
    } catch {
      return null
    }
  }
  const apply = (dark) => {
    root.classList.toggle('dark', dark)
    for (const b of document.querySelectorAll('[data-theme-toggle]')) b.setAttribute('aria-pressed', String(dark))
  }

  apply(root.classList.contains('dark'))
  media.addEventListener('change', (e) => {
    if (!stored()) apply(e.matches)
  })
  document.addEventListener('click', (e) => {
    if (!e.target.closest('[data-theme-toggle]')) return
    const dark = !root.classList.contains('dark')
    try {
      if (dark === media.matches) localStorage.removeItem('theme')
      else localStorage.setItem('theme', dark ? 'dark' : 'light')
    } catch {}
    apply(dark)
  })
})()
