<template>
  <section
    ref="sectionRef"
    class="relative flex items-center overflow-hidden bg-slate-50 py-10 font-sans"
  >
    <div
      class="animate-pulse-slow absolute top-0 left-0 -z-10 h-125 w-125 rounded-full bg-blue-100/60 mix-blend-multiply blur-[100px]"
    ></div>

    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>

    <div class="container mx-auto px-6 lg:px-12">
      <div
        class="grid grid-cols-1 items-center gap-16 lg:grid-cols-2 lg:gap-20"
      >
        <!-- 左侧-->
        <div class="relative z-10 space-y-8">
          <div class="left-content-group">
            <h2
              class="gsap-text-reveal mb-6 text-4xl leading-tight font-black text-slate-900 lg:text-5xl"
            >
              为什么选择 <br />
              <span
                class="relative inline-block bg-linear-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent"
              >
                星实科技
                <span
                  class="gsap-highlight absolute bottom-1 left-0 -z-10 h-3 w-full origin-left scale-x-0 bg-blue-200/30"
                ></span>
              </span>
              ?
            </h2>

            <p
              class="gsap-text-reveal border-l-4 border-blue-200 pl-6 text-justify text-lg leading-relaxed text-slate-600"
            >
              公司专注于非标领域的自动化的需求发展，我们致力于引进国内外高精尖技术力量，向中国高校及科研技术企业提供顶尖的非标自动化设备和完善的技术支持。
            </p>
          </div>

          <div class="flex flex-wrap gap-4 py-4">
            <div
              v-for="(val, index) in systemConfig.chooseUs.values"
              :key="index"
              class="gsap-tag flex cursor-default items-center gap-2 rounded-full border border-slate-200 bg-white px-4 py-2 text-sm font-bold text-slate-700 shadow-sm transition-all duration-300 hover:-translate-y-1 hover:border-blue-400 hover:text-blue-600"
            >
              <component :is="val.icon" class="h-4 w-4 text-blue-500" />
              {{ val.text }}
            </div>
          </div>

          <div class="gsap-btn-wrapper pt-4">
            <button
              class="group relative overflow-hidden rounded-lg bg-blue-600 px-8 py-3 font-bold text-white shadow-lg shadow-blue-600/30 transition-all hover:-translate-y-1 hover:shadow-blue-600/50"
            >
              <div
                class="animate-shimmer absolute inset-0 h-full w-full -translate-x-full bg-linear-to-r from-transparent via-white/20 to-transparent transition-transform duration-1000 group-hover:translate-x-full group-hover:animate-none"
              ></div>
              <span class="relative z-10 flex items-center gap-2">
                了解更多
                <ArrowRight
                  class="h-4 w-4 transition-transform group-hover:translate-x-1"
                />
              </span>
            </button>
          </div>
        </div>

        <!-- 右侧 -->
        <div
          class="perspective-container relative grid grid-cols-1 gap-5 sm:grid-cols-2"
        >
          <div
            class="gsap-bg-blur absolute inset-0 -z-10 -rotate-3 transform rounded-[3rem] bg-linear-to-tr from-blue-50 to-indigo-50 blur-xl"
          ></div>

          <div
            v-for="(feature, index) in systemConfig.chooseUs.merit"
            :key="index"
            class="gsap-card-wrapper"
          >
            <div
              class="group relative h-full transform overflow-hidden rounded-2xl border border-slate-100 bg-white p-6 shadow-lg transition-all duration-300 hover:-translate-y-2 hover:shadow-2xl"
            >
              <div
                class="absolute top-0 left-0 h-1 w-full origin-left scale-x-0 bg-linear-to-r from-blue-400 to-indigo-400 transition-transform duration-500 group-hover:scale-x-100"
              ></div>

              <span
                class="absolute -right-2 -bottom-4 text-6xl font-black text-slate-50 transition-colors select-none group-hover:text-blue-50/80"
              >
                {{ feature.id }}
              </span>

              <div
                class="mb-4 flex h-12 w-12 items-center justify-center rounded-xl text-white shadow-md transition-transform duration-300 group-hover:scale-110 group-hover:rotate-6"
                :class="feature.color"
              >
                <component :is="feature.icon" class="h-6 w-6" />
              </div>

              <h3
                class="mb-2 text-xl font-bold text-slate-900 transition-colors group-hover:text-blue-600"
              >
                {{ feature.title }}
              </h3>
              <p class="text-sm leading-relaxed text-slate-500">
                {{ feature.desc }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import gsap from "gsap"
import { ScrollTrigger } from "gsap/ScrollTrigger"
import { ArrowRight } from "lucide-vue-next"
import { onMounted, onUnmounted, ref } from "vue"

import { systemConfig } from "@/config/header"

gsap.registerPlugin(ScrollTrigger)

let ctx: gsap.Context
const sectionRef = ref<HTMLElement | null>(null)

onMounted(() => {
  if (!sectionRef.value) return
  ctx = gsap.context(() => {
    // 左侧文本动画
    const tlText = gsap.timeline({
      scrollTrigger: {
        trigger: ".left-content-group",
        start: "top 85%",
        end: "bottom 15%",
        toggleActions: "play reverse play reverse",
      },
    })

    tlText
      .from(".gsap-text-reveal", {
        y: 30,
        opacity: 0,
        duration: 0.8,
        stagger: 0.1,
        ease: "power3.out",
      })
      .to(
        ".gsap-line",
        {
          width: "3rem",
          duration: 0.6,
          ease: "power2.out",
        },
        "<0.2"
      )
      .to(
        ".gsap-highlight",
        {
          scaleX: 1,
          duration: 0.8,
          ease: "expo.out",
        },
        "-=0.4"
      )

    // 标签和按钮动画
    gsap.from(".gsap-tag", {
      scrollTrigger: {
        trigger: ".gsap-tag",
        start: "top 90%",
        end: "bottom 10%",
        toggleActions: "play reverse play reverse",
      },
      y: 20,
      opacity: 0,
      duration: 0.5,
      stagger: 0.1,
      ease: "back.out(1.7)",
    })

    gsap.from(".gsap-btn-wrapper", {
      scrollTrigger: {
        trigger: ".gsap-btn-wrapper",
        start: "top 95%",
        end: "bottom 5%",
        toggleActions: "play reverse play reverse",
      },
      y: 20,
      opacity: 0,
      duration: 0.6,
      delay: 0.1,
      ease: "power2.out",
    })

    // 右侧卡片动画
    const tlCards = gsap.timeline({
      scrollTrigger: {
        trigger: ".perspective-container",
        start: "top 80%",
        end: "bottom top",
        toggleActions: "play reverse play reverse",
      },
    })

    tlCards.from(".gsap-bg-blur", {
      scale: 0.8,
      opacity: 0,
      duration: 1,
      ease: "power2.out",
    })

    tlCards.from(
      ".gsap-card-wrapper",
      {
        y: 80,
        rotationX: 25,
        opacity: 0,
        scale: 0.85,
        duration: 0.8,
        stagger: 0.1,
        ease: "back.out(1.0)",
        transformOrigin: "center top",
        clearProps: "all",
      },
      "<0.1"
    )
  }, sectionRef.value)
})

onUnmounted(() => {
  ctx.revert()
})
</script>

<style scoped>
.perspective-container {
  perspective: 1000px;
}

@keyframes pulse-slow {
  0%,
  100% {
    opacity: 0.5;
    transform: scale(1);
  }
  50% {
    opacity: 0.3;
    transform: scale(1.1);
  }
}
.animate-pulse-slow {
  animation: pulse-slow 8s infinite ease-in-out;
}

@keyframes shimmer {
  0% {
    transform: translateX(-150%);
  }
  50% {
    transform: translateX(150%);
  }
  100% {
    transform: translateX(150%);
  }
}
.animate-shimmer {
  animation: shimmer 3s infinite;
}
</style>
