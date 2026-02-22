<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="isOpen"
        class="fixed inset-0 z-50 flex items-center justify-center bg-slate-900/60 p-4 backdrop-blur-sm"
        @click="HandleClose"
      >
        <div
          class="relative flex max-h-[90vh] w-full max-w-3xl flex-col overflow-hidden rounded-2xl bg-white shadow-2xl"
          @click.stop
        >
          <div
            class="flex items-center justify-between border-b border-gray-100 bg-white px-6 py-4"
          >
            <h3 class="line-clamp-1 pr-8 text-lg font-bold text-slate-800">
              {{ honor?.title }}
            </h3>
            <button
              @click="HandleClose"
              class="absolute top-4 right-4 flex h-8 w-8 items-center justify-center rounded-full bg-slate-100 text-slate-500 transition-colors hover:bg-slate-200 hover:text-slate-800"
            >
              <X class="h-5 w-5" />
            </button>
          </div>

          <div class="overflow-y-auto p-6 md:p-8">
            <div class="flex flex-col gap-8 md:flex-row">
              <div
                class="flex w-full items-center justify-center rounded-xl border border-slate-100 bg-slate-50 p-4 md:w-1/2"
              >
                <img
                  v-if="honor?.image"
                  :src="BaseUrl + honor.image"
                  :alt="honor.title"
                  class="max-h-80 rounded-lg object-contain drop-shadow-sm"
                />
              </div>

              <div class="flex w-full flex-col space-y-6 md:w-1/2">
                <div
                  class="space-y-4 rounded-xl border border-blue-50 bg-blue-50/30 p-5"
                >
                  <div class="flex flex-col gap-1">
                    <span class="text-xs font-medium text-slate-500"
                      >颁发日期</span
                    >
                    <span
                      class="flex items-center text-sm font-semibold text-slate-800"
                    >
                      <Calendar class="mr-2 h-4 w-4 text-blue-500" />
                      {{ honor?.issue_date }}
                    </span>
                  </div>
                  <div class="h-px w-full bg-blue-100/50"></div>
                  <div class="flex flex-col gap-1">
                    <span class="text-xs font-medium text-slate-500"
                      >证书编号</span
                    >
                    <span
                      class="flex items-center font-mono text-sm font-semibold text-slate-800"
                    >
                      <ShieldCheck class="mr-2 h-4 w-4 text-blue-500" />
                      {{ honor?.cert_no || "无" }}
                    </span>
                  </div>
                  <div class="h-px w-full bg-blue-100/50"></div>
                  <div class="flex flex-col gap-1" v-if="honor?.issuer">
                    <span class="text-xs font-medium text-slate-500"
                      >颁发机构</span
                    >
                    <span
                      class="flex items-center text-sm font-semibold text-slate-800"
                    >
                      <Award class="mr-2 h-4 w-4 text-blue-500" />
                      {{ honor?.issuer }}
                    </span>
                  </div>
                </div>

                <div>
                  <h4
                    class="mb-3 flex items-center text-sm font-bold text-slate-800"
                  >
                    <span class="mr-2 h-4 w-1.5 rounded-full bg-blue-600"></span
                    >荣誉介绍
                  </h4>
                  <p
                    class="text-justify text-sm leading-relaxed text-slate-600"
                  >
                    {{ honor?.description }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { Award, Calendar, ShieldCheck, X } from "lucide-vue-next"

import type { CompanyHonor } from "@/types/honor"

defineProps<{
  isOpen: boolean
  honor: CompanyHonor | null
}>()

const emit = defineEmits(["close"])
const BaseUrl = import.meta.env.VITE_BASE_URL
function HandleClose() {
  emit("close")
}
</script>

<style scoped lang="scss">
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;

  .relative {
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;

  .relative {
    opacity: 0;
    transform: scale(0.95) translateY(15px);
  }
}
</style>
