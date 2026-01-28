<template>
  <section
    id="process-section"
    class="flex-col-center relative min-h-screen overflow-hidden bg-slate-50 py-24 font-sans"
  >
    <!-- 背景 -->
    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>

    <div class="relative z-10 mx-auto max-w-360 px-6 sm:px-8 lg:px-12">
      <!-- 标题区域 -->
      <div class="mb-20 space-y-3 text-center">
        <h3
          class="text-base font-bold tracking-[0.2em] text-blue-600 uppercase"
        >
          Advantage
        </h3>
        <h2 class="text-4xl font-extrabold text-slate-900 md:text-5xl">
          标准化生产流程展示
        </h2>
        <div class="mx-auto mt-6 h-1.5 w-20 rounded-full bg-blue-600"></div>
      </div>

      <!-- 流程网格 -->
      <div
        class="grid grid-cols-2 gap-x-10 gap-y-16 md:grid-cols-4 lg:grid-cols-7"
      >
        <div
          v-for="(step, index) in steps"
          :key="step.id"
          class="group relative flex flex-col items-center"
          :class="{
            'translate-y-10 opacity-0': !isVisible,
            'animate-fade-up': isVisible,
          }"
          :style="{
            animationDelay: `${index * 50}ms`,
            animationFillMode: 'forwards',
          }"
        >
          <!-- 连接线 (PC端显示) - 位置随卡片尺寸调整 -->
          <div
            v-if="index < steps.length - 1 && (index + 1) % 7 !== 0"
            class="absolute top-20 left-1/2 -z-10 hidden h-0.5 w-full bg-slate-200 transition-colors duration-300 group-hover:bg-blue-300 lg:block"
          ></div>

          <!-- 圆形卡片容器 -->
          <div
            class="relative flex h-40 w-40 cursor-default flex-col items-center justify-center rounded-full border border-slate-200 bg-white p-4 shadow-sm transition-all duration-300 group-hover:-translate-y-2 group-hover:border-blue-300 group-hover:shadow-lg"
          >
            <!-- 序号球  -->
            <div
              class="absolute -top-4 left-1/2 z-10 flex h-9 w-9 -translate-x-1/2 items-center justify-center rounded-full border-[3px] border-white bg-blue-600 text-sm font-bold text-white shadow-md"
            >
              {{ step.id }}
            </div>

            <!-- 虚线圆环 -->
            <div
              class="absolute inset-2 rounded-full border border-dashed border-slate-200 transition-colors duration-300 group-hover:border-blue-200"
            ></div>

            <!-- 内容  -->
            <h4
              class="mb-1.5 text-center text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-600"
            >
              {{ step.title }}
            </h4>
            <p
              class="px-2 text-center text-xs leading-relaxed font-medium text-slate-500"
            >
              {{ step.description }}
            </p>
          </div>

          <!-- 移动端箭头 -->
          <div
            class="mt-6 text-slate-300 lg:hidden"
            v-if="index < steps.length - 1"
          >
            <ArrowRight class="h-6 w-6 rotate-90" />
          </div>
        </div>
      </div>

      <!-- 底部行动按钮 -->
      <div class="mt-24 text-center">
        <button
          class="inline-flex transform items-center justify-center rounded-full bg-blue-600 px-12 py-4 text-xl font-bold text-white shadow-xl shadow-blue-200 transition-all duration-300 hover:-translate-y-1 hover:bg-blue-700 hover:shadow-blue-400"
        >
          立即联系我们
          <CheckCircle2 class="ml-3 h-6 w-6" />
        </button>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { ArrowRight, CheckCircle2 } from "lucide-vue-next"
import { onMounted, ref } from "vue"

interface ProcessStep {
  id: number
  title: string
  description: string
}

const steps: ProcessStep[] = [
  { id: 1, title: "需求沟通", description: "深入沟通了解" },
  { id: 2, title: "工程师", description: "出具方案图" },
  { id: 3, title: "方案图", description: "团队严格审核" },
  { id: 4, title: "报价", description: "确认预算方案" },
  { id: 5, title: "签订合同", description: "预付定金" },
  { id: 6, title: "项目书", description: "确认最终细节" },
  { id: 7, title: "制造交期", description: "最晚45天发货" },
  { id: 8, title: "大件运输", description: "指定地点收货" },
  { id: 9, title: "设备吊装", description: "免费专业吊装" },
  { id: 10, title: "设备安装", description: "7天效果调试" },
  { id: 11, title: "预投试产", description: "确保运行稳定" },
  { id: 12, title: "客户确认", description: "生产确认验收" },
  { id: 13, title: "售后服务", description: "7*12小时响应" },
  { id: 14, title: "项目完结", description: "收受尾款" },
]

const isVisible = ref(false)

onMounted(() => {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          isVisible.value = true
        }
      })
    },
    { threshold: 0.1 }
  )

  const section = document.getElementById("process-section")
  if (section) observer.observe(section)
})
</script>

<style scoped>
@keyframes fade-up {
  from {
    opacity: 0;
    transform: translateY(40px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-up {
  animation-name: fade-up;
  animation-duration: 0.6s;
  animation-timing-function: cubic-bezier(0.16, 1, 0.3, 1);
}
</style>
