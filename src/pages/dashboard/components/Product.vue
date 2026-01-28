<template>
  <section
    ref="containerRef"
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
      class="parallax-item pointer-events-none absolute top-0 right-0 -z-10 h-1/2 w-1/2 rounded-full bg-blue-100/50 blur-[120px]"
      data-depth="0.05"
    ></div>
    <div
      class="parallax-item pointer-events-none absolute bottom-0 left-0 -z-10 h-1/3 w-1/3 rounded-full bg-indigo-100/50 blur-[100px]"
      data-depth="-0.05"
    ></div>

    <div class="relative z-10 container mx-auto px-6 lg:px-12">
      <div class="relative mb-16">
        <h2
          class="flex items-center gap-3 text-4xl font-black tracking-tight text-slate-900 md:text-5xl"
        >
          <span
            class="bg-linear-to-r from-blue-700 to-indigo-600 bg-clip-text text-transparent"
          >
            热门产品系列
          </span>
          <span class="h-3 w-3 animate-pulse rounded-full bg-blue-600"></span>
        </h2>
        <div
          class="mt-4 h-1 w-24 rounded-full bg-linear-to-r from-blue-600 to-transparent"
        ></div>
      </div>

      <div
        class="grid grid-cols-1 items-center gap-16 lg:grid-cols-2 lg:gap-24"
        @mouseenter="handleMouseEnter"
        @mouseleave="handleMouseLeave"
        @mousemove="handleMouseMove"
      >
        <!-- 左侧-->
        <div class="group perspective-1000 relative">
          <div
            ref="decorationRef"
            class="anim-decoration absolute -top-4 -left-4 z-20 h-20 w-20 border-t-2 border-l-2 border-blue-600"
          ></div>
          <div
            class="anim-decoration absolute -right-4 -bottom-4 z-20 h-20 w-20 border-r-2 border-b-2 border-blue-600"
          ></div>

          <Crosshair
            class="anim-decoration absolute -top-6 -right-6 h-8 w-8 text-slate-300 opacity-50"
          />
          <Crosshair
            class="anim-decoration absolute -bottom-6 -left-6 h-8 w-8 text-slate-300 opacity-50"
          />

          <div
            class="relative overflow-hidden rounded-lg border border-white/50 bg-white/40 p-4 shadow-xl backdrop-blur-md"
          >
            <div
              class="animate-scan pointer-events-none absolute inset-0 z-20 h-[20%] w-full bg-linear-to-b from-transparent via-blue-400/10 to-transparent"
            ></div>

            <div
              class="relative aspect-4/3 overflow-hidden rounded bg-slate-200"
            >
              <img
                ref="productImageRef"
                :src="currentProduct?.image"
                :alt="currentProduct?.name"
                class="h-full w-full object-cover will-change-transform"
              />
              <div
                class="absolute inset-0 bg-blue-900/10 mix-blend-overlay"
              ></div>
            </div>
          </div>
        </div>

        <!-- 右侧-->
        <div ref="rightSideRef" class="relative min-h-100">
          <div class="space-y-6">
            <div class="relative overflow-hidden">
              <div
                class="anim-text stroke-text absolute -top-16 -left-4 -z-10 bg-linear-to-b from-slate-200 to-white bg-clip-text font-mono text-9xl font-black tracking-tighter text-transparent opacity-60 select-none"
              >
                {{ currentProduct?.id }}
              </div>

              <div
                class="anim-text mb-2 flex items-center gap-3 font-mono text-sm font-bold tracking-wider text-blue-600"
              >
                <Zap class="h-4 w-4 fill-current" />
                PRODUCT_{{ currentProduct?.id }}
              </div>

              <h3
                class="anim-text text-4xl leading-tight font-bold text-slate-900 lg:text-5xl"
              >
                {{ currentProduct?.name }}
              </h3>
              <p
                class="anim-text mt-1 text-lg font-medium tracking-wide text-slate-400 uppercase"
              >
                {{ currentProduct?.subtitle }}
              </p>
            </div>
            <ul class="space-y-4 pt-4">
              <li
                v-for="(feature, idx) in currentProduct?.features"
                :key="`${currentProduct?.id}-feat-${idx}`"
                class="anim-text flex items-center text-lg text-slate-700"
              >
                <span
                  class="mr-4 h-1.5 w-1.5 rotate-45 bg-blue-500 transition-all duration-300 group-hover:bg-indigo-500"
                ></span>
                {{ feature }}
              </li>
            </ul>
            <div class="anim-text flex items-center gap-6 pt-8">
              <button
                class="group relative overflow-hidden rounded-sm bg-slate-900 px-8 py-3.5 font-bold text-white transition-all hover:shadow-[0_0_20px_rgba(59,130,246,0.5)]"
              >
                <span class="relative z-10 flex items-center gap-2">
                  查看详情
                  <ArrowRight
                    class="h-4 w-4 transition-transform group-hover:translate-x-1"
                  />
                </span>
                <div
                  class="absolute inset-0 -translate-x-full bg-linear-to-r from-blue-600 to-indigo-600 transition-transform duration-300 ease-out group-hover:translate-x-0"
                ></div>
              </button>
            </div>
          </div>

          <!-- 底部控制栏 -->
          <div
            class="absolute right-0 -bottom-20 left-0 mt-12 flex items-center justify-between border-t border-slate-200 pt-6 lg:left-0"
          >
            <div
              class="relative mr-8 h-1 flex-1 overflow-hidden rounded-full bg-slate-200"
            >
              <div
                ref="progressBarRef"
                class="absolute top-0 left-0 h-full w-full origin-left scale-x-0 bg-blue-600"
              ></div>
            </div>

            <div class="flex gap-2">
              <button
                @click="manualChangeSlide(-1)"
                class="flex h-10 w-10 items-center justify-center rounded-sm border border-slate-300 text-slate-600 transition-all hover:border-blue-600 hover:bg-blue-50 hover:text-blue-600"
              >
                <ChevronLeft class="h-5 w-5" />
              </button>
              <button
                @click="manualChangeSlide(1)"
                class="flex h-10 w-10 items-center justify-center rounded-sm border border-slate-300 text-slate-600 transition-all hover:border-blue-600 hover:bg-blue-50 hover:text-blue-600"
              >
                <ChevronRight class="h-5 w-5" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import gsap from "gsap"
import {
  ArrowRight,
  ChevronLeft,
  ChevronRight,
  Crosshair,
  Zap,
} from "lucide-vue-next"
import { computed, nextTick, onMounted, onUnmounted, ref } from "vue"

import { systemConfig } from "@/config/header"

let ctx: gsap.Context | null = null
let autoPlayTween: gsap.core.Tween | null = null
const AUTO_PLAY_TIME = 5

const currentIndex = ref(0)
const isAnimating = ref(false)
const containerRef = ref<HTMLElement | null>(null)
const rightSideRef = ref<HTMLElement | null>(null)
const progressBarRef = ref<HTMLElement | null>(null)
const productImageRef = ref<HTMLElement | null>(null)
const decorationRef = ref<HTMLElement | null>(null)

const currentProduct = computed(() => systemConfig.products[currentIndex.value])

function animateIn() {
  if (!containerRef.value || !rightSideRef.value || !productImageRef.value) {
    return
  }

  const animTexts = rightSideRef.value.querySelectorAll(".anim-text")

  if (animTexts.length === 0) {
    console.warn("No .anim-text elements found")
  }

  const tl = gsap.timeline({
    onComplete: () => {
      isAnimating.value = false
      startAutoPlay()
    },
  })

  tl.set(productImageRef.value, {
    clipPath: "polygon(0% 0%, 0% 0%, 0% 100%, 0% 100%)",
    filter: "grayscale(0%) blur(5px)",
  })

  tl.set(animTexts, { y: 30, opacity: 0 })

  tl.to(productImageRef.value, {
    clipPath: "polygon(0% 0%, 100% 0%, 100% 100%, 0% 100%)",
    scale: 1,
    filter: "grayscale(0%) blur(0px)",
    duration: 1,
    ease: "power4.out",
  })

  tl.to(
    animTexts,
    {
      y: 0,
      opacity: 1,
      duration: 0.6,
      stagger: 0.05,
      ease: "back.out(1.2)",
    },
    "-=0.8"
  )

  return tl
}

function animateOut() {
  if (!rightSideRef.value || !productImageRef.value) {
    return gsap.timeline()
  }

  const animTexts = rightSideRef.value.querySelectorAll(".anim-text")

  const tl = gsap.timeline()
  tl.to(productImageRef.value, {
    clipPath: "polygon(100% 0%, 100% 0%, 100% 100%, 100% 100%)",
    scale: 1.1,
    filter: "grayscale(100%)",
    duration: 0.5,
    ease: "power2.in",
  })
  tl.to(
    animTexts,
    {
      y: -20,
      opacity: 0,
      duration: 0.4,
      stagger: 0.02,
      ease: "power2.in",
    },
    "<"
  )

  return tl
}

function manualChangeSlide(dir: number) {
  const next =
    (currentIndex.value + dir + systemConfig.products.length) %
    systemConfig.products.length
  switchSlide(next)
}

function startAutoPlay() {
  if (!progressBarRef.value) return
  if (autoPlayTween) autoPlayTween.kill()
  autoPlayTween = gsap.fromTo(
    progressBarRef.value,
    { scaleX: 0 },
    {
      scaleX: 1,
      duration: AUTO_PLAY_TIME,
      ease: "none",
      onComplete: () => {
        const next = (currentIndex.value + 1) % systemConfig.products.length
        switchSlide(next)
      },
    }
  )
}

function handleMouseEnter() {
  if (!isAnimating.value && autoPlayTween) autoPlayTween.pause()
}

function handleMouseLeave() {
  if (!isAnimating.value && autoPlayTween) autoPlayTween.play()
}

function handleMouseMove(e: MouseEvent) {
  if (!containerRef.value) return
  const x = (e.clientX / window.innerWidth - 0.5) * 2
  const y = (e.clientY / window.innerHeight - 0.5) * 2

  const parallaxItems = containerRef.value.querySelectorAll(".parallax-item")

  gsap.to(parallaxItems, {
    x: (_, target) => {
      const depth = parseFloat((target as HTMLElement).dataset.depth || "0")
      return x * 30 * depth
    },
    y: (_, target) => {
      const depth = parseFloat((target as HTMLElement).dataset.depth || "0")
      return y * 30 * depth
    },
    duration: 0.5,
    overwrite: "auto",
  })
}

async function switchSlide(newIndex: number) {
  if (isAnimating.value) return
  isAnimating.value = true
  if (autoPlayTween) autoPlayTween.pause()
  await animateOut().then()
  currentIndex.value = newIndex
  await nextTick()
  animateIn()
}

onMounted(() => {
  nextTick(() => {
    if (!containerRef.value || !decorationRef.value) return

    ctx = gsap.context(() => {
      const decorations =
        containerRef.value!.querySelectorAll(".anim-decoration")

      if (decorations.length > 0) {
        gsap.from(decorations, {
          scale: 0,
          opacity: 0,
          duration: 1,
          stagger: 0.2,
          ease: "elastic.out(1, 0.5)",
        })
      }

      animateIn()
    }, containerRef.value)
  })
})

onUnmounted(() => {
  if (ctx) {
    ctx.revert()
  }
  if (autoPlayTween) {
    autoPlayTween.kill()
  }
})
</script>

<style scoped>
.perspective-1000 {
  perspective: 1000px;
}

@keyframes scan {
  0% {
    transform: translateY(-100%);
  }
  100% {
    transform: translateY(500%);
  }
}
.animate-scan {
  animation: scan 4s linear infinite;
}

.stroke-text {
  -webkit-text-stroke: 2px #e2e8f0;
}

.will-change-transform {
  will-change: transform, clip-path;
}
</style>
