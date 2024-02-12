<template>
  <div class="items-top relative flex min-h-[calc(100vh-3rem)] justify-center bg-gray-100 sm:items-center sm:pt-0">
    <div class="mx-auto sm:px-6 lg:px-8">
      <div class="container px-6">
        <!-- Section: Design Block -->
        <section class="mb-32 text-center text-gray-800 lg:text-left">
          <style>
            @media (min-width: 992px) {
              .rotate-lg-6 {
                transform: rotate(6deg);
              }
            }
          </style>

          <div class="py-12 md:px-12">
            <div class="container mx-auto xl:px-32">
              <div class="grid min-h-[calc(100vh-3rem)] items-center lg:grid-cols-2">
                <div class="mb-12 md:mt-12 lg:mt-0 lg:mb-0">
                  <div
                    class="relative block rounded-lg px-6 pt-12 pb-6 shadow-lg md:px-12 lg:-mr-14"
                    style="background: hsla(0, 0%, 100%, 0.55); backdrop-filter: blur(30px); z-index: 1"
                  >
                    <div class="mb-6 flex items-center">
                      <h1 class="text-3xl font-bold">Badgerd nRFNANO</h1>
                      <NuxtLink
                        to="/nrfnano/"
                        class="ml-4 text-2xl font-bold text-blue-500 transition-all ease-in-out after:content-[attr(after)] hover:text-[1.6rem] hover:text-blue-600"
                        >Buy ></NuxtLink
                      >
                    </div>
                    <p class="mb-6 text-gray-500">
                      The Badger'd nRFNANO, is a tiny BLE/IoT development board that enables users to work with nRF52840
                      ic, superfast with the help of UF2 bootloader.
                    </p>

                    <div class="grid gap-x-6 md:grid-cols-3 lg:grid-cols-2 xl:grid-cols-3">
                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >nice!nano compatible
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >ProMicro Form Factor
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >UF2 Bootloader
                        </p>
                      </div>
                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >Built-in Charger
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >1MB Flash
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >TC Debugging
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="md:mb-12 lg:mb-0">
                  <img
                    :src="require('~/assets/nrfnano/nrfnano_0.jpg')"
                    class="rotate-lg-6 w-full rounded-lg shadow-lg"
                    alt="nRFNANO"
                    title="nRFNANO"
                    loading="eager"
                  />
                  <!-- <img
                src="https://mdbootstrap.com/img/new/ecommerce/vertical/121.jpg"
                class="rotate-lg-6 w-full rounded-lg shadow-lg"
                alt=""
              /> -->
                </div>
              </div>
              <div class="mt-12 grid items-center lg:grid-cols-2">
                <div class="mb-12 md:mt-12 lg:mt-0 lg:mb-0">
                  <div
                    class="relative block rounded-lg px-6 pt-12 pb-6 shadow-lg md:px-12 lg:-mr-14"
                    style="background: hsla(0, 0%, 100%, 0.55); backdrop-filter: blur(30px); z-index: 1"
                  >
                    <div class="mb-6 flex items-center">
                      <h1 class="text-3xl font-bold">Badgerd SDWireC</h1>
                      <NuxtLink
                        to="/sdwirec/"
                        class="ml-4 text-2xl font-bold text-blue-500 transition-all ease-in-out after:content-[attr(after)] hover:text-[1.6rem] hover:text-blue-600"
                        >Buy ></NuxtLink
                      >
                    </div>
                    <p class="mb-6 text-gray-500">
                      The Badger'd way of the infamous SDWire. In our version of SDWire, we used USB-C interface to
                      communicate with host.
                    </p>

                    <div class="grid gap-x-6 md:grid-cols-3 lg:grid-cols-2 xl:grid-cols-3">
                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >Hardware Lab
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >Flash SDCard
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >Modern Interface
                        </p>
                      </div>
                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >SDWire v1.4
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >3D Printing
                        </p>
                      </div>

                      <div class="mb-6">
                        <p class="flex items-center">
                          <svg
                            class="mr-2 h-4 w-4 text-gray-900"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 512 512"
                          >
                            <path
                              fill="currentColor"
                              d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"
                            ></path></svg
                          >No Wiring
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="md:mb-12 lg:mb-0">
                  <img
                    :src="require('~/assets/sdwirec/sdwirec-1.jpg')"
                    class="rotate-lg-6 w-full rounded-lg shadow-lg"
                    alt="SDWireC"
                    title="SDWireC"
                    loading="eager"
                  />
                  <!-- <img
                src="https://mdbootstrap.com/img/new/ecommerce/vertical/121.jpg"
                class="rotate-lg-6 w-full rounded-lg shadow-lg"
                alt=""
              /> -->
                </div>
              </div>
            </div>
          </div>
        </section>
        <!-- Section: Design Block -->
      </div>
    </div>
  </div>
  <!-- <div class="items-top relative flex min-h-[calc(100vh-3rem)] justify-center bg-gray-100 sm:items-center sm:pt-0">
    <div class="mx-auto max-w-4xl sm:px-6 lg:px-8">
      <div class="flex justify-center space-x-2 pt-4">


        <img :src="require('~/assets/sdwirec/badgerd-sdwirec-cad.png')" class="w-full" alt="KiCAD" />
      </div>
    </div>
  </div> -->
</template>

<script>
export default {
  name: 'IndexPage',

  async asyncData({ $content }) {
    const page = await $content('hello').fetch()
    return {
      page,
    }
  },
  head() {
    return {
      title: 'Sett | Badgerd Technologies',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content:
            'Discover our innovative solutions for interesting problems in our life. Buy one of our products to support us!',
        },
        {
          hid: 'og:title',
          property: 'og:title',
          content: 'Badgerd Technologies',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content:
            'Discover our innovative solutions for interesting problems in our life. Buy one of our products to support us!',
        },
        {
          hid: 'og:type',
          property: 'og:type',
          content: 'website',
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content: '',
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
        '@id': 'https://badgerd.nl/#webpage',
        '@type': 'WebPage',
        url: 'https://badgerd.nl/',
        name: 'Badgerd Technologies',
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
