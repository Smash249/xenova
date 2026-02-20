<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
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
      <nav class="mb-8 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <el-icon class="mx-2"><ArrowRight /></el-icon>
        <span class="font-medium text-slate-800">软件中心</span>
      </nav>

      <div
        class="mb-8 flex flex-col items-start justify-between gap-4 md:flex-row md:items-center"
      >
        <div
          class="scrollbar-hide flex w-full overflow-x-auto rounded-xl border border-gray-100 bg-white p-1.5 shadow-sm md:w-auto"
        >
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="SelectSoftWareSeries(tab.id)"
            class="rounded-lg px-6 py-2 text-sm font-medium whitespace-nowrap transition-all duration-300"
            :class="[
              activeTab === tab.id
                ? 'bg-blue-600 text-white shadow-md'
                : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
            ]"
          >
            {{ tab.label }}
          </button>
        </div>

        <div class="w-full shrink-0 md:w-64">
          <el-input
            v-model="searchKeyword"
            placeholder="搜索软件名称或版本..."
            class="w-full"
            size="large"
            :prefix-icon="ElSearch"
            clearable
          />
        </div>
      </div>

      <div
        v-if="activeSeriesDescription"
        class="mb-8 flex items-start gap-3 rounded-xl border border-blue-100 bg-blue-50/50 p-4 text-sm leading-relaxed text-slate-600"
      >
        <el-icon class="mt-0.5 text-lg text-blue-500"><InfoFilled /></el-icon>
        <span>{{ activeSeriesDescription }}</span>
      </div>

      <div v-if="loading" class="space-y-4">
        <div
          v-for="n in 4"
          :key="n"
          class="rounded-2xl border border-gray-100 bg-white p-6"
        >
          <el-skeleton animated>
            <template #template>
              <div class="flex items-center gap-5">
                <el-skeleton-item
                  variant="rect"
                  class="h-14 w-14 rounded-2xl"
                />
                <div class="flex-1">
                  <el-skeleton-item variant="h3" class="mb-3 w-1/3" />
                  <el-skeleton-item variant="text" class="mb-2 w-2/3" />
                  <el-skeleton-item variant="text" class="w-1/2" />
                </div>
              </div>
            </template>
          </el-skeleton>
        </div>
      </div>

      <div v-else-if="filteredDownloads.length > 0" class="space-y-4">
        <div
          v-for="(item, index) in filteredDownloads"
          :key="item.id"
          class="group transform rounded-2xl border border-gray-100 bg-white p-6 shadow-sm transition-all duration-300 hover:-translate-y-1 hover:border-blue-100 hover:shadow-xl"
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
                class="flex h-14 w-14 shrink-0 items-center justify-center rounded-2xl bg-blue-50 text-blue-600 transition-colors duration-300 group-hover:bg-blue-600 group-hover:text-white"
              >
                <span class="text-2xl font-bold uppercase">{{
                  item.name ? item.name.charAt(0) : ""
                }}</span>
              </div>

              <div>
                <div class="mb-1 flex items-center gap-3">
                  <h3
                    class="text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-600"
                  >
                    {{ item.name }}
                  </h3>

                  <el-tag
                    v-if="item.is_hot"
                    type="danger"
                    size="small"
                    effect="light"
                    round
                  >
                    Hot
                  </el-tag>
                </div>

                <p
                  class="mb-3 max-w-2xl text-sm leading-relaxed text-slate-500"
                >
                  {{ item.description }}
                </p>

                <div
                  class="flex flex-wrap items-center gap-4 text-xs font-medium text-slate-400"
                >
                  <span class="flex items-center rounded bg-gray-50 px-2 py-1">
                    <el-icon class="mr-1.5"><Setting /></el-icon>
                    {{ item.file_type }}
                  </span>
                  <span class="flex items-center rounded bg-gray-50 px-2 py-1">
                    <el-icon class="mr-1.5"><Monitor /></el-icon>
                    {{ item.file_size }}
                  </span>
                  <span class="flex items-center">
                    更新于:
                    {{ dayjs(item.created_at).format("YYYY-MM-DD") }}
                  </span>
                </div>
              </div>
            </div>

            <div class="flex shrink-0 items-center justify-end">
              <el-button
                type="primary"
                class="w-full md:w-auto"
                size="large"
                @click="HandleAction(item)"
              >
                <el-icon class="mr-2"><Download /></el-icon>
                立即下载
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <el-empty
        v-else-if="!loading"
        description="暂无相关资源，请切换分类或联系客服。"
        :image-size="120"
      />

      <div v-if="canShowPaginate" class="mt-8 flex justify-center">
        <el-pagination
          v-model:current-page="currentPage"
          :page-size="paginate?.page_size ?? 0"
          :total="paginate?.total_count ?? 0"
          layout="prev, pager, next"
          background
          @current-change="HandlePageChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetSoftWareListApi, GetSoftWareSeriesApi } from "@/api/software"
import {
  ArrowRight,
  Download,
  Search as ElSearch,
  InfoFilled,
  Monitor,
  Setting,
} from "@element-plus/icons-vue"
import dayjs from "dayjs"
import { ElMessage } from "element-plus"
import { computed, onMounted, ref, watch } from "vue"

import type { Paginate } from "@/types/common"
import type { SoftWare, SoftWareSeries } from "@/types/software"

interface TabItem {
  id: number
  label: string
  description: string
}

const loading = ref(false)
const activeTab = ref<number>(0)
const searchKeyword = ref("")
const isLoaded = ref(false)
const currentPage = ref(1)

const seriesList = ref<SoftWareSeries[]>([])
const downloads = ref<SoftWare[]>([])
const paginate = ref<Paginate | null>(null)

const tabs = computed<TabItem[]>(() => {
  return seriesList.value.map((item) => ({
    id: item.id,
    label: item.name,
    description: item.description,
  }))
})

const activeSeriesDescription = computed(() => {
  const found = tabs.value.find((tab) => tab.id === activeTab.value)
  return found ? found.description : ""
})

const filteredDownloads = computed(() => {
  if (!searchKeyword.value.trim()) return downloads.value
  const keyword = searchKeyword.value.trim().toLowerCase()
  return downloads.value.filter((item) =>
    item.name.toLowerCase().includes(keyword)
  )
})

const canShowPaginate = computed(() => {
  return !loading.value && downloads.value.length !== 0 && paginate.value
})

function SelectSoftWareSeries(id: number) {
  activeTab.value = id
  currentPage.value = 1
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetSoftWareList()
}

function HandleAction(item: SoftWare) {
  if (!item.file_url) {
    ElMessage.warning("下载链接暂不可用")
    return
  }

  const link = document.createElement("a")
  link.href = item.file_url
  link.download = item.name || ""
  link.target = "_blank"
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

async function GetSoftWareSeries() {
  try {
    const result = await GetSoftWareSeriesApi()
    seriesList.value = result
    if (result.length > 0) {
      activeTab.value = result[0].id
    }
  } catch (error) {
    console.error(error)
    ElMessage.error("获取下载软件系列失败")
  }
}

async function GetSoftWareList() {
  loading.value = true
  try {
    const result = await GetSoftWareListApi(
      currentPage.value,
      searchKeyword.value,
      activeTab.value
    )
    downloads.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.error(error)
    ElMessage.error("获取下载列表失败")
  } finally {
    loading.value = false
  }
}

watch(searchKeyword, (newVal) => {
  if (!newVal.trim()) return
  currentPage.value = 1
  GetSoftWareList()
})

watch(activeTab, () => {
  if (!activeTab.value) return
  GetSoftWareList()
})

onMounted(async () => {
  await GetSoftWareSeries()
  isLoaded.value = true
})
</script>

<style scoped></style>
