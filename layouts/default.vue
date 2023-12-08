<template>
  <div class="bg-zinc-100">
    <div
      class="flex flex-wrap items-center justify-start bg-zinc-700 text-sm font-normal text-zinc-200 subpixel-antialiased"
    >
      <div class="w-1/6" />
      <NuxtLink to="/" class="mx-4 p-1"
        ><Logo class="h-10 w-10 fill-zinc-200 transition-all hover:fill-zinc-50"
      /></NuxtLink>
      <div class="flex flex-grow flex-wrap justify-center">
        <NuxtLink
          to="/"
          class="mx-4 flex items-center justify-center p-2 transition-all hover:font-medium hover:text-zinc-50"
          >STORE</NuxtLink
        >
        <NuxtLink
          to="/sdwirec/"
          class="mx-4 flex items-center justify-center p-2 transition-all ease-in-out hover:font-medium hover:text-zinc-50"
          >SDWireC</NuxtLink
        >
        <NuxtLink
          to="/usbmux/"
          class="mx-4 flex items-center justify-center p-2 transition-all ease-in-out hover:font-medium hover:text-zinc-50"
          >USBMux</NuxtLink
        >
      </div>
      <div class="w-1/6" />
      <div style="christmas-message flex-basis: 100%; height: 0" v-show="new Date().getFullYear() < 2024"></div>
      <div class="christmas-message flex flex-grow flex-col" v-show="new Date().getFullYear() < 2024">
        <ul class="lightrope">
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
        </ul>
        <div class="min-h-[5rem] flex-grow items-center justify-center bg-red-700 font-bold text-zinc-300">
          <p class="mt-12 p-3 text-center">
            Merry Christmas! <br />
            We will be on Christmas break until the 01/01/2024. Please be aware that the orders you made will be shipped
            after 01/01/2024
          </p>
        </div>
      </div>
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

<style>
body {
  background: #000;
}

.lightrope {
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  position: absolute;
  z-index: 1;
  margin: -15px 0 0 0;
  padding: 0;
  pointer-events: none;
  width: 100%;
}
.lightrope li {
  position: relative;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation-iteration-count: infinite;
  animation-iteration-count: infinite;
  list-style: none;
  margin: 0;
  padding: 0;
  display: block;
  width: 12px;
  height: 28px;
  border-radius: 50%;
  margin: 20px;
  display: inline-block;
  background: #00f7a5;
  box-shadow: 0px 4.6666666667px 24px 3px #00f7a5;
  -webkit-animation-name: flash-1;
  animation-name: flash-1;
  -webkit-animation-duration: 2s;
  animation-duration: 2s;
}
.lightrope li:nth-child(2n + 1) {
  background: aqua;
  box-shadow: 0px 4.6666666667px 24px 3px rgba(0, 255, 255, 0.5);
  -webkit-animation-name: flash-2;
  animation-name: flash-2;
  -webkit-animation-duration: 0.4s;
  animation-duration: 0.4s;
}
.lightrope li:nth-child(4n + 2) {
  background: #f70094;
  box-shadow: 0px 4.6666666667px 24px 3px #f70094;
  -webkit-animation-name: flash-3;
  animation-name: flash-3;
  -webkit-animation-duration: 1.1s;
  animation-duration: 1.1s;
}
.lightrope li:nth-child(odd) {
  -webkit-animation-duration: 1.8s;
  animation-duration: 1.8s;
}
.lightrope li:nth-child(3n + 1) {
  -webkit-animation-duration: 1.4s;
  animation-duration: 1.4s;
}
.lightrope li:before {
  content: '';
  position: absolute;
  background: #222;
  width: 10px;
  height: 9.3333333333px;
  border-radius: 3px;
  top: -4.6666666667px;
  left: 1px;
}
.lightrope li:after {
  content: '';
  top: -14px;
  left: 9px;
  position: absolute;
  width: 52px;
  height: 18.6666666667px;
  border-bottom: solid #222 2px;
  border-radius: 50%;
}
.lightrope li:last-child:after {
  content: none;
}
.lightrope li:first-child {
  margin-left: -40px;
}

@-webkit-keyframes flash-1 {
  0%,
  100% {
    background: #00f7a5;
    box-shadow: 0px 4.6666666667px 24px 3px #00f7a5;
  }
  50% {
    background: rgba(0, 247, 165, 0.4);
    box-shadow: 0px 4.6666666667px 24px 3px rgba(0, 247, 165, 0.2);
  }
}

@keyframes flash-1 {
  0%,
  100% {
    background: #00f7a5;
    box-shadow: 0px 4.6666666667px 24px 3px #00f7a5;
  }
  50% {
    background: rgba(0, 247, 165, 0.4);
    box-shadow: 0px 4.6666666667px 24px 3px rgba(0, 247, 165, 0.2);
  }
}
@-webkit-keyframes flash-2 {
  0%,
  100% {
    background: aqua;
    box-shadow: 0px 4.6666666667px 24px 3px aqua;
  }
  50% {
    background: rgba(0, 255, 255, 0.4);
    box-shadow: 0px 4.6666666667px 24px 3px rgba(0, 255, 255, 0.2);
  }
}
@keyframes flash-2 {
  0%,
  100% {
    background: aqua;
    box-shadow: 0px 4.6666666667px 24px 3px aqua;
  }
  50% {
    background: rgba(0, 255, 255, 0.4);
    box-shadow: 0px 4.6666666667px 24px 3px rgba(0, 255, 255, 0.2);
  }
}
@-webkit-keyframes flash-3 {
  0%,
  100% {
    background: #f70094;
    box-shadow: 0px 4.6666666667px 24px 3px #f70094;
  }
  50% {
    background: rgba(247, 0, 148, 0.4);
    box-shadow: 0px 4.6666666667px 24px 3px rgba(247, 0, 148, 0.2);
  }
}
@keyframes flash-3 {
  0%,
  100% {
    background: #f70094;
    box-shadow: 0px 4.6666666667px 24px 3px #f70094;
  }
  50% {
    background: rgba(247, 0, 148, 0.4);
    box-shadow: 0px 4.6666666667px 24px 3px rgba(247, 0, 148, 0.2);
  }
}
</style>
