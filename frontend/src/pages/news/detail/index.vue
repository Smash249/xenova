<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <div
      class="relative flex h-[35vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/90 to-slate-900/90"
      ></div>

      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('https://images.unsplash.com/photo-1504711434969-e33886168f5c?auto=format&fit=crop&q=80')] bg-cover bg-center mix-blend-overlay"
      ></div>

      <div class="relative z-20 container mx-auto mt-16 px-4 text-center">
        <h1
          class="line-clamp-2 text-3xl font-bold tracking-wider text-white drop-shadow-md md:text-4xl"
        >
          {{ newsDetail?.title || "新闻详情" }}
        </h1>

        <div
          v-if="newsDetail"
          class="mt-6 flex flex-wrap justify-center gap-4 text-sm text-blue-100/80 sm:gap-6"
        >
          <div class="flex items-center">
            <Tag class="mr-1.5 h-4 w-4" />
            <span>{{ newsDetail.series_name }}</span>
          </div>
          <div class="flex items-center">
            <Calendar class="mr-1.5 h-4 w-4" />
            <span>发布于：{{ FormatDate(newsDetail.created_at) }}</span>
          </div>
          <div class="flex items-center">
            <Clock class="mr-1.5 h-4 w-4" />
            <span>最后更新：{{ FormatDate(newsDetail.updated_at) }}</span>
          </div>
          <div class="flex items-center">
            <Eye class="mr-1.5 h-4 w-4" />
            <span>{{ newsDetail.view_count }} 阅读</span>
          </div>
        </div>
      </div>
    </div>

    <div class="container mx-auto px-4 py-8 sm:px-6 lg:px-8">
      <nav class="mb-8 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <a href="/news" class="transition-colors hover:text-blue-600"
          >企业动态</a
        >
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="max-w-60 truncate font-medium text-slate-800">{{
          newsDetail?.title ?? "加载中..."
        }}</span>
      </nav>

      <div
        v-if="loading"
        class="mx-auto max-w-4xl rounded-3xl bg-white p-6 shadow-sm lg:p-10"
      >
        <el-skeleton animated>
          <template #template>
            <div class="space-y-6">
              <el-skeleton-item variant="h1" class="mx-auto h-10 w-2/3" />
              <div class="flex justify-center gap-4">
                <el-skeleton-item variant="text" class="w-20" />
                <el-skeleton-item variant="text" class="w-20" />
                <el-skeleton-item variant="text" class="w-20" />
              </div>
              <el-skeleton-item variant="p" class="h-24 w-full" />
              <el-skeleton-item variant="image" class="h-64 w-full" />
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
          class="overflow-hidden rounded-3xl border border-gray-100 bg-white shadow-sm"
        >
          <div
            v-if="newsDetail.description"
            class="border-b border-gray-100 bg-blue-50/30 p-8"
          >
            <div class="flex gap-4">
              <Quote class="h-8 w-8 shrink-0 text-blue-300" />
              <p
                class="text-base leading-relaxed font-medium text-slate-600 italic"
              >
                {{ newsDetail.description }}
              </p>
            </div>
          </div>

          <div class="w-full p-4">
            <MdPreview
              :modelValue="newsDetail.content"
              :codeTheme="theme"
              class="news-content"
              previewTheme="github"
              style="background-color: none"
            />
          </div>
        </article>
      </main>

      <div
        v-else
        class="flex h-[50vh] flex-col items-center justify-center text-slate-500"
      >
        <el-empty description="未找到该新闻信息" :image-size="160">
          <template #extra>
            <el-button type="primary" plain @click="router.back()"
              >返回列表</el-button
            >
          </template>
        </el-empty>
      </div>
    </div>

    <el-backtop :right="40" :bottom="40" />
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

<style scoped></style>
