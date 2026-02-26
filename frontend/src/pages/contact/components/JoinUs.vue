<template>
  <div class="flex flex-col gap-10">
    <div class="flex flex-col gap-2">
      <h3 class="flex items-center gap-3 text-2xl font-bold text-slate-900">
        <div class="h-8 w-1.5 rounded-full bg-blue-600"></div>
        诚聘英才
      </h3>
      <p class="pl-4 text-sm leading-relaxed text-slate-500">
        加入星实科技，与我们一起深耕非标自动化领域，用技术驱动智造未来。
      </p>
    </div>

    <div
      v-if="isLoading"
      class="flex min-h-75 flex-col items-center justify-center gap-4"
    >
      <Loader2Icon class="h-8 w-8 animate-spin text-blue-600" />
      <span class="text-sm font-medium text-slate-500"
        >正在加载职位列表...</span
      >
    </div>

    <div
      v-else-if="jobList.length === 0"
      class="flex min-h-75 flex-col items-center justify-center gap-4 rounded-2xl border border-dashed border-slate-200 bg-slate-50 p-8 text-center"
    >
      <div
        class="flex h-16 w-16 items-center justify-center rounded-full bg-slate-100"
      >
        <BriefcaseIcon class="h-8 w-8 text-slate-400" />
      </div>
      <p class="font-medium text-slate-600">抱歉，当前暂无开放的招聘职位</p>
      <p class="text-sm text-slate-400">
        您可以稍后回来查看，或直接将简历发送至我们的邮箱
      </p>
    </div>

    <div v-else class="flex flex-col gap-6">
      <div
        v-for="job in jobList"
        :key="job.id"
        class="group relative overflow-hidden rounded-2xl border border-slate-100 bg-white p-6 shadow-sm transition-all duration-300 hover:-translate-y-1 hover:border-blue-200 hover:shadow-xl hover:shadow-blue-500/5 sm:p-8"
      >
        <div
          class="mb-6 flex flex-col justify-between gap-4 border-b border-slate-100 pb-6 sm:flex-row sm:items-center"
        >
          <div class="flex items-center gap-3">
            <h4
              class="text-xl font-bold text-slate-900 transition-colors group-hover:text-blue-600"
            >
              {{ job.title }}
            </h4>
            <span
              v-if="job.status === 1"
              class="rounded-full bg-emerald-50 px-2.5 py-1 text-xs font-bold text-emerald-600 ring-1 ring-emerald-600/20 ring-inset"
            >
              招聘中
            </span>
          </div>
          <span class="text-lg font-bold text-orange-500">{{
            job.salaryRange || "面议"
          }}</span>
        </div>

        <div class="mb-8 flex flex-wrap gap-3">
          <div
            class="flex items-center gap-1.5 rounded-lg bg-slate-50 px-3 py-1.5 text-sm font-medium text-slate-600"
          >
            <MapPinIcon class="h-4 w-4 text-slate-400" />
            {{ job.location || "厦门" }}
          </div>
          <div
            v-if="job.department"
            class="flex items-center gap-1.5 rounded-lg bg-slate-50 px-3 py-1.5 text-sm font-medium text-slate-600"
          >
            <BriefcaseIcon class="h-4 w-4 text-slate-400" />
            {{ job.department }}
          </div>
          <div
            class="flex items-center gap-1.5 rounded-lg bg-slate-50 px-3 py-1.5 text-sm font-medium text-slate-600"
          >
            <UsersIcon class="h-4 w-4 text-slate-400" />
            招 {{ job.headcount || "若干" }}
          </div>
          <div
            class="flex items-center gap-1.5 rounded-lg bg-slate-50 px-3 py-1.5 text-sm font-medium text-slate-600"
          >
            <ClockIcon class="h-4 w-4 text-slate-400" />
            {{ job.experience || "经验不限" }}
          </div>
          <div
            class="flex items-center gap-1.5 rounded-lg bg-slate-50 px-3 py-1.5 text-sm font-medium text-slate-600"
          >
            <GraduationCapIcon class="h-4 w-4 text-slate-400" />
            {{ job.education || "学历不限" }}
          </div>
        </div>

        <div class="grid grid-cols-1 gap-8 md:grid-cols-2">
          <div class="space-y-3">
            <h5 class="flex items-center gap-2 font-bold text-slate-900">
              <div class="h-4 w-1 rounded-full bg-blue-600"></div>
              岗位职责
            </h5>
            <div
              class="text-sm leading-relaxed whitespace-pre-line text-slate-600"
              v-html="job.responsibilities"
            ></div>
          </div>
          <div class="space-y-3">
            <h5 class="flex items-center gap-2 font-bold text-slate-900">
              <div class="h-4 w-1 rounded-full bg-blue-600"></div>
              任职要求
            </h5>
            <div
              class="text-sm leading-relaxed whitespace-pre-line text-slate-600"
              v-html="job.requirements"
            ></div>
          </div>
        </div>

        <div
          class="mt-8 flex items-center justify-end border-t border-slate-100 pt-6"
        >
          <a
            :href="'mailto:admin@xsxm.tech?subject=应聘：' + job.title"
            class="inline-flex items-center justify-center gap-2 rounded-xl bg-blue-50 px-6 py-2.5 text-sm font-bold text-blue-600 transition-all hover:bg-blue-600 hover:text-white"
          >
            投递简历
          </a>
        </div>
      </div>
      <div v-if="totalItems > 0" class="mt-8 flex items-center justify-center">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="totalItems"
          :background="true"
          layout="total, prev, pager, next, jumper"
          @current-change="HandlePageChange"
          class="custom-tech-pagination"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetSystemJobPositionList } from "@/api/system"
import {
  BriefcaseIcon,
  ClockIcon,
  GraduationCapIcon,
  Loader2Icon,
  MapPinIcon,
  UsersIcon,
} from "lucide-vue-next"
import { onMounted, ref } from "vue"

import type { JobPosition } from "@/types/system"

const jobList = ref<JobPosition[]>([])
const isLoading = ref(true)
const currentPage = ref(1)
const pageSize = ref(10)
const totalItems = ref(0)

const FetchJobPositions = async () => {
  isLoading.value = true
  try {
    const result = await GetSystemJobPositionList(
      currentPage.value,
      pageSize.value
    )
    jobList.value = result?.data || []
    if (result?.paginate) {
      totalItems.value = result.paginate.total_count || 0
      currentPage.value = result.paginate.page || 1
      pageSize.value = result.paginate.page_size || 10
    }
  } catch (error) {
    console.error("获取职位列表失败:", error)
  } finally {
    isLoading.value = false
  }
}

const HandlePageChange = (page: number) => {
  currentPage.value = page
  FetchJobPositions()
  window.scrollTo({ top: 0, behavior: "smooth" })
}

onMounted(() => {
  FetchJobPositions()
})
</script>

<style scoped>
:deep(.custom-tech-pagination.el-pagination.is-background .btn-next),
:deep(.custom-tech-pagination.el-pagination.is-background .btn-prev),
:deep(.custom-tech-pagination.el-pagination.is-background .el-pager li) {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 0.75rem;
  color: #64748b;
  font-weight: 600;
  height: 2.5rem;
  min-width: 2.5rem;
  transition: all 0.3s ease;
}

:deep(.custom-tech-pagination.el-pagination.is-background .btn-next:hover),
:deep(.custom-tech-pagination.el-pagination.is-background .btn-prev:hover),
:deep(.custom-tech-pagination.el-pagination.is-background .el-pager li:hover) {
  color: #2563eb;
  border-color: #2563eb;
}

:deep(
  .custom-tech-pagination.el-pagination.is-background .el-pager li.is-active
) {
  background-color: #2563eb;
  border-color: #2563eb;
  color: #ffffff;
  box-shadow: 0 4px 6px -1px rgba(59, 130, 246, 0.3);
}

:deep(.custom-tech-pagination .el-pagination__total),
:deep(.custom-tech-pagination .el-pagination__jump) {
  color: #64748b;
  font-size: 0.875rem;
}

:deep(.custom-tech-pagination .el-input__wrapper) {
  border-radius: 0.5rem;
  background-color: #f8fafc;
  box-shadow: 0 0 0 1px #e2e8f0 inset;
}

:deep(.custom-tech-pagination .el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #2563eb inset !important;
}
</style>
