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

    <div class="relative z-10 container mx-auto px-6 py-16 lg:px-12">
      <div class="header-group relative mb-8">
        <div class="mb-1 flex items-end gap-2">
          <span class="text-3xl font-medium tracking-wide">Quality</span>
          <div class="mb-2 h-0.5 w-12"></div>
        </div>
        <h2 class="mb-3 text-4xl font-bold tracking-tight md:text-5xl">
          星实设备·以质取胜
        </h2>
        <p
          class="text-sm font-medium tracking-wide text-slate-800 md:text-base"
        >
          用数据捍卫标准，以专业定义可靠——因为信任始于真实，卓越源于对品质的执着
        </p>
      </div>

      <div class="mb-10 grid grid-cols-2 md:grid-cols-4">
        <button
          v-for="(item, index) in systemConfig.advantages"
          :key="index"
          @click="HandleTabClick(index)"
          class="gsap-tab-btn group relative flex h-24 cursor-pointer flex-col items-center justify-center gap-1 border-r border-white/50 transition-all duration-300 last:border-r-0 md:h-28"
          :class="
            activeTab === index
              ? 'bg-[#4773c4]'
              : 'bg-[#999999] hover:bg-[#8c8c8c]'
          "
        >
          <span
            class="text-3xl font-bold transition-colors duration-300 md:text-4xl"
            :class="
              activeTab === index
                ? 'text-white'
                : 'text-[#7cb0e5] group-hover:text-[#6ba1d9]'
            "
          >
            {{ String(index + 1).padStart(2, "0") }}
          </span>
          <span
            class="px-2 text-center text-xs font-medium transition-colors duration-300 md:text-sm"
            :class="
              activeTab === index
                ? 'text-white'
                : 'text-slate-900 group-hover:text-black'
            "
          >
            {{ tabTitles[index] || item.tabTitle }}
          </span>
        </button>
      </div>

      <div
        ref="contentContainerRef"
        class="relative min-h-125 overflow-hidden rounded-2xl border border-slate-100 bg-white shadow-sm"
      >
        <div class="grid h-full grid-cols-1 lg:grid-cols-12">
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
              {{ tabTitles[activeTab] || currentItem?.tabTitle }} VIEW
            </div>
          </div>

          <div
            class="right-panel-block relative flex flex-col justify-center bg-white p-8 lg:col-span-7 lg:p-14"
          >
            <span
              class="absolute top-6 right-6 z-0 text-9xl font-black text-slate-50 select-none"
            >
              {{ String(activeTab + 1).padStart(2, "0") }}
            </span>

            <div class="relative z-10">
              <span
                class="mb-2 block text-sm font-bold tracking-widest text-[#4773c4] uppercase"
              >
                {{ currentItem?.subTitle }}
              </span>

              <h3
                class="mb-6 text-3xl leading-tight font-black text-slate-900 lg:text-4xl"
              >
                {{ currentItem?.headline }}
              </h3>
              <div class="mb-8 h-1.5 w-20 rounded-full bg-[#4773c4]"></div>

              <ul class="mb-12 space-y-6">
                <li
                  v-for="(desc, dIdx) in currentItem?.descriptions"
                  :key="dIdx"
                  class="group/item flex items-start gap-4"
                >
                  <div
                    class="mt-1 flex h-5 w-5 shrink-0 items-center justify-center rounded-full bg-blue-50 transition-colors group-hover/item:bg-blue-100"
                  >
                    <div class="h-2 w-2 rounded-full bg-[#4773c4]"></div>
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
                      class="flex h-12 w-12 items-center justify-center rounded-lg border border-slate-200 bg-white text-[#4773c4] shadow-sm"
                    >
                      <component :is="stat.icon" class="h-6 w-6" />
                    </div>
                    <div>
                      <div
                        class="text-2xl leading-none font-black text-slate-900"
                      >
                        <span
                          class="gsap-counter"
                          :data-target="ParseNumber(stat.value)"
                          >{{ ParseNumber(stat.value) }}</span
                        >{{ GetSuffix(stat.value) }}
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
                    class="text-xs font-bold tracking-widest text-[#4773c4]/80"
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

const tabTitles = [
  "卓越的工程设计质量",
  "可靠的零部件与制造质量",
  "稳定的过程控制与调试质量",
  "卓越的长期运行与服务质量",
]

function StartAutoTab() {
  if (timer) clearInterval(timer)
  timer = setInterval(() => {
    let nextIndex = activeTab.value + 1
    if (nextIndex >= systemConfig.advantages.length) {
      nextIndex = 0
    }
    HandleTabClick(nextIndex)
  }, 4000)
}

function ParseNumber(str: string) {
  const match = str.match(/\d+(\.\d+)?/)
  return match ? parseFloat(match[0]) : 0
}

function GetSuffix(str: string) {
  const match = str.replace(/[\d\.]/g, "")
  return match
}

function AnimateContentIn() {
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

async function AnimateContentOut() {
  if (!contentContainerRef.value) return
  await gsap.to([".left-panel-block", ".right-panel-block"], {
    opacity: 0,
    y: -10,
    duration: 0.2,
    ease: "power1.in",
    overwrite: true,
  })
}

async function HandleTabClick(index: number) {
  if (activeTab.value === index || isAnimating.value) return
  isAnimating.value = true
  StartAutoTab()
  await AnimateContentOut()
  activeTab.value = index
  await nextTick()
  AnimateContentIn()
}

onMounted(() => {
  ctx = gsap.context(() => {
    AnimateContentIn()
  }, contentContainerRef.value || undefined)
  StartAutoTab()
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
