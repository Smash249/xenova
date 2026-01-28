<template>
  <header
    class="fixed inset-0 z-50 h-16 w-full transition-shadow duration-100 ease-linear"
    :class="{ 'shadow-sm': fadeProgress > 0.9 }"
    :style="{
      backgroundColor: `rgba(255, 255, 255, ${headerOpacity})`,
      backdropFilter: `blur(${(1 - fadeProgress) * 12}px)`,
      borderBottom: `1px solid rgba(229, 231, 235, ${headerOpacity})`,
    }"
  >
    <div
      class="container mx-auto flex h-full items-center justify-between px-4"
    >
      <!-- 左侧 Logo -->
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

      <!-- 中间菜单 -->
      <nav
        ref="navRef"
        class="absolute left-1/2 hidden -translate-x-1/2 items-center gap-2 rounded-full p-1 md:flex"
        :style="{
          backgroundColor: `rgba(243, 244, 246, ${0.3 + fadeProgress * 0.5})`,
          backdropFilter: `blur(${(1 - fadeProgress) * 8}px)`,
        }"
      >
        <div
          class="absolute h-10.5 rounded-full bg-[#4F61AD] transition-all duration-300 ease-out"
          :style="{
            left: `${activeSlider.left}px`,
            width: `${activeSlider.width}px`,
            top: '4px',
          }"
        ></div>

        <div
          v-for="(item, index) in systemConfig.navItems"
          :key="index"
          ref="menuItemsRef"
          class="relative z-10 cursor-pointer rounded-full px-6 py-2.5 text-sm font-medium transition-colors duration-300"
          @click="handleClickItem(item, index)"
          :style="{
            color: activeItem === index ? '#ffffff' : dynamicTextColor,
          }"
        >
          {{ item.title }}
        </div>
      </nav>

      <!-- 右侧用户头像 -->
      <div class="hidden items-center gap-3 md:flex">
        <div class="cursor-pointer" @click="toggeleDesktopDrawer">
          <img
            src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix"
            alt="用户头像"
            class="h-10 w-10 rounded-full border-2 border-gray-200 transition-transform duration-200 hover:scale-110"
          />
        </div>
      </div>
    </div>
  </header>
  <HeaderDrawer :open="showDesktopDrawer" />
</template>

<script setup>
import { useWindowScroll, useWindowSize } from "@vueuse/core"
import { computed, nextTick, onMounted, ref } from "vue"
import { useRouter } from "vue-router"

import { systemConfig } from "@/config/header"
import HeaderDrawer from "@/components/layout/HeaderDeskDrawer.vue"

const router = useRouter()
const { y: scrollY } = useWindowScroll()
const { height: windowHeight } = useWindowSize()
const activeItem = ref(0)

const navRef = ref(null)
const menuItemsRef = ref([])
const activeSlider = ref({
  left: 0,
  width: 0,
})
const showDesktopDrawer = ref(false)

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

// 更新滑块位置
function updateSliderPosition(element) {
  if (!element) return
  const rect = element.getBoundingClientRect()
  const navRect = navRef.value.getBoundingClientRect()
  activeSlider.value = {
    left: rect.left - navRect.left,
    width: rect.width,
  }
}

// 点击菜单项
function handleClickItem(item, index) {
  const element = menuItemsRef.value[index]
  updateSliderPosition(element)
  activeItem.value = index
  router.push(item.path)
}

// 处理抽屉菜单点击
function handleDrawerMenuClick(item, index) {
  activeItem.value = index
  router.push(item.path)
}

function toggeleDesktopDrawer() {
  showDesktopDrawer.value = !showDesktopDrawer.value
}

onMounted(() => {
  nextTick(() => {
    if (menuItemsRef.value && menuItemsRef.value.length > 0) {
      const firstItem = menuItemsRef.value[0]
      updateSliderPosition(firstItem)
    }
  })
  window.addEventListener("resize", () => {
    if (menuItemsRef.value && menuItemsRef.value[activeItem.value]) {
      updateSliderPosition(menuItemsRef.value[activeItem.value])
    }
  })
})
</script>
