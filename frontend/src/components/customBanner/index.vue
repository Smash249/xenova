<template>
  <div
    class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
  >
    <div
      class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
    ></div>

    <div class="relative z-20 px-4 text-center">
      <h1
        class="mb-4 text-4xl font-bold tracking-wider text-white drop-shadow-md md:text-5xl"
      >
        {{ title }}
      </h1>
    </div>
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  title: string
}>()
</script>

<style lang="scss" scoped></style>
