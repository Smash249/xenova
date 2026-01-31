<template>
  <section class="relative h-screen w-full overflow-hidden">
    <div class="absolute inset-0 z-0">
      <img
        src="@/assets/images/banner.webp"
        alt="Industrial Background"
        class="h-full w-full object-cover"
      />

      <div
        class="absolute inset-0 bg-linear-to-b from-blue-900/40 via-blue-900/60 to-blue-900/80 mix-blend-multiply"
      ></div>
    </div>
    <div
      class="relative z-10 container mx-auto flex h-full flex-col items-center justify-center px-4 pt-20 text-center text-white"
    >
      <div
        v-motion
        :initial="{ opacity: 0, y: 50 }"
        :enter="{
          opacity: 1,
          y: 0,
          transition: { duration: 800, type: 'spring' },
        }"
      >
        <h1
          class="mb-6 text-4xl font-bold tracking-wider drop-shadow-lg md:text-6xl"
        >
          星实(厦门)科技有限公司
        </h1>
      </div>

      <div
        v-motion
        :initial="{ opacity: 0, y: 50 }"
        :enter="{
          opacity: 1,
          y: 0,
          transition: { duration: 800, delay: 200, type: 'spring' },
        }"
      >
        <p
          class="mb-12 text-xl font-light tracking-wide text-blue-100 md:text-2xl"
        >
          非标自动化生产制造 · 专业企业咨询
        </p>
      </div>

      <div
        v-motion
        :initial="{ opacity: 0, y: 30 }"
        :enter="{ opacity: 1, y: 0, transition: { duration: 800, delay: 400 } }"
        class="mx-auto mb-16 max-w-4xl"
      >
        <p
          class="text-justify text-base leading-loose text-blue-50/90 md:text-center md:text-lg"
        >
          星实是一家年轻的自动化公司，公司成立于2021年，经过5年的摸爬滚打，走出了一条适合自己的发展道路，我们专注于包装机、激光机、医疗、食品等领域的自动化专业生产制造设备，同时为大型企业、矿山、海洋等领域提供科学技术咨询方案。
        </p>
      </div>
    </div>

    <div
      class="group absolute bottom-12 left-1/2 z-20 -translate-x-1/2 cursor-pointer"
      @click="scrollToContent"
    >
      <div
        v-motion
        :initial="{ opacity: 0, y: -20 }"
        :enter="{
          opacity: 1,
          y: [0, 10, 0],
          scale: [1, 1.1, 1],
          transition: {
            opacity: { duration: 800, delay: 1000 },
            y: { repeat: Infinity, duration: 1500, ease: 'easeInOut' },
            scale: { repeat: Infinity, duration: 1500, ease: 'easeInOut' },
          },
        }"
        class="flex flex-col items-center gap-2 opacity-80 transition-opacity hover:opacity-100"
      >
        <span
          class="text-[10px] font-bold tracking-[0.3em] text-blue-200 uppercase"
          >Scroll Down</span
        >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="h-10 w-10 text-white drop-shadow-md"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M19.5 8.25l-7.5 7.5-7.5-7.5"
          />
        </svg>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
function scrollToContent() {
  window.scrollTo({
    top: window.innerHeight,
    behavior: "smooth",
  })
}
</script>

<style scoped lang="scss"></style>
