<script setup lang="ts">
import { useMutation } from '@pinia/colada'
import { NForm, NFormItem, NInput, NButton, NCarousel, useMessage } from 'naive-ui'
import {
  computed,
  defineAsyncComponent,
  onMounted,
  onUnmounted,
  reactive,
  ref,
  toRaw,
  useTemplateRef,
} from 'vue'

import topographySvg from '@/assets/topography.svg'
import { useInjection } from '@/composables'
import { mediaQueryInjectionKey } from '@/injection'
import router from '@/router'
import { toRefsPreferencesStore, useUserStore } from '@/stores'
import { twColor } from '@/utils/colors'

import type { FormItemRule } from 'naive-ui'

defineOptions({
  name: 'SignIn',
})

const message = useMessage()
const { isMaxSm } = useInjection(mediaQueryInjectionKey)

const { isDark } = toRefsPreferencesStore()

const { Login } = useUserStore()

const illustrations = [
  defineAsyncComponent(() => import('./component/Illustration1.vue')),
  defineAsyncComponent(() => import('./component/Illustration2.vue')),
  defineAsyncComponent(() => import('./component/Illustration3.vue')),
]

const isNavigating = ref(false)
const errorMessage = ref('')

const textureMaskParams = reactive({
  size: '666px 666px',
  x: 0,
  y: 0,
})

const textureStyle = computed(() => {
  return {
    filter: isDark.value ? 'invert(0.18)' : 'invert(0.9)',
    maskImage: `radial-gradient(circle 250px at ${textureMaskParams.x}px ${textureMaskParams.y}px, #000 0%, transparent 100%)`,
    WebkitMaskImage: `radial-gradient(circle 250px at ${textureMaskParams.x}px ${textureMaskParams.y}px, #000 0%, transparent 100%)`,
  }
})

const signInFormRef = useTemplateRef<InstanceType<typeof NForm>>('signInFormRef')

const signInForm = reactive({
  email: '',
  password: '',
})

const signInFormRules: Record<string, FormItemRule[]> = {
  email: [{ required: true, message: '请输入邮箱账号', trigger: ['input', 'blur'] }],
  password: [{ required: true, message: '请输入密码', trigger: ['input', 'blur'] }],
}

const { isLoading: isSignInLoading, mutate: signInMutation } = useMutation({
  mutation: Login,
  onSuccess: () => {
    errorMessage.value = ''
    message.success('登录成功')
    ToLayout()
  },
  onError: (error: any) => {
    // 适配常见的 axios 或者 fetch 错误响应格式
    const msg = error?.response?.data?.message || error?.message || '登录失败，请检查账号和密码'
    errorMessage.value = msg
    message.error(msg)
  },
})

const mergedLoading = computed(() => isSignInLoading.value || isNavigating.value)

function ToLayout() {
  const { r } = router.currentRoute.value.query

  isNavigating.value = true
  router
    .replace({
      path: (r as string) || '/',
    })
    .finally(() => {
      isNavigating.value = false
    })
}

function HandleSubmitClick() {
  errorMessage.value = '' // 提交前清空错误信息
  signInFormRef.value?.validate((errors) => {
    if (!errors) {
      signInMutation(toRaw(signInForm))
    }
  })
}

function UpdateTexturePosition(x: number, y: number) {
  textureMaskParams.x = x
  textureMaskParams.y = y
}

function OnMouseMove(e: MouseEvent) {
  UpdateTexturePosition(e.clientX, e.clientY)
}

function OnTouchMove(e: TouchEvent) {
  if (!e.touches[0]) return
  UpdateTexturePosition(e.touches[0].clientX, e.touches[0].clientY)
}

onMounted(() => {
  window.addEventListener('mousemove', OnMouseMove)
  window.addEventListener('touchmove', OnTouchMove)
})
onUnmounted(() => {
  window.removeEventListener('mousemove', OnMouseMove)
  window.removeEventListener('touchmove', OnTouchMove)
})
</script>

<template>
  <div
    class="relative flex h-svh items-center justify-center overflow-hidden bg-slate-50 p-6 transition-colors duration-300 dark:bg-neutral-950"
  >
    <div
      class="pointer-events-none absolute top-1/2 left-1/2 h-200 w-200 -translate-x-1/2 -translate-y-1/2 rounded-full bg-blue-400/10 blur-[120px] dark:bg-blue-600/10"
    ></div>
    <div
      class="pointer-events-none absolute top-0 right-0 h-125 w-125 translate-x-1/3 -translate-y-1/3 rounded-full bg-indigo-400/10 blur-[100px] dark:bg-indigo-600/10"
    ></div>

    <div
      class="absolute top-0 left-0 size-full bg-blue-50/40 transition-[background-color] dark:bg-neutral-900/50"
      :style="{
        'mask-image': `url(${topographySvg})`,
        '-webkit-mask-image': `url(${topographySvg})`,
        'mask-repeat': 'repeat',
        'mask-size': textureMaskParams.size,
        '-webkit-mask-repeat': 'repeat',
        '-webkit-mask-size': textureMaskParams.size,
      }"
    />
    <div
      class="pointer-events-none absolute inset-0 z-10 transition-[filter]"
      :style="{
        'background-image': `url(${topographySvg})`,
        'background-size': textureMaskParams.size,
        '-webkit-mask-repeat': 'no-repeat',
        'mask-repeat': 'no-repeat',
        ...textureStyle,
      }"
    />

    <div
      class="relative z-50 flex h-130 w-full max-w-225 justify-center overflow-hidden rounded-3xl border border-white/60 bg-white/60 shadow-[0_20px_60px_-15px_rgba(37,99,235,0.1)] backdrop-blur-xl transition-all sm:w-225 dark:border-neutral-800/60 dark:bg-neutral-900/60 dark:shadow-none"
    >
      <div
        v-if="!isMaxSm"
        class="relative flex-1 overflow-hidden bg-linear-to-br from-blue-50/50 to-indigo-50/50 p-8 transition-colors dark:from-neutral-800/50 dark:to-neutral-900/50"
      >
        <div
          class="absolute inset-0 bg-[linear-gradient(to_right,#80808012_1px,transparent_1px),linear-gradient(to_bottom,#80808012_1px,transparent_1px)] bg-size-[24px_24px]"
        ></div>

        <NCarousel
          draggable
          :show-dots="true"
          :show-arrow="false"
          dot-type="line"
          loop
          autoplay
          class="relative z-10 h-full drop-shadow-sm"
        >
          <div
            v-for="(illustration, index) in illustrations"
            :key="index"
            class="flex h-full items-center justify-center p-4"
          >
            <component
              :is="illustration"
              class="scale-110 transform transition-transform duration-700 hover:scale-105"
            />
          </div>
        </NCarousel>
      </div>

      <div
        class="relative flex w-full flex-col bg-white/80 px-12 py-10 transition-colors sm:w-105 dark:bg-neutral-850/80"
      >
        <div class="mt-8 flex flex-1 flex-col justify-center">
          <div class="mb-10">
            <h2 class="text-3xl font-bold tracking-tight text-neutral-800 dark:text-neutral-100">
              欢迎登录
            </h2>
            <p
              class="mt-2 text-xs font-semibold tracking-[0.2em] text-blue-500/80 uppercase transition-colors dark:text-blue-400/80"
            >
              Sign in to system
            </p>
          </div>

          <div class="w-full">
            <NForm
              ref="signInFormRef"
              :model="signInForm"
              :show-require-mark="false"
              :rules="signInFormRules"
              size="large"
            >
              <NFormItem
                path="email"
                class="mb-2"
              >
                <NInput
                  v-model:value="signInForm.email"
                  placeholder="请输入账号"
                  clearable
                  class="rounded-xl!"
                  :theme-overrides="
                    isDark
                      ? {
                          color: twColor('neutral', 800),
                          border: `1px solid ${twColor('neutral', 700)}`,
                          borderRadius: '12px',
                        }
                      : {
                          borderRadius: '12px',
                          borderHover: `1px solid ${twColor('blue', 400)}`,
                          borderFocus: `1px solid ${twColor('blue', 500)}`,
                        }
                  "
                  :input-props="{
                    autocomplete: 'off',
                  }"
                  @keydown.enter.prevent="HandleSubmitClick"
                >
                  <template #prefix>
                    <span class="mr-2 iconify size-5.5 text-neutral-400 ph--user-circle" />
                  </template>
                </NInput>
              </NFormItem>
              <NFormItem path="password">
                <NInput
                  v-model:value="signInForm.password"
                  placeholder="请输入密码"
                  type="password"
                  show-password-on="click"
                  clearable
                  class="rounded-xl!"
                  :theme-overrides="
                    isDark
                      ? {
                          color: twColor('neutral', 800),
                          border: `1px solid ${twColor('neutral', 700)}`,
                          borderRadius: '12px',
                        }
                      : {
                          borderRadius: '12px',
                          borderHover: `1px solid ${twColor('blue', 400)}`,
                          borderFocus: `1px solid ${twColor('blue', 500)}`,
                        }
                  "
                  :input-props="{
                    autocomplete: 'off',
                  }"
                  @keydown.enter.prevent="HandleSubmitClick"
                >
                  <template #prefix>
                    <span class="mr-2 iconify size-5.5 text-neutral-400 ph--lock-key" />
                  </template>
                </NInput>
              </NFormItem>

              <div class="mt-2">
                <NButton
                  type="primary"
                  :loading="mergedLoading"
                  :disabled="mergedLoading"
                  attr-type="button"
                  block
                  size="large"
                  class="h-12 rounded-xl! border-0! bg-linear-to-r! from-blue-500! to-indigo-500! font-medium! text-white! shadow-lg shadow-blue-500/30 transition-all hover:-translate-y-0.5 hover:shadow-blue-500/40"
                  @click="HandleSubmitClick"
                >
                  <span class="tracking-[0.15em]">登 录</span>
                </NButton>
              </div>
            </NForm>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
