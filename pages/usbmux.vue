<template>
  <div>
    <div class="relative w-full bg-blue-400 p-4 text-center font-medium text-slate-200">
      10% of your purchase of USBMux will be donated to
      <a href="https://oceanfdn.org/" rel="nofollow" target="_blank">The Ocean Foundation</a>
      <span class="block text-sm font-light">Thank you for supporting efforts to save our oceans!</span>
    </div>
    <div
      class="items-top relative flex min-h-[calc(100vh-10rem)] flex-col justify-center bg-zinc-100 subpixel-antialiased sm:items-center sm:pt-0"
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
                <button
                  type="button"
                  data-bs-target="#carouselIndicators"
                  data-bs-slide-to="3"
                  aria-label="Slide 4"
                ></button>
                <button
                  type="button"
                  data-bs-target="#carouselIndicators"
                  data-bs-slide-to="4"
                  aria-label="Slide 5"
                ></button>
                <button
                  type="button"
                  data-bs-target="#carouselIndicators"
                  data-bs-slide-to="5"
                  aria-label="Slide 6"
                ></button>
              </div>
              <div class="carousel-inner relative w-full overflow-hidden">
                <div class="carousel-item active relative float-left w-full">
                  <img
                    :src="require('~/assets/usbmux/usbmux_2.jpg')"
                    class="block w-full"
                    alt="USBMux Cad View"
                    title="USBMux Cad View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/usbmux/usbmux_1.jpg')"
                    class="block w-full"
                    alt="USBMux Cad View"
                    title="USBMux Cad View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/usbmux/usbmux_3.jpg')"
                    class="block w-full"
                    alt="USBMux Cad View"
                    title="USBMux Cad View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/usbmux/usbmux_4.jpg')"
                    class="block w-full"
                    alt="USBMux Cad View"
                    title="USBMux Cad View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/usbmux/usbmux_5.jpg')"
                    class="block w-full"
                    alt="USBMux Cad View"
                    title="USBMux Cad View"
                    loading="eager"
                    width="828"
                    height="828"
                  />
                </div>
                <div class="carousel-item relative float-left w-full">
                  <img
                    :src="require('~/assets/usbmux/usbmux_6.jpg')"
                    class="block w-full"
                    alt="USBMux Cad View"
                    title="USBMux Cad View"
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
              <a href="https://github.com/Badger-Embedded/muxer" rel="nofollow,external" target="_blank"
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
            </div>
          </div>
          <div class="flex flex-col p-4">
            <form action="/api/stripe/checkout" method="post">
              <div class="flex flex-wrap space-y-5 sm:flex-nowrap sm:space-y-2 sm:space-x-8">
                <div class="flex flex-col">
                  <div class="flex items-center text-lg font-semibold">
                    <p>USBMux</p>
                    <div
                      v-if="product.current_stock < 10 && product.current_stock > 0"
                      class="pl-2 text-xs font-normal text-orange-600"
                    >
                      Last {{ product.current_stock }} items
                    </div>
                    <div v-else-if="product.current_stock != 0" class="pl-2 text-xs font-normal text-green-600">
                      Last {{ product.current_stock }} items
                    </div>
                  </div>

                  <input type="hidden" name="codename" value="badgerd_usbmux" />
                  <div class="font-semibold">
                    €180.00 <span class="text-xs font-light">(incl. VAT)</span> + Shipping
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
                <span v-if="product.ignore_stock" class="block font-semibold text-orange-600"
                  >Use <span class="font-bold">PREORDER30</span> at checkout to get 30% discount.</span
                >
                <span class="font-semibold">Looking for a bulk order or B2B(VAT-free)?</span>
                <br />
                <a href="mailto:info+quoteusbmux@badgerd.nl" rel="nofollow" class="text-blue-500 hover:text-blue-700"
                  >Contact</a
                >
                us to get a quote.
              </p>
              <hr class="mt-4" />
            </div>
            <div class="flex flex-col justify-around text-sm">
              <p class="">
                USBMux is a device that helps you to control connection of three different usb ports. USB-mux has three
                USB 2.0 ports, Host, DUT and Device.
                <br />
                <br />
                The Host port is used to control the USBMux with our cli software tool.
                <br />
                <br />
                The Device port is used to connect USB devices like USB flash storage devices, mouses etc.
                <br />
                <br />
                The DUT port is used to connect the device under test. It allows device itself to act as USB host or USB
                device. You can also control USB ID pin using our usb-mux-ctrl cli tool.
                <br />
                <br />
                The USBMux uses the analog switches and power transistors to create connections between all these three
                ports. It allows you to connect
                <span class="font-semibold">Host to Device</span>, <span class="font-semibold">Device to DUT</span> and
                <span class="font-semibold">Host to DUT</span> using cli tool.
                <br />
                <br />
                You can connect your host to the device port to write data to a connected flash storage before
                connecting DUT to Device port. Or you can connect your DUT to the host port to test DUT device in its
                device functionality.
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
        <h1 id="quick-start" class="my-8 px-4 text-center text-lg font-bold lg:my-12">How to use Badgerd's USBMux?</h1>

        <p class="">
          Let's install the command line application that will help us control the Badgerd USBMux device. We have
          already created a PPA to make the installation of the application easy for you. So let's execute the following
          commands to install the application. If you want to build the application from the source please follow the
          instructions
          <a href="#build-from-source" class="text-blue-500 hover:text-blue-700">here</a>.
        </p>
        <!-- prettier-ignore -->
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo add-apt-repository ppa:tchavadar/badgerd
sudo apt update
sudo apt install badgerd-muxer</code>
        <p>
          Once the installation is completed successfully. Now you can see the demonstrations steps to see what you can
          do with the Badgerd USBMux.
          <a href="#demonstration" class="text-blue-500 hover:text-blue-700">Jump the demonstration steps here.</a>
        </p>
        <p id="build-from-source" class="">Building the badgerd-muxer from source.</p>
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
          >sudo apt install build-essential libftdi1-dev libpopt-dev cmake pkg-config</code
        >
        <p>Now it is time to install the software to control the USBMux.</p>
        <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
          >git clone https://github.com/Badger-Embedded/muxer.git badgerd-muxer && cd badgerd-muxer</code
        >

        <p>Build and install the badgerd-muxer application by executing standard cmake build commands.</p>
        <!-- prettier-ignore -->
        <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">mkdir build && cd build
cmake ..
make
sudo make install</code>
        <p>Commands above will install `badgerd-muxer` into /usr/local/bin</p>

        <p id="demonstration">
          Since we have the tools we need, now we can continue on actually using the USBMux. For the purpose of
          demonstration, we will use the USBMux with two different USB storage connected to DUT and DEVICE ports. We
          also need a USB Type B cable to connect USBMux to our computer. (In this case I am using my
          <span class="text-orange-600">Ubuntu</span> host)
        </p>
        <ol class="list-decimal space-y-2 pl-4">
          <li>First connect 2 USB Storage devices to DUT and DEVICE ports</li>
          <li>Connect USB Type B cable to USBMux</li>
          <li>
            Let's check if our USBMux is already configured or not (we are sending them pre-configured but it is best to
            check)
            <ul class="mt-2 list-disc space-y-2 pl-8">
              <li>
                <p>In your computer's terminal run the following command:</p>
                <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">sudo dmesg -w</code>
              </li>
              <li>Connect your computer to USBMux using USB cable connected to HOST port of the USBMux</li>
              <li>
                Then you should be able to see 'idVendor', 'idProduct', 'SerialNumber'(alphanumeric) of the USBMux. We
                are going to use these values to configure if our USBMux is not configured yet. Let's check that!
              </li>
              <li>
                <p>In another terminal window, run the following command:</p>
                <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
                  >sudo badgerd-muxer --list</code
                >
                <p>If you see following output in terminal, this means that your USBMux is not configured.</p>
                <!-- prettier-ignore -->
                <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">Number of FTDI devices found: 0</code>
                <p>Let's configure it otherwise you can continue on step 4</p>
                <ul class="mt-2 list-[square] space-y-2 pl-12">
                  <li>
                    <p>Run the following command to configure USBMux:</p>
                    <!-- prettier-ignore -->
                    <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">badgerd-muxer --device-serial={SerialNumber} --vendor=0x{idVendor} --product=0x{idProduct} --set-serial=bdgrd_usbmux_101</code>
                    <p>Remember that all these values are coming from 'dmesg' output.</p>
                  </li>
                  <li>
                    <p>
                      After the configuration of the USBMux, you might need to re-connect the USBMux by plugging out and
                      in the USB cable. Now we should be able to see our USBMux is recognized by 'badgerd-muxer'
                    </p>
                    <!-- prettier-ignore -->
                    <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">badgerd-muxer --list
Number of FTDI devices found: 1
Dev: 0, Manufacturer: Badgerd Technologies, Serial: bdgrd_usbmux_101, Description: USBMUX</code>
                  </li>
                </ul>
              </li>
            </ul>
          </li>
          <li>
            <p>To connect the USB Storage Device on `DEVICE` port to your computer</p>
            <!-- prettier-ignore -->
            <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">badgerd-muxer --device-serial=bdgrd_usbmux_101 --device</code>
          </li>
          <li>
            Do whatever you want to do with this USB Device, you can change the content of it, you are free to do
            anything that USB2.0 is capable of.
          </li>
          <li>
            <p>To connect USB Storage Device on `DUT` port to your computer</p>
            <!-- prettier-ignore -->
            <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">badgerd-muxer --device-serial=bdgrd_usbmux_101 --dut</code>
          </li>
          <li>
            <p>
              This scenario will help you to switch between usb devices easily. But USBMux is not just about this. With
              the help of USBMux, you can also create connection between `DEVICE` and `DUT` ports.
            </p>
          </li>
          <li>
            <p>
              Let's say you have a USB Device that you want to test connected to `DUT` port of USBMux, you can connect
              the USB Device on `DEVICE` port of USBMux to `DUT` port by simply running the following command
            </p>
            <!-- prettier-ignore -->
            <code class=" block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4">badgerd-muxer --device-serial=bdgrd_usbmux_101 --dut-device</code>
          </li>
          <li>
            <p>
              This way you can simply change the content in USB Storage Device on `DEVICE` port of USBMux by connecting
              it to host as explained in step 4. And then connect it back to DUT.
            </p>
          </li>
          <li>
            <p>
              This is not all, there is a reason why we picked micro usb port for DUT port of USBMux. With the help of
              USBMux, you can manipulate the id value of the micro usb port on `DUT` so that you can test the behaviour
              of the `DUT` on different conditions
            </p>
          </li>
          <li>
            <p>To manupulate the id pin of micro usb port (DUT), you can run the following command</p>
            <!-- prettier-ignore -->
            <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
              >badgerd-muxer --device-serial=bdgrd_usbmux_101 --set-id-pin=1 # 1 to set it high, 0 to set it low</code
            >
          </li>
          <li>
            <p>To check the status of USBMux you can run the following command.</p>
            <!-- prettier-ignore -->
            <code class="block whitespace-pre-wrap border border-zinc-200 bg-[#fafafa] p-4"
              >$ badgerd-muxer --device-serial=bdgrd_usbmux_101 --status
DUT USB-ID State: 0 and DUT is connected to DEVICE [0:1]
</code>
          </li>
        </ol>

        <br />
        <p>Hope this quick start guide helps you to give our USBMux a kickstart. See you in another challenge!</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const products = (await $axios.get(`/api/stripe/products`)).data
    const product = products.find((p) => p.codename === 'badgerd_usbmux')
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
      title: 'Buy Badgerd USBMux | Quickstart With USBMux',
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
            'USBMux is a tool that helps you to switch connections between 3 USB ports. You can control the connection of the ports from a computer.',
        },
        {
          hid: 'og:title',
          property: 'og:title',
          content: 'Badgerd USBMux',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content:
            'USBMux is a tool that helps you to switch connections between 3 USB ports. You can control the connection of the ports from a computer.',
        },
        {
          hid: 'og:type',
          property: 'og:type',
          content: 'website',
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content: 'https://www.badgerd.nl/usbmux/badgerd-usbmux-og.png',
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
        '@id': 'https://badgerd.nl/#usbmux',
        '@type': 'Product',
        url: 'https://badgerd.nl/usbmux/',
        name: 'Badgerd USBMux | USBMux',
        mainEntityOfPage: 'https://badgerd.nl/usbmux/',
        image: 'https://www.badgerd.nl/usbmux/usbmux-1.jpg',
        description: 'USBMux is a USB switch that you can control the connection of 3 different USB Ports',
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
            datePublished: '2023-05-19',
            reviewBody: 'It does what you expect from a USBMux/Switch.',
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
            price: 180,
            sku: 'badgerd_usbmux',
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
