<script setup lang="ts">
import CustomFooter from "./footer/index.vue"
import CustomHeader from "./header/index.vue"
</script>

<template>
  <div class="flex-col-start h-full w-full">
    <CustomHeader />
    <main class="flex-col-center min-h-[calc(100vh-4rem)] w-full px-4 py-6">
      <slot />
    </main>
    <CustomFooter />
  </div>
</template>

<style scoped lang="scss"></style>
