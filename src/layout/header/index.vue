<template>
  <header
    class="fixed inset-0 z-50 h-16 w-full transition-shadow duration-100 ease-linear"
    :class="{ 'shadow-sm': fadeProgress > 0.9 }"
    :style="{
      backgroundColor: `rgba(255, 255, 255, ${headerOpacity})`,
      backdropFilter: `blur(${headerOpacity * 12}px)`,
      borderBottom: `1px solid rgba(229, 231, 235, ${headerOpacity})`,
    }"
  >
    <div
      class="container mx-auto flex h-full items-center justify-between px-4"
    >
      <div class="flex items-center gap-3">
        <div class="relative h-10 w-10 shrink-0">
          <svg viewBox="0 0 100 100" class="h-full w-full" fill="none">
            <path
              d="M20 10 L45 50 L20 90 L35 90 L60 50 L35 10 Z"
              fill="#EAB308"
            />
            <path
              d="M80 10 L55 50 L80 90 L65 90 L40 50 L65 10 Z"
              fill="#52525B"
            />
          </svg>
        </div>
        <div class="flex flex-col">
          <h1 class="text-2xl font-bold tracking-wide text-orange-500">
            星实科技
          </h1>
          <div
            class="hidden origin-left scale-90 items-center gap-2 text-xs transition-colors md:flex"
            :style="{ color: dynamicTextColor }"
          >
            <span>创新科技</span>
            <span>星耀未来</span>
          </div>
        </div>
      </div>

      <nav class="hidden items-center space-x-1 md:flex">
        <div
          v-for="(item, index) in systemConfig.navItems"
          :key="index"
          class="cursor-pointer rounded-full px-5 py-2 text-base font-bold transition-all duration-300"
          @click="handelClickItem(item, index)"
          :style="getItemStyle(index)"
        >
          {{ item.title }}
        </div>
      </nav>

      <button class="p-2 md:hidden" :style="{ color: dynamicTextColor }">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h16"
          />
        </svg>
      </button>
    </div>
  </header>
</template>

<script setup>
import { useWindowScroll, useWindowSize } from "@vueuse/core"
import { computed, ref } from "vue"
import { useRouter } from "vue-router"

import { systemConfig } from "@/config/header"

const router = useRouter()
const { y: scrollY } = useWindowScroll()
const { height: windowHeight } = useWindowSize()
const activeItem = ref(0)

const fadeProgress = computed(() => {
  if (windowHeight.value === 0) return 0
  const start = windowHeight.value * 0.5
  const end = windowHeight.value
  return Math.min(Math.max((scrollY.value - start) / (end - start), 0), 1)
})
const headerOpacity = computed(() => fadeProgress.value * 0.95)
const dynamicTextColor = computed(() => {
  const p = fadeProgress.value
  const v = Math.round(255 - 200 * p)
  return `rgb(${v}, ${v}, ${v})`
})

function getItemStyle(index) {
  const isActive = activeItem.value === index
  const p = fadeProgress.value
  if (!isActive) {
    return {
      color: dynamicTextColor.value,
      backgroundColor: "transparent",
    }
  }
  return {
    color: "#d97706",
  }
}

function handelClickItem(item, index) {
  activeItem.value = index
  router.push(item.path)
}
</script>
