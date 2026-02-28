<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <div
      class="relative flex h-[30vh] items-center justify-center overflow-hidden bg-slate-900 sm:h-[35vh]"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/90 to-slate-900/90"
      ></div>

      <div
        class="relative z-20 container mx-auto mt-12 px-4 text-center sm:mt-16"
      >
        <h1
          class="line-clamp-2 text-2xl font-bold tracking-wider text-white drop-shadow-md sm:text-3xl md:text-4xl"
        >
          {{ newsDetail?.title || "新闻详情" }}
        </h1>

        <div
          v-if="newsDetail"
          class="mt-4 flex flex-wrap justify-center gap-3 text-xs text-blue-100/80 sm:mt-6 sm:gap-6 sm:text-sm"
        >
          <div class="flex items-center">
            <Tag class="mr-1 h-3.5 w-3.5 sm:mr-1.5 sm:h-4 sm:w-4" />
            <span>{{ newsDetail.series_name }}</span>
          </div>
          <div class="flex items-center">
            <Calendar class="mr-1 h-3.5 w-3.5 sm:mr-1.5 sm:h-4 sm:w-4" />
            <span>发布于：{{ FormatDate(newsDetail.created_at) }}</span>
          </div>
          <div class="hidden items-center sm:flex">
            <Clock class="mr-1 h-3.5 w-3.5 sm:mr-1.5 sm:h-4 sm:w-4" />
            <span>最后更新：{{ FormatDate(newsDetail.updated_at) }}</span>
          </div>
          <div class="flex items-center">
            <Eye class="mr-1 h-3.5 w-3.5 sm:mr-1.5 sm:h-4 sm:w-4" />
            <span>{{ newsDetail.view_count }} 阅读</span>
          </div>
        </div>
      </div>
    </div>

    <div class="container mx-auto px-4 py-6 sm:px-6 sm:py-8 lg:px-8">
      <nav
        class="mb-4 flex items-center text-xs text-slate-500 sm:mb-8 sm:text-sm"
      >
        <a href="/" class="shrink-0 transition-colors hover:text-blue-600"
          >首页</a
        >
        <ChevronRight class="mx-1 h-3 w-3 shrink-0 sm:mx-2 sm:h-4 sm:w-4" />
        <a href="/news" class="shrink-0 transition-colors hover:text-blue-600"
          >企业动态</a
        >
        <ChevronRight class="mx-1 h-3 w-3 shrink-0 sm:mx-2 sm:h-4 sm:w-4" />
        <span
          class="max-w-30 truncate font-medium text-slate-800 sm:max-w-60"
          >{{ newsDetail?.title ?? "加载中..." }}</span
        >
      </nav>

      <div
        v-if="loading"
        class="mx-auto max-w-4xl rounded-2xl bg-white p-4 shadow-sm sm:rounded-3xl sm:p-6 lg:p-10"
      >
        <el-skeleton animated>
          <template #template>
            <div class="space-y-6">
              <el-skeleton-item
                variant="h1"
                class="mx-auto h-8 w-3/4 sm:h-10 sm:w-2/3"
              />
              <div class="flex justify-center gap-4">
                <el-skeleton-item variant="text" class="w-16 sm:w-20" />
                <el-skeleton-item variant="text" class="w-16 sm:w-20" />
                <el-skeleton-item variant="text" class="w-16 sm:w-20" />
              </div>
              <el-skeleton-item variant="p" class="h-16 w-full sm:h-24" />
              <el-skeleton-item variant="image" class="h-48 w-full sm:h-64" />
              <div class="space-y-2">
                <el-skeleton-item variant="text" class="w-full" />
                <el-skeleton-item variant="text" class="w-full" />
                <el-skeleton-item variant="text" class="w-4/5" />
              </div>
            </div>
          </template>
        </el-skeleton>
      </div>

      <main v-else-if="newsDetail">
        <article
          class="overflow-hidden rounded-2xl border border-gray-100 bg-white shadow-sm sm:rounded-3xl"
        >
          <div
            v-if="newsDetail.description"
            class="border-b border-gray-100 bg-blue-50/30 p-4 sm:p-6 lg:p-8"
          >
            <div class="flex gap-3 sm:gap-4">
              <Quote class="h-6 w-6 shrink-0 text-blue-300 sm:h-8 sm:w-8" />
              <p
                class="text-sm leading-relaxed font-medium text-slate-600 italic sm:text-base"
              >
                {{ newsDetail.description }}
              </p>
            </div>
          </div>

          <div class="w-full p-2 sm:p-4 lg:p-8">
            <MdPreview
              :modelValue="newsDetail.content"
              :codeTheme="theme"
              class="news-content"
              previewTheme="github"
              style="background-color: transparent"
            />
          </div>
        </article>
      </main>

      <div
        v-else
        class="flex h-[40vh] flex-col items-center justify-center text-slate-500 sm:h-[50vh]"
      >
        <el-empty description="未找到该新闻信息" :image-size="120">
          <template #extra>
            <el-button type="primary" plain @click="router.back()"
              >返回列表</el-button
            >
          </template>
        </el-empty>
      </div>
    </div>

    <el-backtop :right="20" :bottom="20" class="sm:right-10! sm:bottom-10!" />
  </div>
</template>

<script setup lang="ts">
import { ElMessage } from "element-plus"
import { MdPreview } from "md-editor-v3"
import { onMounted, ref } from "vue"
import { useRoute, useRouter } from "vue-router"

import "md-editor-v3/lib/preview.css"

import { GetNewsDetailApi } from "@/api/news"
import dayjs from "dayjs"
import { Calendar, ChevronRight, Clock, Eye, Quote, Tag } from "lucide-vue-next"

import type { NewsDetail } from "@/types/new"

const route = useRoute()
const router = useRouter()
const loading = ref(true)
const newsDetail = ref<NewsDetail | null>(null)
const theme = ref("atom")

function FormatDate(dateStr: string, full: boolean = false) {
  if (!dateStr) return "-"
  return dayjs(dateStr).format(full ? "YYYY-MM-DD HH:mm" : "YYYY-MM-DD")
}

async function GetNewsDetail() {
  const newsId = route.params.newsId
  if (!newsId) return
  loading.value = true
  try {
    const result = await GetNewsDetailApi(Number(newsId))
    newsDetail.value = result
  } catch (error) {
    console.error("Failed to fetch news detail:", error)
    ElMessage.error("获取新闻详情失败")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  GetNewsDetail()
})
</script>

<style scoped>
:deep(.news-content) {
  background-color: transparent !important;
}
</style>
