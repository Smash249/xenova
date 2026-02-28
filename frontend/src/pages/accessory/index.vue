<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <CustomBanner title="配件商城" />
    <div class="container mx-auto px-4 py-8 sm:px-6 sm:py-12 lg:px-8">
      <nav class="mb-6 flex items-center text-sm text-slate-500 sm:mb-10">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">配件商城</span>
      </nav>

      <div class="grid grid-cols-1 gap-6 lg:grid-cols-4 lg:gap-10">
        <aside class="space-y-4 lg:col-span-1 lg:space-y-8">
          <div
            class="flex items-center rounded-2xl border border-gray-100 bg-white p-3 shadow-sm sm:p-4"
          >
            <Search class="mr-3 h-5 w-5 text-slate-400" />
            <input
              v-model="accessoryName"
              type="text"
              placeholder="搜索配件型号..."
              class="w-full border-none bg-transparent text-sm text-slate-700 placeholder-slate-400 focus:ring-0"
            />
          </div>

          <div class="rounded-2xl border border-gray-100 bg-white shadow-sm">
            <div
              class="flex items-center justify-between border-b border-gray-50 bg-gray-50/50 p-5"
            >
              <h2 class="flex items-center font-bold text-slate-900">
                <Filter class="mr-2 h-4 w-4 text-blue-600" />
                配件系列
              </h2>
            </div>
            <ul
              class="scrollbar-hide flex cursor-pointer flex-row gap-2 overflow-x-auto p-2 lg:flex-col lg:gap-1 lg:p-3"
            >
              <li
                v-for="cat in categories"
                :key="cat.id"
                class="shrink-0 lg:w-full"
              >
                <button
                  @click="SelectAccessorySeries(cat.id)"
                  class="group flex w-full items-center justify-between rounded-xl px-4 py-2 text-left transition-all duration-200 lg:py-3"
                  :class="[
                    activeAccessorySeriesId === cat.id
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
              全部配件
              <span
                class="ml-2 text-sm font-normal text-slate-400 sm:text-base"
              >
                共 {{ paginate?.total_count ?? 0 }} 款设备
              </span>
            </h2>
          </div>

          <div
            v-if="loading"
            class="grid grid-cols-1 gap-4 sm:gap-6 md:grid-cols-2 lg:grid-cols-3"
          >
            <div
              v-for="n in 6"
              :key="n"
              class="overflow-hidden rounded-2xl border border-gray-100 bg-white"
            >
              <el-skeleton animated>
                <template #template>
                  <el-skeleton-item
                    variant="image"
                    class="h-48 w-full sm:h-56"
                  />
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
            <el-empty description="暂无相关配件" :image-size="160">
              <template #description>
                <span class="text-sm text-slate-400"
                  >暂无相关配件，请尝试其他搜索条件</span
                >
              </template>
            </el-empty>
          </div>

          <TransitionGroup
            v-else
            appear
            name="list"
            tag="div"
            class="grid flex-1 grid-cols-1 gap-4 sm:gap-6 md:grid-cols-2 lg:grid-cols-3"
          >
            <div
              v-for="(accessory, index) in accessorys"
              :key="accessory.id"
              class="group flex h-full max-h-100 transform cursor-pointer flex-col overflow-hidden rounded-2xl border border-gray-100 bg-white transition-all duration-500 hover:shadow-2xl hover:shadow-blue-900/10"
              :style="{ animationDelay: `${index * 50}ms` }"
              @click="HandleGotoDetail(accessory)"
            >
              <div class="relative h-48 overflow-hidden bg-gray-100 sm:h-56">
                <img
                  :src="BASE_URL + accessory.cover"
                  :alt="accessory.name"
                  class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-110"
                />

                <div
                  class="absolute inset-0 bg-blue-900/0 transition-colors duration-300 group-hover:bg-blue-900/10"
                ></div>
              </div>

              <div class="flex flex-1 flex-col p-4 sm:p-6">
                <h3
                  class="mb-2 text-base font-bold text-slate-800 transition-colors group-hover:text-blue-600 sm:text-lg"
                >
                  {{ accessory.name }}
                </h3>
                <p
                  class="mb-4 line-clamp-2 flex-1 text-xs text-slate-500 sm:mb-6 sm:text-sm"
                >
                  {{ accessory.description }}
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
          </TransitionGroup>

          <div v-if="canShowPaginate" class="mt-8 flex justify-center">
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

<script setup lang="ts">
import { GetAccessoryListApi, GetAccessorySeriesApi } from "@/api/accessorys"
import router from "@/router"
import { Sleep } from "@/utils/common"
import { ElMessage } from "element-plus"
import { ArrowRight, ChevronRight, Filter, Search } from "lucide-vue-next"
import { computed, onMounted, ref, watch } from "vue"

import type { Accessory, AccessorySeries } from "@/types/accessorys"
import type { Paginate } from "@/types/common"
import ContactUs from "@/components/contactUs/index.vue"
import CustomBanner from "@/components/customBanner/index.vue"

const BASE_URL = import.meta.env.VITE_BASE_URL

const loading = ref(true)

const accessoryName = ref("")
const categories = ref<AccessorySeries[]>([])
const activeAccessorySeriesId = ref<null | number>(null)
const accessorys = ref<Accessory[]>([])
const paginate = ref<Paginate | null>(null)
const currentPage = ref(1)

const isEmpty = computed(() => {
  return !loading.value && accessorys.value.length === 0
})

const canShowPaginate = computed(() => {
  return !loading.value && accessorys.value.length > 0 && !!paginate.value
})

function SelectAccessorySeries(id: number) {
  activeAccessorySeriesId.value = id
  currentPage.value = 1
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetAccessoryList()
}

function HandleGotoDetail(accessory: Accessory) {
  router.push({
    name: "accessoryDetail",
    params: { accessory_id: accessory.id },
  })
}

async function GetAccessorySeries() {
  try {
    const result = await GetAccessorySeriesApi()
    categories.value = result.data
    if (result.data.length > 0) {
      activeAccessorySeriesId.value = result.data[0].id
    } else {
      loading.value = false
    }
  } catch (error) {
    console.log(error)
    ElMessage.error("获取配件系列失败")
    loading.value = false
  }
}

async function GetAccessoryList() {
  loading.value = true
  try {
    const [result] = await Promise.all([
      GetAccessoryListApi(
        currentPage.value,
        accessoryName.value,
        activeAccessorySeriesId.value
      ),
      Sleep(500),
    ])
    accessorys.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.log(error)
    ElMessage.error("获取配件列表失败")
  } finally {
    loading.value = false
  }
}

watch(accessoryName, (newVal) => {
  if (!newVal.trim()) return
  currentPage.value = 1
  GetAccessoryList()
})

watch(activeAccessorySeriesId, () => {
  if (!activeAccessorySeriesId.value) return
  GetAccessoryList()
})

onMounted(() => {
  GetAccessorySeries()
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
