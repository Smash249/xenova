<template>
  <section
    class="relative flex min-h-screen items-center overflow-hidden bg-slate-50 py-24 font-sans"
  >
    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>

    <div class="relative z-10 container mx-auto px-6">
      <div class="mx-auto mb-16 max-w-3xl text-center">
        <h2
          class="mb-6 text-4xl font-extrabold tracking-tight text-slate-900 md:text-5xl"
        >
          我们的<span
            class="bg-linear-to-r from-blue-600 to-cyan-500 bg-clip-text text-transparent"
            >业务领域</span
          >
        </h2>
        <p class="text-lg leading-relaxed text-slate-600">
          我们专注于自动化设备制造和企业咨询服务两大核心业务，致力于为客户提供全方位的智能制造解决方案，助力产业升级。
        </p>
      </div>

      <div
        class="grid grid-cols-1 gap-x-6 gap-y-12 pt-12 md:grid-cols-2 lg:grid-cols-4"
      >
        <div
          v-for="(service, index) in services"
          :key="index"
          class="group relative flex flex-col rounded-2xl border border-slate-200 bg-white p-6 shadow-sm transition-all duration-300 hover:-translate-y-2 hover:border-blue-300 hover:shadow-xl hover:shadow-blue-500/10"
        >
          <div class="-mt-14 mb-6 flex justify-center">
            <div
              class="flex h-16 w-16 items-center justify-center rounded-full border border-slate-100 bg-white shadow-md ring-8 ring-slate-50 transition-transform duration-300 group-hover:scale-110"
            >
              <div
                class="h-0 w-0 border-r-10 border-b-18 border-l-10 border-r-transparent border-b-blue-600 border-l-transparent"
              ></div>
            </div>
          </div>

          <div class="mb-6 text-center">
            <h3
              class="mb-3 text-2xl font-bold text-slate-800 transition-colors group-hover:text-blue-600"
            >
              {{ service.title }}
            </h3>
            <p class="min-h-14 text-sm leading-relaxed text-slate-500">
              {{ service.description }}
            </p>
          </div>

          <ul
            class="mt-auto cursor-pointer space-y-3 rounded-xl border border-slate-100 bg-slate-50 p-4 transition-colors group-hover:border-blue-50"
          >
            <li
              v-for="(feature, idx) in service.features"
              :key="idx"
              class="flex items-center gap-2 text-sm text-slate-700"
            >
              <Check class="h-4 w-4 shrink-0 text-cyan-500" />
              <span>{{ feature.text }}</span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { Check } from "lucide-vue-next"

interface ServiceFeature {
  text: string
}

interface ServiceCard {
  title: string
  description: string
  features: ServiceFeature[]
}

const services: ServiceCard[] = [
  {
    title: "包装自动化",
    description: "为食品、医药、日化等行业提供高效、智能的包装自动化解决方案。",
    features: [
      { text: "全自动包装生产线" },
      { text: "智能分拣与码垛系统" },
      { text: "定制化包装解决方案" },
      { text: "视觉检测集成" },
    ],
  },
  {
    title: "精密组装",
    description: "针对3C电子、汽车零部件提供高精度的自动组装与测试设备。",
    features: [
      { text: "高速自动组装机" },
      { text: "精密点胶与固化" },
      { text: "在线气密性测试" },
      { text: "柔性供料系统" },
    ],
  },
  {
    title: "物流输送",
    description: "构建智能工厂物流体系，实现物料的高效流转与仓储管理。",
    features: [
      { text: "AGV智能调度系统" },
      { text: "立体仓库堆垛机" },
      { text: "倍速链输送线" },
      { text: "WMS系统集成" },
    ],
  },
  {
    title: "视觉检测",
    description: "利用深度学习算法，解决复杂的表面缺陷检测与尺寸测量难题。",
    features: [
      { text: "AI外观缺陷检测" },
      { text: "高精度尺寸测量" },
      { text: "机器人引导定位" },
      { text: "OCR字符识别" },
    ],
  },
  {
    title: "激光应用",
    description: "提供激光打标、焊接、切割等工艺的自动化集成方案。",
    features: [
      { text: "精密激光焊接机" },
      { text: "在线激光打标" },
      { text: "激光切割工作站" },
      { text: "光路设计与优化" },
    ],
  },
  {
    title: "数字化咨询",
    description: "助力企业进行数字化转型，提供MES、SCADA等系统开发与实施。",
    features: [
      { text: "生产数据采集" },
      { text: "设备联网监控" },
      { text: "数字化车间规划" },
      { text: "产能分析与优化" },
    ],
  },
  {
    title: "非标定制",
    description: "根据客户特殊工艺需求，量身定制专属的自动化设备。",
    features: [
      { text: "工艺流程分析" },
      { text: "原型机开发验证" },
      { text: "复杂机构设计" },
      { text: "交钥匙工程服务" },
    ],
  },
  {
    title: "售后运维",
    description: "提供设备全生命周期管理，确保生产线稳定运行。",
    features: [
      { text: "设备预防性维护" },
      { text: "备件快速响应" },
      { text: "远程故障诊断" },
      { text: "操作人员培训" },
    ],
  },
]
</script>
