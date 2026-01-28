<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-linear-to-r from-blue-900 to-slate-800"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>
      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('./banner/news.webp')] bg-cover bg-center mix-blend-overlay"
      ></div>
      <h1
        class="relative z-10 text-4xl font-bold tracking-wide text-white md:text-5xl"
      >
        企业动态
      </h1>
    </div>

    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <!-- 面包屑导航 -->
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a class="transition-colors hover:text-blue-600" href="/">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">企业动态</span>
      </nav>

      <div class="grid grid-cols-1 gap-10 lg:grid-cols-4">
        <!-- 左侧侧边栏 -->
        <aside class="space-y-8 lg:col-span-1">
          <div
            class="rounded-2xl border border-gray-100 bg-white p-6 shadow-sm"
          >
            <h2 class="mb-6 flex items-center text-xl font-bold text-slate-900">
              <span class="mr-3 h-6 w-1 rounded-full bg-blue-600"></span>
              新闻分类
            </h2>
            <ul class="space-y-2">
              <li v-for="(cat, index) in categories" :key="index">
                <a
                  :class="[
                    cat.active
                      ? 'bg-blue-50 font-semibold text-blue-700 shadow-sm'
                      : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
                  ]"
                  class="group block rounded-xl px-4 py-3 transition-all duration-200"
                  href="#"
                >
                  <div class="flex items-center justify-between">
                    {{ cat.name }}
                    <ChevronRight
                      :class="
                        cat.active
                          ? 'opacity-100'
                          : 'opacity-0 group-hover:opacity-100'
                      "
                      class="h-4 w-4 transition-transform"
                    />
                  </div>
                </a>
              </li>
            </ul>
          </div>

          <!-- 联系方式卡片 (已更新为亮色风格) -->
          <div
            class="rounded-2xl border border-blue-100 bg-linear-to-br from-blue-50 to-white p-6 shadow-lg shadow-blue-900/5"
          >
            <h2 class="mb-6 flex items-center text-xl font-bold text-blue-800">
              联系我们
              <span
                class="ml-2 h-2 w-2 animate-pulse rounded-full bg-green-400"
              ></span>
            </h2>
            <div class="space-y-5">
              <div class="flex items-start text-slate-700">
                <div
                  class="mt-1 mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
                >
                  <MapPin class="h-4 w-4 text-blue-600" />
                </div>
                <div>
                  <h3
                    class="mb-1 text-xs tracking-wider text-slate-500 uppercase"
                  >
                    公司地址
                  </h3>
                  <p class="text-sm font-medium">星实（厦门）科技有限公司</p>
                  <p class="mt-1 text-sm text-slate-500">集美区软件园</p>
                </div>
              </div>

              <div class="flex items-start text-slate-700">
                <div
                  class="mt-1 mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
                >
                  <Phone class="h-4 w-4 text-blue-600" />
                </div>
                <div>
                  <h3
                    class="mb-1 text-xs tracking-wider text-slate-500 uppercase"
                  >
                    咨询热线
                  </h3>
                  <p class="text-lg font-bold tracking-wide text-blue-900">
                    0592-6818878
                  </p>
                </div>
              </div>

              <div class="flex items-start text-slate-700">
                <div
                  class="mt-1 mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
                >
                  <Mail class="h-4 w-4 text-blue-600" />
                </div>
                <div>
                  <h3
                    class="mb-1 text-xs tracking-wider text-slate-500 uppercase"
                  >
                    电子邮箱
                  </h3>
                  <p class="text-sm font-medium break-all">1888@xingshi.com</p>
                </div>
              </div>
            </div>
          </div>
        </aside>

        <!-- 右侧新闻列表 -->
        <main class="lg:col-span-3">
          <div class="space-y-6">
            <article
              v-for="(news, index) in newsData"
              :key="news.id"
              class="group rounded-2xl border border-gray-100 bg-white p-6 shadow-sm transition-all duration-500 hover:-translate-y-1 hover:shadow-xl hover:shadow-blue-900/5 md:p-8"
              :class="[
                isLoaded
                  ? 'translate-y-0 opacity-100'
                  : 'translate-y-8 opacity-0',
              ]"
              :style="{ transitionDelay: `${index * 100}ms` }"
            >
              <div
                class="mb-4 flex flex-col justify-between gap-4 md:flex-row md:items-start"
              >
                <div class="flex-1">
                  <div class="mb-3 flex items-center gap-3">
                    <span
                      class="rounded-full bg-blue-50 px-3 py-1 text-xs font-semibold text-blue-600"
                    >
                      {{ news.category }}
                    </span>
                    <div class="flex items-center text-xs text-slate-400">
                      <Calendar class="mr-1 h-3.5 w-3.5" />
                      {{ news.date }}
                    </div>
                  </div>
                  <h3
                    class="mb-3 line-clamp-2 text-xl font-bold text-slate-800 transition-colors group-hover:text-blue-700 md:text-2xl"
                  >
                    <a :href="news.link">{{ news.title }}</a>
                  </h3>
                </div>
              </div>

              <p class="mb-6 line-clamp-3 leading-relaxed text-slate-600">
                {{ news.summary }}
              </p>

              <div
                class="flex items-center justify-between border-t border-gray-50 pt-6"
              >
                <div class="flex items-center text-sm text-slate-400">
                  <Eye class="mr-2 h-4 w-4" />
                  <span>{{ news.views }} 次浏览</span>
                </div>

                <a
                  :href="news.link"
                  class="group/btn inline-flex items-center justify-center rounded-lg bg-blue-600 px-6 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700"
                >
                  阅读更多
                  <ChevronRight
                    class="ml-1 h-4 w-4 transition-transform group-hover/btn:translate-x-1"
                  />
                </a>
              </div>
            </article>
          </div>

          <!-- 分页 -->
          <div class="mt-12 flex justify-center">
            <nav class="flex items-center gap-2">
              <button
                class="rounded-lg p-2 text-slate-400 transition-colors hover:bg-blue-50 hover:text-blue-600 disabled:opacity-50"
              >
                Previous
              </button>
              <button
                class="h-10 w-10 rounded-lg bg-blue-600 font-medium text-white shadow-md shadow-blue-200"
              >
                1
              </button>
              <button
                class="h-10 w-10 rounded-lg font-medium text-slate-600 transition-colors hover:bg-gray-100 hover:text-blue-600"
              >
                2
              </button>
              <button
                class="h-10 w-10 rounded-lg font-medium text-slate-600 transition-colors hover:bg-gray-100 hover:text-blue-600"
              >
                3
              </button>
              <span class="px-2 text-slate-400">...</span>
              <button
                class="rounded-lg p-2 text-slate-400 transition-colors hover:bg-blue-50 hover:text-blue-600"
              >
                Next
              </button>
            </nav>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {
  Calendar,
  ChevronRight,
  Eye,
  Mail,
  MapPin,
  Phone,
} from "lucide-vue-next"
import { onMounted, ref } from "vue"

interface NewsItem {
  id: number
  title: string
  date: string
  summary: string
  views: number
  category: string
  link: string
}

interface Category {
  name: string
  active: boolean
}

const newsData: NewsItem[] = [
  {
    id: 1,
    title: "星实的视觉打标机与其他行业内有什么区别?",
    date: "2025-06-18",
    summary:
      "星实制造的激光器设备在行业上首次采用了AI视觉辅助功能，皆知无需人工干预即可实现自动打标。",
    views: 1800,
    category: "公司讯息",
    link: "#",
  },
  {
    id: 2,
    title: "装盒机厂家与智能包装设备的未来趋势",
    date: "2025-06-18",
    summary:
      "装盒机厂家与其他包装设备在功能、适用范围、操作便捷性、智能化程度以及设计细节等方面存在显著的区别。以下是对这些区别的详细分析...",
    views: 1650,
    category: "行业讯息",
    link: "#",
  },
  {
    id: 3,
    title: "自动化产线升：如何提升生产效率",
    date: "2025-06-18",
    summary:
      "深入探讨自动化产线在现代制造业中的应用，以及如何通过技术升级来降���成本并大幅提升生产效率。",
    views: 1520,
    category: "技术讯息",
    link: "#",
  },
]

const categories: Category[] = [
  { name: "公司讯息", active: true },
  { name: "行业讯息", active: false },
  { name: "技术讯息", active: false },
]

const isLoaded = ref(false)

onMounted(() => {
  setTimeout(() => {
    isLoaded.value = true
  }, 100)
})
</script>

<style scoped></style>
