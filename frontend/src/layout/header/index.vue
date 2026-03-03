<template>
  <header
    class="fixed inset-0 z-50 h-16 w-full border-b border-gray-200 bg-white py-4 shadow-sm"
  >
    <div class="mx-auto flex h-full w-[90%] items-center justify-between">
      <img src="/images/logo.png" alt="logo" class="h-auto w-28 sm:w-48" />

      <nav
        ref="navRef"
        class="relative hidden items-center gap-2 rounded-full bg-gray-100 p-1 md:flex"
      >
        <div
          v-show="isHovering"
          :style="{
            left: `${hoverSlider.left}px`,
            width: `${hoverSlider.width}px`,
            top: '4px',
          }"
          class="absolute z-0 h-10.5 rounded-full bg-blue-600/30 transition-all duration-300 ease-out"
        ></div>

        <div
          v-if="activeItem !== -1"
          :style="{
            left: `${activeSlider.left}px`,
            width: `${activeSlider.width}px`,
            top: '4px',
          }"
          class="absolute z-10 h-10.5 rounded-full bg-blue-600 transition-all duration-300 ease-out"
        ></div>

        <div
          v-for="(item, index) in systemConfig.navItems"
          :key="index"
          ref="menuItemsRef"
          :style="{
            color: activeItem === index ? '#ffffff' : '#333333',
          }"
          class="relative z-20 cursor-pointer rounded-full px-6 py-3 text-xs font-medium transition-colors duration-300"
          @mouseenter="HandleMouseEnter(index)"
          @mouseleave="HandleMouseLeave"
          @click="HandleClickItem(item, index)"
        >
          {{ item.title }}
        </div>
      </nav>

      <div class="flex items-center gap-3">
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

        <div class="flex cursor-pointer md:hidden" @click="ToggleMobileMenu">
          <el-icon :size="28" color="#333333">
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
import { computed, nextTick, onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"

import { systemConfig } from "@/config/header"
import DeskDrawer from "@/components/layout/DeskDrawer.vue"
import MobileDrawer from "@/components/layout/MobileDrawer.vue"

const router = useRouter()
const route = useRoute()

const user = userStore()

const activeItem = ref(-1)
const navRef = ref(null)
const menuItemsRef = ref([])

const activeSlider = ref({ left: 0, width: 0 })
const hoverSlider = ref({ left: 0, width: 0 })
const isHovering = ref(false)
let hoverTimer = null

const drawerVisible = ref(false)
const mobileMenuVisible = ref(false)

const avatarLetter = computed(() => {
  return user.email ? user.email.charAt(0).toUpperCase() : ""
})

function UpdateSliderPosition(element, isHover = false) {
  if (!element || !navRef.value) return
  const rect = element.getBoundingClientRect()
  const navRect = navRef.value.getBoundingClientRect()
  const positionData = {
    left: rect.left - navRect.left,
    width: rect.width,
  }
  if (isHover) {
    hoverSlider.value = positionData
  } else {
    activeSlider.value = positionData
  }
}

function HandleMouseEnter(index) {
  const element = menuItemsRef.value[index]
  if (!element) return

  isHovering.value = true
  UpdateSliderPosition(element, true)

  if (hoverTimer) {
    clearTimeout(hoverTimer)
  }

  hoverTimer = setTimeout(() => {
    activeItem.value = index
    UpdateSliderPosition(element, false)
    isHovering.value = false
    hoverTimer = null
  }, 500)
}

function HandleMouseLeave() {
  if (hoverTimer) {
    clearTimeout(hoverTimer)
    hoverTimer = null
  }
  isHovering.value = false
}

function HandleClickItem(item, index) {
  if (hoverTimer) {
    clearTimeout(hoverTimer)
    hoverTimer = null
  }

  activeItem.value = index
  if (menuItemsRef.value[index]) {
    UpdateSliderPosition(menuItemsRef.value[index], false)
  }
  isHovering.value = false

  if (item.isNav) {
    router.push(item.path)
  } else {
    router.push(item.path || "/about")
  }
}

function SyncActiveItemWithRoute() {
  nextTick(() => {
    const currentPath = route.path
    const currentHash = route.hash

    const index = systemConfig.navItems.findIndex((item) => {
      if (item.path === currentPath) return true
      if (item.title === "关于我们")
        return currentPath === "/about" || currentHash === "#about"
      return false
    })

    if (index !== -1) {
      activeItem.value = index
      if (menuItemsRef.value[index]) {
        UpdateSliderPosition(menuItemsRef.value[index], false)
      }
    } else {
      activeItem.value = -1
    }
  })
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
      UpdateSliderPosition(menuItemsRef.value[activeItem.value], false)
    }
  })
})
</script>
