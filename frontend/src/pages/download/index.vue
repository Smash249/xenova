<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <!-- 头部 Banner -->
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>

      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('./banner/download.webp')] bg-cover bg-center"
      ></div>

      <div class="relative z-20 px-4 text-center">
        <h1
          class="mb-4 text-4xl font-bold tracking-wider text-white drop-shadow-md md:text-5xl"
        >
          软件中心
        </h1>
      </div>
    </div>

    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">软件中心</span>
      </nav>

      <div
        class="mb-10 flex flex-col items-center justify-between gap-6 md:flex-row"
      >
        <div
          class="flex max-w-full overflow-x-auto rounded-xl border border-gray-100 bg-white p-1 shadow-sm"
        >
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id as any"
            class="rounded-lg px-6 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-300"
            :class="[
              activeTab === tab.id
                ? 'bg-blue-600 text-white shadow-md'
                : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
            ]"
          >
            {{ tab.label }}
          </button>
        </div>

        <div class="relative w-full md:w-80">
          <input
            type="text"
            placeholder="搜索软件名称或版本..."
            class="w-full rounded-xl border border-gray-200 bg-white py-3 pr-4 pl-10 text-sm shadow-sm transition-all focus:border-blue-500 focus:ring-2 focus:ring-blue-100"
          />
          <Search
            class="absolute top-1/2 left-3.5 h-4 w-4 -translate-y-1/2 text-slate-400"
          />
        </div>
      </div>

      <div class="space-y-4">
        <div
          v-for="(item, index) in filteredDownloads"
          :key="item.id"
          class="group transform rounded-2xl border border-gray-100 bg-white p-6 shadow-sm transition-all duration-300 hover:border-blue-100 hover:shadow-xl"
          :class="[
            isLoaded ? 'translate-y-0 opacity-100' : 'translate-y-8 opacity-0',
          ]"
          :style="{ transitionDelay: `${index * 80}ms` }"
        >
          <div
            class="flex flex-col justify-between gap-6 md:flex-row md:items-center"
          >
            <div class="flex items-start gap-5">
              <div
                class="flex h-14 w-14 shrink-0 items-center justify-center rounded-2xl transition-colors duration-300"
                :class="[
                  item.type === 'software'
                    ? 'bg-blue-50 text-blue-600 group-hover:bg-blue-600 group-hover:text-white'
                    : item.type === 'driver'
                      ? 'bg-indigo-50 text-indigo-600 group-hover:bg-indigo-600 group-hover:text-white'
                      : 'bg-orange-50 text-orange-600 group-hover:bg-orange-600 group-hover:text-white',
                ]"
              >
                <component :is="getIcon(item.type)" class="h-7 w-7" />
              </div>

              <div>
                <div class="mb-1 flex items-center gap-3">
                  <h3
                    class="text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-600"
                  >
                    {{ item.title }}
                  </h3>

                  <span
                    v-if="item.isPopular"
                    class="rounded-md border border-red-100 bg-red-50 px-2 py-0.5 text-[10px] font-bold tracking-wider text-red-600 uppercase"
                  >
                    Hot
                  </span>
                </div>

                <p
                  class="mb-3 max-w-2xl text-sm leading-relaxed text-slate-500"
                >
                  {{ item.description }}
                </p>

                <div
                  class="flex items-center gap-4 text-xs font-medium text-slate-400"
                >
                  <span class="flex items-center rounded bg-gray-50 px-2 py-1">
                    <Settings class="mr-1.5 h-3 w-3" /> {{ item.version }}
                  </span>
                  <span class="flex items-center rounded bg-gray-50 px-2 py-1">
                    <HardDrive class="mr-1.5 h-3 w-3" /> {{ item.size }}
                  </span>
                  <span class="flex items-center">
                    更新于: {{ item.date }}
                  </span>
                </div>
              </div>
            </div>

            <div class="flex items-center justify-end">
              <button
                class="group/btn flex w-full items-center justify-center rounded-xl px-6 py-3 font-medium transition-all duration-300 active:scale-95 md:w-auto"
                :class="[
                  item.type === 'manual'
                    ? 'border border-slate-200 bg-white text-slate-700 hover:border-blue-600 hover:text-blue-600'
                    : 'bg-blue-600 text-white shadow-lg shadow-blue-200 hover:bg-blue-700',
                ]"
              >
                <Download
                  class="mr-2 h-4 w-4 transition-transform group-hover/btn:-translate-y-1"
                />
                <span v-if="item.type === 'manual'">预览文档</span>
                <span v-else>立即下载</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="filteredDownloads.length === 0" class="py-20 text-center">
        <div class="mb-4 inline-flex rounded-full bg-gray-100 p-4">
          <Search class="h-8 w-8 text-gray-400" />
        </div>
        <p class="text-slate-500">暂无相关资源，请切换分类或联系客服。</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  ChevronRight,
  Cpu,
  Download,
  FileCode,
  FileText,
  HardDrive,
  Search,
  Settings,
} from "lucide-vue-next"
import { computed, onMounted, ref } from "vue"

interface DownloadItem {
  id: number
  title: string
  version: string
  size: string
  date: string
  type: "driver" | "software" | "manual"
  description: string
  isPopular?: boolean
}

const activeTab = ref<"all" | "driver" | "software" | "manual">("all")
const tabs = [
  { id: "all", label: "全部资源" },
  { id: "software", label: "控制软件" },
  { id: "driver", label: "驱动程序" },
  { id: "manual", label: "产品手册" },
]

const downloads: DownloadItem[] = [
  {
    id: 1,
    title: "XS-Control Center Pro",
    version: "v4.2.1",
    size: "156 MB",
    date: "2025-06-15",
    type: "software",
    description: "星实全系列自动化设备通用主控平台，支持Windows 10/11。",
    isPopular: true,
  },
  {
    id: 2,
    title: "Vision-X 激光器USB驱动",
    version: "v2.0.4",
    size: "12 MB",
    date: "2025-05-20",
    type: "driver",
    description: "适用于Vision-X系列的底层通讯驱动，解决连接不稳定问题。",
  },
  {
    id: 3,
    title: "智能装盒机 XS-2000 操作手册",
    version: "PDF",
    size: "8.5 MB",
    date: "2025-04-10",
    type: "manual",
    description: "包含设备安装、调试、日常维护及故障排查指南。",
  },
  {
    id: 4,
    title: "Robo-Arm S3 固件升级包",
    version: "v1.12.0",
    size: "45 MB",
    date: "2025-06-01",
    type: "driver",
    description: "提升机械臂运动平滑度，修复特定路径下的抖动bug。",
  },
  {
    id: 5,
    title: "AirTest 数据分析工具",
    version: "v3.0.0",
    size: "89 MB",
    date: "2025-03-22",
    type: "software",
    description: "用于导出和分析气密测试仪的历史数据，生成质量报表。",
  },
]

const isLoaded = ref(false)

onMounted(() => {
  setTimeout(() => {
    isLoaded.value = true
  }, 100)
})

const filteredDownloads = computed(() => {
  if (activeTab.value === "all") return downloads
  return downloads.filter((item) => item.type === activeTab.value)
})

const getIcon = (type: string) => {
  switch (type) {
    case "driver":
      return Cpu
    case "software":
      return HardDrive
    case "manual":
      return FileText
    default:
      return FileCode
  }
}
</script>

<style scoped></style>
