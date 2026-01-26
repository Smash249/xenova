<template>
  <div class="flex min-h-screen w-full flex-col gap-2">
    <CustomHeader />
    <main class="flex w-full flex-1 flex-col">
      <router-view />
    </main>
    <CustomFooter />
  </div>
</template>

<script setup lang="ts">
import CustomFooter from "./footer/index.vue"
import CustomHeader from "./header/index.vue"
</script>

<style scoped lang="scss"></style>
