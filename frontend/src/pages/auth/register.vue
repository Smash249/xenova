<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 via-white to-blue-50 px-4">
    <div class="w-full max-w-md">
      <div class="bg-white rounded-2xl shadow-xl p-8 space-y-6">
        <!-- Header -->
        <div class="text-center space-y-2">
          <h1 class="text-3xl font-bold text-gray-900">创建账户</h1>
          <p class="text-gray-600">注册新用户</p>
        </div>

        <!-- Form -->
        <el-form
          ref="registerFormRef"
          :model="registerForm"
          :rules="registerRules"
          class="space-y-4"
          @submit.prevent="handleRegister"
        >
          <el-form-item prop="UserName">
            <el-input
              v-model="registerForm.UserName"
              placeholder="用户名"
              size="large"
              :prefix-icon="User"
            />
          </el-form-item>

          <el-form-item prop="email">
            <el-input
              v-model="registerForm.email"
              placeholder="邮箱地址"
              size="large"
              :prefix-icon="Message"
            />
          </el-form-item>

          <el-form-item prop="code">
            <div class="flex gap-2">
              <el-input
                v-model="registerForm.code"
                placeholder="验证码"
                size="large"
                :prefix-icon="Key"
                class="flex-1"
              />
              <el-button
                type="primary"
                size="large"
                :disabled="codeDisabled"
                @click="handleSendCode"
              >
                {{ codeText }}
              </el-button>
            </div>
          </el-form-item>

          <el-form-item prop="password">
            <el-input
              v-model="registerForm.password"
              type="password"
              placeholder="密码 (8-20个字符)"
              size="large"
              :prefix-icon="Lock"
              show-password
            />
          </el-form-item>

          <el-form-item prop="confirmPassword">
            <el-input
              v-model="registerForm.confirmPassword"
              type="password"
              placeholder="确认密码"
              size="large"
              :prefix-icon="Lock"
              show-password
            />
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              size="large"
              class="w-full"
              :loading="loading"
              @click="handleRegister"
            >
              注册
            </el-button>
          </el-form-item>
        </el-form>

        <!-- Footer -->
        <div class="text-center text-sm text-gray-600">
          已有账户？
          <router-link
            to="/login"
            class="text-blue-600 hover:text-blue-700 font-medium"
          >
            立即登录
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onBeforeUnmount } from "vue"
import { useRouter } from "vue-router"
import { ElMessage } from "element-plus"
import { User, Lock, Message, Key } from "@element-plus/icons-vue"
import userStore from "@/store/modules/user"
import type { FormInstance, FormRules } from "element-plus"

const router = useRouter()
const useUserStore = userStore()

const registerFormRef = ref<FormInstance>()
const loading = ref(false)
const codeDisabled = ref(false)
const codeText = ref("发送验证码")
const countdown = ref(60)
const countdownTimer = ref<NodeJS.Timeout | null>(null)

const registerForm = reactive({
  UserName: "",
  email: "",
  code: "",
  password: "",
  confirmPassword: "",
})

const validateConfirmPassword = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请再次输入密码"))
  } else if (value !== registerForm.password) {
    callback(new Error("两次输入的密码不一致"))
  } else {
    callback()
  }
}

const registerRules: FormRules = {
  UserName: [
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

const handleSendCode = async () => {
  if (!registerForm.email) {
    ElMessage.warning("请先输入邮箱地址")
    return
  }

  // TODO: 实现发送验证码的 API 调用
  // 这里暂时模拟发送验证码
  ElMessage.info("验证码发送功能需要后端支持，请联系管理员获取验证码")
  
  codeDisabled.value = true
  countdownTimer.value = setInterval(() => {
    countdown.value--
    codeText.value = `${countdown.value}秒后重发`
    if (countdown.value <= 0) {
      if (countdownTimer.value) {
        clearInterval(countdownTimer.value)
        countdownTimer.value = null
      }
      codeDisabled.value = false
      codeText.value = "发送验证码"
      countdown.value = 60
    }
  }, 1000)
}

onBeforeUnmount(() => {
  if (countdownTimer.value) {
    clearInterval(countdownTimer.value)
  }
})

const handleRegister = async () => {
  if (!registerFormRef.value) return

  await registerFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        await useUserStore.register(
          registerForm.UserName,
          registerForm.email,
          registerForm.code,
          registerForm.password
        )
        ElMessage.success("注册成功，即将跳转到登录页面")
        setTimeout(() => {
          router.push("/login")
        }, 1500)
      } catch (error: any) {
        ElMessage.error(error.response?.data?.message || "注册失败，请稍后重试")
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
:deep(.el-input__wrapper) {
  padding: 12px 15px;
}

:deep(.el-button) {
  padding: 12px 20px;
  font-size: 16px;
}
</style>
