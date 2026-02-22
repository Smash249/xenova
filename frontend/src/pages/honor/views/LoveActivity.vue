<template>
  <div class="h-full space-y-6">
    <div
      v-if="loading"
      class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3"
    >
      <div
        v-for="n in 6"
        :key="n"
        class="overflow-hidden rounded-2xl border border-slate-200 bg-white"
      >
        <el-skeleton animated>
          <template #template>
            <el-skeleton-item variant="image" class="h-52 w-full" />
            <div class="p-5">
              <el-skeleton-item variant="h3" class="mb-3 w-3/4" />
              <el-skeleton-item variant="text" class="mb-2 w-full" />
              <el-skeleton-item variant="text" class="w-1/2" />
            </div>
          </template>
        </el-skeleton>
      </div>
    </div>

    <div
      v-else-if="isEmpty"
      class="flex h-full flex-col items-center justify-center rounded-2xl border border-dashed border-slate-300 bg-slate-50/50 py-20"
    >
      <div
        class="mb-4 flex h-20 w-20 items-center justify-center rounded-full bg-slate-100"
      >
        <Heart class="h-10 w-10 text-slate-300" />
      </div>
      <p class="mb-1 text-lg font-semibold text-slate-400">暂无活动数据</p>
      <p class="text-sm text-slate-300">爱心活动记录将在此处展示</p>
    </div>

    <div v-else class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="item in activities"
        :key="item.id"
        @click="OpenModal(item)"
        class="group relative flex cursor-pointer flex-col overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm transition-all duration-300 hover:-translate-y-1.5 hover:border-rose-300 hover:shadow-xl hover:shadow-rose-900/5"
      >
        <div class="relative h-52 w-full overflow-hidden bg-slate-100">
          <img
            :src="BaseUrl + item.cover"
            :alt="item.title"
            class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-110"
          />

          <div
            class="absolute top-4 left-4 flex flex-col items-center justify-center rounded-lg bg-white/95 px-3 py-1.5 text-center shadow-sm backdrop-blur-sm"
          >
            <span class="text-xs font-bold text-rose-600">{{
              GetMonth(item.activity_date)
            }}</span>
            <span class="text-lg leading-none font-black text-slate-800">{{
              GetDay(item.activity_date)
            }}</span>
          </div>

          <div
            class="absolute inset-0 z-10 flex items-center justify-center bg-rose-900/20 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
          >
            <div
              class="flex h-12 w-12 scale-50 items-center justify-center rounded-full bg-white/90 text-rose-500 shadow-lg backdrop-blur-md transition-all duration-300 group-hover:scale-100"
            >
              <Heart class="h-6 w-6 fill-current" />
            </div>
          </div>
        </div>

        <div class="flex flex-1 flex-col p-5">
          <h3
            class="line-clamp-2 text-lg leading-snug font-bold text-slate-800 transition-colors group-hover:text-rose-600"
          >
            {{ item.title }}
          </h3>

          <p class="mt-2 line-clamp-2 text-sm leading-relaxed text-slate-500">
            {{ item.summary }}
          </p>

          <div
            class="mt-auto flex items-center justify-between border-t border-slate-100 pt-4 text-xs text-slate-400"
          >
            <div class="flex items-center">
              <MapPin class="mr-1.5 h-3.5 w-3.5" />
              <span class="max-w-25 truncate">{{ item.location }}</span>
            </div>
            <div class="flex items-center">
              <Users class="mr-1.5 h-3.5 w-3.5" />
              <span>{{ item.participants }}人参与</span>
            </div>
          </div>
        </div>
      </div>
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

    <LoveActivityModal
      :is-open="isModalOpen"
      :activity="selectedActivity"
      @close="CloseModal"
    />
  </div>
</template>

<script setup lang="ts">
import { GetLoveActivityListApi } from "@/api/honor"
import { ElMessage } from "element-plus"
import { Heart, MapPin, Users } from "lucide-vue-next"
import { computed, onMounted, ref } from "vue"

import type { Paginate } from "@/types/common"
import type { LoveActivity } from "@/types/honor"

import LoveActivityModal from "../components/LoveActivityModal.vue"

const BaseUrl = import.meta.env.VITE_BASE_URL

const loading = ref(false)
const isModalOpen = ref(false)
const activities = ref<LoveActivity[]>([])
const selectedActivity = ref<LoveActivity | null>(null)
const paginate = ref<Paginate | null>(null)
const currentPage = ref(1)

const isEmpty = computed(() => {
  return !loading.value && activities.value.length === 0
})

const canShowPaginate = computed(() => {
  return (
    !loading.value && activities.value.length !== 0 && paginate.value !== null
  )
})

function GetMonth(dateStr: string): string {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}月`
}

function GetDay(dateStr: string): string {
  const date = new Date(dateStr)
  return date.getDate().toString().padStart(2, "0")
}

function OpenModal(item: LoveActivity) {
  selectedActivity.value = item
  isModalOpen.value = true
  document.body.style.overflow = "hidden"
}

function CloseModal() {
  isModalOpen.value = false
  document.body.style.overflow = ""
  setTimeout(() => {
    selectedActivity.value = null
  }, 300)
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetLoveActivity()
}

async function GetLoveActivity() {
  loading.value = true
  try {
    const result = await GetLoveActivityListApi(currentPage.value)
    activities.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.error(error)
    ElMessage.error("获取活动列表失败")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  GetLoveActivity()
})
</script>

<style scoped lang="scss"></style>
