<template>
  <div class="space-y-10">
    <!-- 顶部愿景与使命 Banner -->
    <div
      class="relative overflow-hidden rounded-3xl bg-linear-to-br from-blue-600 via-indigo-700 to-blue-900 px-6 py-16 text-center shadow-xl shadow-blue-900/20 sm:px-12"
    >
      <!-- 动态背景装饰 -->
      <div
        class="absolute -top-24 -left-24 h-64 w-64 rounded-full bg-white/10 blur-3xl"
      ></div>
      <div
        class="absolute -right-24 -bottom-24 h-64 w-64 rounded-full bg-blue-400/20 blur-3xl"
      ></div>
      <div
        class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAiIGhlaWdodD0iMjAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGNpcmNsZSBjeD0iMSIgY3k9IjEiIHI9IjEiIGZpbGw9InJnYmEoMjU1LDI1NSwyNTUsMC4xKSIvPjwvc3ZnPg==')] opacity-50"
      ></div>

      <div class="relative z-10 flex flex-col items-center justify-center">
        <div
          class="mb-6 flex h-16 w-16 items-center justify-center rounded-2xl bg-white/10 text-white ring-1 ring-white/20 backdrop-blur-md"
        >
          <Leaf class="h-8 w-8" />
        </div>
        <h2
          class="text-3xl font-black tracking-widest text-white drop-shadow-md sm:text-4xl md:text-5xl"
        >
          赓续初心 探索可持续发展
        </h2>
        <div class="mt-8 flex items-center justify-center space-x-2">
          <span class="h-px w-12 bg-white/30"></span>
          <p
            class="text-sm font-medium tracking-widest text-blue-100 uppercase"
          >
            OUR CULTURE & VISION
          </p>
          <span class="h-px w-12 bg-white/30"></span>
        </div>
      </div>
    </div>

    <!-- 核心价值观标题 -->
    <div class="text-center">
      <h3 class="relative inline-block text-2xl font-bold text-slate-800">
        核心价值观
        <span
          class="absolute -bottom-2 left-1/2 h-1 w-1/2 -translate-x-1/2 rounded-full bg-blue-600"
        ></span>
      </h3>
    </div>

    <!-- 核心价值观网格 -->
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="(item, index) in cultureList"
        :key="index"
        class="group relative overflow-hidden rounded-2xl border border-slate-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:shadow-xl"
        :class="item.hoverBorder"
      >
        <!-- 背景渐变光晕 -->
        <div
          class="absolute -top-8 -right-8 h-32 w-32 rounded-full opacity-0 blur-2xl transition-opacity duration-500 group-hover:opacity-100"
          :class="item.glowColor"
        ></div>

        <div class="relative z-10">
          <!-- 图标区域 -->
          <div
            class="mb-6 flex h-14 w-14 items-center justify-center rounded-xl transition-transform duration-500 group-hover:scale-110 group-hover:rotate-3"
            :class="[item.bgClass, item.textClass]"
          >
            <component :is="item.icon" class="h-7 w-7" />
          </div>

          <!-- 文本区域 -->
          <h4
            class="mb-3 text-xl font-bold text-slate-800 transition-colors duration-300"
            :class="`group-hover:${item.textClass}`"
          >
            {{ item.title }}
          </h4>

          <p class="text-sm leading-relaxed text-slate-500">
            {{ item.description }}
          </p>
        </div>

        <!-- 底部装饰线 -->
        <div
          class="absolute bottom-0 left-0 h-1 w-0 transition-all duration-500 group-hover:w-full"
          :class="item.lineColor"
        ></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Compass,
  Flame,
  Gem,
  HeartHandshake,
  Leaf,
  ShieldCheck,
  Users,
} from "lucide-vue-next"
import { markRaw } from "vue"

const cultureList = [
  {
    title: "客户至上",
    description: "以客户需求为导向，持续创造价值",
    icon: markRaw(HeartHandshake),
    bgClass: "bg-blue-50",
    textClass: "text-blue-600",
    hoverBorder: "hover:border-blue-300 hover:shadow-blue-900/5",
    glowColor: "bg-blue-400/20",
    lineColor: "bg-blue-500",
  },
  {
    title: "激情奋斗",
    description: "永葆创业激情，敢于挑战不可能",
    icon: markRaw(Flame),
    bgClass: "bg-rose-50",
    textClass: "text-rose-600",
    hoverBorder: "hover:border-rose-300 hover:shadow-rose-900/5",
    glowColor: "bg-rose-400/20",
    lineColor: "bg-rose-500",
  },
  {
    title: "拥抱变化",
    description: "敏锐洞察趋势，主动寻求突破",
    icon: markRaw(Compass),
    bgClass: "bg-violet-50",
    textClass: "text-violet-600",
    hoverBorder: "hover:border-violet-300 hover:shadow-violet-900/5",
    glowColor: "bg-violet-400/20",
    lineColor: "bg-violet-500",
  },
  {
    title: "诚信正直",
    description: "坚守商业道德，坦荡做人做事",
    icon: markRaw(ShieldCheck),
    bgClass: "bg-emerald-50",
    textClass: "text-emerald-600",
    hoverBorder: "hover:border-emerald-300 hover:shadow-emerald-900/5",
    glowColor: "bg-emerald-400/20",
    lineColor: "bg-emerald-500",
  },
  {
    title: "精益求精",
    description: "追求卓越品质，关注每一个细节",
    icon: markRaw(Gem),
    bgClass: "bg-amber-50",
    textClass: "text-amber-600",
    hoverBorder: "hover:border-amber-300 hover:shadow-amber-900/5",
    glowColor: "bg-amber-400/20",
    lineColor: "bg-amber-500",
  },
  {
    title: "团队合作",
    description: "共享共担，协同作战达成目标",
    icon: markRaw(Users),
    bgClass: "bg-cyan-50",
    textClass: "text-cyan-600",
    hoverBorder: "hover:border-cyan-300 hover:shadow-cyan-900/5",
    glowColor: "bg-cyan-400/20",
    lineColor: "bg-cyan-500",
  },
]
</script>

<style scoped lang="scss"></style>
