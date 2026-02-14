<template>
  <el-drawer v-model="visible" direction="rtl" size="300px">
    <div class="flex h-full flex-col">
      <div class="mb-6 flex items-center gap-4 border-b border-gray-100 pb-6">
        <el-avatar
          :size="50"
          class="bg-blue-100 text-xl font-bold text-blue-600"
        >
          {{ avatarLetter }}
        </el-avatar>
        <div class="flex flex-col">
          <span class="text-lg font-bold text-gray-800">{{
            user.userName || "用户"
          }}</span>
          <span class="text-sm text-gray-500">{{ user.email }}</span>
        </div>
      </div>

      <div class="flex-1 space-y-2">
        <div
          v-for="(item, index) in menuItems"
          :key="index"
          class="flex cursor-pointer items-center rounded-lg py-2 text-gray-600 transition-colors hover:bg-gray-50 hover:text-blue-600"
          @click="HandleMenuClick(item)"
        >
          <i :class="item.icon" class="mr-3 text-lg"></i>
          <span class="font-medium">{{ item.title }}</span>
        </div>
      </div>

      <div class="border-t border-gray-100 pt-4">
        <button
          class="flex w-full items-center justify-center rounded-lg bg-red-50 px-4 py-3 text-red-600 transition-colors hover:bg-red-100"
          @click="HandleLogout"
        >
          <span class="font-bold">退出登录</span>
        </button>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import userStore from "@/store/modules/user"
import { computed } from "vue"

const menuItems = [
  { title: "个人中心", icon: "el-icon-user", action: "profile" },
  { title: "设置", icon: "el-icon-setting", action: "settings" },
]

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(["update:modelValue"])
const user = userStore()

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
})

const avatarLetter = computed(() => {
  return user.email ? user.email.charAt(0).toUpperCase() : ""
})

function HandleMenuClick(item) {
  visible.value = false
}

function HandleLogout() {
  visible.value = false
  user.Logout()
}
</script>

<style scoped>
:deep(.el-drawer__body) {
  padding: 0;
}
</style>
