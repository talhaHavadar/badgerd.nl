<template>
  <div class="items-top relative flex min-h-[calc(100vh-3rem)] justify-center bg-zinc-100">
    <div class="mx-auto w-full text-zinc-700 sm:px-6 lg:px-8">
      <div class="flex flex-col items-center space-x-2 p-4 pt-16">
        <img :src="celebrationGif" class="w-full max-w-4xl md:w-[80%] 2xl:w-[50%]" alt="" srcset="" />
        <h2 class="mt-10 text-center text-4xl font-extrabold">Thank You!</h2>
        <h2 class="my-2 text-center text-lg font-bold">Order# {{ orderId }}</h2>
        <p class="mt-2 w-[80%] text-center">
          You made us this happy by just buying this <a :href="product.url" class="text-blue-500">beauty</a>.
          <br />
          Please keep your order number safe. Your shipping confirmation will be send to your email.
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    const orderId = this.$route.query.order_id
    const product = {}
    const celebrationGif = ''
    product.url = '#'
    return { product, orderId, celebrationGif }
  },
  head() {
    return {
      link: [
        {
          hid: 'canonical',
          rel: 'canonical',
          href: `${this.$config.baseURL}${this.$route.path}`,
        },
      ],
    }
  },

  async mounted() {
    const qCodename = this.$route.query.codename

    if (qCodename) {
      console.log(`looking for ${qCodename}`)
      const products = (await this.$axios.get(`/api/stripe/products`)).data
      this.product = products.find((p) => p.codename === qCodename)
    }
    this.celebrationGif = this.getRandomGif()
  },
  methods: {
    getRandomGif() {
      const gifs = [
        '/gif/media/l0amJzVHIAfl7jMDos/giphy.gif',
        '/gif/media/BpGWitbFZflfSUYuZ9/giphy.gif',
        '/gif/media/cXblnKXr2BQOaYnTni/giphy.gif',
        '/gif/media/IwAZ6dvvvaTtdI8SD5/giphy.gif',
        '/gif/media/6R2mLi910HL4VXFwOG/giphy.gif',
        '/gif/media/MZocLC5dJprPTcrm65/giphy.gif',
        '/gif/media/DhstvI3zZ598Nb1rFf/giphy.gif',
        '/gif/media/Hm3rh1nMYe9BR20ThG/giphy.gif',
        '/gif/media/XQf2LxzdXBt8yo6NcA/giphy.gif',
        '/gif/media/w0MYyUAvYCS64/giphy.gif',
        '/gif/media/sgZdwNkOH43wZOEUJ2/giphy.gif',
        '/gif/media/3ohhwemK8gvyPkHVzq/giphy.gif',
        '/gif/media/xTiTnLWl6ftNuZAe8E/giphy.gif',
        '/gif/media/i1hiQy3uVZ0KQ/giphy.gif',
        '/gif/media/NOSHrLbvT589ttrt8I/giphy.gif',
        '/gif/media/yqtpq8rqqXBh6/giphy.gif',
        '/gif/media/7ZvPlxBHwfK1y/giphy.gif',
        '/gif/media/ghutdpgRkhkxq/giphy.gif',
        '/gif/media/fE32ZJgwTysQo/giphy.gif',
      ]
      return gifs[Math.floor(Math.random() * gifs.length)]
    },
  },
}
</script>
