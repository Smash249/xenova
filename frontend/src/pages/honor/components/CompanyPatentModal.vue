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
            <div class="flex items-center gap-3 pr-8">
              <span
                class="shrink-0 rounded-md px-2 py-1 text-xs font-bold"
                :class="GetTypeColor(patent?.type || '')"
              >
                {{ patent?.type }}
              </span>
              <h3 class="line-clamp-1 text-lg font-bold text-slate-800">
                {{ patent?.title }}
              </h3>
            </div>
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
                class="relative flex w-full items-center justify-center overflow-hidden rounded-xl border border-slate-100 bg-slate-50 p-4 md:w-1/2"
              >
                <div
                  class="absolute inset-0 bg-[linear-gradient(to_right,#e2e8f0_1px,transparent_1px),linear-gradient(to_bottom,#e2e8f0_1px,transparent_1px)] bg-size-[24px_24px] opacity-50"
                ></div>
                <img
                  v-if="patent?.image"
                  :src="patent.image"
                  :alt="patent.title"
                  class="relative z-10 max-h-80 rounded-lg object-contain mix-blend-multiply drop-shadow-sm"
                />
              </div>

              <div class="flex w-full flex-col space-y-6 md:w-1/2">
                <div
                  class="space-y-4 rounded-xl border border-indigo-50 bg-indigo-50/30 p-5"
                >
                  <div class="flex flex-col gap-1">
                    <span class="text-xs font-medium text-slate-500"
                      >专利申请号/专利号</span
                    >
                    <span
                      class="flex items-center font-mono text-sm font-semibold text-slate-800"
                    >
                      <FileCheck class="mr-2 h-4 w-4 text-indigo-500" />
                      {{ patent?.patentNo }}
                    </span>
                  </div>
                  <div class="h-px w-full bg-indigo-100/50"></div>
                  <div class="flex flex-col gap-1">
                    <span class="text-xs font-medium text-slate-500"
                      >授权公告日</span
                    >
                    <span
                      class="flex items-center text-sm font-semibold text-slate-800"
                    >
                      <Calendar class="mr-2 h-4 w-4 text-indigo-500" />
                      {{ patent?.date }}
                    </span>
                  </div>
                  <div class="h-px w-full bg-indigo-100/50"></div>
                  <div class="flex flex-col gap-1">
                    <span class="text-xs font-medium text-slate-500"
                      >发明人/设计人</span
                    >
                    <span
                      class="flex items-center text-sm font-semibold text-slate-800"
                    >
                      <Users class="mr-2 h-4 w-4 text-indigo-500" />
                      {{ patent?.inventor }}
                    </span>
                  </div>
                </div>

                <div>
                  <h4
                    class="mb-3 flex items-center text-sm font-bold text-slate-800"
                  >
                    <span
                      class="mr-2 h-4 w-1.5 rounded-full bg-indigo-600"
                    ></span
                    >摘要说明
                  </h4>
                  <p
                    class="text-justify text-sm leading-relaxed text-slate-600"
                  >
                    {{ patent?.summary }}
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
import { Calendar, FileCheck, Users, X } from "lucide-vue-next"

import type { CompanyPatent } from "@/types/honor"

defineProps<{
  isOpen: boolean
  patent: CompanyPatent | null
}>()

const emit = defineEmits(["close"])

function HandleClose() {
  emit("close")
}

function GetTypeColor(type: string): string {
  if (type.includes("发明")) {
    return "bg-orange-100 text-orange-700 border border-orange-200"
  } else if (type.includes("实用新型")) {
    return "bg-indigo-100 text-indigo-700 border border-indigo-200"
  } else if (type.includes("外观设计")) {
    return "bg-teal-100 text-teal-700 border border-teal-200"
  }
  return "bg-slate-100 text-slate-700 border border-slate-200"
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
