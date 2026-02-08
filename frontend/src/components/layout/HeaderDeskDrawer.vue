<template>
  <el-drawer
    :model-value="open"
    direction="rtl"
    size="320px"
    :show-close="true"
    @close="handleClose"
    append-to-body
    lock-scroll
    destroy-on-close
  >
    <!-- 抽屉内容 -->
    <div class="flex h-full flex-col"></div>
  </el-drawer>
</template>

<script setup lang="ts">
import { ElDrawer } from "element-plus"

const props = withDefaults(
  defineProps<{
    open: boolean
  }>(),
  {
    open: false,
  }
)

const emit = defineEmits<{
  close: []
}>()

function handleClose() {
  emit("close")
}
</script>

<style scoped></style>
