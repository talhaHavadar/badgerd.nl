<template>
  <div class="bg-zinc-100">
    <div class="flex items-center justify-start bg-zinc-700 p-1 text-sm font-normal text-zinc-200 subpixel-antialiased">
      <div class="w-1/6" />
      <NuxtLink to="/" class="mx-4"
        ><Logo class="h-10 w-10 fill-zinc-200 transition-all hover:fill-zinc-50"
      /></NuxtLink>
      <div class="flex flex-grow flex-wrap justify-center">
        <NuxtLink
          to="/"
          class="mx-4 flex items-center justify-center p-1 transition-all hover:font-medium hover:text-zinc-50"
          >STORE</NuxtLink
        >
        <NuxtLink
          to="/sdwirec/"
          class="mx-4 flex items-center justify-center p-1 transition-all ease-in-out hover:font-medium hover:text-zinc-50"
          >SDWireC</NuxtLink
        >
        <NuxtLink
          to="/usbmux/"
          class="mx-4 flex items-center justify-center p-1 transition-all ease-in-out hover:font-medium hover:text-zinc-50"
          >USBMux</NuxtLink
        >
      </div>
      <div class="w-1/6" />
    </div>
    <Nuxt />
    <div
      class="flex min-h-[20vh] w-full flex-row flex-wrap justify-between border border-t-gray-300 bg-gray-100 p-8 text-zinc-600 subpixel-antialiased"
    >
      <div class="flex flex-col items-start justify-end font-thin subpixel-antialiased">
        <p>KvK# 87317427</p>
        <p>Copyright © 2023 Badgerd Technologies B.V. All rights reserved.</p>
      </div>
      <div class="flex flex-grow flex-col items-end">
        <div class="flex flex-grow items-end">
          <a href="https://github.com/Badger-Embedded" rel="nofollow,external" target="_blank"
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
        </div>
        <div class="flex items-end">
          <a href="mailto:info@badgerd.nl" rel="nofollow" class="p-1 hover:text-zinc-800">Contact</a>
          <NuxtLink to="/refund/" class="p-1 hover:text-zinc-800">Refund</NuxtLink>
          <NuxtLink to="/terms/" class="p-1 hover:text-zinc-800"
            >Terms <span class="hidden sm:inline">of Service</span></NuxtLink
          >
          <NuxtLink to="/privacy/" class="p-1 hover:text-zinc-800">Privacy</NuxtLink>
          <NuxtLink to="/legal/" class="p-1 hover:text-zinc-800">Legal Notice</NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>
