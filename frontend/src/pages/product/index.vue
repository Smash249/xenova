<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>
      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('./banner/product.webp')] bg-cover bg-center"
      ></div>

      <div class="relative z-20 px-4 text-center">
        <h1
          class="mb-4 text-4xl font-bold tracking-wider text-white drop-shadow-md md:text-5xl"
        >
          产品中心
        </h1>
      </div>
    </div>

    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <!-- 面包屑导航 -->
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">产品中心</span>
      </nav>

      <div class="grid grid-cols-1 gap-10 lg:grid-cols-4">
        <!-- 左侧侧边栏 -->
        <aside class="space-y-8 lg:col-span-1">
          <!-- 产品搜索 -->
          <div
            class="flex items-center rounded-2xl border border-gray-100 bg-white p-4 shadow-sm"
          >
            <Search class="mr-3 h-5 w-5 text-slate-400" />
            <input
              type="text"
              placeholder="搜索产品型号..."
              class="w-full border-none bg-transparent text-sm text-slate-700 placeholder-slate-400 focus:ring-0"
            />
          </div>

          <!-- 产品分类 -->
          <div
            class="overflow-hidden rounded-2xl border border-gray-100 bg-white shadow-sm"
          >
            <div
              class="flex items-center justify-between border-b border-gray-50 bg-gray-50/50 p-5"
            >
              <h2 class="flex items-center font-bold text-slate-900">
                <Filter class="mr-2 h-4 w-4 text-blue-600" />
                产品系列
              </h2>
            </div>
            <ul class="space-y-1 p-3">
              <li v-for="cat in categories" :key="cat.id">
                <button
                  @click="selectCategory(cat.id)"
                  class="group flex w-full items-center justify-between rounded-xl px-4 py-3 text-left transition-all duration-200"
                  :class="[
                    cat.active
                      ? 'bg-blue-600 text-white shadow-md shadow-blue-200'
                      : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
                  ]"
                >
                  <span class="font-medium">{{ cat.name }}</span>
                  <span
                    class="rounded-full px-2 py-0.5 text-xs transition-colors"
                    :class="
                      cat.active
                        ? 'bg-white/20 text-white'
                        : 'bg-gray-100 text-slate-400'
                    "
                  >
                    {{ cat.count }}
                  </span>
                </button>
              </li>
            </ul>
          </div>

          <!-- 联系方式-->
          <div
            class="rounded-2xl border border-blue-100 bg-linear-to-br from-blue-50 to-white p-6 shadow-lg shadow-blue-900/5"
          >
            <h3 class="mb-4 flex items-center font-bold text-blue-800">
              联系销售顾问
              <span
                class="ml-2 h-2 w-2 animate-pulse rounded-full bg-green-400"
              ></span>
            </h3>
            <p class="mb-6 text-sm leading-relaxed text-slate-600">
              需要定制方案或获取报价？我们的工程师随时为您服务。
            </p>
            <div class="space-y-4 text-sm">
              <div
                class="group flex cursor-pointer items-center text-slate-700"
              >
                <div
                  class="mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100 transition-colors group-hover:bg-blue-600"
                >
                  <Phone
                    class="h-4 w-4 text-blue-600 transition-colors group-hover:text-white"
                  />
                </div>
                <span
                  class="font-medium transition-colors group-hover:text-blue-600"
                  >0592-6818878</span
                >
              </div>
              <div
                class="group flex cursor-pointer items-center text-slate-700"
              >
                <div
                  class="mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100 transition-colors group-hover:bg-blue-600"
                >
                  <Mail
                    class="h-4 w-4 text-blue-600 transition-colors group-hover:text-white"
                  />
                </div>
                <span
                  class="font-medium transition-colors group-hover:text-blue-600"
                  >1888@xingshi.com</span
                >
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

        <!-- 右侧产品网格 -->
        <main class="lg:col-span-3">
          <!-- 顶部工具栏 -->
          <div
            class="mb-8 flex flex-col items-center justify-between border-b border-gray-200 pb-4 sm:flex-row"
          >
            <h2 class="mb-4 text-2xl font-bold text-slate-800 sm:mb-0">
              全部产品
              <span class="ml-2 text-base font-normal text-slate-400"
                >展示 6 款设备</span
              >
            </h2>
            <div class="flex items-center space-x-2">
              <span class="text-sm text-slate-500">排序:</span>
              <select
                class="cursor-pointer border-none bg-transparent text-sm font-medium text-slate-700 focus:ring-0"
              >
                <option>推荐排序</option>
                <option>最新发布</option>
              </select>
            </div>
          </div>

          <!-- 产品卡片网格 -->
          <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
            <div
              v-for="(product, index) in products"
              :key="product.id"
              class="group flex h-full transform cursor-pointer flex-col overflow-hidden rounded-2xl border border-gray-100 bg-white transition-all duration-500 hover:shadow-2xl hover:shadow-blue-900/10"
              :class="[
                isLoaded
                  ? 'translate-y-0 opacity-100'
                  : 'translate-y-8 opacity-0',
              ]"
              :style="{ transitionDelay: `${index * 100}ms` }"
            >
              <!-- 图片区域 -->
              <div class="relative h-56 overflow-hidden bg-gray-100">
                <div class="absolute top-4 left-4 z-10">
                  <span
                    class="rounded-full bg-white/90 px-3 py-1 text-xs font-bold text-blue-800 shadow-sm backdrop-blur-sm"
                  >
                    {{ product.category }}
                  </span>
                </div>

                <img
                  :src="product.image"
                  :alt="product.name"
                  class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-110"
                />

                <div
                  class="absolute inset-0 bg-blue-900/0 transition-colors duration-300 group-hover:bg-blue-900/10"
                ></div>
              </div>

              <!-- 内容区域 -->
              <div class="flex flex-1 flex-col p-6">
                <h3
                  class="mb-2 text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-600"
                >
                  {{ product.name }}
                </h3>
                <p class="mb-6 line-clamp-2 flex-1 text-sm text-slate-500">
                  {{ product.description }}
                </p>

                <div class="mt-auto flex items-center justify-between">
                  <span
                    class="text-xs font-semibold text-slate-400 transition-colors group-hover:text-blue-500"
                  >
                    查看详情
                  </span>
                  <div
                    class="flex h-8 w-8 transform items-center justify-center rounded-full bg-gray-50 text-slate-400 transition-all duration-300 group-hover:-rotate-45 group-hover:bg-blue-600 group-hover:text-white"
                  >
                    <ArrowRight class="h-4 w-4" />
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 分页 -->
          <div class="mt-16 flex justify-center">
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
        </main>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  ArrowRight,
  ChevronRight,
  Filter,
  Mail,
  MapPin,
  Phone,
  Search,
} from "lucide-vue-next"
import { onMounted, ref } from "vue"

interface Product {
  id: number
  name: string
  category: string
  image: string
  description: string
}

interface Category {
  id: string
  name: string
  count: number
  active: boolean
}

const categories = ref<Category[]>([
  { id: "all", name: "全部产品", count: 12, active: true },
  { id: "vision", name: "视觉激光打标机", count: 4, active: false },
  { id: "robot", name: "协作机械手", count: 3, active: false },
  { id: "test", name: "气密测试仪", count: 2, active: false },
  { id: "box", name: "智能装盒机", count: 3, active: false },
])

const products: Product[] = [
  {
    id: 1,
    name: "XS-2000 智能装盒机",
    category: "智能装盒机",
    image:
      "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?q=80&w=1000&auto=format&fit=crop",
    description: "全自动高速装盒，支持多种尺寸规格切换，配备智能视觉检测系统。",
  },
  {
    id: 2,
    name: "Vision-X 激光打标机",
    category: "视觉激光打标机",
    image:
      "https://images.unsplash.com/photo-1616401784845-180882ba9ba8?q=80&w=1000&auto=format&fit=crop",
    description: "高精度UV激光，搭载AI视觉定位，实现微米级打标精度。",
  },
  {
    id: 3,
    name: "Robo-Arm S3 协作臂",
    category: "协作机械手",
    image:
      "https://images.unsplash.com/photo-1531746790731-6c087fecd65a?q=80&w=1000&auto=format&fit=crop",
    description: "6轴自由度，安全碰撞检测，适用于精密组装与搬运场景。",
  },
  {
    id: 4,
    name: "AirTest Pro 气密测试仪",
    category: "气密测试仪",
    image:
      "https://images.unsplash.com/photo-1581092160562-40aa08e78837?q=80&w=1000&auto=format&fit=crop",
    description: "差压式检漏，支持多通道并行测试，数据实时上传云端。",
  },
  {
    id: 5,
    name: "XS-3000 高速装箱线",
    category: "智能装盒机",
    image:
      "https://images.unsplash.com/photo-1504917595217-d4dc5ebe6122?q=80&w=1000&auto=format&fit=crop",
    description: "一体化开箱、装箱、封箱解决方案，产能高达1200箱/小时。",
  },
  {
    id: 6,
    name: "Vision-Mini 便携打标",
    category: "视觉激光打标机",
    image:
      "https://images.unsplash.com/photo-1563770095-39d468f9a51d?q=80&w=1000&auto=format&fit=crop",
    description: "手持式设计，内置电池，适用于大型设备无法移动的场景。",
  },
]

const isLoaded = ref(false)

onMounted(() => {
  setTimeout(() => {
    isLoaded.value = true
  }, 100)
})

const selectCategory = (id: string) => {
  categories.value.forEach((c) => (c.active = c.id === id))
}
</script>

<style scoped></style>
