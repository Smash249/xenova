<template>
  <div class="relative min-h-screen w-full overflow-hidden">
    <div
      class="fixed inset-0 z-0 h-full w-full bg-cover bg-center bg-no-repeat transition-transform duration-[20s] hover:scale-110"
      style="
        background-image: url(&quot;https://images.unsplash.com/photo-1451187580459-43490279c0fa?q=80&w=2072&auto=format&fit=crop&quot;);
      "
    ></div>

    <div
      class="fixed inset-0 z-10 bg-linear-to-br from-blue-900/90 via-blue-800/80 to-indigo-900/80 mix-blend-multiply"
    ></div>
    <div class="fixed top-6 left-4 z-50 text-white sm:top-10 sm:left-10">
      <button
        @click="router.push('/')"
        class="group flex cursor-pointer items-center gap-2 rounded-full text-sm font-medium backdrop-blur-md transition-all hover:scale-105"
      >
        <el-icon class="transition-transform group-hover:-translate-x-1"
          ><Back
        /></el-icon>
        返回首页
      </button>
    </div>

    <div
      class="relative z-20 grid h-screen w-full grid-cols-1 items-center justify-evenly lg:grid-cols-2"
    >
      <div
        class="hidden h-full flex-col justify-evenly px-12 py-6 text-white lg:flex"
      >
        <div class="space-y-12">
          <div class="flex items-center gap-4">
            <div
              class="flex h-12 w-12 items-center justify-center rounded-xl bg-white/20 shadow-lg shadow-black/10 backdrop-blur-md"
            >
              <svg
                class="h-7 w-7 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13 10V3L4 14h7v7l9-11h-7z"
                />
              </svg>
            </div>
            <div>
              <h1 class="text-2xl font-bold tracking-wide">星实科技</h1>
              <p class="text-xs text-blue-200">StarTech Industries</p>
            </div>
          </div>
          <h2 class="text-5xl leading-tight font-bold tracking-tight">
            连接未来<br />
            <span
              class="bg-linear-to-r from-blue-200 to-indigo-200 bg-clip-text text-transparent"
              >重塑数字世界</span
            >
          </h2>
          <p class="max-w-md text-lg font-light text-blue-100/80">
            我们致力于提供最前沿的云计算解决方案，助您在数字化转型的浪潮中抢占先机。
          </p>
        </div>

        <div class="flex gap-6">
          <div
            class="rounded-2xl border border-white/10 bg-white/5 p-4 backdrop-blur-sm transition-colors hover:bg-white/10"
          >
            <div class="text-3xl font-bold">10k+</div>
            <div class="text-xs text-blue-200">企业用户</div>
          </div>
          <div
            class="rounded-2xl border border-white/10 bg-white/5 p-4 backdrop-blur-sm transition-colors hover:bg-white/10"
          >
            <div class="text-3xl font-bold">99.9%</div>
            <div class="text-xs text-blue-200">服务可用性</div>
          </div>
        </div>
      </div>

      <div
        class="relative flex h-full w-full flex-col items-center justify-center px-4 sm:px-8"
      >
        <div
          class="glass-form animate-fade-in-up w-full max-w-md rounded-2xl p-6 sm:rounded-3xl sm:p-8 md:p-10"
        >
          <!-- 移动端顶部品牌信息展示 -->
          <div class="mb-6 flex flex-col items-center justify-center lg:hidden">
            <div
              class="mb-3 flex h-12 w-12 items-center justify-center rounded-xl bg-white/20 shadow-lg shadow-black/10 backdrop-blur-md"
            >
              <svg
                class="h-7 w-7 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13 10V3L4 14h7v7l9-11h-7z"
                />
              </svg>
            </div>
            <h1 class="text-2xl font-bold tracking-wide text-white">
              星实科技
            </h1>
          </div>

          <div class="mb-8 text-center">
            <h2 class="text-2xl font-bold text-white">欢迎回来</h2>
            <p class="mt-2 text-sm text-blue-200">登录您的账户以继续</p>
          </div>

          <el-form
            ref="loginFormRef"
            :model="loginForm"
            :rules="loginRules"
            class="space-y-5"
            size="large"
            @submit.prevent
          >
            <el-form-item prop="email">
              <div class="group relative w-full">
                <el-input
                  v-model="loginForm.email"
                  placeholder="邮箱地址"
                  class="custom-input h-12"
                  autocomplete="off"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><User /></el-icon>
                  </template>
                </el-input>
              </div>
            </el-form-item>

            <el-form-item prop="password">
              <div class="group relative w-full">
                <el-input
                  v-model="loginForm.password"
                  type="password"
                  placeholder="密码"
                  show-password
                  class="custom-input h-12"
                  autocomplete="off"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><Lock /></el-icon>
                  </template>
                </el-input>
              </div>
            </el-form-item>

            <div class="flex items-center justify-between px-2 text-sm">
              <el-checkbox
                v-model="rememberMe"
                label="记住我"
                class="custom-checkbox"
              />
            </div>

            <el-button
              :loading="loading"
              class="h-12! w-full rounded-xl! border-0! bg-blue-600! font-bold! text-white! shadow-lg shadow-blue-900/50 transition-all hover:-translate-y-1 hover:bg-blue-500! active:scale-95"
              @click="HandleLogin"
            >
              立即登录
            </el-button>
          </el-form>

          <div class="mt-8 text-center text-sm text-blue-200/80">
            还没有账号?
            <router-link
              to="/register"
              class="font-bold text-white hover:text-blue-200 hover:underline"
            >
              免费注册
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import userStore from "@/store/modules/user"
import { Back, Lock, User } from "@element-plus/icons-vue"
import { ElMessage } from "element-plus"
import type { FormInstance, FormRules } from "element-plus"
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"

const router = useRouter()
const useUserStore = userStore()

const loginFormRef = ref<FormInstance>()
const loading = ref(false)
const rememberMe = ref(false)

const loginForm = reactive({
  email: "",
  password: "",
})

const loginRules: FormRules = {
  email: [
    { required: true, message: "请输入邮箱地址", trigger: "blur" },
    { type: "email", message: "请输入正确的邮箱格式", trigger: "blur" },
  ],
  password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    { min: 6, max: 20, message: "密码长度为 6-20 个字符", trigger: "blur" },
  ],
}

async function HandleLogin() {
  if (!loginFormRef.value) return

  await loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        await useUserStore.Login(loginForm)
        ElMessage.success("登录成功")
        router.push("/dashboard")
      } catch (error: any) {
        ElMessage.error(error.message || "登录失败，请检查您的凭据")
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
.glass-form {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

:deep(.custom-input .el-input__wrapper) {
  background-color: rgba(0, 0, 0, 0.2) !important;
  box-shadow: none !important;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 0.75rem;
  padding: 1px 15px;
  color: white;
  transition: all 0.3s ease;
}

:deep(.custom-input .el-input__wrapper:hover) {
  background-color: rgba(0, 0, 0, 0.3) !important;
  border-color: rgba(255, 255, 255, 0.3);
}

:deep(.custom-input .el-input__wrapper.is-focus) {
  background-color: rgba(0, 0, 0, 0.4) !important;
  border-color: #409eff;
  box-shadow: 0 0 0 1px #409eff !important;
}

:deep(.el-input__inner) {
  color: white !important;
}

:deep(.el-input__inner::placeholder) {
  color: rgba(255, 255, 255, 0.4);
}

:deep(.custom-checkbox .el-checkbox__label) {
  color: rgba(255, 255, 255, 0.8) !important;
}
:deep(.custom-checkbox .el-checkbox__inner) {
  background-color: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.4);
}
:deep(.custom-checkbox .el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: #409eff;
  border-color: #409eff;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
}
</style>
