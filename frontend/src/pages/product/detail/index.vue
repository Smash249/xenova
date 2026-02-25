<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>
      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?auto=format&fit=crop&q=80')] bg-cover bg-center"
      ></div>

      <div class="relative z-20 px-4 text-center">
        <h1
          class="text-3xl font-bold tracking-wider text-white drop-shadow-md md:text-4xl"
        >
          {{ productDetail?.name || "产品详情" }}
        </h1>
      </div>
    </div>

    <div class="container mx-auto px-4 py-8 sm:px-6 lg:px-8">
      <nav class="mb-8 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <a href="/product" class="transition-colors hover:text-blue-600"
          >产品中心</a
        >
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="max-w-50 truncate font-medium text-slate-800">{{
          productDetail?.name ?? "加载中..."
        }}</span>
      </nav>

      <div v-if="loading" class="rounded-3xl bg-white p-6 shadow-sm">
        <el-skeleton animated>
          <template #template>
            <div class="grid grid-cols-1 gap-10 lg:grid-cols-2">
              <el-skeleton-item
                variant="image"
                class="h-100 w-full rounded-2xl"
              />
              <div class="flex flex-col justify-center space-y-6">
                <el-skeleton-item variant="h1" class="h-10 w-3/4" />
                <div class="space-y-2">
                  <el-skeleton-item variant="text" class="w-full" />
                  <el-skeleton-item variant="text" class="w-full" />
                  <el-skeleton-item variant="text" class="w-2/3" />
                </div>
                <div class="mt-4 flex gap-4">
                  <el-skeleton-item variant="text" class="h-12 w-1/3" />
                  <el-skeleton-item variant="text" class="h-12 w-1/3" />
                </div>
                <div class="pt-6">
                  <el-skeleton-item variant="button" class="h-14 w-40" />
                </div>
              </div>
            </div>
          </template>
        </el-skeleton>
      </div>

      <div v-else-if="productDetail" class="animate-fade-in space-y-12">
        <div
          class="rounded-3xl border border-gray-100 bg-white p-6 shadow-sm lg:p-10"
        >
          <div class="grid grid-cols-1 gap-10 lg:grid-cols-2">
            <div
              class="group relative overflow-hidden rounded-2xl border border-gray-100 bg-gray-50 p-4"
            >
              <div class="absolute top-4 left-4 z-10">
                <span
                  class="rounded-full bg-blue-600/90 px-3 py-1 text-xs font-bold text-white shadow-sm backdrop-blur-sm"
                >
                  官方正品
                </span>
              </div>
              <el-image
                :src="Baseurl + (productDetail.cover ?? '')"
                :alt="productDetail.name ?? ''"
                fit="contain"
                class="h-full w-full rounded-lg transition-transform duration-500 group-hover:scale-105"
              >
                <template #error>
                  <div
                    class="flex h-full w-full flex-col items-center justify-center bg-gray-100 text-slate-400"
                  >
                    <ImageOff class="mb-2 h-10 w-10 opacity-50" />
                    <span class="text-xs">图片加载失败</span>
                  </div>
                </template>
              </el-image>
            </div>

            <!-- 右侧信息 -->
            <div class="flex flex-col justify-between space-y-6">
              <div>
                <div
                  class="mb-6 grid grid-cols-2 gap-4 border-b border-gray-100 pb-6 text-sm"
                >
                  <div class="space-y-1">
                    <span class="block text-slate-400">产品名称</span>
                    <span class="block font-medium text-slate-800">{{
                      productDetail.name
                    }}</span>
                  </div>
                  <div class="space-y-1">
                    <span class="block text-slate-400">产品分类</span>
                    <span class="block font-medium text-blue-600">
                      {{ productDetail.series_name ?? "默认分类" }}
                    </span>
                  </div>
                </div>

                <div
                  class="space-y-4 rounded-xl border border-gray-100 bg-gray-50 p-6 text-slate-600"
                >
                  <p class="text-justify text-sm leading-relaxed">
                    {{ productDetail.description ?? "暂无产品描述" }}
                  </p>

                  <div class="mt-4 rounded-lg bg-blue-50/50 p-4">
                    <h4 class="mb-3 font-bold text-slate-800">
                      有意向购买？请联系我们
                    </h4>
                    <div class="space-y-2 text-sm">
                      <div class="flex items-center text-slate-700">
                        <Phone class="mr-2 h-4 w-4 text-blue-600" />
                        <span class="font-medium">咨询热线：</span>
                        <span>400-123-4567</span>
                      </div>
                      <div class="flex items-center text-slate-700">
                        <Mail class="mr-2 h-4 w-4 text-blue-600" />
                        <span class="font-medium">电子邮箱：</span>
                        <span>sales@example.com</span>
                      </div>
                    </div>
                    <p class="mt-2 text-xs text-slate-500">
                      * 工作时间：周一至周五 9:00 - 18:00
                    </p>
                  </div>

                  <div
                    class="mt-4 flex items-center space-x-4 text-xs font-medium"
                  >
                    <div
                      class="flex items-center rounded bg-green-50 px-2 py-1 text-green-600"
                    >
                      <CheckCircle class="mr-1 h-3 w-3" />
                      <span>现货供应</span>
                    </div>
                    <div
                      class="flex items-center rounded bg-blue-50 px-2 py-1 text-blue-600"
                    >
                      <ShieldCheck class="mr-1 h-3 w-3" />
                      <span>产品质保</span>
                    </div>
                  </div>
                </div>
              </div>

              <div class="border-t border-gray-100 pt-6">
                <div class="mb-6 flex items-baseline gap-2">
                  <span class="text-sm text-slate-400">订购数量</span>
                </div>
                <div class="flex flex-col gap-4 sm:flex-row">
                  <el-input-number
                    v-model="quantity"
                    :min="1"
                    controls-position="right"
                    class="w-36!"
                    size="large"
                  />
                  <el-button
                    :disabled="true"
                    type="primary"
                    size="large"
                    class="flex! h-auto! flex-1! items-center justify-center rounded-xl! px-8! py-3! text-base! font-bold! shadow-lg! shadow-blue-200! hover:shadow-blue-300! sm:w-64! sm:flex-initial!"
                  >
                    <ShoppingCart class="mr-2 h-5 w-5" />
                    立即下单
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="space-y-6 pb-20">
          <div
            class="flex items-center space-x-4 border-b border-gray-200 pb-4"
          >
            <div class="h-8 w-1 rounded-r bg-blue-600"></div>
            <h2 class="text-2xl font-bold text-slate-800">产品详情展示</h2>
          </div>

          <div class="space-y-4">
            <template
              v-if="productDetail.previews && productDetail.previews.length > 0"
            >
              <div
                v-for="(img, idx) in productDetail.previews"
                :key="idx"
                class="overflow-hidden rounded-2xl bg-white shadow-sm ring-1 ring-gray-100"
              >
                <el-image
                  :src="img"
                  lazy
                  class="block w-full object-cover"
                  fit="cover"
                >
                  <template #placeholder>
                    <div
                      class="flex h-64 w-full items-center justify-center bg-gray-50"
                    >
                      <span class="text-slate-400">加载中...</span>
                    </div>
                  </template>
                  <template #error>
                    <div
                      class="flex h-40 w-full flex-col items-center justify-center bg-gray-50 text-slate-400"
                    >
                      <ImageOff class="mb-2 h-8 w-8 opacity-40" />
                      <span class="text-xs">图片无法显示</span>
                    </div>
                  </template>
                </el-image>
              </div>
            </template>

            <div v-else class="flex justify-center py-12">
              <el-empty description="暂无产品详情图" :image-size="160" />
            </div>
          </div>
        </div>
      </div>

      <!-- 404 状态 -->
      <div
        v-else
        class="flex h-[50vh] flex-col items-center justify-center text-slate-500"
      >
        <el-empty description="未找到该产品信息" :image-size="160">
          <template #extra>
            <el-button type="primary" plain @click="router.back()"
              >返回产品列表</el-button
            >
          </template>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetProductDetailApi } from "@/api/product"
import { ElMessage } from "element-plus"
import {
  CheckCircle,
  ChevronRight,
  ImageOff,
  Mail,
  Phone,
  ShieldCheck,
  ShoppingCart,
} from "lucide-vue-next"
import { onMounted, ref } from "vue"
import { useRoute, useRouter } from "vue-router"

import type { ProductDetail } from "@/types/product"

const Baseurl = import.meta.env.VITE_BASE_URL

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const productDetail = ref<ProductDetail | null>(null)
const quantity = ref(1)

async function GetProductDetail() {
  if (!route.params.product_id) return
  try {
    loading.value = true
    const result = await GetProductDetailApi(Number(route.params.product_id))
    productDetail.value = result.data ?? result
  } catch (error) {
    console.error(error)
    ElMessage.error("获取产品详情失败")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  GetProductDetail()
})
</script>

<style scoped>
:deep(.el-input-number .el-input__wrapper) {
  background-color: transparent;
  box-shadow: none;
  border: 1px solid #e2e8f0;
  border-radius: 0.75rem;
}
:deep(.el-input-number .el-input__wrapper.is-focus) {
  border-color: #2563eb;
  box-shadow: 0 0 0 1px #2563eb;
}
</style>
