<template>
  <section
    class="flex-col-center relative min-h-screen overflow-hidden bg-slate-50 font-sans select-none"
  >
    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>

    <div
      class="absolute top-0 left-0 -z-10 h-64 w-full bg-linear-to-b from-slate-50 to-transparent"
    ></div>

    <div class="relative z-10 container mx-auto px-6 lg:px-12">
      <!-- 移除了初始的 opacity-0 和 translate-y-8 -->
      <div class="header-group relative mb-16">
        <h2
          class="flex items-center gap-3 text-4xl font-black tracking-tight text-slate-900 md:text-5xl"
        >
          <span
            class="bg-linear-to-r from-blue-700 to-indigo-600 bg-clip-text text-transparent"
          >
            产品优势
          </span>
          <span class="h-3 w-3 animate-pulse rounded-full bg-blue-600"></span>
        </h2>
        <div
          class="mt-4 h-1 w-24 rounded-full bg-linear-to-r from-blue-600 to-transparent"
        ></div>
      </div>

      <div class="mb-8 grid grid-cols-2 gap-4 md:grid-cols-4">
        <!-- 移除了初始的 opacity-0 和 translate-y-4 -->
        <button
          v-for="(item, index) in systemConfig.advantages"
          :key="index"
          @click="handleTabClick(index)"
          class="gsap-tab-btn group relative flex h-32 cursor-pointer flex-col items-center justify-center gap-3 overflow-hidden rounded-xl border transition-all duration-300"
          :class="
            activeTab === index
              ? 'z-10 -translate-y-1 border-blue-600 bg-blue-600 shadow-xl shadow-blue-600/30'
              : 'border-slate-200 bg-white hover:border-blue-300 hover:bg-blue-50/50'
          "
        >
          <div
            v-if="activeTab === index"
            class="pointer-events-none absolute inset-0 bg-linear-to-br from-white/10 to-transparent"
          ></div>
          <span
            class="absolute top-2 right-4 text-4xl font-black italic opacity-20 transition-colors select-none"
            :class="activeTab === index ? 'text-white' : 'text-slate-300'"
          >
            {{ item.id }}
          </span>
          <component
            :is="item.icon"
            class="relative z-10 h-8 w-8 transition-colors duration-300"
            :class="
              activeTab === index
                ? 'text-white'
                : 'text-slate-500 group-hover:text-blue-500'
            "
          />
          <span
            class="relative z-10 text-lg font-bold transition-colors duration-300"
            :class="
              activeTab === index
                ? 'text-white'
                : 'text-slate-600 group-hover:text-blue-600'
            "
          >
            {{ item.tabTitle }}
          </span>
        </button>
      </div>

      <div
        ref="contentContainerRef"
        class="relative min-h-125 overflow-hidden rounded-2xl border border-slate-100 bg-white"
      >
        <div class="grid h-full grid-cols-1 lg:grid-cols-12">
          <!-- 移除了初始的 opacity-0 -->
          <div
            class="left-panel-block group relative h-64 overflow-hidden lg:col-span-5 lg:h-full"
          >
            <div
              class="absolute inset-0 z-10 bg-blue-900/10 transition-colors duration-500 group-hover:bg-transparent"
            ></div>
            <img
              :src="currentItem?.image"
              :alt="currentItem?.headline"
              class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-105"
            />
            <div
              class="absolute top-6 left-6 z-20 rounded bg-white/90 px-4 py-1 text-sm font-bold text-blue-800 shadow-sm backdrop-blur"
            >
              {{ currentItem?.tabTitle }} VIEW
            </div>
          </div>

          <!-- 移除了初始的 opacity-0 -->
          <div
            class="right-panel-block relative flex flex-col justify-center bg-white p-8 lg:col-span-7 lg:p-14"
          >
            <span
              class="absolute top-6 right-6 z-0 text-9xl font-black text-slate-50 select-none"
            >
              {{ currentItem?.id }}
            </span>

            <div class="relative z-10">
              <span
                class="mb-2 block text-sm font-bold tracking-widest text-blue-600 uppercase"
              >
                {{ currentItem?.subTitle }}
              </span>

              <h3
                class="mb-6 text-3xl leading-tight font-black text-slate-900 lg:text-4xl"
              >
                {{ currentItem?.headline }}
              </h3>
              <div class="mb-8 h-1.5 w-20 rounded-full bg-blue-600"></div>

              <ul class="mb-12 space-y-6">
                <li
                  v-for="(desc, dIdx) in currentItem?.descriptions"
                  :key="dIdx"
                  class="group/item flex items-start gap-4"
                >
                  <div
                    class="mt-1 flex h-5 w-5 shrink-0 items-center justify-center rounded-full bg-blue-50 transition-colors group-hover/item:bg-blue-100"
                  >
                    <div class="h-2 w-2 rounded-full bg-blue-600"></div>
                  </div>
                  <p
                    class="text-justify text-lg leading-relaxed text-slate-600 transition-colors group-hover/item:text-slate-900"
                  >
                    {{ desc }}
                  </p>
                </li>
              </ul>

              <div
                class="flex flex-col items-center justify-between gap-8 rounded-xl border border-slate-200 bg-slate-50 p-6 transition-colors hover:border-blue-200 sm:flex-row"
              >
                <div class="flex w-full gap-8 sm:w-auto">
                  <div
                    v-for="(stat, sIdx) in currentItem?.stats"
                    :key="sIdx"
                    class="flex items-center gap-3"
                  >
                    <div
                      class="flex h-12 w-12 items-center justify-center rounded-lg border border-slate-200 bg-white text-blue-600 shadow-sm"
                    >
                      <component :is="stat.icon" class="h-6 w-6" />
                    </div>
                    <div>
                      <div
                        class="text-2xl leading-none font-black text-slate-900"
                      >
                        <span
                          class="gsap-counter"
                          :data-target="parseNumber(stat.value)"
                          >{{ parseNumber(stat.value) }}</span
                        >{{ getSuffix(stat.value) }}
                      </div>
                      <div class="mt-1 text-xs text-slate-500 uppercase">
                        {{ stat.label }}
                      </div>
                    </div>
                  </div>
                </div>
                <div
                  class="mx-4 hidden h-px flex-1 bg-slate-200 sm:block"
                ></div>
                <div class="hidden text-right sm:block">
                  <span
                    class="text-xs font-bold tracking-widest text-blue-600/80"
                    >XINGSHI · TECH</span
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import gsap from "gsap"
import { computed, nextTick, onMounted, onUnmounted, ref } from "vue"

import { systemConfig } from "@/config/header"

let ctx: gsap.Context
let timer: number | null = null

const contentContainerRef = ref<HTMLElement | null>(null)
const activeTab = ref(0)
const isAnimating = ref(false)
const currentItem = computed(() => systemConfig.advantages[activeTab.value])

function startAutoTab() {
  if (timer) clearInterval(timer)
  timer = setInterval(() => {
    let nextIndex = activeTab.value + 1
    if (nextIndex >= systemConfig.advantages.length) {
      nextIndex = 0
    }
    handleTabClick(nextIndex)
  }, 4000)
}

function parseNumber(str: string) {
  const match = str.match(/\d+(\.\d+)?/)
  return match ? parseFloat(match[0]) : 0
}

function getSuffix(str: string) {
  const match = str.replace(/[\d\.]/g, "")
  return match
}

function animateContentIn() {
  if (!contentContainerRef.value) return
  gsap.killTweensOf(".gsap-counter")

  const tl = gsap.timeline({
    onComplete: () => {
      isAnimating.value = false
    },
  })

  tl.set([".left-panel-block", ".right-panel-block"], {
    opacity: 0,
    y: 15,
  })

  tl.to([".left-panel-block", ".right-panel-block"], {
    opacity: 1,
    y: 0,
    duration: 0.4,
    stagger: 0.1,
    ease: "power2.out",
  })

  tl.from(
    ".gsap-counter",
    {
      textContent: 0,
      duration: 1,
      ease: "power1.out",
      snap: { textContent: 1 },
      onUpdate: function () {
        // @ts-ignore
        this.targets().forEach((target) => {
          const val = Math.ceil(target.textContent)
          target.textContent = val
        })
      },
    },
    "<"
  )
}

async function animateContentOut() {
  if (!contentContainerRef.value) return
  await gsap.to([".left-panel-block", ".right-panel-block"], {
    opacity: 0,
    y: -10,
    duration: 0.2,
    ease: "power1.in",
    overwrite: true,
  })
}

async function handleTabClick(index: number) {
  if (activeTab.value === index || isAnimating.value) return
  isAnimating.value = true
  startAutoTab()
  await animateContentOut()
  activeTab.value = index
  await nextTick()
  animateContentIn()
}

onMounted(() => {
  ctx = gsap.context(() => {
    animateContentIn()
  }, contentContainerRef.value || undefined)
  startAutoTab()
})

onUnmounted(() => {
  ctx?.revert()
  if (timer) {
    clearInterval(timer)
    timer = null
  }
})
</script>

<style scoped>
.gsap-counter {
  font-variant-numeric: tabular-nums;
}
</style>
