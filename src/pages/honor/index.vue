<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <!-- 头部 Banner -->
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <!-- 背景装饰 -->
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>

      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('./banner/about.webp')] bg-cover bg-center"
      ></div>

      <div class="relative z-20 px-4 text-center">
        <h1
          class="mb-4 text-4xl font-bold tracking-wider text-white drop-shadow-md md:text-5xl"
        >
          荣誉与文化
        </h1>
      </div>
    </div>

    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <!-- 面包屑导航 -->
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">荣誉与文化</span>
      </nav>

      <div class="grid grid-cols-1 gap-10 lg:grid-cols-4">
        <!-- 左侧侧边栏 -->
        <aside class="space-y-8 lg:col-span-1">
          <div
            class="rounded-2xl border border-gray-100 bg-white p-2 shadow-sm"
          >
            <ul class="space-y-1">
              <li v-for="cat in categories" :key="cat.id">
                <button
                  @click="activeCategory = cat.id"
                  class="group relative flex w-full items-center overflow-hidden rounded-xl px-5 py-4 transition-all duration-200"
                  :class="[
                    activeCategory === cat.id
                      ? 'bg-blue-600 text-white shadow-md shadow-blue-200'
                      : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
                  ]"
                >
                  <component
                    :is="cat.icon"
                    class="z-10 mr-3 h-5 w-5 transition-transform duration-300"
                    :class="
                      activeCategory === cat.id
                        ? 'scale-110'
                        : 'group-hover:scale-110'
                    "
                  />
                  <span class="z-10 font-medium">{{ cat.name }}</span>

                  <div
                    v-if="activeCategory === cat.id"
                    class="pointer-events-none absolute top-0 right-0 h-full w-20 bg-linear-to-l from-white/10 to-transparent"
                  ></div>
                </button>
              </li>
            </ul>
          </div>

          <!-- 联系方式 -->
          <div
            class="rounded-2xl border border-blue-100 bg-linear-to-br from-blue-50 to-white p-6 shadow-lg shadow-blue-900/5"
          >
            <h3 class="mb-4 flex items-center font-bold text-blue-800">
              联系我们
              <span
                class="ml-2 h-2 w-2 animate-pulse rounded-full bg-green-400"
              ></span>
            </h3>
            <div class="space-y-4 text-sm">
              <div class="flex items-center text-slate-700">
                <div
                  class="mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
                >
                  <Phone class="h-4 w-4 text-blue-600" />
                </div>
                <span class="font-medium">0592-6818878</span>
              </div>
              <div class="flex items-center text-slate-700">
                <div
                  class="mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
                >
                  <Mail class="h-4 w-4 text-blue-600" />
                </div>
                <span class="font-medium">1888@xingshi.com</span>
              </div>
              <div class="flex items-center text-slate-700">
                <div
                  class="mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
                >
                  <MapPin class="h-4 w-4 text-blue-600" />
                </div>
                <span>集美区软件园</span>
              </div>
            </div>
          </div>
        </aside>

        <!-- 右侧内容区 -->
        <main class="lg:col-span-3">
          <div
            class="mb-8 flex items-center justify-between border-b border-gray-200 pb-4"
          >
            <h2 class="relative pl-4 text-2xl font-bold text-slate-800">
              <span
                class="absolute top-1/2 left-0 h-6 w-1 -translate-y-1/2 rounded-full bg-blue-600"
              ></span>
              {{ currentCategoryName }}
            </h2>

            <div
              v-if="['patent', 'honor'].includes(activeCategory)"
              class="relative hidden sm:block"
            >
              <input
                type="text"
                placeholder="搜索证书名称..."
                class="w-64 rounded-lg border border-gray-200 bg-white py-2 pr-4 pl-9 text-sm transition-all focus:border-blue-500 focus:ring-2 focus:ring-blue-100"
              />
              <Search
                class="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-gray-400"
              />
            </div>
          </div>

          <div
            v-if="['honor', 'patent'].includes(activeCategory)"
            class="space-y-12"
          >
            <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
              <div
                v-for="(item, index) in currentData"
                :key="item.id"
                class="group relative flex flex-col overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm transition-all duration-300 hover:border-blue-300 hover:shadow-xl hover:shadow-blue-900/5"
                :class="[
                  isLoaded
                    ? 'translate-y-0 opacity-100'
                    : 'translate-y-8 opacity-0',
                ]"
                :style="{ transitionDelay: `${index * 80}ms` }"
              >
                <div class="h-1 w-full bg-blue-600"></div>

                <div class="border-b border-gray-50 bg-gray-50/30 p-5">
                  <h3
                    class="line-clamp-2 flex min-h-14 items-center justify-center text-center text-lg leading-snug font-bold text-slate-800 transition-colors group-hover:text-blue-600"
                  >
                    {{ item.title }}
                  </h3>
                </div>

                <div
                  class="relative flex h-56 items-center justify-center overflow-hidden bg-white p-4 transition-colors group-hover:bg-gray-50"
                >
                  <img
                    :src="item.image"
                    :alt="item.title"
                    class="max-h-full max-w-full object-contain transition-all duration-500 group-hover:scale-105"
                  />

                  <div
                    class="pointer-events-none absolute inset-0 flex items-center justify-center opacity-0 transition-opacity duration-300 group-hover:opacity-100"
                  >
                    <div
                      class="flex h-10 w-10 translate-y-4 transform items-center justify-center rounded-full bg-blue-600/90 text-white shadow-lg backdrop-blur-sm transition-transform duration-300 group-hover:translate-y-0"
                    >
                      <ZoomIn class="h-5 w-5" />
                    </div>
                  </div>

                  <div
                    v-if="item.type"
                    class="absolute top-2 right-2 rounded border border-slate-200 bg-slate-100 px-2 py-0.5 text-[10px] text-slate-500"
                  >
                    {{ item.type }}
                  </div>
                </div>

                <div
                  class="relative mt-auto space-y-2 overflow-hidden border-t border-gray-100 bg-slate-50 p-4"
                >
                  <div
                    class="absolute right-0 bottom-0 -mr-4 -mb-4 h-16 w-16 rounded-full bg-blue-100/30 blur-xl"
                  ></div>

                  <div class="relative z-10" v-if="item.certNo">
                    <p class="text-xs font-medium text-slate-500">
                      {{ activeCategory === "patent" ? "专利号" : "证书编号" }}
                    </p>
                    <p class="font-mono text-sm font-bold text-slate-700">
                      {{ item.certNo }}
                    </p>
                  </div>

                  <div
                    class="relative z-10 flex items-center justify-between pt-1"
                  >
                    <div>
                      <p class="text-xs font-medium text-slate-500">
                        {{
                          activeCategory === "patent" ? "申请日期" : "颁发日期"
                        }}
                      </p>
                      <p class="text-xs text-slate-700">{{ item.date }}</p>
                    </div>
                    <div class="text-blue-200">
                      <Award class="h-6 w-6 opacity-20" />
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="mt-12 flex justify-center">
              <nav class="flex items-center gap-2">
                <button
                  class="flex h-10 w-10 items-center justify-center rounded-lg text-slate-400 transition-colors hover:bg-gray-100"
                >
                  <ChevronRight class="h-5 w-5 rotate-180" />
                </button>
                <button
                  class="h-10 w-10 rounded-lg bg-blue-600 font-medium text-white shadow-lg shadow-blue-200"
                >
                  1
                </button>
                <button
                  class="h-10 w-10 rounded-lg font-medium text-slate-600 transition-colors hover:bg-gray-100"
                >
                  2
                </button>
                <button
                  class="flex h-10 w-10 items-center justify-center rounded-lg text-slate-400 transition-colors hover:bg-gray-100 hover:text-blue-600"
                >
                  <ChevronRight class="h-5 w-5" />
                </button>
              </nav>
            </div>
          </div>

          <div
            v-else
            class="flex flex-col items-center justify-center rounded-2xl border border-dashed border-gray-200 bg-white py-20"
          >
            <div
              class="mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-blue-50 text-blue-400"
            >
              <component
                :is="categories.find((c) => c.id === activeCategory)?.icon"
                class="h-8 w-8"
              />
            </div>
            <h3 class="text-lg font-bold text-slate-400">
              {{ currentCategoryName }} 内容正在建设中...
            </h3>
            <p class="mt-2 text-sm text-slate-400">请稍后访问</p>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Award,
  ChevronRight,
  FileCheck,
  Heart,
  Mail,
  MapPin,
  Phone,
  Search,
  Users,
  ZoomIn,
} from "lucide-vue-next"
import { computed, onMounted, ref } from "vue"

const categories = [
  { id: "honor", name: "公司荣誉", icon: Award },
  { id: "activity", name: "爱心活动", icon: Heart },
  { id: "patent", name: "公司专利", icon: FileCheck },
  { id: "culture", name: "企业文化", icon: Users },
]

const activeCategory = ref("honor")

interface CertItem {
  id: number
  title: string
  image: string
  certNo?: string
  date: string
  type?: string
}

// 荣誉数据
const honors: CertItem[] = [
  {
    id: 1,
    title: "国家级高新技术企业",
    image:
      "https://images.unsplash.com/photo-1635350736475-c8cef4b21906?q=80&w=1000&auto=format&fit=crop",
    certNo: "GR202435001234",
    date: "2024年12月",
    type: "国家级",
  },
  {
    id: 2,
    title: "福建省科技小巨人领军企业",
    image:
      "https://images.unsplash.com/photo-1579548122080-c35fd6820ecb?q=80&w=1000&auto=format&fit=crop",
    certNo: "FJ-LJQY-2024-056",
    date: "2024年08月",
    type: "省级",
  },
  {
    id: 3,
    title: "ISO9001 质量管理体系认证",
    image:
      "https://images.unsplash.com/photo-1607619056574-7b8d3ee536b2?q=80&w=1000&auto=format&fit=crop",
    certNo: "ISO-2023-CN-8899",
    date: "2023年05月",
    type: "国际认证",
  },
  {
    id: 4,
    title: "厦门市“专精特新”中小企业",
    image:
      "https://images.unsplash.com/photo-1577401239170-897942555fb3?q=80&w=1000&auto=format&fit=crop",
    certNo: "XM-ZJTX-2023-102",
    date: "2023年11月",
    type: "市级",
  },
]

// 专利数据
const patents: CertItem[] = [
  {
    id: 1,
    title: "一种智能设备的节能控制方法",
    image:
      "https://images.unsplash.com/photo-1581092580497-e0d23cbdf1dc?q=80&w=1000&auto=format&fit=crop",
    certNo: "ZL202410123456.7",
    date: "2024年10月20日",
    type: "发明专利",
  },
  {
    id: 2,
    title: "自动化产线视觉检测装置",
    image:
      "https://images.unsplash.com/photo-1531746790731-6c087fecd65a?q=80&w=1000&auto=format&fit=crop",
    certNo: "ZL202420556789.X",
    date: "2024年09月15日",
    type: "实用新型",
  },
  {
    id: 3,
    title: "基于AI的激光打标定位算法",
    image:
      "https://images.unsplash.com/photo-1555255707-c07966088b7b?q=80&w=1000&auto=format&fit=crop",
    certNo: "ZL202410998877.1",
    date: "2024年08月01日",
    type: "软件著作权",
  },
]

const isLoaded = ref(false)

onMounted(() => {
  setTimeout(() => {
    isLoaded.value = true
  }, 100)
})

const currentCategoryName = computed(() => {
  return categories.find((c) => c.id === activeCategory.value)?.name || ""
})

// 根据当前分类返回对应数据
const currentData = computed(() => {
  if (activeCategory.value === "honor") return honors
  if (activeCategory.value === "patent") return patents
  return []
})
</script>

<style scoped>
@keyframes pulse-slow {
  0%,
  100% {
    opacity: 0.8;
  }
  50% {
    opacity: 0.4;
  }
}
.animate-pulse-slow {
  animation: pulse-slow 8s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>
