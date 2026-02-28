<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <CustomBanner title="配件详情" />

    <div class="container mx-auto px-4 py-6 sm:px-6 sm:py-8 lg:px-8">
      <nav
        class="mb-4 flex items-center text-xs text-slate-500 sm:mb-8 sm:text-sm"
      >
        <a href="/" class="shrink-0 transition-colors hover:text-blue-600"
          >首页</a
        >
        <ChevronRight class="mx-1 h-3 w-3 shrink-0 sm:mx-2 sm:h-4 sm:w-4" />
        <a
          href="/accessory"
          class="shrink-0 transition-colors hover:text-blue-600"
          >配件商场</a
        >
        <ChevronRight class="mx-1 h-3 w-3 shrink-0 sm:mx-2 sm:h-4 sm:w-4" />
        <span
          class="max-w-30 truncate font-medium text-slate-800 sm:max-w-50"
          >{{ accessoryDetail?.name ?? "加载中..." }}</span
        >
      </nav>

      <div
        v-if="loading"
        class="rounded-2xl bg-white p-4 shadow-sm sm:rounded-3xl sm:p-6 lg:p-10"
      >
        <el-skeleton animated>
          <template #template>
            <div class="grid grid-cols-1 gap-6 lg:grid-cols-2 lg:gap-10">
              <el-skeleton-item
                variant="image"
                class="h-64 w-full rounded-2xl sm:h-100"
              />
              <div class="flex flex-col justify-center space-y-4 sm:space-y-6">
                <el-skeleton-item variant="h1" class="h-8 w-3/4 sm:h-10" />
                <div class="space-y-2">
                  <el-skeleton-item variant="text" class="w-full" />
                  <el-skeleton-item variant="text" class="w-full" />
                  <el-skeleton-item variant="text" class="w-2/3" />
                </div>
                <div class="mt-2 flex gap-4 sm:mt-4">
                  <el-skeleton-item variant="text" class="h-10 w-1/3 sm:h-12" />
                  <el-skeleton-item variant="text" class="h-10 w-1/3 sm:h-12" />
                </div>
                <div class="pt-4 sm:pt-6">
                  <el-skeleton-item
                    variant="button"
                    class="h-12 w-full sm:h-14 sm:w-40"
                  />
                </div>
              </div>
            </div>
          </template>
        </el-skeleton>
      </div>

      <div
        v-else-if="accessoryDetail"
        class="animate-fade-in space-y-8 sm:space-y-12"
      >
        <div
          class="rounded-2xl border border-gray-100 bg-white p-4 shadow-sm sm:rounded-3xl sm:p-6 lg:p-10"
        >
          <div class="grid grid-cols-1 gap-6 lg:grid-cols-2 lg:gap-10">
            <div
              class="group relative overflow-hidden rounded-2xl border border-gray-100 bg-gray-50 p-2 sm:p-4"
            >
              <div class="absolute top-2 left-2 z-10 sm:top-4 sm:left-4">
                <span
                  class="rounded-full bg-blue-600/90 px-2 py-1 text-[10px] font-bold text-white shadow-sm backdrop-blur-sm sm:px-3 sm:text-xs"
                >
                  官方正品
                </span>
              </div>
              <el-image
                :src="Baseurl + (accessoryDetail.cover ?? '')"
                :alt="accessoryDetail.name ?? ''"
                fit="contain"
                class="h-64 w-full rounded-lg transition-transform duration-500 group-hover:scale-105 sm:h-96 lg:h-full lg:min-h-100"
              >
                <template #error>
                  <div
                    class="flex h-full w-full flex-col items-center justify-center bg-gray-100 text-slate-400"
                  >
                    <ImageOff class="mb-2 h-8 w-8 opacity-50 sm:h-10 sm:w-10" />
                    <span class="text-xs">图片加载失败</span>
                  </div>
                </template>
              </el-image>
            </div>
            <div class="flex h-full flex-col">
              <div
                class="mb-4 flex flex-col items-start justify-between gap-4 border-b border-gray-100 pb-4 sm:mb-6 sm:flex-row sm:gap-6 sm:pb-6"
              >
                <div class="flex-1">
                  <h2
                    class="mb-2 text-xl font-bold tracking-tight text-slate-900 sm:mb-3 sm:text-3xl"
                  >
                    {{ accessoryDetail.name }}
                  </h2>
                  <div class="flex flex-wrap items-center gap-2">
                    <span
                      class="rounded bg-slate-100 px-2 py-0.5 text-[10px] font-medium text-slate-600 sm:px-2.5 sm:py-1 sm:text-xs"
                    >
                      {{ accessoryDetail.series_name ?? "默认分类" }}
                    </span>
                    <div
                      class="flex items-center rounded bg-green-50 px-2 py-0.5 text-[10px] font-medium text-green-600 sm:py-1 sm:text-xs"
                    >
                      <CheckCircle class="mr-1 h-3 w-3" />
                      现货供应
                    </div>
                    <div
                      class="flex items-center rounded bg-blue-50 px-2 py-0.5 text-[10px] font-medium text-blue-600 sm:py-1 sm:text-xs"
                    >
                      <ShieldCheck class="mr-1 h-3 w-3" />
                      原厂质保
                    </div>
                  </div>
                </div>
                <div class="w-full shrink-0 sm:w-auto sm:text-right">
                  <span class="mb-1 block text-xs text-slate-400">参考价</span>
                  <div class="text-red-600">
                    <span class="text-base font-bold sm:text-lg">¥</span>
                    <span
                      class="text-2xl font-black tracking-tight sm:text-3xl"
                      >{{ accessoryDetail.price ?? "面议" }}</span
                    >
                  </div>
                </div>
              </div>

              <div class="mb-6 flex-1">
                <h3
                  class="mb-2 text-xs font-bold tracking-widest text-slate-400 uppercase sm:mb-3 sm:text-sm"
                >
                  配件参数与描述
                </h3>
                <div
                  class="relative rounded-xl bg-slate-50/80 p-4 text-slate-600 ring-1 ring-slate-100/50 ring-inset sm:rounded-2xl sm:p-5"
                >
                  <p class="text-justify text-xs leading-relaxed sm:text-sm">
                    {{ accessoryDetail.description ?? "暂无配件描述" }}
                  </p>
                </div>
              </div>

              <div class="mb-6 sm:mb-8">
                <h3
                  class="mb-2 text-xs font-bold tracking-widest text-slate-400 uppercase sm:mb-3 sm:text-sm"
                >
                  采购与技术支持
                </h3>
                <div class="grid grid-cols-1 gap-3 sm:grid-cols-2 sm:gap-4">
                  <div
                    class="flex flex-col items-start rounded-xl border border-blue-100 bg-blue-50/30 p-3 transition-colors hover:bg-blue-50 sm:rounded-2xl sm:p-4"
                  >
                    <div
                      class="mb-2 flex h-7 w-7 items-center justify-center rounded-full bg-blue-100/50 text-blue-600 sm:h-8 sm:w-8"
                    >
                      <Phone class="h-3.5 w-3.5 sm:h-4 sm:w-4" />
                    </div>
                    <span class="text-[10px] text-slate-500 sm:text-xs"
                      >咨询热线</span
                    >
                    <span
                      class="mt-0.5 font-mono text-xs font-bold text-slate-800 sm:text-sm"
                      >400-123-4567</span
                    >
                  </div>
                  <div
                    class="flex flex-col items-start rounded-xl border border-blue-100 bg-blue-50/30 p-3 transition-colors hover:bg-blue-50 sm:rounded-2xl sm:p-4"
                  >
                    <div
                      class="mb-2 flex h-7 w-7 items-center justify-center rounded-full bg-blue-100/50 text-blue-600 sm:h-8 sm:w-8"
                    >
                      <Mail class="h-3.5 w-3.5 sm:h-4 sm:w-4" />
                    </div>
                    <span class="text-[10px] text-slate-500 sm:text-xs"
                      >电子邮箱</span
                    >
                    <span
                      class="mt-0.5 font-mono text-xs font-bold text-slate-800 sm:text-sm"
                      >sales@example.com</span
                    >
                  </div>
                </div>
              </div>

              <div class="mt-auto border-t border-gray-100 pt-4 sm:pt-6">
                <div class="mb-3 flex items-baseline gap-2 sm:mb-4">
                  <span class="text-xs font-medium text-slate-800 sm:text-sm"
                    >选择订购数量</span
                  >
                </div>
                <div class="flex flex-col gap-3 sm:flex-row sm:gap-4">
                  <el-input-number
                    v-model="quantity"
                    :min="1"
                    controls-position="right"
                    class="w-full! sm:w-36!"
                    size="large"
                  />
                  <el-button
                    :disabled="true"
                    type="primary"
                    size="large"
                    class="flex! h-10! w-full! flex-1! items-center justify-center rounded-xl! px-6! text-sm! font-bold! shadow-lg! shadow-blue-200! hover:shadow-blue-300! sm:h-auto! sm:w-64! sm:flex-initial! sm:px-8! sm:py-3! sm:text-base!"
                  >
                    <ShoppingCart
                      class="mr-1.5 h-4 w-4 sm:mr-2 sm:h-5 sm:w-5"
                    />
                    加入采购单
                  </el-button>
                </div>

                <div
                  class="mt-3 flex items-center text-xs text-slate-500 sm:mt-4 sm:text-sm"
                >
                  <Info
                    class="mr-1.5 h-3.5 w-3.5 shrink-0 text-blue-500 sm:h-4 sm:w-4"
                  />
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

        <div class="space-y-4 pb-12 sm:space-y-6 sm:pb-20">
          <div
            class="flex items-center space-x-3 border-b border-gray-200 pb-3 sm:space-x-4 sm:pb-4"
          >
            <div class="h-6 w-1 rounded-r bg-blue-600 sm:h-8"></div>
            <h2 class="text-lg font-bold text-slate-800 sm:text-2xl">
              配件详情展示
            </h2>
          </div>

          <div class="space-y-3 sm:space-y-4">
            <template
              v-if="
                accessoryDetail.previews && accessoryDetail.previews.length > 0
              "
            >
              <div
                v-for="(img, idx) in accessoryDetail.previews"
                :key="idx"
                class="overflow-hidden rounded-xl bg-white shadow-sm ring-1 ring-gray-100 sm:rounded-2xl"
              >
                <el-image
                  :src="Baseurl + img"
                  lazy
                  class="block w-full object-cover"
                  fit="cover"
                >
                  <template #placeholder>
                    <div
                      class="flex h-48 w-full items-center justify-center bg-gray-50 sm:h-64"
                    >
                      <span class="text-xs text-slate-400 sm:text-sm"
                        >加载中...</span
                      >
                    </div>
                  </template>
                  <template #error>
                    <div
                      class="flex h-32 w-full flex-col items-center justify-center bg-gray-50 text-slate-400 sm:h-40"
                    >
                      <ImageOff class="mb-2 h-6 w-6 opacity-40 sm:h-8 sm:w-8" />
                      <span class="text-xs">图片无法显示</span>
                    </div>
                  </template>
                </el-image>
              </div>
            </template>

            <div v-else class="flex justify-center py-8 sm:py-12">
              <el-empty description="暂无配件详情图" :image-size="120" />
            </div>
          </div>
        </div>
      </div>

      <div
        v-else
        class="flex h-[40vh] flex-col items-center justify-center text-slate-500 sm:h-[50vh]"
      >
        <el-empty description="未找到该配件信息" :image-size="120">
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
