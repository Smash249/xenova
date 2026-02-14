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
    <div class="mx-auto flex h-full w-[90%] items-center justify-between">
      <div class="flex items-center">
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
          <h1 class="text-2xl font-bold tracking-wide text-orange-400">
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
        class="hidden items-center gap-2 rounded-full p-1 md:flex"
      >
        <div
          v-if="activeItem !== -1"
          :style="{
            left: `${activeSlider.left}px`,
            width: `${activeSlider.width}px`,
            top: '4px',
          }"
          class="absolute h-10.5 rounded-full bg-blue-600 transition-all duration-300 ease-out"
        ></div>

        <div
          v-for="(item, index) in systemConfig.navItems"
          :key="index"
          ref="menuItemsRef"
          :style="{
            color: activeItem === index ? '#ffffff' : dynamicTextColor,
          }"
          class="relative z-10 cursor-pointer rounded-full px-6 py-3 text-xs font-medium transition-colors duration-300"
          @click="HandleClickItem(item, index)"
        >
          {{ item.title }}
        </div>
      </nav>

      <div class="hidden items-center gap-3 md:flex">
        <div
          v-if="user.isLogin"
          class="cursor-pointer transition-transform hover:scale-105 active:scale-95"
          @click="ToggleDesktopDrawer"
        >
          <el-avatar class="bg-blue-100 font-bold text-blue-600">
            {{ avatarLetter }}
          </el-avatar>
        </div>
        <el-button
          v-else
          type="primary"
          round
          class="border-0! bg-blue-600! px-6! py-5! text-sm font-medium hover:bg-blue-500!"
          @click="router.push('/login')"
          :icon="Position"
        >
          登录
        </el-button>
      </div>
    </div>
  </header>
  <DeskDrawer v-model="drawerVisible" />
</template>

<script setup>
import userStore from "@/store/modules/user"
import { Position } from "@element-plus/icons-vue"
import { useWindowScroll, useWindowSize } from "@vueuse/core"
import { computed, nextTick, onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"

import { systemConfig } from "@/config/header"
import DeskDrawer from "@/components/layout/DeskDrawer.vue"

const router = useRouter()
const route = useRoute()
const { y: scrollY } = useWindowScroll()
const { height: windowHeight } = useWindowSize()

const user = userStore()

const activeItem = ref(-1)
const navRef = ref(null)
const menuItemsRef = ref([])
const activeSlider = ref({ left: 0, width: 0 })
const drawerVisible = ref(false)

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

const avatarLetter = computed(() => {
  return user.email ? user.email.charAt(0).toUpperCase() : ""
})

function UpdateSliderPosition(element) {
  if (!element || !navRef.value) return
  const rect = element.getBoundingClientRect()
  const navRect = navRef.value.getBoundingClientRect()
  activeSlider.value = {
    left: rect.left - navRect.left,
    width: rect.width,
  }
}

function SyncActiveItemWithRoute() {
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
        UpdateSliderPosition(menuItemsRef.value[index])
      }
    } else {
      activeItem.value = -1
    }
  })
}

function HandleClickItem(item, index) {
  if (item.isNav) {
    router.push(item.path)
  } else {
    router.push("/dashboard#about")
  }
}

function ToggleDesktopDrawer() {
  drawerVisible.value = true
}

watch(() => route.fullPath, SyncActiveItemWithRoute, { immediate: true })

onMounted(() => {
  window.addEventListener("resize", () => {
    if (activeItem.value !== -1 && menuItemsRef.value[activeItem.value]) {
      UpdateSliderPosition(menuItemsRef.value[activeItem.value])
    }
  })
})
</script>
