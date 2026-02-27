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
            <el-skeleton-item variant="image" class="h-56 w-full" />
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
        <Trophy class="h-10 w-10 text-slate-300" />
      </div>
      <p class="mb-1 text-lg font-semibold text-slate-400">暂无荣誉数据</p>
      <p class="text-sm text-slate-300">荣誉记录将在此处展示</p>
    </div>

    <div v-else class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="item in honors"
        :key="item.id"
        @click="OpenModal(item)"
        class="group relative flex cursor-pointer flex-col overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm transition-all duration-300 hover:-translate-y-1.5 hover:border-blue-400 hover:shadow-xl hover:shadow-blue-900/10"
      >
        <div
          class="relative flex h-56 w-full items-center justify-center overflow-hidden bg-slate-50 p-6"
        >
          <div
            class="absolute inset-0 bg-[radial-gradient(#cbd5e1_1px,transparent_1px)] bg-size-[16px_16px] opacity-40"
          ></div>

          <img
            :src="BaseUrl + item.image"
            :alt="item.title"
            class="relative z-10 max-h-full max-w-full object-contain drop-shadow-md transition-transform duration-500 group-hover:scale-105"
          />

          <div
            class="absolute inset-0 z-20 flex items-center justify-center bg-slate-900/40 opacity-0 backdrop-blur-[2px] transition-all duration-300 group-hover:opacity-100"
          >
            <div
              class="flex h-12 w-12 translate-y-4 items-center justify-center rounded-full bg-white text-blue-600 shadow-lg transition-transform duration-300 group-hover:translate-y-0"
            >
              <ZoomIn class="h-6 w-6" />
            </div>
          </div>
        </div>

        <div class="flex flex-1 flex-col p-5">
          <div class="mb-3 flex items-start justify-between gap-3">
            <h3
              class="line-clamp-2 text-lg leading-snug font-bold text-slate-800 transition-colors group-hover:text-blue-600"
            >
              {{ item.title }}
            </h3>
            <span
              class="shrink-0 rounded-md border border-blue-100/50 bg-blue-50 px-2 py-1 text-xs font-semibold text-blue-600"
            >
              {{ GetYear(item.issue_date) }}
            </span>
          </div>

          <div class="mt-auto space-y-2.5 border-t border-slate-100 pt-4">
            <div
              class="flex items-center text-sm text-slate-500"
              v-if="item.cert_no"
            >
              <ShieldCheck class="mr-2 h-4 w-4 shrink-0 text-slate-400" />
              <span class="truncate font-mono">{{ item.cert_no }}</span>
            </div>
            <div class="flex items-center text-sm text-slate-500">
              <Calendar class="mr-2 h-4 w-4 shrink-0 text-slate-400" />
              <span>{{ item.issue_date }}</span>
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

    <CompanyHonorModal
      :is-open="isModalOpen"
      :honor="selectedHonor"
      @close="CloseModal"
    />
  </div>
</template>

<script setup lang="ts">
import { GetCompanyHonorApi } from "@/api/honor"
import { Sleep } from "@/utils/common"
import { ElMessage } from "element-plus"
import { Calendar, ShieldCheck, Trophy, ZoomIn } from "lucide-vue-next"
import { computed, onMounted, ref } from "vue"

import type { Paginate } from "@/types/common"
import type { CompanyHonor } from "@/types/honor"

import CompanyHonorModal from "../components/CompanyHonorModal.vue"

const BaseUrl = import.meta.env.VITE_BASE_URL

const loading = ref(true)
const isModalOpen = ref(false)
const honors = ref<CompanyHonor[]>([])
const selectedHonor = ref<CompanyHonor | null>(null)
const paginate = ref<Paginate | null>(null)
const currentPage = ref(1)

const isEmpty = computed(() => {
  return !loading.value && honors.value.length === 0
})

const canShowPaginate = computed(() => {
  return !loading.value && honors.value.length !== 0 && paginate.value !== null
})

function GetYear(dateStr: string): string {
  return dateStr.substring(0, 4)
}

function OpenModal(item: CompanyHonor) {
  selectedHonor.value = item
  isModalOpen.value = true
  document.body.style.overflow = "hidden"
}

function CloseModal() {
  isModalOpen.value = false
  document.body.style.overflow = ""
  setTimeout(() => {
    selectedHonor.value = null
  }, 300)
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetCompanyHonor()
}

async function GetCompanyHonor() {
  loading.value = true
  try {
    const [result] = await Promise.all([
      GetCompanyHonorApi(currentPage.value),
      Sleep(500),
    ])
    honors.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.error(error)
    ElMessage.error("获取荣誉列表失败")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  GetCompanyHonor()
})
</script>

<style scoped lang="scss"></style>
