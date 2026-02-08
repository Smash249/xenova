<script setup lang="ts">
import { ArrowRight } from "lucide-vue-next"

// 模拟合作伙伴 Logo 数据
const partners = [
  {
    name: "宏发股份",
    logo: "https://placehold.co/200x80/white/red?text=HONGFA",
  },
  {
    name: "宁德时代",
    logo: "https://placehold.co/200x80/white/blue?text=CATL",
  },
  { name: "比亚迪", logo: "https://placehold.co/200x80/white/black?text=BYD" },
  {
    name: "汇川技术",
    logo: "https://placehold.co/200x80/white/green?text=INOVANCE",
  },
  {
    name: "立讯精密",
    logo: "https://placehold.co/200x80/white/blue?text=LUXSHARE",
  },
  { name: "大族激光", logo: "https://placehold.co/200x80/white/red?text=HANS" },
]

// 左侧展示图片
const showcaseImages = [
  "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?auto=format&fit=crop&q=80&w=600",
  "https://images.unsplash.com/photo-1565514020176-db710609b552?auto=format&fit=crop&q=80&w=600",
  "https://images.unsplash.com/photo-1531297461136-82lw9b44d940?auto=format&fit=crop&q=80&w=600",
  "https://images.unsplash.com/photo-1504917595217-d4dc5ebe6122?auto=format&fit=crop&q=80&w=600",
]
</script>

<template>
  <section
    id="partners"
    class="relative w-full overflow-hidden bg-slate-50 py-24 font-sans"
  >
    <!-- 背景点阵纹理 -->
    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>

    <div class="relative z-10 container mx-auto px-6">
      <div class="mb-16 text-center">
        <h2 class="mb-4 text-4xl font-bold tracking-tight text-slate-900">
          长期合作伙伴
        </h2>
        <p class="mx-auto max-w-2xl text-lg text-slate-600">
          长为客户持续提供坚实稳定的设备，助力制造业高质量发展
        </p>
      </div>

      <div class="grid grid-cols-1 gap-8 lg:grid-cols-12">
        <!-- 左侧：图片展示墙 -->
        <div class="grid grid-cols-2 gap-4 lg:col-span-5">
          <div
            v-for="(img, index) in showcaseImages"
            :key="index"
            class="group relative aspect-video w-full overflow-hidden rounded-xl bg-white shadow-sm ring-1 ring-slate-900/5 lg:aspect-auto lg:h-48"
          >
            <div
              class="absolute inset-0 z-10 bg-slate-900/0 transition-colors duration-300 group-hover:bg-slate-900/10"
            ></div>
            <img
              :src="img"
              alt="Showcase"
              class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-110"
            />
          </div>
        </div>

        <!-- 右侧：合作伙伴列表 (添加磨砂玻璃背景) -->
        <div
          class="flex flex-col justify-between gap-3 rounded-2xl border border-white/60 bg-white/40 p-6 shadow-sm backdrop-blur-md lg:col-span-7 lg:p-8"
        >
          <div class="grid grid-cols-2 gap-4 sm:grid-cols-3">
            <div
              v-for="(partner, index) in partners"
              :key="index"
              class="group flex aspect-3/2 cursor-pointer items-center justify-center rounded-xl border border-white/50 bg-white/80 p-6 shadow-sm transition-all duration-300 hover:-translate-y-1 hover:border-blue-200 hover:bg-white hover:shadow-lg hover:shadow-blue-500/5"
            >
              <img
                :src="partner.logo"
                :alt="partner.name"
                class="max-h-[80%] max-w-[80%] object-contain opacity-60 grayscale filter transition-all duration-300 group-hover:opacity-100 group-hover:grayscale-0"
              />
            </div>
          </div>

          <!-- 底部操作栏 (添加按钮背景样式) -->
          <div class="mt-6 flex justify-end">
            <button
              class="group flex items-center gap-2 rounded-full bg-blue-600 px-6 py-2.5 text-sm font-semibold text-white shadow-md shadow-blue-600/20 transition-all hover:bg-blue-700 hover:shadow-lg hover:shadow-blue-600/30 active:scale-95"
            >
              <span>查看更多伙伴</span>
              <ArrowRight
                class="h-4 w-4 transition-transform group-hover:translate-x-1"
              />
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
