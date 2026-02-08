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
              关于<span
                class="bg-linear-to-r from-blue-600 to-cyan-500 bg-clip-text text-transparent"
                >我们</span
              >
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
              class="group flex items-center gap-2 rounded-full bg-blue-600 px-8 py-3.5 font-medium text-white shadow-lg transition-all duration-300 hover:bg-blue-700 hover:shadow-blue-500/30"
            >
              了解更多
              <ArrowRight
                class="h-5 w-5 transition-transform group-hover:translate-x-1"
              />
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

      <div class="mt-20">
        <div class="grid grid-cols-2 gap-8 md:grid-cols-3 lg:grid-cols-5">
          <div
            v-for="(item, index) in stats"
            :key="index"
            class="group rounded-xl border border-slate-100 bg-white p-6 text-center shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-xl"
          >
            <div
              class="mb-2 inline-block origin-center text-3xl font-bold text-blue-600 transition-transform duration-300 group-hover:scale-110 md:text-4xl"
            >
              {{ item.value }}
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
import { ArrowRight, CheckCircle } from "lucide-vue-next"

interface StatItem {
  value: string
  label: string
}

const stats: StatItem[] = [
  { value: "15年", label: "行业经验" },
  { value: "1500m²", label: "制造基地" },
  { value: "100+", label: "客户群体" },
  { value: "500+", label: "累计发货" },
  { value: "5+", label: "每月发货非标设备" },
]
</script>
