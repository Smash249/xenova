<template>
  <div class="flex flex-col gap-10">
    <div class="flex flex-col gap-2">
      <h3 class="flex items-center gap-3 text-2xl font-bold text-slate-900">
        <div class="h-8 w-1.5 rounded-full bg-blue-600"></div>
        联系方式
      </h3>
      <p class="pl-4 text-sm text-slate-500">
        欢迎您通过以下方式与我们取得联系，我们将竭诚为您提供专业的自动化解决方案与支持。
      </p>
    </div>

    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:col-span-2">
      <div
        class="group flex items-start gap-4 rounded-xl border border-slate-100 bg-slate-50 p-5 transition-all hover:border-blue-200 hover:bg-white hover:shadow-md hover:shadow-blue-50"
      >
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white text-blue-600 shadow-sm ring-1 ring-slate-100 transition-colors group-hover:bg-blue-600 group-hover:text-white group-hover:ring-blue-600"
        >
          <Building2Icon class="h-6 w-6" />
        </div>
        <div>
          <div
            class="mb-1 text-xs font-bold tracking-wider text-slate-400 uppercase"
          >
            公司名称
          </div>
          <div class="text-sm leading-relaxed font-medium text-slate-800">
            星实（厦门）科技有限公司
          </div>
        </div>
      </div>

      <div
        class="group flex items-start gap-4 rounded-xl border border-slate-100 bg-slate-50 p-5 transition-all hover:border-blue-200 hover:bg-white hover:shadow-md hover:shadow-blue-50"
      >
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white text-blue-600 shadow-sm ring-1 ring-slate-100 transition-colors group-hover:bg-blue-600 group-hover:text-white group-hover:ring-blue-600"
        >
          <MapPinIcon class="h-6 w-6" />
        </div>
        <div>
          <div
            class="mb-1 text-xs font-bold tracking-wider text-slate-400 uppercase"
          >
            公司地址
          </div>
          <div class="text-sm leading-relaxed font-medium text-slate-800">
            福建省厦门市集美区<br />软件园三期F18
          </div>
        </div>
      </div>

      <div
        class="group flex items-start gap-4 rounded-xl border border-slate-100 bg-slate-50 p-5 transition-all hover:border-blue-200 hover:bg-white hover:shadow-md hover:shadow-blue-50"
      >
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white text-blue-600 shadow-sm ring-1 ring-slate-100 transition-colors group-hover:bg-blue-600 group-hover:text-white group-hover:ring-blue-600"
        >
          <PhoneIcon class="h-6 w-6" />
        </div>
        <div>
          <div
            class="mb-1 text-xs font-bold tracking-wider text-slate-400 uppercase"
          >
            联系电话
          </div>
          <div class="font-mono text-sm font-bold text-slate-800">
            0592-6818878<br />18064470280
          </div>
        </div>
      </div>

      <div
        class="group flex items-start gap-4 rounded-xl border border-slate-100 bg-slate-50 p-5 transition-all hover:border-blue-200 hover:bg-white hover:shadow-md hover:shadow-blue-50"
      >
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white text-blue-600 shadow-sm ring-1 ring-slate-100 transition-colors group-hover:bg-blue-600 group-hover:text-white group-hover:ring-blue-600"
        >
          <PrinterIcon class="h-6 w-6" />
        </div>
        <div>
          <div
            class="mb-1 text-xs font-bold tracking-wider text-slate-400 uppercase"
          >
            企业传真
          </div>
          <div class="font-mono text-sm font-medium text-slate-800">
            +86 0592-6818878
          </div>
        </div>
      </div>

      <div
        class="group flex items-start gap-4 rounded-xl border border-slate-100 bg-slate-50 p-5 transition-all hover:border-blue-200 hover:bg-white hover:shadow-md hover:shadow-blue-50"
      >
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white text-blue-600 shadow-sm ring-1 ring-slate-100 transition-colors group-hover:bg-blue-600 group-hover:text-white group-hover:ring-blue-600"
        >
          <MailIcon class="h-6 w-6" />
        </div>
        <div>
          <div
            class="mb-1 text-xs font-bold tracking-wider text-slate-400 uppercase"
          >
            邮箱地址
          </div>
          <a
            href="mailto:admin@xsxm.tech"
            class="text-sm font-bold text-slate-800 transition-colors hover:text-blue-600"
          >
            admin@xsxm.tech
          </a>
        </div>
      </div>

      <div
        class="group flex items-start gap-4 rounded-xl border border-slate-100 bg-slate-50 p-5 transition-all hover:border-blue-200 hover:bg-white hover:shadow-md hover:shadow-blue-50"
      >
        <div
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-white text-blue-600 shadow-sm ring-1 ring-slate-100 transition-colors group-hover:bg-blue-600 group-hover:text-white group-hover:ring-blue-600"
        >
          <ClockIcon class="h-6 w-6" />
        </div>
        <div>
          <div
            class="mb-1 text-xs font-bold tracking-wider text-slate-400 uppercase"
          >
            工作时间
          </div>
          <div class="text-sm font-medium text-slate-800">
            周一至周五<br />9:00 - 18:00
          </div>
        </div>
      </div>
    </div>

    <div
      class="group relative h-75 w-full overflow-hidden rounded-2xl border border-slate-200 bg-slate-100 shadow-sm transition-all hover:shadow-md sm:h-100"
    >
      <div
        class="pointer-events-none absolute inset-0 z-10 flex items-center justify-center bg-slate-900/10 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
      >
        <span
          class="rounded-full bg-white/90 px-4 py-2 text-sm font-bold text-blue-600 shadow-sm backdrop-blur-xs"
          >点击查看大图或导航</span
        >
      </div>

      <img
        src="https://via.placeholder.com/1200x400?text=Baidu+Map+Placeholder"
        alt="公司地图"
        class="h-full w-full object-cover transition-transform duration-700 group-hover:scale-105"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Building2Icon,
  ClockIcon,
  MailIcon,
  MapPinIcon,
  PhoneIcon,
  PrinterIcon,
} from "lucide-vue-next"
</script>
