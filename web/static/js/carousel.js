// Minimal replacement for the tw-elements carousel: prev/next controls,
// indicators, autoplay (paused on hover / hidden tab) and the same slide
// classes as before (carousel-item-next/prev/start/end, see input.css).
;(() => {
  const DURATION = 600

  for (const root of document.querySelectorAll('[data-carousel]')) {
    const items = [...root.querySelectorAll('.carousel-item')]
    const dots = [...root.querySelectorAll('.carousel-indicators > button')]
    if (items.length < 2) continue

    const interval = Number(root.dataset.interval) || 5000
    const reduceMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    let index = Math.max(0, items.findIndex((el) => el.classList.contains('active')))
    let busy = false
    let timer = null
    let hovering = false

    const show = (to, dir) => {
      to = (to + items.length) % items.length
      if (busy || to === index) return
      dir = dir || (to > index ? 'next' : 'prev')
      const from = items[index]
      const next = items[to]
      dots.forEach((d, i) => {
        d.classList.toggle('active', i === to)
        if (i === to) d.setAttribute('aria-current', 'true')
        else d.removeAttribute('aria-current')
      })
      index = to

      if (reduceMotion) {
        from.classList.remove('active')
        next.classList.add('active')
        return
      }
      const order = dir === 'next' ? 'carousel-item-next' : 'carousel-item-prev'
      const edge = dir === 'next' ? 'carousel-item-start' : 'carousel-item-end'
      busy = true
      next.classList.add(order)
      void next.offsetWidth // force reflow so the transition runs
      from.classList.add(edge)
      next.classList.add(edge)

      let done = false
      const finish = () => {
        if (done) return
        done = true
        next.removeEventListener('transitionend', onEnd)
        next.classList.remove(order, edge)
        next.classList.add('active')
        from.classList.remove('active', edge)
        busy = false
      }
      const onEnd = (e) => {
        if (e.target === next && e.propertyName === 'transform') finish()
      }
      next.addEventListener('transitionend', onEnd)
      setTimeout(finish, DURATION + 100) // fallback if transitionend never fires
    }

    const restart = () => {
      clearInterval(timer)
      timer = null
      if (!hovering && !document.hidden) timer = setInterval(() => show(index + 1, 'next'), interval)
    }

    root.querySelector('[data-slide="prev"]')?.addEventListener('click', () => {
      show(index - 1, 'prev')
      restart()
    })
    root.querySelector('[data-slide="next"]')?.addEventListener('click', () => {
      show(index + 1, 'next')
      restart()
    })
    dots.forEach((d) =>
      d.addEventListener('click', () => {
        show(Number(d.dataset.slideTo))
        restart()
      })
    )
    root.addEventListener('mouseenter', () => {
      hovering = true
      restart()
    })
    root.addEventListener('mouseleave', () => {
      hovering = false
      restart()
    })
    document.addEventListener('visibilitychange', restart)
    restart()
  }
})()
