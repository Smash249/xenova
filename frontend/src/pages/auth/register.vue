<template>
  <div class="relative min-h-screen w-full overflow-hidden">
    <div
      class="fixed inset-0 z-0 h-full w-full bg-cover bg-center bg-no-repeat transition-transform duration-[30s] hover:scale-110"
      style="
        background-image: url(&quot;https://images.unsplash.com/photo-1518770660439-4636190af475?q=80&w=2070&auto=format&fit=crop&quot;);
      "
    ></div>

    <div
      class="fixed inset-0 z-10 bg-linear-to-br from-blue-900/90 via-blue-800/80 to-indigo-900/80 mix-blend-multiply"
    ></div>

    <div
      class="relative z-20 grid h-screen w-full grid-cols-[3fr_2fr] items-center justify-center"
    >
      <div
        class="hidden h-full flex-col justify-between p-12 text-white lg:flex"
      >
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

        <div class="space-y-6">
          <h2 class="text-5xl leading-tight font-bold tracking-tight">
            加入我们<br />
            <span
              class="bg-linear-to-r from-blue-200 to-indigo-200 bg-clip-text text-transparent"
              >开启数字之旅</span
            >
          </h2>
          <p class="max-w-md text-lg font-light text-blue-100/80">
            创建一个账户，立即访问全球领先的数据分析平台和云计算资源。
          </p>
        </div>

        <div class="flex gap-4">
          <div
            class="flex items-center gap-3 rounded-xl border border-white/10 bg-white/5 px-4 py-3 backdrop-blur-sm"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-500/30"
            >
              <el-icon><DataLine /></el-icon>
            </div>
            <span class="text-sm">实时数据</span>
          </div>
          <div
            class="flex items-center gap-3 rounded-xl border border-white/10 bg-white/5 px-4 py-3 backdrop-blur-sm"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-500/30"
            >
              <el-icon><Monitor /></el-icon>
            </div>
            <span class="text-sm">多端同步</span>
          </div>
        </div>
      </div>

      <div class="flex h-full w-full items-center justify-center px-4">
        <div
          class="glass-form animate-fade-in-up custom-scrollbar max-h-[90vh] w-full max-w-120 overflow-y-auto rounded-3xl p-8 md:p-10"
        >
          <div class="mb-6 text-center">
            <h2 class="text-2xl font-bold text-white">创建新账户</h2>
            <p class="mt-2 text-sm text-blue-200">只需几步即可完成注册</p>
          </div>

          <el-form
            ref="registerFormRef"
            :model="registerForm"
            :rules="registerRules"
            class="space-y-5"
            size="large"
            autocomplete="off"
            @submit.prevent
          >
            <el-form-item prop="userName">
              <div class="group relative w-full">
                <el-input
                  v-model="registerForm.userName"
                  placeholder="用户名"
                  class="custom-input h-12"
                  autocomplete="new-username"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><User /></el-icon>
                  </template>
                </el-input>
              </div>
            </el-form-item>

            <el-form-item prop="email">
              <div class="group relative w-full">
                <el-input
                  v-model="registerForm.email"
                  placeholder="邮箱地址"
                  class="custom-input h-12"
                  autocomplete="new-email"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><Message /></el-icon>
                  </template>
                </el-input>
              </div>
            </el-form-item>

            <el-form-item prop="code">
              <div class="flex w-full gap-3">
                <el-input
                  v-model="registerForm.code"
                  placeholder="验证码"
                  class="custom-input h-12 flex-1"
                  autocomplete="one-time-code"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><Key /></el-icon>
                  </template>
                </el-input>
                <el-button
                  type="primary"
                  class="h-12! w-30 rounded-lg! border-0! bg-blue-600/80! text-white! hover:bg-blue-500!"
                  :disabled="codeDisabled"
                  @click="HandleSendCode"
                >
                  {{ codeText }}
                </el-button>
              </div>
            </el-form-item>

            <el-form-item prop="password">
              <div class="group relative w-full">
                <el-input
                  v-model="registerForm.password"
                  type="password"
                  placeholder="设置密码 (8-20个字符)"
                  show-password
                  class="custom-input h-12"
                  autocomplete="new-password"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><Lock /></el-icon>
                  </template>
                </el-input>
              </div>
            </el-form-item>

            <el-form-item prop="confirmPassword">
              <div class="group relative w-full">
                <el-input
                  v-model="registerForm.confirmPassword"
                  type="password"
                  placeholder="确认密码"
                  show-password
                  class="custom-input h-12"
                  autocomplete="new-password"
                >
                  <template #prefix>
                    <el-icon class="text-white/60"><CircleCheck /></el-icon>
                  </template>
                </el-input>
              </div>
            </el-form-item>

            <el-button
              :loading="loading"
              class="mt-2! h-12! w-full rounded-xl! border-0! bg-blue-600! font-bold! text-white! shadow-lg shadow-blue-900/50 transition-all hover:-translate-y-1 hover:bg-blue-500! active:scale-95"
              @click="HandleRegister"
            >
              立即注册
            </el-button>
          </el-form>

          <div class="mt-8 text-center text-sm text-blue-200/80">
            已有账户?
            <router-link
              to="/login"
              class="font-bold text-white hover:text-blue-200 hover:underline"
            >
              立即登录
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RegisterApi, SendEmailApi } from "@/api/auth"
import {
  CircleCheck,
  DataLine,
  Key,
  Lock,
  Message,
  Monitor,
  User,
} from "@element-plus/icons-vue"
import { ElMessage } from "element-plus"
import type { FormInstance, FormRules } from "element-plus"
import { onBeforeUnmount, reactive, ref } from "vue"
import { useRouter } from "vue-router"

const router = useRouter()

const registerFormRef = ref<FormInstance>()
const loading = ref(false)
const codeDisabled = ref(false)
const codeText = ref("发送验证码")
const countdown = ref(60)
const countdownTimer = ref<number | null>(null)

const registerForm = reactive({
  userName: "",
  email: "",
  code: "",
  password: "",
  confirmPassword: "",
})

const validateConfirmPassword = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请再次输入密码"))
  } else if (value !== registerForm.password) {
    callback(new Error("两次输入的密码不一致"))
  } else {
    callback()
  }
}

const registerRules: FormRules = {
  userName: [
    { required: true, message: "请输入用户名", trigger: "blur" },
    { min: 2, max: 20, message: "用户名长度为 2-20 个字符", trigger: "blur" },
  ],
  email: [
    { required: true, message: "请输入邮箱地址", trigger: "blur" },
    { type: "email", message: "请输入正确的邮箱格式", trigger: "blur" },
  ],
  code: [{ required: true, message: "请输入验证码", trigger: "blur" }],
  password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    { min: 8, max: 20, message: "密码长度为 8-20 个字符", trigger: "blur" },
  ],
  confirmPassword: [
    { required: true, message: "请再次输入密码", trigger: "blur" },
    { validator: validateConfirmPassword, trigger: "blur" },
  ],
}

async function HandleSendCode() {
  if (!registerForm.email) {
    ElMessage.warning("请先输入邮箱地址")
    return
  }
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(registerForm.email)) {
    ElMessage.warning("请输入有效的邮箱格式")
    return
  }
  try {
    await SendEmailApi(registerForm.email)
    ElMessage.success("验证码已发送，请查收邮箱")
    codeDisabled.value = true
    codeText.value = `重新发送 (${countdown.value}s)`
    countdownTimer.value = window.setInterval(() => {
      countdown.value -= 1
      if (countdown.value <= 0) {
        codeDisabled.value = false
        codeText.value = "发送验证码"
        countdown.value = 60
        if (countdownTimer.value) {
          clearInterval(countdownTimer.value)
        }
      } else {
        codeText.value = `重新发送 (${countdown.value}s)`
      }
    }, 1000)
  } catch (error: any) {
    ElMessage.error(error.message || "验证码发送失败，请稍后重试")
  }
}

async function HandleRegister() {
  if (!registerFormRef.value) return
  await registerFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        await RegisterApi(registerForm)
        ElMessage.success("注册成功，即将跳转到登录页面")
        setTimeout(() => {
          router.push("/login")
        }, 1500)
      } catch (error: any) {
        ElMessage.error(error.message || "注册失败，请稍后重试")
      } finally {
        loading.value = false
      }
    }
  })
}

onBeforeUnmount(() => {
  if (countdownTimer.value) {
    clearInterval(countdownTimer.value)
  }
})
</script>

<style scoped>
.glass-form {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}

.custom-scrollbar::-webkit-scrollbar {
  width: 0px;
  background: transparent;
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
  height: 100%;
}

:deep(.el-input__inner::placeholder) {
  color: rgba(255, 255, 255, 0.4);
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
