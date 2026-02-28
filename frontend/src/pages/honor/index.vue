<template>
  <div class="min-h-screen bg-gray-50 font-sans text-slate-800">
    <CustomBanner title="荣誉与文化" />

    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <nav class="mb-10 flex items-center text-sm text-slate-500">
        <a href="/" class="transition-colors hover:text-blue-600">首页</a>
        <ChevronRight class="mx-2 h-4 w-4" />
        <span class="font-medium text-slate-800">荣誉与文化</span>
      </nav>

      <div class="grid grid-cols-1 gap-10 lg:grid-cols-4">
        <aside class="lg:col-span-1">
          <div class="sticky top-8 space-y-8">
            <div
              class="rounded-2xl border border-gray-100 bg-white p-2 shadow-sm"
            >
              <ul class="space-y-1">
                <li v-for="cat in categories" :key="cat.id">
                  <button
                    @click="SetActiveCategory(cat.id)"
                    class="group relative flex w-full items-center overflow-hidden rounded-xl px-5 py-4 transition-all duration-200"
                    :class="[
                      activeCategory === cat.id
                        ? 'bg-blue-600 text-white shadow-md shadow-blue-200'
                        : 'text-slate-600 hover:bg-gray-50 hover:text-blue-600',
                    ]"
                  >
                    <component
                      :is="cat.icon"
                      class="z-10 mr-3 h-5 w-5 transition-transform duration-300"
                      :class="
                        activeCategory === cat.id
                          ? 'scale-110'
                          : 'group-hover:scale-110'
                      "
                    />
                    <span class="z-10 font-medium">{{ cat.name }}</span>

                    <div
                      v-if="activeCategory === cat.id"
                      class="pointer-events-none absolute top-0 right-0 h-full w-20 bg-linear-to-l from-white/10 to-transparent"
                    ></div>
                  </button>
                </li>
              </ul>
            </div>
            <div class="hidden lg:block">
              <ContactUs />
            </div>
          </div>
        </aside>

        <main class="flex min-h-150 flex-col lg:col-span-3">
          <div
            class="mb-8 flex items-center justify-between border-b border-gray-200 pb-4"
          >
            <h2 class="relative pl-4 text-2xl font-bold text-slate-800">
              <span
                class="absolute top-1/2 left-0 h-6 w-1 -translate-y-1/2 rounded-full bg-blue-600"
              ></span>
              {{ currentCategoryName }}
            </h2>
          </div>

          <div class="relative w-full flex-1">
            <Suspense>
              <template #default>
                <Transition name="fade-slide" mode="out-in">
                  <KeepAlive>
                    <component :is="currentComponent" :key="activeCategory" />
                  </KeepAlive>
                </Transition>
              </template>
              <template #fallback>
                <div
                  class="flex h-100 w-full flex-col items-center justify-center rounded-2xl border border-dashed border-gray-200 bg-white/50"
                >
                  <div
                    class="h-10 w-10 animate-spin rounded-full border-4 border-blue-100 border-t-blue-600"
                  ></div>
                  <span class="mt-4 text-sm font-medium text-slate-500"
                    >内容加载中...</span
                  >
                </div>
              </template>
            </Suspense>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Award, ChevronRight, FileCheck, Heart, Users } from "lucide-vue-next"
import { computed, defineAsyncComponent, markRaw, ref } from "vue"

import ContactUs from "@/components/contactUs/index.vue"
import CustomBanner from "@/components/customBanner/index.vue"

const CompanyHonor = defineAsyncComponent(
  () => import("./views/CompanyHonor.vue")
)
const LoveActivity = defineAsyncComponent(
  () => import("./views/LoveActivity.vue")
)
const CompanyPatent = defineAsyncComponent(
  () => import("./views/CompanyPatent.vue")
)
const CorporateCulture = defineAsyncComponent(
  () => import("./views/CorporateCulture.vue")
)

const categories = [
  { id: "honor", name: "公司荣誉", icon: markRaw(Award) },
  { id: "activity", name: "爱心活动", icon: markRaw(Heart) },
  { id: "patent", name: "公司专利", icon: markRaw(FileCheck) },
  { id: "culture", name: "企业文化", icon: markRaw(Users) },
]
const componentMap: Record<string, any> = {
  honor: CompanyHonor,
  activity: LoveActivity,
  patent: CompanyPatent,
  culture: CorporateCulture,
}

const activeCategory = ref("honor")

const currentComponent = computed(() => {
  return componentMap[activeCategory.value]
})

const currentCategoryName = computed(() => {
  return categories.find((c) => c.id === activeCategory.value)?.name || ""
})

function SetActiveCategory(id: string) {
  activeCategory.value = id
}
</script>

<style scoped>
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(15px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-15px);
}
</style>
