<script setup lang="tsx">
import { NAvatar } from 'naive-ui'
import { computed } from 'vue'

import { useUserStore } from '@/stores'

const userStore = useUserStore()

const userName = computed(() => userStore.user?.userName ?? '匿名用户')
</script>
<template>
  <div class="aspect-square size-full">
    <n-avatar
      round
      class="flex size-full items-center justify-center bg-primary font-bold text-white"
    >
      {{ userName.charAt(0).toUpperCase() }}
    </n-avatar>
  </div>
</template>
