<template>
  <div
    class="rounded-2xl border border-blue-100 bg-linear-to-br from-blue-50 to-white p-6 shadow-lg shadow-blue-900/5"
  >
    <h2 class="mb-6 flex items-center text-xl font-bold text-blue-800">
      联系我们
      <span class="ml-2 h-2 w-2 animate-pulse rounded-full bg-green-400"></span>
    </h2>
    <div class="space-y-5">
      <div class="flex items-start text-slate-700">
        <div
          class="mt-1 mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
        >
          <MapPin class="h-4 w-4 text-blue-600" />
        </div>
        <div>
          <h3 class="mb-1 text-xs tracking-wider text-slate-500 uppercase">
            公司地址
          </h3>
          <p class="text-sm font-medium">星实（厦门）科技有限公司</p>
          <p class="mt-1 text-sm text-slate-500">集美区软件园</p>
        </div>
      </div>

      <div class="flex items-start text-slate-700">
        <div
          class="mt-1 mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
        >
          <Phone class="h-4 w-4 text-blue-600" />
        </div>
        <div>
          <h3 class="mb-1 text-xs tracking-wider text-slate-500 uppercase">
            咨询热线
          </h3>
          <p class="text-lg font-bold tracking-wide text-blue-900">
            0592-6818878
          </p>
        </div>
      </div>

      <div class="flex items-start text-slate-700">
        <div
          class="mt-1 mr-3 flex h-8 w-8 items-center justify-center rounded-full bg-blue-100"
        >
          <Mail class="h-4 w-4 text-blue-600" />
        </div>
        <div>
          <h3 class="mb-1 text-xs tracking-wider text-slate-500 uppercase">
            电子邮箱
          </h3>
          <p class="text-sm font-medium break-all">1888@xingshi.com</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Mail, MapPin, Phone } from "lucide-vue-next"
</script>
