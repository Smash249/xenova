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
        <FileCheck class="h-10 w-10 text-slate-300" />
      </div>
      <p class="mb-1 text-lg font-semibold text-slate-400">暂无专利数据</p>
      <p class="text-sm text-slate-300">专利记录将在此处展示</p>
    </div>

    <div v-else class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="item in patents"
        :key="item.id"
        @click="OpenModal(item)"
        class="group relative flex cursor-pointer flex-col overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm transition-all duration-300 hover:-translate-y-1.5 hover:border-indigo-400 hover:shadow-xl hover:shadow-indigo-900/10"
      >
        <div
          class="relative flex h-52 w-full items-center justify-center overflow-hidden bg-slate-50 p-4"
        >
          <div
            class="absolute inset-0 bg-[linear-gradient(to_right,#e2e8f0_1px,transparent_1px),linear-gradient(to_bottom,#e2e8f0_1px,transparent_1px)] bg-size-[24px_24px] opacity-50"
          ></div>

          <img
            :src="BaseUrl + item.image"
            :alt="item.title"
            class="relative z-10 max-h-full max-w-full object-contain drop-shadow-sm transition-transform duration-500 group-hover:scale-105"
          />

          <div
            class="absolute top-3 left-3 z-20 rounded-md bg-white/90 px-2 py-1 text-xs font-bold shadow-sm backdrop-blur-sm"
            :class="GetTypeColor(item.type)"
          >
            {{ item.type }}
          </div>

          <div
            class="absolute inset-0 z-20 flex items-center justify-center bg-slate-900/40 opacity-0 backdrop-blur-[2px] transition-all duration-300 group-hover:opacity-100"
          >
            <div
              class="flex h-12 w-12 translate-y-4 items-center justify-center rounded-full bg-white text-indigo-600 shadow-lg transition-transform duration-300 group-hover:translate-y-0"
            >
              <ZoomIn class="h-6 w-6" />
            </div>
          </div>
        </div>

        <div class="flex flex-1 flex-col p-5">
          <h3
            class="line-clamp-2 text-base leading-snug font-bold text-slate-800 transition-colors group-hover:text-indigo-600"
          >
            {{ item.title }}
          </h3>

          <div class="mt-auto space-y-2.5 border-t border-slate-100 pt-4">
            <div class="flex items-center text-sm text-slate-500">
              <FileCheck class="mr-2 h-4 w-4 shrink-0 text-slate-400" />
              <span class="truncate font-mono">{{ item.patent_no }}</span>
            </div>
            <div class="flex items-center text-sm text-slate-500">
              <Calendar class="mr-2 h-4 w-4 shrink-0 text-slate-400" />
              <span>授权日：{{ item.auth_date }}</span>
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

    <CompanyPatentModal
      :is-open="isModalOpen"
      :patent="selectedPatent"
      @close="CloseModal"
    />
  </div>
</template>

<script setup lang="ts">
import { GetCompanyPatentApi } from "@/api/honor"
import { Sleep } from "@/utils/common"
import { ElMessage } from "element-plus"
import { Calendar, FileCheck, ZoomIn } from "lucide-vue-next"
import { computed, onMounted, ref } from "vue"

import type { Paginate } from "@/types/common"
import type { CompanyPatent } from "@/types/honor"

import CompanyPatentModal from "../components/CompanyPatentModal.vue"

const BaseUrl = import.meta.env.VITE_BASE_URL

const loading = ref(true)
const isModalOpen = ref(false)
const patents = ref<CompanyPatent[]>([])
const selectedPatent = ref<CompanyPatent | null>(null)
const paginate = ref<Paginate | null>(null)
const currentPage = ref(1)

const isEmpty = computed(() => {
  return !loading.value && patents.value.length === 0
})

const canShowPaginate = computed(() => {
  return !loading.value && patents.value.length !== 0 && paginate.value !== null
})

function GetTypeColor(type: string): string {
  if (type.includes("发明")) {
    return "bg-orange-100 text-orange-700 border border-orange-200"
  } else if (type.includes("实用新型")) {
    return "bg-indigo-100 text-indigo-700 border border-indigo-200"
  } else if (type.includes("外观设计")) {
    return "bg-teal-100 text-teal-700 border border-teal-200"
  }
  return "bg-slate-100 text-slate-700 border border-slate-200"
}

function OpenModal(item: CompanyPatent) {
  selectedPatent.value = item
  isModalOpen.value = true
  document.body.style.overflow = "hidden"
}

function CloseModal() {
  isModalOpen.value = false
  document.body.style.overflow = ""
  setTimeout(() => {
    selectedPatent.value = null
  }, 300)
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetCompanyPatent()
}

async function GetCompanyPatent() {
  loading.value = true
  try {
    const [result] = await Promise.all([
      GetCompanyPatentApi(currentPage.value),
      Sleep(500),
    ])
    patents.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.error(error)
    ElMessage.error("获取专利列表失败")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  GetCompanyPatent()
})
</script>

<style scoped lang="scss"></style>
