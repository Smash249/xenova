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
          <div class="relative h-64 w-full shrink-0 sm:h-80">
            <img
              v-if="activity?.cover"
              :src="BaseUrl + activity.cover"
              :alt="activity.title"
              class="h-full w-full object-cover"
            />
            <div
              class="absolute inset-0 bg-linear-to-t from-slate-900/80 via-slate-900/20 to-transparent"
            ></div>

            <button
              @click="HandleClose"
              class="absolute top-4 right-4 flex h-8 w-8 items-center justify-center rounded-full bg-black/20 text-white backdrop-blur-md transition-colors hover:bg-black/40"
            >
              <X class="h-5 w-5" />
            </button>

            <div class="absolute bottom-0 left-0 w-full p-6">
              <div class="mb-3 flex flex-wrap gap-2">
                <span
                  class="inline-flex items-center rounded-full bg-rose-500/90 px-2.5 py-1 text-xs font-medium text-white backdrop-blur-sm"
                >
                  <Calendar class="mr-1.5 h-3.5 w-3.5" />
                  {{ activity?.activity_date }}
                </span>
                <span
                  class="inline-flex items-center rounded-full bg-white/20 px-2.5 py-1 text-xs font-medium text-white backdrop-blur-sm"
                >
                  <MapPin class="mr-1.5 h-3.5 w-3.5" />
                  {{ activity?.location }}
                </span>
                <span
                  class="inline-flex items-center rounded-full bg-white/20 px-2.5 py-1 text-xs font-medium text-white backdrop-blur-sm"
                >
                  <Users class="mr-1.5 h-3.5 w-3.5" />
                  {{ activity?.participants }}人参与
                </span>
              </div>
              <h3
                class="text-2xl font-bold text-white shadow-black/10 drop-shadow-md sm:text-3xl"
              >
                {{ activity?.title }}
              </h3>
            </div>
          </div>

          <div class="overflow-y-auto p-6 md:p-8">
            <div class="prose prose-slate prose-rose max-w-none">
              <h4
                class="mb-4 flex items-center text-lg font-bold text-slate-800"
              >
                <Heart class="mr-2 h-5 w-5 text-rose-500" />
                活动纪实
              </h4>
              <div
                class="space-y-4 text-sm leading-loose text-slate-600 md:text-base"
              >
                <p
                  v-for="(paragraph, index) in FormatContent(activity?.content)"
                  :key="index"
                >
                  {{ paragraph }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { Calendar, Heart, MapPin, Users, X } from "lucide-vue-next"

import type { LoveActivity } from "@/types/honor"

defineProps<{
  isOpen: boolean
  activity: LoveActivity | null
}>()

const emit = defineEmits(["close"])
const BaseUrl = import.meta.env.VITE_BASE_URL

function HandleClose() {
  emit("close")
}

function FormatContent(content?: string): string[] {
  if (!content) return []
  return content.split("\n").filter((p) => p.trim() !== "")
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
