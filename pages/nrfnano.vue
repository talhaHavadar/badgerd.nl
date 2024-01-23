<template>
  <div>
    <div
      class="items-top relative flex min-h-[calc(100vh-3rem)] flex-col justify-center bg-zinc-100 subpixel-antialiased sm:items-center sm:pt-0"
    >
      <div class="mx-auto max-w-[1656px] sm:px-6 lg:px-8">
        <div class="flex flex-row flex-wrap items-center p-4 lg:flex-nowrap">
          <div class="flex flex-col px-2">
            <div
              id="carouselIndicators"
              class="slide carousel relative max-h-[828px] max-w-[828px] lg:w-[50vw]"
              data-bs-ride="carousel"
            >
              <div class="carousel-indicators absolute right-0 bottom-0 left-0 mb-4 flex justify-center p-0">
                <button
                  type="button"
                  data-bs-target="#carouselIndicators"
                  data-bs-slide-to="0"
                  class="active"
                  aria-current="true"
                  aria-label="Slide 1"
                ></button>
                <button
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
                ></button>
              </div>
              <div class="carousel-inner relative w-full overflow-hidden">
                <div class="carousel-item active relative float-left w-full">
                  <img
                    :src="require('~/assets/sdwirec/sdwirec-1.jpg')"
                    class="block w-full"
                    alt="SDWireC Top View"
                    title="SDWireC Top View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/sdwirec/sdwirec-2.jpg')"
                    class="block w-full"
                    alt="SDWireC Bottom View"
                    title="SDWireC Bottom View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/sdwirec/sdwirec-3.jpg')"
                    class="block w-full"
                    alt="SDWireC Being Used in Remote"
                    title="SDWireC Being Used in Remote"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
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
            <div class="flex flex-row space-x-4 pt-4 font-light">
              <a href="https://github.com/Badger-Embedded/badgerd-nrfnano" rel="nofollow,external" target="_blank"
                ><svg
                  class="button--github h-6 w-6 text-zinc-600 hover:text-zinc-800"
                  xmlns="http://www.w3.org/2000/svg"
                  xmlns:xlink="http://www.w3.org/1999/xlink"
                  aria-hidden="true"
                  role="img"
                  width="32"
                  height="32"
                  preserveAspectRatio="xMidYMid meet"
                  viewBox="0 0 24 24"
                >
                  <path
                    d="M12 2.247a10 10 0 0 0-3.162 19.487c.5.088.687-.212.687-.475c0-.237-.012-1.025-.012-1.862c-2.513.462-3.163-.613-3.363-1.175a3.636 3.636 0 0 0-1.025-1.413c-.35-.187-.85-.65-.013-.662a2.001 2.001 0 0 1 1.538 1.025a2.137 2.137 0 0 0 2.912.825a2.104 2.104 0 0 1 .638-1.338c-2.225-.25-4.55-1.112-4.55-4.937a3.892 3.892 0 0 1 1.025-2.688a3.594 3.594 0 0 1 .1-2.65s.837-.262 2.75 1.025a9.427 9.427 0 0 1 5 0c1.912-1.3 2.75-1.025 2.75-1.025a3.593 3.593 0 0 1 .1 2.65a3.869 3.869 0 0 1 1.025 2.688c0 3.837-2.338 4.687-4.563 4.937a2.368 2.368 0 0 1 .675 1.85c0 1.338-.012 2.413-.012 2.75c0 .263.187.575.687.475A10.005 10.005 0 0 0 12 2.247z"
                    fill="currentColor"
                  /></svg
              ></a>
              <a href="#faq" class="text-zinc-700 hover:text-zinc-900">FAQ</a>
              <a href="#pinout" class="text-zinc-700 hover:text-zinc-900">Pinout</a>
            </div>
          </div>
          <div class="flex flex-col p-4">
            <form action="/api/stripe/checkout" method="post">
              <div class="flex flex-wrap space-y-5 sm:flex-nowrap sm:space-y-2 sm:space-x-8">
                <div class="flex flex-col">
                  <div class="flex items-center text-lg font-semibold">
                    <p>nRFNANO</p>
                    <div
                      v-if="product.current_stock < 10 && product.current_stock > 0"
                      class="pl-2 text-xs font-normal text-orange-600"
                    >
                      Last {{ product.current_stock }} items
                    </div>
                    <div v-else-if="product.current_stock > 50" class="pl-2 text-xs font-normal text-green-600">
                      50+ items
                    </div>
                    <div v-else class="pl-2 text-xs font-normal text-green-600">
                      Last {{ product.current_stock }} items
                    </div>
                  </div>

                  <input type="hidden" name="codename" value="badgerd_sdwirec" />
                  <div class="font-semibold">
                    €19.00 <span class="text-xs font-light">(incl. VAT*)</span> + Shipping
                  </div>
                  <div
                    v-if="product.current_stock == 0 && !product.ignore_stock"
                    class="flex text-sm font-light text-red-600"
                  >
                    Out of Stock
                  </div>
                  <div
                    v-else-if="!product.ignore_stock && product.current_stock > 0"
                    class="flex text-sm font-bold text-green-600"
                  >
                    In Stock
                  </div>
                  <div v-else-if="product.ignore_stock" class="flex text-sm font-light text-orange-600">
                    Open for Pre-Order
                  </div>
                </div>
                <div class="flex flex-col space-y-2">
                  <div class="flex space-x-2">
                    <select
                      class="form-select-sm form-select m-0 block appearance-none rounded border border-solid border-gray-300 bg-white bg-clip-padding bg-no-repeat px-6 py-2 text-sm font-normal text-gray-700 transition ease-in-out focus:border-blue-500 focus:bg-white focus:text-zinc-700 focus:outline-none"
                      aria-label=".form-select-sm"
                      name="quantity"
                      :disabled="product.current_stock == 0 && !product.ignore_stock"
                      required
                    >
                      <template v-if="product.current_stock < 20 && product.current_stock > 0">
                        <option v-for="n in product.current_stock" :key="n" :selected="n == 1" :value="n">
                          {{ n }}
                        </option>
                      </template>
                      <template v-else-if="product.current_stock > 20 || product.ignore_stock">
                        <option v-for="n in 20" :key="n" :selected="n == 1" :value="n">
                          {{ n }}
                        </option>
                      </template>
                    </select>
                    <select
                      class="form-select-sm form-select m-0 block appearance-none rounded border border-solid border-gray-300 bg-white bg-clip-padding bg-no-repeat px-8 py-2 text-sm font-normal text-gray-700 transition ease-in-out focus:border-blue-500 focus:bg-white focus:text-zinc-700 focus:outline-none"
                      aria-label=".form-select-sm"
                      name="country_code"
                      :disabled="product.current_stock == 0 && !product.ignore_stock"
                      required
                    >
                      <option value="" disabled selected>Your Country</option>
                      <option v-for="country in countries" :key="country.code" :value="country.code">
                        {{ country.name }}
                      </option>
                    </select>
                  </div>
                  <button
                    type="submit"
                    data-mdb-ripple="true"
                    data-mdb-ripple-color="light"
                    :disabled="product.current_stock == 0 && !product.ignore_stock"
                    class="inline-block w-full min-w-[6rem] rounded border border-zinc-700 bg-zinc-100 px-6 py-2.5 text-xs font-semibold uppercase leading-tight text-zinc-700 shadow-md transition duration-150 ease-in-out hover:bg-zinc-200 hover:shadow-lg focus:bg-zinc-700 focus:text-zinc-200 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-zinc-800 active:shadow-lg disabled:border-zinc-200 disabled:text-zinc-400"
                  >
                    Buy
                  </button>
                </div>
              </div>
            </form>

            <div class="mt-16 mb-8">
              <div v-if="product.reserved_stock > 0" class="text-xs font-normal text-red-600">
                {{ product.reserved_stock }} items reserved, might be yours, try again in 30 minutes
              </div>
              <p v-if="product.current_stock > 0 || product.ignore_stock" class="text-xs font-light">
                <span class="font-semibold">Looking for a bulk order or B2B?</span>
                <br />
                <a href="mailto:info+quote@badgerd.nl" rel="nofollow" class="text-blue-500 hover:text-blue-700"
                  >Contact</a
                >
                us to get a quote.
              </p>
              <hr class="mt-4" />
            </div>
            <div class="flex flex-col justify-around text-sm">
              <p class="">
                The Badger<span class="font-bold text-red-700">'</span>d nRFNANO is an Arduino ProMicro form factor
                nrf52840 development board.
                <br />
                <br />
                nRFNANO comes with fully compatible pinout with nice!nano. In addition to that you will have an exposed
                pad to add extra safety for your charging circuit which allows you to connect a temperature sensor.
                <br />
                <br />
                It comes with UF2 drag and drop bootloader which makes life way much easier for developers. With extra
                isolated PCB antenna, it will provide relatively powerful signal which allows more stable connections
                when there are obstacles in between.
                <br />
                <br />
                <span class="font-medium">Tech Specs: </span> <br />
                <span class="block pl-4">
                  Microcontroller: nRF52840 <br />
                  Flash: 1MB <br />
                  RAM: 256KB <br />
                  GPIO Count: 19<br />
                  Battery Charger: 4.2V Lithium @ 100/500mA rate <br />
                  Voltage 3.3V
                </span>
                <br />
                <br />
                Check the <a href="#pinout" class="text-blue-500 hover:text-blue-700">Pinout</a> or go to
                <a href="#faq" class="text-blue-500 hover:text-blue-700">FAQ</a> to start using it immediately!
                <br />
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <hr class="lg:mt-16" />
    <div class="my-4 flex justify-center">
      <div class="max-w-7xl px-6 sm:px-16 lg:px-32">
        <h1 id="pinout" class="my-8 px-4 text-center text-lg font-bold lg:my-12">Pinout of Badgerd nRFNANO</h1>
        <img
          :src="require('~/assets/nrfnano/diagram.svg')"
          class="block w-full"
          alt="SDWireC Top View"
          title="SDWireC Top View"
          loading="eager"
          width="828"
          height="828"
        />
        <hr class="lg:pb-16" />

        <h1 id="faq" class="my-8 px-4 text-center text-lg font-bold lg:my-12">Frequently Asked Questions</h1>

        <p class="mt-4 font-semibold">Where can I find the example application that the board came with?</p>
        <p>
          We are publishing some of the materials we use during development of nRFNANO in
          <a
            href="https://github.com/Badger-Embedded/badgerd-nrfnano"
            target="_blank"
            class="text-blue-500 hover:text-blue-700"
            >this</a
          >
          github repo. You can find the example applications in the repository.
        </p>
        <p class="mt-4 font-semibold">How can I put nRFNANO into bootloader mode?</p>
        <p>
          To activate UF2 bootloader simply double tap nRESET pin. You can use a pair of tweezers to connect nRESET and
          GND pins of the board twice, this will activate bootloader mode so that you can reflash the UF2 firmware you
          have.
        </p>
        <p class="mt-4 font-semibold">Where can I find the original UF2 firmware for nRFNANO?</p>
        <p>
          nRFNANO has compatible hardware design and pinout with nice!nano so the firmwares compatible with nice!nano
          should work with nRFNANO as well. But if you want to download the firmware we ship, you can get it from
          <a
            href="https://badgerd.nl/nrfnano/nrfnano_bootloader-0.7.0_s140_6.1.1.hex"
            rel="nofollow"
            target="_blank"
            class="text-blue-500 hover:text-blue-700"
            >here</a
          >. <br />
          To download the circuitpython UF2 file we ship please use
          <a
            href="https://badgerd.nl/nrfnano/adafruit-circuitpython-nice_nano-en_US-8.2.9.uf2"
            rel="nofollow"
            target="_blank"
            class="text-blue-500 hover:text-blue-700"
            >this</a
          >
          link.
        </p>
        <p class="mt-4 font-semibold">Is there a specific connector for debug pins?</p>
        <p>
          Yes, footprint of the debug pins are compatible with Tag-Connect cables, we used
          <a
            href="https://www.tag-connect.com/product/tc2030-ctx-nl-6-pin-no-legs-cable-with-10-pin-micro-connector-for-cortex-processors"
            rel="nofollow"
            target="_blank"
            class="text-blue-500 hover:text-blue-700"
            >TC2030-CTX-NL</a
          >
          footprint to prevent cable mess and speed-up the debugging/testing of the boards.
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const products = (await $axios.get(`/api/stripe/products`)).data
    const product = products.find((p) => p.codename === 'badgerd_nrfnano')
    const shippingRates = (await $axios.get(`/api/shippingrates`)).data
    const countries = []
    for (const countryCode in shippingRates) {
      countries.push({ code: countryCode, name: shippingRates[countryCode].name })
    }
    countries.sort((a, b) => (a.name > b.name ? 1 : -1))
    return { product, countries }
  },
  head() {
    return {
      title: 'Buy nRFNANO IoT Development Board | Quickstart With Badgerd nRFNANO (nrf52840)',
      link: [
        {
          hid: 'canonical',
          rel: 'canonical',
          href: `${this.$config.baseURL}${this.$route.path}`,
        },
      ],
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: `The Badgerd nRFNANO is an Arduino ProMicro form factor nrf52840 development board. nRFNANO comes with fully compatible pinout with nice!nano. In addition to that you will have an exposed pad to add extra safety for your charging circuit which allows you to connect a temperature sensor.`,
        },
        {
          hid: 'og:title',
          property: 'og:title',
          content: 'Badgerd nRFNANO',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content: `The Badgerd nRFNANO is an Arduino ProMicro form factor nrf52840 development board. nRFNANO comes with fully compatible pinout with nice!nano. In addition to that you will have an exposed pad to add extra safety for your charging circuit which allows you to connect a temperature sensor.`,
        },
        {
          hid: 'og:type',
          property: 'og:type',
          content: 'website',
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content: 'https://www.badgerd.nl/nrfnano/nrfnano-0.jpg',
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
    return [
      {
        '@context': 'http://schema.org',
        '@id': 'https://badgerd.nl/#sdwirec',
        '@type': 'Product',
        url: 'https://badgerd.nl/sdwirec/',
        name: 'Badgerd SDWireC | SDWire',
        mainEntityOfPage: 'https://badgerd.nl/nrfnano/',
        image: 'https://www.badgerd.nl/nrfnano/nrfnano-0.jpg',
        description: 'The Badgerd nRFNANO is an Arduino ProMicro form factor nrf52840 development board.',
        brand: {
          logo: 'https://badgerd.nl/logo-nobg.png',
          name: 'Badgerd',
          slogan: 'Creating software for hardware, hardware for humanity',
        },
        aggregateRating: {
          '@type': 'AggregateRating',
          ratingValue: 5,
          reviewCount: 1,
        },
        review: [
          {
            '@type': 'Review',
            author: {
              '@type': 'Person',
              name: 'Talha',
            },
            datePublished: '2024-01-19',
            reviewBody:
              'nRFNANO is super easy to learn embedded development for iot by providing circuitpython as a default firmware',
            name: 'First User',
            reviewRating: {
              '@type': 'Rating',
              bestRating: 5,
              ratingValue: 5,
              worstRating: 1,
            },
          },
        ],
        offers: [
          {
            '@type': 'Offer',
            priceCurrency: 'EUR',
            availability: 'https://schema.org/InStock',
            price: 19,
            sku: 'badgerd_nrfnano',
            priceValidUntil: '2024-12-31',
          },
        ],
      },
      {
        '@context': 'https://schema.org/',
        '@type': 'BreadcrumbList',
        itemListElement: [
          {
            '@type': 'ListItem',
            position: 1,
            name: 'Store',
            item: 'https://badgerd.nl/',
          },
          {
            '@type': 'ListItem',
            position: 2,
            name: 'SDWireC',
            item: 'https://badgerd.nl/sdwirec/',
          },
          {
            '@type': 'ListItem',
            position: 3,
            name: 'USBMux',
            item: 'https://badgerd.nl/usbmux/',
          },
          {
            '@type': 'ListItem',
            position: 4,
            name: 'nRFNANO',
            item: 'https://badgerd.nl/nrfnano/',
          },
        ],
      },
    ]
  },
}
</script>
