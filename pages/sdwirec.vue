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
              <div class="flex flex-wrap space-y-5 sm:flex-nowrap sm:space-y-2 sm:space-x-8">
                <div class="flex flex-col">
                  <div class="flex items-center text-lg font-semibold">
                    <p>SDWireC</p>
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
                  <div class="font-semibold">€85.00 <span class="text-xs font-light">(incl. VAT)</span> + Shipping</div>
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
                Check out the new generation SDWire3 available
                <a href="/sdwire3" target="_blank" class="text-blue-500 hover:text-blue-700">here</a>.
                <br />
                <br />
                The Badger<span class="font-bold text-red-700">'</span>d way of the infamous SDWire. It can be used to
                automate testing sdcard interface or in 3D printing or to create a fully remote working place. In our
                version of SDWire, we used USB-C interface to communicate with host.
                <br />
                <br />
                SDWireC allows to flash SD card connected to the DUT (Device Under Test), without physical contact with
                the device. There is only one USB-C socket for connecting to host PC. Both USB mass storage and MUX
                control are served through the same USB connection. The PCB board is designed in such way that it fits
                into micro SD card slots. Thanks to this, there is no need for special cables with a micro SD adapter,
                like in the muxPi product. It can even be used as SD card reader on steroids.
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
                <br />
                All the SDWireC products that will be shipped is tested on micro SD card slot of Raspberry Pi 3 Model B
                Version 1.2, using Gnome Disk Benchmark tool.
                <br />
                <br />
                SDWireC is tested and verified on following SD cards in house:
                <br />
                <br />
              </p>
              <ul class="list-disc pl-10">
                <li>SanDisk Ultra A1 Micro SDHC U1 Class 10</li>
                <li>Samsung PRO Plus A2 Micro SDXC U3 Class V30</li>
                <li>SanDisk Extreme A2 Micro SDXC U3 Class V30</li>
                <li>Samsung Evo A1 Micro SDHC</li>
              </ul>
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

        <hr class="lg:mt-16" />
        <h4 class="my-8 px-4 text-center text-lg lg:my-12">Python CLI App</h4>
        <p class="">
          For the ones who wants to try the board easily (even in windows), we have a python application which you can
          install using pip as shown below. If you are in a system where you can use apt (Ubuntu/Debian/Raspbian etc.)
          then we would recommend you to use apt to install the sdwire as shown below. Source code of the CLI
          application can be found
          <a
            href="https://github.com/badger-Embedded/sdwire-cli/"
            class="text-blue-500 hover:text-blue-700"
            rel="nofollow,external"
            >here</a
          >.
        </p>

        <!-- prettier-ignore -->
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo pip install sdwire
# use --help to see how to use the application
sdwire --help
sdwire switch --help
sudo sdwire list # to list the devices detected

# For Ubuntu users please use apt instead of pip
sudo add-apt-repository ppa:tchavadar/badgerd # for stable builds
sudo add-apt-repository ppa:tchavadar/badgerd-unstable # for daily builds
sudo apt install python3-sdwire

# For Rasbian/Debian users please use apt instead of pip directly
echo -e "deb https://ppa.launchpadcontent.net/tchavadar/badgerd/ubuntu noble main" | sudo tee -a /etc/apt/sources.list.d/badgerd.list
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B35AA5B4E400B24A
sudo apt install python3-sdwire
# Do not hesitate to open issue in our github repo: https://github.com/Badger-Embedded/sdwire-cli/issues</code>
        <p class="">
          Please keep in mind that new CLI doesnt have all the functionalities of old CLI application yet if you need
          those functionalities, please the guide below to unstall old CLI application.
        </p>
        <hr class="lg:mt-16" />
        <h4 class="my-8 px-4 text-center text-lg lg:my-12">Oldschool CLI Application</h4>
        <p class="">
          Our first step is to install dependencies to build the software we need to control SDWire. So let's execute
          the following command in shell.
        </p>
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
          >sudo apt install build-essential libftdi1-dev libpopt-dev cmake pkg-config</code
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
    const shippingRates = (await $axios.get(`/api/shippingrates`)).data
    let countries = []

    if (Object.prototype.hasOwnProperty.call(shippingRates, 'countries')) {
      countries = shippingRates.countries
    } else {
      for (const countryCode in shippingRates) {
        countries.push({ code: countryCode, name: shippingRates[countryCode].name })
      }
    }

    countries.sort((a, b) => (a.name > b.name ? 1 : -1))
    return { product, countries }
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
        name: 'Badgerd SDWireC | SDWire',
        mainEntityOfPage: 'https://badgerd.nl/sdwirec/',
        image: 'https://www.badgerd.nl/sdwirec/sdwirec-5.jpg',
        description: 'USB-C Variant of standard SDWire product, now we are calling it as SDWireC',
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
            datePublished: '2023-01-19',
            reviewBody:
              'It does what you expect from an SDWire device. And having it with USB-C is another plus for me.',
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
            price: 85,
            sku: 'badgerd_sdwirec',
            priceValidUntil: '2024-06-31',
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
