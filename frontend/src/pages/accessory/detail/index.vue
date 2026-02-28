<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <CustomBanner title="配件详情" />

    <div class="container mx-auto px-4 py-8 sm:px-6 lg:px-8">
      <nav class="mb-8 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <a href="/accessory" class="transition-colors hover:text-blue-600"
          >配件商场</a
        >
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="max-w-50 truncate font-medium text-slate-800">{{
          accessoryDetail?.name ?? "加载中..."
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

      <div v-else-if="accessoryDetail" class="animate-fade-in space-y-12">
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
                :src="Baseurl + (accessoryDetail.cover ?? '')"
                :alt="accessoryDetail.name ?? ''"
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
            <div class="flex h-full flex-col">
              <div
                class="mb-6 flex flex-wrap items-start justify-between gap-6 border-b border-gray-100 pb-6"
              >
                <div class="flex-1">
                  <h2
                    class="mb-3 text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl"
                  >
                    {{ accessoryDetail.name }}
                  </h2>
                  <div class="flex flex-wrap items-center gap-2">
                    <span
                      class="rounded bg-slate-100 px-2.5 py-1 text-xs font-medium text-slate-600"
                    >
                      {{ accessoryDetail.series_name ?? "默认分类" }}
                    </span>
                    <div
                      class="flex items-center rounded bg-green-50 px-2 py-1 text-xs font-medium text-green-600"
                    >
                      <CheckCircle class="mr-1 h-3 w-3" />
                      现货供应
                    </div>
                    <div
                      class="flex items-center rounded bg-blue-50 px-2 py-1 text-xs font-medium text-blue-600"
                    >
                      <ShieldCheck class="mr-1 h-3 w-3" />
                      原厂质保
                    </div>
                  </div>
                </div>
                <div class="shrink-0 text-right">
                  <span class="mb-1 block text-xs text-slate-400">参考价</span>
                  <div class="text-red-600">
                    <span class="text-lg font-bold">¥</span>
                    <span class="text-3xl font-black tracking-tight">{{
                      accessoryDetail.price ?? "面议"
                    }}</span>
                  </div>
                </div>
              </div>

              <!-- 独立商品描述区块 -->
              <div class="mb-6 flex-1">
                <h3
                  class="mb-3 text-sm font-bold tracking-widest text-slate-400 uppercase"
                >
                  配件参数与描述
                </h3>
                <div
                  class="relative rounded-2xl bg-slate-50/80 p-5 text-slate-600 ring-1 ring-slate-100/50 ring-inset"
                >
                  <p class="text-justify text-sm leading-relaxed">
                    {{ accessoryDetail.description ?? "暂无配件描述" }}
                  </p>
                </div>
              </div>

              <!-- 双栏网格联系方式卡片 (与产品页不同) -->
              <div class="mb-8">
                <h3
                  class="mb-3 text-sm font-bold tracking-widest text-slate-400 uppercase"
                >
                  采购与技术支持
                </h3>
                <div class="grid grid-cols-2 gap-4">
                  <div
                    class="flex flex-col items-start rounded-2xl border border-blue-100 bg-blue-50/30 p-4 transition-colors hover:bg-blue-50"
                  >
                    <div
                      class="mb-2 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100/50 text-blue-600"
                    >
                      <Phone class="h-4 w-4" />
                    </div>
                    <span class="text-xs text-slate-500">咨询热线</span>
                    <span
                      class="mt-0.5 font-mono text-sm font-bold text-slate-800"
                      >400-123-4567</span
                    >
                  </div>
                  <div
                    class="flex flex-col items-start rounded-2xl border border-blue-100 bg-blue-50/30 p-4 transition-colors hover:bg-blue-50"
                  >
                    <div
                      class="mb-2 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100/50 text-blue-600"
                    >
                      <Mail class="h-4 w-4" />
                    </div>
                    <span class="text-xs text-slate-500">电子邮箱</span>
                    <span
                      class="mt-0.5 font-mono text-sm font-bold text-slate-800"
                      >sales@example.com</span
                    >
                  </div>
                </div>
              </div>

              <!-- 下单操作区 -->
              <div class="mt-auto border-t border-gray-100 pt-6">
                <div class="mb-4 flex items-baseline gap-2">
                  <span class="text-sm font-medium text-slate-800"
                    >选择订购数量</span
                  >
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
                    加入采购单
                  </el-button>
                </div>

                <!-- 新增购买提示 -->
                <div class="mt-4 flex items-center text-sm text-slate-500">
                  <Info class="mr-1.5 h-4 w-4 shrink-0 text-blue-500" />
                  <span
                    >如需购买，请联系我们的电话：<a
                      href="tel:4001234567"
                      class="font-medium text-blue-600 hover:underline"
                      >400-123-4567</a
                    ></span
                  >
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
            <h2 class="text-2xl font-bold text-slate-800">配件详情展示</h2>
          </div>

          <div class="space-y-4">
            <template
              v-if="
                accessoryDetail.previews && accessoryDetail.previews.length > 0
              "
            >
              <div
                v-for="(img, idx) in accessoryDetail.previews"
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
              <el-empty description="暂无配件详情图" :image-size="160" />
            </div>
          </div>
        </div>
      </div>

      <div
        v-else
        class="flex h-[50vh] flex-col items-center justify-center text-slate-500"
      >
        <el-empty description="未找到该配件信息" :image-size="160">
          <template #extra>
            <el-button type="primary" plain @click="router.back()"
              >返回配件列表</el-button
            >
          </template>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetAccessoryDetailApi } from "@/api/accessorys"
import { ElMessage } from "element-plus"
import {
  CheckCircle,
  ChevronRight,
  ImageOff,
  Info,
  Mail,
  Phone,
  ShieldCheck,
  ShoppingCart,
} from "lucide-vue-next"
import { onMounted, ref } from "vue"
import { useRoute, useRouter } from "vue-router"

import type { AccessoryDetail } from "@/types/accessorys"
import CustomBanner from "@/components/customBanner/index.vue"

const Baseurl = import.meta.env.VITE_BASE_URL

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const accessoryDetail = ref<AccessoryDetail | null>(null)
const quantity = ref(1)

async function GetAccessoryDetail() {
  if (!route.params.accessory_id) return
  try {
    loading.value = true
    const result = await GetAccessoryDetailApi(
      Number(route.params.accessory_id)
    )
    accessoryDetail.value = result.data ?? result
  } catch (error) {
    console.error(error)
    ElMessage.error("获取配件详情失败")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  GetAccessoryDetail()
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
