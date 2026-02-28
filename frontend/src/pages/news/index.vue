<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <CustomBanner title="企业动态" />
    <div class="container mx-auto px-4 py-8 sm:px-6 sm:py-12 lg:px-8">
      <nav class="mb-6 flex items-center text-sm text-slate-500 sm:mb-10">
        <a class="transition-colors hover:text-blue-600" href="/">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">企业动态</span>
      </nav>

      <div class="grid grid-cols-1 gap-6 lg:grid-cols-4 lg:gap-10">
        <aside class="space-y-4 lg:col-span-1 lg:space-y-8">
          <div
            class="flex items-center rounded-2xl border border-gray-100 bg-white p-3 shadow-sm sm:p-4"
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
            <ul
              class="flex cursor-pointer flex-row gap-2 overflow-x-auto p-2 lg:flex-col lg:gap-1 lg:p-3"
            >
              <li
                v-for="cat in categories"
                :key="cat.id"
                class="shrink-0 lg:w-full"
              >
                <button
                  @click="SelectProductSeries(cat.id)"
                  class="group flex w-full items-center justify-between rounded-xl px-4 py-2 text-left transition-all duration-200 lg:py-3"
                  :class="[
                    activeProductSeriesId === cat.id
                      ? 'bg-blue-600 text-white shadow-md shadow-blue-200'
                      : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
                  ]"
                >
                  <span
                    class="text-sm font-medium whitespace-nowrap lg:whitespace-normal"
                    >{{ cat.name }}</span
                  >
                </button>
              </li>
            </ul>
          </div>

          <div class="hidden lg:block">
            <ContactUs />
          </div>
        </aside>

        <main class="flex flex-col lg:col-span-3">
          <div
            class="mb-6 flex flex-col items-start justify-between border-b border-gray-200 pb-4 sm:mb-8 sm:flex-row sm:items-center"
          >
            <h2
              class="mb-2 text-xl font-bold text-slate-800 sm:mb-0 sm:text-2xl"
            >
              全部动态
              <span
                class="ml-2 text-sm font-normal text-slate-400 sm:text-base"
              >
                共 {{ paginate?.total_count ?? 0 }} 篇文章
              </span>
            </h2>
          </div>

          <div v-if="loading" class="space-y-4 sm:space-y-6">
            <div
              v-for="n in 4"
              :key="n"
              class="overflow-hidden rounded-2xl border border-gray-100 bg-white"
            >
              <el-skeleton animated>
                <template #template>
                  <div class="p-4 sm:p-6">
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
            class="flex h-64 flex-col items-center justify-center sm:h-full"
          >
            <el-empty description="暂无相关动态" :image-size="160">
              <template #description>
                <span class="text-sm text-slate-400">
                  暂无相关动态，请尝试其他搜索条件
                </span>
              </template>
            </el-empty>
          </div>

          <TransitionGroup
            v-else
            appear
            name="list"
            tag="div"
            class="flex-1 space-y-4 sm:space-y-6"
          >
            <article
              v-for="(news, index) in newsData"
              :key="news.id"
              class="group cursor-pointer rounded-2xl border border-gray-100 bg-white p-4 shadow-sm transition-all duration-500 hover:-translate-y-1 hover:shadow-xl hover:shadow-blue-900/5 sm:p-6 md:p-8"
              :style="{ animationDelay: `${index * 50}ms` }"
              @click="HandleGotoNewsDetail(news)"
            >
              <div
                class="mb-3 flex flex-col justify-between gap-3 sm:mb-4 sm:gap-4 md:flex-row md:items-start"
              >
                <div class="flex-1">
                  <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                    <span
                      class="rounded-full bg-blue-50 px-2 py-1 text-[10px] font-semibold text-blue-600 sm:px-3 sm:text-xs"
                    >
                      {{ GetCategoryName(news.series_id) }}
                    </span>
                    <div class="flex items-center text-xs text-slate-400">
                      <Calendar class="mr-1 h-3.5 w-3.5" />
                      {{ FormatDate(news.created_at) }}
                    </div>
                  </div>
                  <h3
                    class="mb-2 line-clamp-2 text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-700 sm:mb-3 sm:text-xl md:text-2xl"
                  >
                    <a :href="`/news/detail/${news.id}`">{{ news.title }}</a>
                  </h3>
                </div>
              </div>

              <p
                class="mb-4 line-clamp-3 text-sm leading-relaxed text-slate-600 sm:mb-6 sm:text-base"
              >
                {{ news.content }}
              </p>

              <div
                class="flex items-center justify-between border-t border-gray-50 pt-4 sm:pt-6"
              >
                <div
                  class="flex items-center text-xs text-slate-400 sm:text-sm"
                >
                  <Eye class="mr-1.5 h-4 w-4 sm:mr-2" />
                  <span>{{ news.view_count }} 次浏览</span>
                </div>

                <div
                  class="group/btn inline-flex items-center justify-center rounded-lg bg-blue-600 px-4 py-2 text-xs font-medium text-white transition-colors hover:bg-blue-700 sm:px-6 sm:py-2.5 sm:text-sm"
                >
                  阅读更多
                  <ChevronRight
                    class="ml-1 h-4 w-4 transition-transform group-hover/btn:translate-x-1"
                  />
                </div>
              </div>
            </article>
          </TransitionGroup>

          <div v-if="canShowPaginate" class="mt-8 flex justify-center sm:mt-12">
            <el-pagination
              v-model:current-page="currentPage"
              :page-size="paginate?.page_size ?? 0"
              :total="paginate?.total_count ?? 0"
              layout="prev, pager, next"
              background
              small
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
import { Sleep } from "@/utils/common"
import dayjs from "dayjs"
import { ElMessage } from "element-plus"
import { Calendar, ChevronRight, Eye, Filter, Search } from "lucide-vue-next"
import { computed, onMounted, ref, watch } from "vue"

import type { Paginate } from "@/types/common"
import type { News, NewsSeries } from "@/types/new"
import ContactUs from "@/components/contactUs/index.vue"
import CustomBanner from "@/components/customBanner/index.vue"

const loading = ref(true)
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

function HandleGotoNewsDetail(news: News) {
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
    categories.value = result.data
    if (result.data.length > 0) {
      activeProductSeriesId.value = result.data[0].id
    } else {
      loading.value = false
    }
  } catch (error) {
    console.log(error)
    ElMessage.error("获取新闻系列失败")
    loading.value = false
  }
}

async function GetNewsList() {
  loading.value = true
  try {
    const [result] = await Promise.all([
      GetNewsListApi(
        currentPage.value,
        title.value,
        activeProductSeriesId.value
      ),
      Sleep(500),
    ])

    newsData.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.log(error)
    ElMessage.error("获取新闻列表失败")
  } finally {
    loading.value = false
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

<style scoped>
.list-enter-active {
  animation: list-fade-in-up 0.5s ease-out both;
}

.list-leave-active {
  animation: list-fade-in-up 0.3s ease-in reverse both;
}
</style>
