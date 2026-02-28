<template>
  <div
    class="relative flex min-h-screen w-full flex-col bg-slate-50 font-sans text-slate-600"
  >
    <CustomBanner title="联系我们" />
    <div
      class="mx-auto w-full space-y-4 px-4 py-6 sm:px-6 md:w-4/5 md:py-8 lg:px-8"
    >
      <div
        class="flex flex-wrap items-center text-xs font-medium text-slate-500 sm:text-sm"
      >
        <a class="transition-colors hover:text-blue-600" href="/">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">{{ tabs[activeTab] }}</span>
      </div>

      <div
        class="grid w-full grid-cols-2 gap-2 rounded-2xl bg-white p-2 shadow-sm ring-1 ring-slate-200/60 md:flex md:flex-nowrap"
      >
        <button
          v-for="(tab, index) in tabs"
          :key="index"
          @click="activeTab = index"
          :class="[
            'group relative flex-1 overflow-hidden rounded-xl px-2 py-2.5 text-center text-sm font-bold tracking-widest transition-all duration-300 sm:px-4 sm:py-3.5 sm:text-base md:text-lg',
            activeTab === index
              ? 'bg-linear-to-r from-blue-600 to-indigo-600 text-white shadow-lg ring-1 shadow-blue-500/30 ring-blue-500/50'
              : 'bg-transparent text-slate-500 hover:bg-slate-50 hover:text-blue-600',
          ]"
        >
          <span class="relative z-10">{{ tab }}</span>
          <div
            v-if="activeTab === index"
            class="absolute inset-0 z-0 bg-linear-to-b from-white/20 to-transparent opacity-60"
          ></div>
        </button>
      </div>
      <div
        class="min-h-[50vh] rounded-2xl bg-white p-4 shadow-sm ring-1 ring-slate-100 sm:p-6 md:p-8"
      >
        <transition name="fade" mode="out-in">
          <component :is="currentTabComponent" :key="activeTab" />
        </transition>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ChevronRight } from "lucide-vue-next"
import { computed, ref } from "vue"

import CustomBanner from "@/components/customBanner/index.vue"

import ContactInfo from "./components/ContactInfo.vue"
import CopyrightInfo from "./components/CopyrightInfo.vue"
import JoinUs from "./components/JoinUs.vue"
import OnlineMessage from "./components/OnlineMessage.vue"

const activeTab = ref(0)
const tabs = ["联系我们", "在线留言", "加入我们", "版权说明"]

const components = [ContactInfo, OnlineMessage, JoinUs, CopyrightInfo]
const currentTabComponent = computed(() => components[activeTab.value])
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition:
    opacity 0.3s ease,
    transform 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}
</style>
