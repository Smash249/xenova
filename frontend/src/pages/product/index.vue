<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>
      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('./banner/product.webp')] bg-cover bg-center"
      ></div>

      <div class="relative z-20 px-4 text-center">
        <h1
          class="mb-4 text-4xl font-bold tracking-wider text-white drop-shadow-md md:text-5xl"
        >
          产品中心
        </h1>
      </div>
    </div>

    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">产品中心</span>
      </nav>

      <div class="grid grid-cols-1 gap-10 lg:grid-cols-4">
        <aside class="space-y-8 lg:col-span-1">
          <div
            class="flex items-center rounded-2xl border border-gray-100 bg-white p-4 shadow-sm"
          >
            <Search class="mr-3 h-5 w-5 text-slate-400" />
            <input
              v-model="productName"
              type="text"
              placeholder="搜索产品型号..."
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
                产品系列
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
              全部产品
              <span class="ml-2 text-base font-normal text-slate-400">
                共 {{ paginate?.total_count ?? 0 }} 款设备
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

          <div
            v-if="loading"
            class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3"
          >
            <div
              v-for="n in 6"
              :key="n"
              class="overflow-hidden rounded-2xl border border-gray-100 bg-white"
            >
              <el-skeleton animated>
                <template #template>
                  <el-skeleton-item variant="image" class="h-56 w-full" />
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
            <el-empty description="暂无相关产品" :image-size="160">
              <template #description>
                <span class="text-sm text-slate-400"
                  >暂无相关产品，请尝试其他搜索条件</span
                >
              </template>
            </el-empty>
          </div>

          <TransitionGroup
            v-else
            appear
            name="list"
            tag="div"
            class="grid flex-1 grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3"
          >
            <div
              v-for="(product, index) in products"
              :key="product.id"
              class="group flex h-full max-h-100 transform cursor-pointer flex-col overflow-hidden rounded-2xl border border-gray-100 bg-white transition-all duration-500 hover:shadow-2xl hover:shadow-blue-900/10"
              :style="{ animationDelay: `${index * 50}ms` }"
              @click="HandleGotoDetail(product)"
            >
              <div class="relative h-56 overflow-hidden bg-gray-100">
                <img
                  :src="product.cover"
                  :alt="product.name"
                  class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-110"
                />

                <div
                  class="absolute inset-0 bg-blue-900/0 transition-colors duration-300 group-hover:bg-blue-900/10"
                ></div>
              </div>

              <div class="flex flex-1 flex-col p-6">
                <h3
                  class="mb-2 text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-600"
                >
                  {{ product.name }}
                </h3>
                <p class="mb-6 line-clamp-2 flex-1 text-sm text-slate-500">
                  {{ product.description }}
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
              @current-change="HandlePageChange"
            />
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetProductListApi, GetProductSeriesApi } from "@/api/product"
import router from "@/router"
import { ElMessage } from "element-plus"
import { ArrowRight, ChevronRight, Filter, Search } from "lucide-vue-next"
import { computed, onMounted, ref, watch } from "vue"

import type { Paginate } from "@/types/common"
import type { Product, ProductSeries } from "@/types/product"
import ContactUs from "@/components/contactUs/index.vue"

const loading = ref(false)

const productName = ref("")
const categories = ref<ProductSeries[]>([])
const activeProductSeriesId = ref<null | number>(null)
const products = ref<Product[]>([])
const paginate = ref<Paginate | null>(null)
const currentPage = ref(1)

const isEmpty = computed(() => {
  return !loading.value && products.value.length === 0
})

const canShowPaginate = computed(() => {
  return !loading.value && products.value.length > 0 && !!paginate.value
})

function SelectProductSeries(id: number) {
  activeProductSeriesId.value = id
  currentPage.value = 1
}

function HandlePageChange(page: number) {
  currentPage.value = page
  GetProductList()
}

function HandleGotoDetail(product: Product) {
  router.push({ name: "productDetail", params: { product_id: product.id } })
}

async function GetProductSeries() {
  try {
    const result = await GetProductSeriesApi()
    categories.value = result
    if (result.length > 0) activeProductSeriesId.value = result[0].id
  } catch (error) {
    console.log(error)
    ElMessage.error("获取产品系列失败")
  }
}

async function GetProductList() {
  loading.value = true
  try {
    const result = await GetProductListApi(
      currentPage.value,
      productName.value,
      activeProductSeriesId.value
    )
    products.value = result.data
    if (result.paginate) paginate.value = result.paginate
  } catch (error) {
    console.log(error)
    ElMessage.error("获取产品列表失败")
  } finally {
    loading.value = false
  }
}

watch(productName, (newVal) => {
  if (!newVal.trim()) return
  currentPage.value = 1
  GetProductList()
})

watch(activeProductSeriesId, () => {
  if (!activeProductSeriesId.value) return
  GetProductList()
})

onMounted(() => {
  GetProductSeries()
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
