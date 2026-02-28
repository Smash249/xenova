<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <CustomBanner title="产品详情" />

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

            <div class="flex flex-col space-y-8">
              <div class="border-b border-gray-100 pb-6">
                <h2
                  class="mb-4 text-3xl font-extrabold tracking-tight text-slate-900"
                >
                  {{ productDetail.name }}
                </h2>

                <div class="flex flex-wrap items-center gap-3">
                  <span
                    class="rounded-full border border-blue-100 bg-blue-50 px-3 py-1 text-xs font-semibold text-blue-600"
                  >
                    {{ productDetail.series_name ?? "默认分类" }}
                  </span>
                  <div
                    class="flex items-center rounded-full border border-green-100 bg-green-50 px-3 py-1 text-xs font-medium text-green-600"
                  >
                    <CheckCircle class="mr-1.5 h-3.5 w-3.5" />
                    <span>现货供应</span>
                  </div>
                  <div
                    class="flex items-center rounded-full border border-blue-100 bg-blue-50 px-3 py-1 text-xs font-medium text-blue-600"
                  >
                    <ShieldCheck class="mr-1.5 h-3.5 w-3.5" />
                    <span>产品质保</span>
                  </div>
                </div>
              </div>

              <div>
                <h3
                  class="mb-3 flex items-center text-base font-bold text-slate-800"
                >
                  <div class="mr-2 h-4 w-1 rounded-sm bg-blue-600"></div>
                  产品简介
                </h3>
                <div
                  class="rounded-xl border border-gray-100 bg-gray-50 p-5 text-slate-600"
                >
                  <p class="text-justify text-sm leading-relaxed">
                    {{ productDetail.description ?? "暂无产品描述" }}
                  </p>
                </div>
              </div>

              <div
                class="mt-auto rounded-2xl border border-blue-100 bg-linear-to-br from-blue-50/80 to-white p-6 shadow-sm"
              >
                <h4 class="mb-5 text-lg font-bold text-slate-800">
                  有意向购买？请联系我们
                </h4>
                <div class="space-y-3 text-sm">
                  <div
                    class="flex items-center rounded-xl border border-gray-50 bg-white p-3 shadow-sm transition-all hover:border-blue-100"
                  >
                    <div
                      class="mr-4 flex h-10 w-10 shrink-0 items-center justify-center rounded-full bg-blue-50"
                    >
                      <Phone class="h-4 w-4 text-blue-600" />
                    </div>
                    <div>
                      <div class="mb-0.5 text-xs text-slate-400">咨询热线</div>
                      <div class="text-base font-bold text-slate-800">
                        400-123-4567
                      </div>
                    </div>
                  </div>
                  <div
                    class="flex items-center rounded-xl border border-gray-50 bg-white p-3 shadow-sm transition-all hover:border-blue-100"
                  >
                    <div
                      class="mr-4 flex h-10 w-10 shrink-0 items-center justify-center rounded-full bg-blue-50"
                    >
                      <Mail class="h-4 w-4 text-blue-600" />
                    </div>
                    <div>
                      <div class="mb-0.5 text-xs text-slate-400">电子邮箱</div>
                      <div class="text-base font-bold text-slate-800">
                        sales@example.com
                      </div>
                    </div>
                  </div>
                </div>
                <p class="mt-4 text-right text-xs text-slate-400">
                  * 工作时间：周一至周五 9:00 - 18:00
                </p>
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
} from "lucide-vue-next"
import { onMounted, ref } from "vue"
import { useRoute, useRouter } from "vue-router"

import type { ProductDetail } from "@/types/product"
import CustomBanner from "@/components/customBanner/index.vue"

const Baseurl = import.meta.env.VITE_BASE_URL

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const productDetail = ref<ProductDetail | null>(null)

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
