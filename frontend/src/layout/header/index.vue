<template>
  <header
    :class="{ 'shadow-sm': fadeProgress > 0.9 }"
    :style="{
      backgroundColor: `rgba(255, 255, 255, ${headerOpacity})`,
      backdropFilter: `blur(${(1 - fadeProgress) * 12}px)`,
      borderBottom: `1px solid rgba(229, 231, 235, ${headerOpacity})`,
    }"
    class="fixed inset-0 z-50 h-16 w-full transition-shadow duration-100 ease-linear"
  >
    <div
      class="container mx-auto flex h-full items-center justify-between px-4"
    >
      <div class="flex items-center gap-3">
        <div class="relative h-10 w-10 shrink-0">
          <svg class="h-full w-full" fill="none" viewBox="0 0 100 100">
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
        </div>
      </div>

      <nav
        ref="navRef"
        :style="{
          backgroundColor: `rgba(243, 244, 246, ${0.3 + fadeProgress * 0.5})`,
          backdropFilter: `blur(${(1 - fadeProgress) * 8}px)`,
        }"
        class="absolute left-1/2 hidden -translate-x-1/2 items-center gap-2 rounded-full p-1 md:flex"
      >
        <div
          v-if="activeItem !== -1"
          :style="{
            left: `${activeSlider.left}px`,
            width: `${activeSlider.width}px`,
            top: '4px',
          }"
          class="absolute h-10.5 rounded-full bg-[#4F61AD] transition-all duration-300 ease-out"
        ></div>

        <div
          v-for="(item, index) in systemConfig.navItems"
          :key="index"
          ref="menuItemsRef"
          :style="{
            color: activeItem === index ? '#ffffff' : dynamicTextColor,
          }"
          class="relative z-10 cursor-pointer rounded-full px-6 py-3 text-xs font-medium transition-colors duration-300"
          @click="handleClickItem(item, index)"
        >
          {{ item.title }}
        </div>
      </nav>

      <div class="hidden items-center gap-3 md:flex">
        <div class="cursor-pointer" @click="toggleDesktopDrawer">
          <img
            alt="用户头像"
            class="h-10 w-10 rounded-full border-2 border-gray-200 transition-transform duration-200 hover:scale-110"
            src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix"
          />
        </div>
      </div>
    </div>
  </header>
</template>

<script setup>
import { useWindowScroll, useWindowSize } from "@vueuse/core"
import { computed, nextTick, onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"

import { systemConfig } from "@/config/header"

const router = useRouter()
const route = useRoute()
const { y: scrollY } = useWindowScroll()
const { height: windowHeight } = useWindowSize()

const activeItem = ref(-1)
const navRef = ref(null)
const menuItemsRef = ref([])
const activeSlider = ref({ left: 0, width: 0 })

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

function updateSliderPosition(element) {
  if (!element || !navRef.value) return
  const rect = element.getBoundingClientRect()
  const navRect = navRef.value.getBoundingClientRect()
  activeSlider.value = {
    left: rect.left - navRect.left,
    width: rect.width,
  }
}

function updateActiveByPath() {
  nextTick(() => {
    const currentPath = route.path
    const currentHash = route.hash

    const index = systemConfig.navItems.findIndex((item) => {
      if (item.title === "关于我们") {
        return currentHash === "#about"
      }
      return item.path === currentPath && !currentHash
    })

    if (index !== -1) {
      activeItem.value = index
      if (menuItemsRef.value[index]) {
        updateSliderPosition(menuItemsRef.value[index])
      }
    } else {
      activeItem.value = -1
    }
  })
}

function handleClickItem(item, index) {
  if (item.isNav) {
    router.push(item.path)
  } else {
    router.push("/dashboard#about")
  }
}

watch(
  () => route.fullPath,
  () => {
    updateActiveByPath()
  },
  { immediate: true }
)

onMounted(() => {
  window.addEventListener("resize", () => {
    if (activeItem.value !== -1 && menuItemsRef.value[activeItem.value]) {
      updateSliderPosition(menuItemsRef.value[activeItem.value])
    }
  })
})
</script>
