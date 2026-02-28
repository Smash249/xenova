<template>
  <header
    :class="{ 'shadow-sm': fadeProgress > 0.9 }"
    :style="{
      backgroundColor: `rgba(255, 255, 255, ${headerOpacity})`,
      backdropFilter: `blur(${(1 - fadeProgress) * 12}px)`,
      borderBottom: `1px solid rgba(229, 231, 235, ${headerOpacity})`,
    }"
    class="fixed inset-0 z-50 h-16 w-full py-4 transition-shadow duration-100 ease-linear"
  >
    <div class="mx-auto flex w-[90%] items-center justify-between">
      <img src="/images/logo.png" alt="logo" class="h-auto w-28 sm:w-48" />

      <!-- 桌面端导航 -->
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

      <div class="flex items-center gap-3">
        <!-- 桌面端操作区 -->
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

        <!-- 移动端菜单按钮 -->
        <div class="flex cursor-pointer md:hidden" @click="ToggleMobileMenu">
          <el-icon :size="28" :style="{ color: dynamicTextColor }">
            <Menu />
          </el-icon>
        </div>
      </div>
    </div>
  </header>

  <DeskDrawer v-model="drawerVisible" />

  <MobileDrawer
    v-model="mobileMenuVisible"
    :active-item="activeItem"
    @item-click="HandleClickItem"
    @login-click="router.push('/login')"
    @user-click="ToggleDesktopDrawer"
  />
</template>

<script setup>
import userStore from "@/store/modules/user"
import { Menu, Position } from "@element-plus/icons-vue"
import { useWindowScroll, useWindowSize } from "@vueuse/core"
import { computed, nextTick, onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"

import { systemConfig } from "@/config/header"
import DeskDrawer from "@/components/layout/DeskDrawer.vue"
import MobileDrawer from "@/components/layout/MobileDrawer.vue"

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
const mobileMenuVisible = ref(false)

const fadeProgress = computed(() => {
  if (windowHeight.value === 0) return 0
  const start = windowHeight.value * 0.3
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

function ToggleMobileMenu() {
  mobileMenuVisible.value = true
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
