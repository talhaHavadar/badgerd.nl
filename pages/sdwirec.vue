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
            <div class="flex flex-row space-x-4 pt-4 font-light">
              <a href="https://github.com/Badger-Embedded/badgerd-sdwirec" rel="nofollow,external" target="_blank"
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
              <a href="#quick-start" class="text-zinc-700 hover:text-zinc-900">How-To</a>
              <a
                href="https://docs.dasharo.com/transparent-validation/sd-wire/usage-validation/#environment-preparation"
                target="_blank"
                rel="nofollow,external"
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
                    <p>SDWireC</p>
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
                  rel="nofollow"
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
                  rel="nofollow"
                  class="text-blue-500 hover:text-blue-700"
                  >Wiki</a
                >
                <br />
                <br />
                Go to <a href="#quick-start" class="text-blue-500 hover:text-blue-700">Quick Start</a> guide to start
                testing it immediately!
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <hr class="lg:mt-16" />
    <div class="my-4 flex justify-center">
      <div class="max-w-7xl space-y-4 px-6 sm:px-16 lg:px-32">
        <h1 id="quick-start" class="my-8 px-4 text-center text-lg font-bold lg:my-12">
          How to use Badgerd's SDWire with USB-C?
        </h1>

        <p class="">
          Our first step is to install dependencies to build the software we need to control SDWire. So let's execute
          the following command in shell.
        </p>
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
          >sudo apt install build-essential libftdi1-dev libpopt-dev cmake</code
        >
        <p>Now it is time to install the software to control the SDWireC.</p>
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
          >git clone https://github.com/Badger-Embedded/badgerd-sdwirec.git && cd badgerd-sdwirec/sdwirec-sw</code
        >

        <p>Build and install the sd-mux-ctrl application by executing standard cmake build commands.</p>
        <!-- prettier-ignore -->
        <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">mkdir build && cd build
cmake ..
make
sudo make install</code>
        <p>Commands above will install `sd-mux-ctrl` into /usr/local/bin</p>

        <p>
          Since we have the tools we need, now we can continue on actually using the SDWireC. To do that we need a
          Raspberry Pi like device that has an SD Card slot so that we can place our SDWireC into it. Plus, we also need
          an USB cable to connect SDWireC to our computer. (In this case I am using my
          <span class="text-orange-600">Ubuntu</span> host)
        </p>
        <ol class="list-decimal space-y-2 pl-4">
          <li>Make sure that Raspberry Pi is not powered</li>
          <li>Place SDWireC into Raspberry Pi's SD Card slot</li>
          <li>
            Let's check if our SDWireC is already configured or not
            <ul class="mt-2 list-disc space-y-2 pl-8">
              <li>
                <p>In your computer's terminal run the following command:</p>
                <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo dmesg -w</code>
              </li>
              <li>Connect your computer (or another host machine) with SDWireC using USB cable</li>
              <li>
                Then you should be able to see 'idVendor', 'idProduct', 'SerialNumber'(alphanumeric) of the SDWireC. We
                are going to use these values to configure if our SDWireC is not configured yet. Let's check that!
              </li>
              <li>
                <p>In another terminal window, run the following command:</p>
                <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
                  >sudo sd-mux-ctrl --list</code
                >
                <p>If you see following output in terminal, this means that your SDWireC is not configured.</p>
                <!-- prettier-ignore -->
                <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">Number of FTDI devices found: 0</code>
                <p>Let's configure it otherwise you can continue on step 4</p>
                <ul class="mt-2 list-[square] space-y-2 pl-12">
                  <li>
                    <p>Run the following command to configure SDWireC:</p>
                    <!-- prettier-ignore -->
                    <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo sd-mux-ctrl --device-serial={SerialNumber} --vendor=0x{idVendor} --product=0x{idProduct} --device-type=sd-wire --set-serial=sdwirec_10</code>
                    <p>Remember that all these values are coming from 'dmesg' output.</p>
                  </li>
                  <li>
                    <p>Now we should be able to see our SDWireC is recognized by 'sd-mux-ctrl'</p>
                    <!-- prettier-ignore -->
                    <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo sd-mux-ctrl --list
Number of FTDI devices found: 1
Dev: 0, Manufacturer: SRPOL, Serial: sdwirec_10, Description: sd-wire</code>
                  </li>
                </ul>
              </li>
            </ul>
          </li>
          <li>
            <p>Connect SDCard to your computer (host|ts)</p>
            <!-- prettier-ignore -->
            <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo sd-mux-ctrl --device-serial=sdwirec_10 --ts</code>
          </li>
          <li>
            Flash the SDCard using
            <a
              href="https://www.raspberrypi.com/software/"
              class="text-blue-500 hover:text-blue-700"
              rel="nofollow,external"
              >RPi Imager</a
            >
            or
            <a href="https://www.balena.io/etcher" class="text-blue-500 hover:text-blue-700" rel="nofollow,external"
              >BalenaEtcher</a
            >
          </li>
          <li>
            <p>Connect SD card to Raspberry Pi (DUT) with sd-mux-ctrl</p>
            <!-- prettier-ignore -->
            <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo sd-mux-ctrl --device-serial=sdwirec_10 --dut</code>
          </li>
          <li>
            <p>Connect power supply of Raspberry Pi (DUT) and then voila!</p>
          </li>
        </ol>

        <br />
        <p>Hope this quick start guide helps you to give our SDWireC a kickstart. See you in another challenge!</p>
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
      title: 'Buy SDWire With USB-C | Quickstart With Badgerd SDWireC',
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
          content:
            'Our way of implementing the infamous sdwire board by Tizen. Explore our variant of sdwire with an USB-C now it is SDWireC.',
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
            'Our way of implementing the infamous sdwire board by Tizen. Explore our variant of sdwire with an USB-C now it is SDWireC.',
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
    return [
      {
        '@context': 'http://schema.org',
        '@id': 'https://badgerd.nl/#sdwirec',
        '@type': 'Product',
        url: 'https://badgerd.nl/sdwirec/',
        name: 'Badgerd SDWireC',
        mainEntityOfPage: 'https://badgerd.nl/sdwirec/',
        image: 'https://www.badgerd.nl/sdwirec/badgerd-sdwirec-cad.png',
        desciption: 'USBC Variant of standard SDWire product, now we are calling it as SDWireC',
        offers: [
          {
            '@type': 'Offer',
            priceCurrency: 'EUR',
            price: 85,
            sku: 'badgerd_sdwirec',
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
        ],
      },
    ]
  },
}
</script>
