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
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a class="transition-colors hover:text-blue-600" href="/">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">企业动态</span>
      </nav>

      <div class="grid grid-cols-1 gap-10 lg:grid-cols-4">
        <aside class="space-y-8 lg:col-span-1">
          <div
            class="flex items-center rounded-2xl border border-gray-100 bg-white p-4 shadow-sm"
          >
            <Search class="mr-3 h-5 w-5 text-slate-400" />
            <input
              v-model="title"
              type="text"
              placeholder="搜索新闻标题..."
              class="w-full border-none bg-transparent text-sm text-slate-700 placeholder-slate-400 focus:ring-0"
            />
          </div>

          <div
            class="overflow-hidden rounded-2xl border border-gray-100 bg-white shadow-sm"
          >
            <div
              class="flex items-center justify-between border-b border-gray-50 bg-gray-50/50 p-5"
            >
              <h2 class="flex items-center font-bold text-slate-900">
                <Filter class="mr-2 h-4 w-4 text-blue-600" />
                新闻分类
              </h2>
            </div>
            <ul class="cursor-pointer space-y-1 p-3">
              <li v-for="cat in categories" :key="cat.id">
                <button
                  @click="SelectProductSeries(cat.id)"
                  class="group flex w-full items-center justify-between rounded-xl px-4 py-3 text-left transition-all duration-200"
                  :class="[
                    activeProductSeriesId === cat.id
                      ? 'bg-blue-600 text-white shadow-md shadow-blue-200'
                      : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
                  ]"
                >
                  <span class="font-medium">{{ cat.name }}</span>
                </button>
              </li>
            </ul>
          </div>

          <ContactUs />
        </aside>

        <main class="flex flex-col lg:col-span-3">
          <div
            class="mb-8 flex flex-col items-center justify-between border-b border-gray-200 pb-4 sm:flex-row"
          >
            <h2 class="mb-4 text-2xl font-bold text-slate-800 sm:mb-0">
              全部动态
              <span class="ml-2 text-base font-normal text-slate-400">
                共 {{ paginate?.total_count ?? 0 }} 篇文章
              </span>
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

          <div v-if="loading" class="space-y-6">
            <div
              v-for="n in 4"
              :key="n"
              class="overflow-hidden rounded-2xl border border-gray-100 bg-white"
            >
              <el-skeleton animated>
                <template #template>
                  <div class="p-6">
                    <el-skeleton-item variant="h3" class="mb-3 w-3/4" />
                    <el-skeleton-item variant="text" class="mb-2 w-full" />
                    <el-skeleton-item variant="text" class="w-2/3" />
                  </div>
                </template>
              </el-skeleton>
            </div>
          </div>

          <div
            v-else-if="isEmpty"
            class="flex h-full items-center justify-center"
          >
            <el-empty description="暂无相关动态" :image-size="160">
              <template #description>
                <span class="text-sm text-slate-400">
                  暂无相关动态，请尝试其他搜索条件
                </span>
              </template>
            </el-empty>
          </div>

          <div v-else class="flex-1 space-y-6">
            <article
              v-for="(news, index) in newsData"
              :key="news.id"
              class="group cursor-pointer rounded-2xl border border-gray-100 bg-white p-6 shadow-sm transition-all duration-500 hover:-translate-y-1 hover:shadow-xl hover:shadow-blue-900/5 md:p-8"
              :class="[
                loaded
                  ? 'translate-y-0 opacity-100'
                  : 'translate-y-8 opacity-0',
              ]"
              :style="{ transitionDelay: `${index * 100}ms` }"
              @click="HandelGotoNewsDetail(news)"
            >
              <div
                class="mb-4 flex flex-col justify-between gap-4 md:flex-row md:items-start"
              >
                <div class="flex-1">
                  <div class="mb-3 flex items-center gap-3">
                    <span
                      class="rounded-full bg-blue-50 px-3 py-1 text-xs font-semibold text-blue-600"
                    >
                      {{ GetCategoryName(news.series_id) }}
                    </span>
                    <div class="flex items-center text-xs text-slate-400">
                      <Calendar class="mr-1 h-3.5 w-3.5" />
                      {{ FormatDate(news.created_at) }}
                    </div>
                  </div>
                  <h3
                    class="mb-3 line-clamp-2 text-xl font-bold text-slate-800 transition-colors group-hover:text-blue-700 md:text-2xl"
                  >
                    <a :href="`/news/detail/${news.id}`">{{ news.title }}</a>
                  </h3>
                </div>
              </div>

              <p class="mb-6 line-clamp-3 leading-relaxed text-slate-600">
                {{ news.content }}
              </p>

              <div
                class="flex items-center justify-between border-t border-gray-50 pt-6"
              >
                <div class="flex items-center text-sm text-slate-400">
                  <Eye class="mr-2 h-4 w-4" />
                  <span>{{ news.view_count }} 次浏览</span>
                </div>

                <div
                  class="group/btn inline-flex items-center justify-center rounded-lg bg-blue-600 px-6 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700"
                >
                  阅读更多
                  <ChevronRight
                    class="ml-1 h-4 w-4 transition-transform group-hover/btn:translate-x-1"
                  />
                </div>
              </div>
            </article>
          </div>

          <div v-if="canShowPaginate" class="mt-12 flex justify-center">
            <el-pagination
              v-model:current-page="currentPage"
              :page-size="paginate?.page_size ?? 0"
              :total="paginate?.total_count ?? 0"
              layout="prev, pager, next"
              background
              @current-change="HandlePageChange"
            />
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { GetNewsListApi, GetNewsSeriesApi } from "@/api/news"
import router from "@/router"
import dayjs from "dayjs"
import { ElMessage } from "element-plus"
import { Calendar, ChevronRight, Eye, Filter, Search } from "lucide-vue-next"
import { computed, onMounted, ref, watch } from "vue"

import type { Paginate } from "@/types/common"
import type { News, NewsSeries } from "@/types/new"
import ContactUs from "@/components/contactUs/index.vue"

const loading = ref(false)
const loaded = ref(false)
const title = ref("")
const categories = ref<NewsSeries[]>([])
const activeProductSeriesId = ref<null | number>(null)
const newsData = ref<News[]>([])
const paginate = ref<Paginate | null>(null)
const currentPage = ref(1)

const isEmpty = computed(() => {
  return !loading.value && newsData.value.length === 0
})

const canShowPaginate = computed(() => {
  return (
    !loading.value && newsData.value.length !== 0 && paginate.value !== null
  )
})

function FormatDate(dateStr: string) {
  if (!dateStr) return ""
  return dayjs(dateStr).format("YYYY-MM-DD")
}

function GetCategoryName(seriesId: number): string {
  const category = categories.value.find((c) => c.id === seriesId)
  return category ? category.name : "未分类"
}

function SelectProductSeries(id: number) {
  activeProductSeriesId.value = id
  currentPage.value = 1
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetNewsList()
}

function HandelGotoNewsDetail(news: News) {
  router.push({
    name: "newsDetail",
    params: {
      newsId: news.id,
    },
  })
}

async function GetNewsSeries() {
  try {
    const result = await GetNewsSeriesApi()
    categories.value = result
    if (result.length > 0) activeProductSeriesId.value = result[0].id
  } catch (error) {
    console.log(error)
    ElMessage.error("获取新闻系列失败")
  }
}

async function GetNewsList() {
  loading.value = true
  loaded.value = false
  try {
    const result = await GetNewsListApi(
      currentPage.value,
      title.value,
      activeProductSeriesId.value
    )

    newsData.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.log(error)
    ElMessage.error("获取新闻列表失败")
  } finally {
    loading.value = false
    loaded.value = true
  }
}

watch(title, (newVal) => {
  if (!newVal.trim()) return
  currentPage.value = 1
  GetNewsList()
})

watch(activeProductSeriesId, () => {
  if (!activeProductSeriesId.value) return
  GetNewsList()
})

onMounted(() => {
  GetNewsSeries()
})
</script>

<style scoped></style>
