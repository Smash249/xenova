<template>
  <div class="flex flex-col gap-10">
    <div class="flex flex-col gap-2">
      <h3 class="flex items-center gap-3 text-2xl font-bold text-slate-900">
        <div class="h-8 w-1.5 rounded-full bg-blue-600"></div>
        版权与法律声明
      </h3>
      <p class="pl-4 text-sm leading-relaxed text-slate-500">
        请在访问和使用星实（厦门）科技有限公司（以下简称“本公司”）网站前仔细阅读以下条款。访问本网站即表示您同意接受以下条款的约束。
      </p>
    </div>

    <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
      <div
        class="group relative overflow-hidden rounded-2xl border border-slate-100 bg-slate-50 p-8 transition-all duration-300 hover:-translate-y-1 hover:border-blue-200 hover:bg-white hover:shadow-xl hover:shadow-blue-500/5"
      >
        <div
          class="mb-5 flex h-12 w-12 items-center justify-center rounded-xl bg-blue-100 text-blue-600 transition-colors group-hover:bg-blue-600 group-hover:text-white"
        >
          <CopyrightIcon class="h-6 w-6" />
        </div>
        <h4 class="mb-3 text-lg font-bold text-slate-900">知识产权声明</h4>
        <p class="text-sm leading-relaxed text-slate-600">
          本网站包含的所有内容（包括但不限于文字描述、产品图片、技术文档、视音频资料、代码及版面设计等）的所有权、著作权及其他知识产权均归<strong
            class="font-bold text-slate-800"
            >星实科技</strong
          >所有，或已取得合法授权。受国家著作权法及相关国际条约保护。
        </p>
      </div>

      <div
        class="group relative overflow-hidden rounded-2xl border border-slate-100 bg-slate-50 p-8 transition-all duration-300 hover:-translate-y-1 hover:border-blue-200 hover:bg-white hover:shadow-xl hover:shadow-blue-500/5"
      >
        <div
          class="mb-5 flex h-12 w-12 items-center justify-center rounded-xl bg-indigo-100 text-indigo-600 transition-colors group-hover:bg-indigo-600 group-hover:text-white"
        >
          <ShieldAlertIcon class="h-6 w-6" />
        </div>
        <h4 class="mb-3 text-lg font-bold text-slate-900">网站使用规范</h4>
        <p class="text-sm leading-relaxed text-slate-600">
          未经本公司事先书面许可，任何单位或个人不得商业目的复制、转载、摘编、修改、镜像、传播或展示本网站的任何内容。仅限在注明出处且不用于商业用途的前提下，方可进行个人学习或研究引用。
        </p>
      </div>

      <div
        class="group relative overflow-hidden rounded-2xl border border-slate-100 bg-slate-50 p-8 transition-all duration-300 hover:-translate-y-1 hover:border-blue-200 hover:bg-white hover:shadow-xl hover:shadow-blue-500/5"
      >
        <div
          class="mb-5 flex h-12 w-12 items-center justify-center rounded-xl bg-emerald-100 text-emerald-600 transition-colors group-hover:bg-emerald-600 group-hover:text-white"
        >
          <AwardIcon class="h-6 w-6" />
        </div>
        <h4 class="mb-3 text-lg font-bold text-slate-900">商标权保护</h4>
        <p class="text-sm leading-relaxed text-slate-600">
          本网站所使用的“星实科技”、“星实自动化”及相关产品名称、企业徽标（Logo）、图形设计等均属本公司的注册商标或商业标识。任何未经授权的擅自使用，本公司都将依法追究其侵权责任。
        </p>
      </div>

      <div
        class="group relative overflow-hidden rounded-2xl border border-slate-100 bg-slate-50 p-8 transition-all duration-300 hover:-translate-y-1 hover:border-blue-200 hover:bg-white hover:shadow-xl hover:shadow-blue-500/5"
      >
        <div
          class="mb-5 flex h-12 w-12 items-center justify-center rounded-xl bg-orange-100 text-orange-600 transition-colors group-hover:bg-orange-600 group-hover:text-white"
        >
          <ScaleIcon class="h-6 w-6" />
        </div>
        <h4 class="mb-3 text-lg font-bold text-slate-900">免责条款</h4>
        <p class="text-sm leading-relaxed text-slate-600">
          我们致力于提供准确的产品与技术信息，但不对内容的绝对准确性、时效性及完整性作出保证。对于因依赖本站信息或因系统维护、网络故障导致的任何直接或间接损失，本公司不承担任何法律责任。
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  AwardIcon,
  CopyrightIcon,
  ScaleIcon,
  ShieldAlertIcon,
} from "lucide-vue-next"
</script>
