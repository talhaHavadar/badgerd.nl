export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Badgerd Technologies B.V.',
    htmlAttrs: {
      lang: 'en',
      style: 'font-family: Montserrat',
    },
    meta: [
      { charset: 'utf-8' },
      { hid: 'viewport', name: 'viewport', content: 'width=device-width, initial-scale=1.0' },

      {
        hid: 'description',
        name: 'description',
        content:
          'Discover our innovative solutions for interesting problems in our life. Buy one of our products to support us!',
      },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'apple-touch-icon', sizes: '180x180', href: '/apple-touch-icon.png' },
      { rel: 'icon', type: 'image/png', sizes: '32x32', href: '/favicon-32x32.png' },
      { rel: 'icon', type: 'image/png', sizes: '16x16', href: '/favicon-16x16.png' },
      { rel: 'manifest', href: '/site.webmanifest' },
      { hid: 'canonical', rel: 'canonical', href: 'https://badgerd.nl' },
    ],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ['@/assets/css/main.css'],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [{ src: '~/node_modules/tw-elements', mode: 'client' }, '~/plugins/jsonld'],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    '@nuxtjs/eslint-module',
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxt/postcss8',
    '@nuxtjs/google-fonts',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    // https://go.nuxtjs.dev/content
    '@nuxt/content',
    '@nuxtjs/proxy',
    '@nuxtjs/sitemap',
    '@nuxtjs/robots',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: '/',
  },

  // Content module configuration: https://go.nuxtjs.dev/config-content
  content: {},

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    postcss: {
      plugins: {
        tailwindcss: {},
        autoprefixer: {},
      },
    },
  },

  googleFonts: {
    families: {
      Montserrat: {
        wght: [100, 200, 300, 400, 500, 600, 700, 800],
      },
    },
    download: false,
  },

  axios: {
    proxy: true,
  },

  proxy: {
    '/gif': {
      target: 'https://media.giphy.com',
      pathRewrite: {
        '^/gif': '/',
      },
    },
    '/api': {
      //target: 'http://0.0.0.0:8000',
      target: 'https://governor.badgerd.nl',
      pathRewrite: {
        '^/api': '',
      },
    },
  },

  robots: {
    UserAgent: '*',
    Disallow: '',
    Sitemap: 'https://badgerd.nl/sitemap.xml',
  },

  sitemap: {
    trailingSlash: true,
  },

  publicRuntimeConfig: {
    baseURL: process.env.BASE_URL || 'https://badgerd.nl',
  },
}
