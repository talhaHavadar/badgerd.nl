<template>
  <div
    class="items-top relative flex min-h-[calc(100vh-3rem)] justify-center bg-zinc-100 subpixel-antialiased sm:items-center sm:pt-0"
  >
    <div class="max-w-9xl mx-auto sm:px-6 lg:px-8">
      <div class="flex flex-row flex-wrap items-center p-4 lg:flex-nowrap">
        <div class="flex flex-col px-2">
          <div id="carouselIndicators" class="slide carousel relative lg:w-[50vw]" data-bs-ride="carousel">
            <div class="carousel-indicators absolute right-0 bottom-0 left-0 mb-4 flex justify-center p-0">
              <button
                type="button"
                data-bs-target="#carouselIndicators"
                data-bs-slide-to="0"
                class="active"
                aria-current="true"
                aria-label="Slide 1"
              ></button>
              <!-- <button
                type="button"
                data-bs-target="#carouselIndicators"
                data-bs-slide-to="1"
                aria-label="Slide 2"
              ></button>
              <button
                type="button"
                data-bs-target="#carouselIndicators"
                data-bs-slide-to="2"
                aria-label="Slide 3"
              ></button> -->
            </div>
            <div class="carousel-inner relative w-full overflow-hidden">
              <div class="carousel-item active float-left w-full">
                <!-- active -->
                <img
                  :src="require('~/assets/sdwirec/badgerd-sdwirec-cad.png')"
                  class="w-full"
                  alt="KiCAD 3D View"
                  title="KiCAD 3D View"
                  loading="lazy"
                />
              </div>
              <!-- <div class="carousel-item float-left w-full">
                <img src="https://mdbcdn.b-cdn.net/img/new/slides/042.webp" class="block w-full" alt="Camera" />
              </div>
              <div class="carousel-item float-left w-full">
                <img src="https://mdbcdn.b-cdn.net/img/new/slides/043.webp" class="block w-full" alt="Exotic Fruits" />
              </div> -->
            </div>
            <button
              class="carousel-control-prev absolute top-0 bottom-0 left-0 flex items-center justify-center border-0 p-0 text-center hover:no-underline hover:outline-none focus:no-underline focus:outline-none"
              type="button"
              data-bs-target="#carouselIndicators"
              data-bs-slide="prev"
            >
              <span class="carousel-control-prev-icon inline-block bg-no-repeat" aria-hidden="true"></span>
              <span class="visually-hidden">Previous</span>
            </button>
            <button
              class="carousel-control-next absolute top-0 bottom-0 right-0 flex items-center justify-center border-0 p-0 text-center hover:no-underline hover:outline-none focus:no-underline focus:outline-none"
              type="button"
              data-bs-target="#carouselIndicators"
              data-bs-slide="next"
            >
              <span class="carousel-control-next-icon inline-block bg-no-repeat" aria-hidden="true"></span>
              <span class="visually-hidden">Next</span>
            </button>
          </div>
          <div class="pt-4 font-light">
            <a
              href="https://docs.dasharo.com/transparent-validation/sd-wire/usage-validation/#environment-preparation"
              target="_blank"
              rel="external"
              class="text-zinc-700 hover:text-zinc-900"
              >Docs</a
            >
          </div>
        </div>
        <div class="flex flex-col p-4">
          <form action="/api/stripe/checkout" method="post">
            <div class="flex">
              <div class="flex flex-col">
                <div class="flex items-center text-lg font-semibold">
                  <h1>SDWireC</h1>
                  <div
                    v-if="product.current_stock < 10 && product.current_stock > 0"
                    class="pl-2 text-xs font-normal text-orange-600"
                  >
                    Last {{ product.current_stock }} items
                  </div>
                </div>

                <input type="hidden" name="codename" value="badgerd_sdwirec" />
                <div class="font-semibold">€85.00 + shipping and handling</div>
                <div
                  v-if="product.current_stock == 0 && !product.ignore_stock"
                  class="flex text-sm font-light text-red-600"
                >
                  Out of Stock
                </div>
                <div v-else-if="product.ignore_stock" class="flex text-sm font-light text-orange-600">
                  Open for Pre-Order
                </div>
              </div>
              <div class="ml-8 flex items-start justify-center">
                <select
                  class="form-select-sm form-select m-0 mb-3 block w-16 appearance-none rounded border border-solid border-gray-300 bg-white bg-clip-padding bg-no-repeat px-4 py-2 text-sm font-normal text-gray-700 transition ease-in-out focus:border-blue-500 focus:bg-white focus:text-zinc-700 focus:outline-none"
                  aria-label=".form-select-sm"
                  name="quantity"
                  :disabled="product.current_stock == 0 && !product.ignore_stock"
                  required
                >
                  <template v-if="product.current_stock < 10 && product.current_stock > 0">
                    <option v-for="n in product.current_stock" :key="n" :selected="n == 1" :value="n">
                      {{ n }}
                    </option>
                  </template>
                  <template v-else-if="product.current_stock > 10 || product.ignore_stock">
                    <option v-for="n in 10" :key="n" :selected="n == 1" :value="n">
                      {{ n }}
                    </option>
                  </template>
                </select>
                <button
                  type="submit"
                  data-mdb-ripple="true"
                  data-mdb-ripple-color="light"
                  :disabled="product.current_stock == 0 && !product.ignore_stock"
                  class="ml-4 inline-block min-w-[6rem] rounded border border-zinc-700 bg-zinc-100 px-6 py-2.5 text-xs font-semibold uppercase leading-tight text-zinc-700 shadow-md transition duration-150 ease-in-out hover:bg-zinc-200 hover:shadow-lg focus:bg-zinc-700 focus:text-zinc-200 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-zinc-800 active:shadow-lg disabled:border-zinc-200 disabled:text-zinc-400"
                >
                  Pre Order
                </button>
              </div>
            </div>
          </form>

          <div class="mt-16 mb-8">
            <div v-if="product.reserved_stock > 0" class="text-xs font-normal text-red-600">
              {{ product.reserved_stock }} items reserved, might be yours, try again in 30 minutes
            </div>
            <p v-if="product.current_stock > 0 || product.ignore_stock" class="text-xs font-light">
              <span class="font-semibold">Living in EU?</span>
              <br />
              Then try this
              <a
                href="https://buy.stripe.com/dR6013dsb0wr8BWbIR"
                rel="external"
                class="text-blue-500 hover:text-blue-700"
                >link</a
              >
              for different shipping rates.
            </p>
            <hr class="mt-4" />
          </div>
          <div class="flex flex-col justify-around text-sm">
            <p class="">
              The Badger<span class="font-bold text-red-700">'</span>d way of the infamous SDWire. It can be used to
              automate testing sdcard interface or in 3D printing or to create a fully remote working place. In our
              version of SDWire, we used USB-C interface to communicate with host.
              <br />
              <br />
              SDWireC allows to flash SD card connected to the DUT (Device Under Test), without physical contact with
              the device. There is only one USB-C socket for connecting to host PC. Both USB mass storage and MUX
              control are served through the same USB connection. The PCB board is designed in such way that it fits
              into micro SD card slots. Thanks to this, there is no need for special cables with a micro SD adapter,
              like in the muxPi product.
              <br />
              <br />
              Based on Tizen SDWire:
              <a
                href="https://wiki.tizen.org/SDWire"
                target="_blank"
                rel="external"
                class="text-blue-500 hover:text-blue-700"
                >Wiki</a
              >
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const products = (await $axios.get(`/api/stripe/products`)).data
    const product = products.find((p) => p.codename === 'badgerd_sdwirec')
    return { product }
  },
  head() {
    return {
      title: 'SDWireC | Badgerd Technologies',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content:
            'Our way of implementing the infamous sdwire board by Tizen. Explore our variant of sdwire with an USBC now it is SDWireC.',
        },
        {
          hid: 'og:title',
          property: 'og:title',
          content: 'Badgerd SDWireC',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content:
            'Our way of implementing the infamous sdwire board by Tizen. Explore our variant of sdwire with an USBC now it is SDWireC.',
        },
        {
          hid: 'og:type',
          property: 'og:type',
          content: 'website',
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content: 'https://www.badgerd.nl/sdwirec/badgerd-sdwirec-cad.png',
        },
        {
          hid: 'og:url',
          property: 'og:url',
          content: `${this.$config.baseURL}${this.$route.path}`,
        },
        {
          hid: 'og:locale',
          property: 'og:locale',
          content: 'en_US',
        },
      ],
    }
  },
  jsonld() {
    return {
      '@context': 'http://schema.org',
      '@id': 'https://badgerd.nl/#sdwirec',
      '@type': 'Product',
      url: 'https://badgerd.nl/sdwirec',
      name: 'Badgerd SDWireC',
      mainEntityOfPage: 'https://badgerd.nl/sdwirec',
      image: 'https://www.badgerd.nl/sdwirec/badgerd-sdwirec-cad.png',
      offers: [
        {
          '@type': 'Offer',
          priceCurrency: 'EUR',
          price: 85,
          sku: 'badgerd_sdwirec',
        },
      ],
    }
  },
}
</script>
