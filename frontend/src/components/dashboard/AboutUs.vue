<template>
  <section
    id="about"
    class="flex-col-center relative min-h-screen w-full overflow-hidden bg-slate-50 py-20"
  >
    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>

    <div class="relative z-10 container mx-auto px-6">
      <div class="flex flex-col items-center gap-16 lg:flex-row">
        <!-- 左侧-->
        <div class="w-full space-y-8 lg:w-1/2">
          <div class="space-y-4">
            <h2
              class="text-4xl font-extrabold tracking-tight text-slate-900 md:text-5xl"
            >
              关于我们
            </h2>
            <p
              class="border-l-4 border-blue-600 pl-4 text-xl font-semibold text-slate-700"
            >
              您在非标设备领域的最佳合作伙伴
            </p>
          </div>

          <p class="text-justify text-lg leading-loose text-slate-600">
            星实（厦门）科技有限公司成立于2021年1月，专业生产各类激光系统的光学元件，为激光系统制造商和终端用户提供高质量的激光光学元件及解决方案。公司位于深圳市宝安区，占地面积9500平方米。我们是一家集光学设计、研发、样品制作、生产和销售于一体的光学元件制造商。我们的生产能力涵盖切割、研磨、抛光、镀膜、测试、组装和最终检验等全流程。
          </p>

          <div class="flex flex-wrap gap-4">
            <button
              class="group relative overflow-hidden rounded-full bg-[#b8dcf2] px-8 py-3 font-bold tracking-wider text-[#0062b1] shadow-sm transition-all duration-300 hover:-translate-y-1 hover:bg-[#a6d1eb] hover:shadow-md"
            >
              <div
                class="absolute inset-0 z-0 h-full w-full -translate-x-full bg-linear-to-r from-transparent via-white/80 to-transparent transition-transform duration-700 ease-in-out group-hover:translate-x-full"
              ></div>
              <span class="relative z-10 flex items-center gap-2 text-lg">
                立即联系我们
                <ArrowRight
                  class="h-5 w-5 transition-transform group-hover:translate-x-1"
                />
              </span>
            </button>
          </div>
        </div>

        <!-- 右侧 -->
        <div class="relative w-full lg:w-1/2">
          <div
            class="absolute -top-10 -right-10 h-32 w-32 bg-[radial-gradient(circle,var(--color-blue-600)_1px,transparent_1px)] bg-size-[12px_12px] opacity-20"
          ></div>

          <div
            class="absolute -bottom-6 -left-6 -z-10 h-2/3 w-2/3 rounded-3xl bg-linear-to-tr from-cyan-100 to-blue-50"
          ></div>

          <div
            class="relative transform overflow-hidden rounded-2xl border-4 border-white shadow-2xl transition-transform duration-500 hover:scale-[1.01]"
          >
            <img
              src="/images/aboutUs.jpg"
              alt="Modern Architecture"
              class="h-100 w-full object-cover"
            />

            <div
              class="absolute right-6 bottom-6 hidden rounded-xl border border-white/50 bg-white/90 p-4 shadow-lg backdrop-blur-md md:block"
            >
              <div class="flex items-center gap-3">
                <div
                  class="flex h-10 w-10 items-center justify-center rounded-full bg-green-100 text-green-600"
                >
                  <CheckCircle :size="20" />
                </div>
                <div>
                  <p class="text-sm text-slate-500">质量认证</p>
                  <p class="font-bold text-slate-900">ISO 9001 标准</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-20" ref="statsContainerRef">
        <div class="grid grid-cols-2 gap-8 md:grid-cols-3 lg:grid-cols-5">
          <div
            v-for="(item, index) in stats"
            :key="index"
            class="group rounded-xl border border-slate-100 bg-white p-6 text-center shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-xl"
          >
            <div
              class="mb-2 inline-block origin-center text-3xl font-bold text-blue-600 transition-transform duration-300 group-hover:scale-110 md:text-4xl"
            >
              {{ item.current }}{{ item.suffix }}
            </div>
            <div class="text-sm font-medium text-slate-500 md:text-base">
              {{ item.label }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { useIntersectionObserver } from "@vueuse/core"
import { CheckCircle } from "lucide-vue-next"
import { reactive, ref } from "vue"

const statsContainerRef = ref(null)

const stats = reactive([
  { target: 15, current: 0, suffix: "年", label: "行业经验" },
  { target: 1500, current: 0, suffix: "m²", label: "制造基地" },
  { target: 100, current: 0, suffix: "+", label: "客户群体" },
  { target: 500, current: 0, suffix: "+", label: "累计发货" },
  { target: 5, current: 0, suffix: "+", label: "每月发货非标设备" },
])

function AnimateNumbers() {
  const duration = 2000
  const startTime = performance.now()

  function Step(currentTime: number) {
    const elapsedTime = currentTime - startTime
    const progress = Math.min(elapsedTime / duration, 1)

    const easeOutProgress = progress === 1 ? 1 : 1 - Math.pow(2, -10 * progress)

    stats.forEach((stat) => {
      stat.current = Math.floor(easeOutProgress * stat.target)
    })

    if (progress < 1) {
      requestAnimationFrame(Step)
    } else {
      stats.forEach((stat) => {
        stat.current = stat.target
      })
    }
  }

  requestAnimationFrame(Step)
}

const { stop } = useIntersectionObserver(
  statsContainerRef,
  (entries) => {
    const entry = entries[0]
    if (entry?.isIntersecting) {
      AnimateNumbers()
      stop()
    }
  },
  { threshold: 0.2 }
)
</script>
