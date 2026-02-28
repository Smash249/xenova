<template>
  <el-drawer
    :model-value="modelValue"
    @update:model-value="$emit('update:modelValue', $event)"
    direction="rtl"
    size="70%"
    :with-header="false"
  >
    <div class="flex flex-col px-4 py-8">
      <div class="mb-8 flex items-center justify-between">
        <img src="/images/logo.png" alt="logo" class="h-auto w-24" />
      </div>
      <div class="flex flex-col gap-6">
        <div
          v-for="(item, index) in systemConfig.navItems"
          :key="index"
          class="cursor-pointer text-lg font-medium transition-colors"
          :class="activeItem === index ? 'text-blue-600' : 'text-gray-800'"
          @click="HandleClickMobileItem(item, index)"
        >
          {{ item.title }}
        </div>
      </div>
      <div class="mt-10 border-t border-gray-200 pt-8">
        <div
          v-if="user.isLogin"
          class="flex cursor-pointer items-center gap-3"
          @click="HandleMobileUserClick"
        >
          <el-avatar class="bg-blue-100 font-bold text-blue-600">
            {{ avatarLetter }}
          </el-avatar>
          <span class="text-base font-medium text-gray-800">个人中心</span>
        </div>
        <el-button
          v-else
          type="primary"
          round
          class="w-full border-0! bg-blue-600! py-5! text-sm font-medium hover:bg-blue-500!"
          @click="HandleMobileLogin"
          :icon="Position"
        >
          登录
        </el-button>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import userStore from "@/store/modules/user"
import { Position } from "@element-plus/icons-vue"
import { computed } from "vue"

import { systemConfig } from "@/config/header"

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
  activeItem: {
    type: Number,
    default: -1,
  },
})

const emit = defineEmits([
  "update:modelValue",
  "item-click",
  "login-click",
  "user-click",
])

const user = userStore()

const avatarLetter = computed(() => {
  return user.email ? user.email.charAt(0).toUpperCase() : ""
})

function HandleClickMobileItem(item, index) {
  emit("update:modelValue", false)
  emit("item-click", item, index)
}

function HandleMobileLogin() {
  emit("update:modelValue", false)
  emit("login-click")
}

function HandleMobileUserClick() {
  emit("update:modelValue", false)
  emit("user-click")
}
</script>
