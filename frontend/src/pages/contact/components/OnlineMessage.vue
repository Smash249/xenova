<template>
  <div class="flex flex-col gap-10">
    <div class="flex flex-col gap-2">
      <h3 class="flex items-center gap-3 text-2xl font-bold text-slate-900">
        <div class="h-8 w-1.5 rounded-full bg-blue-600"></div>
        在线留言
      </h3>
      <p class="pl-4 text-sm text-slate-500">
        请填写以下表单，我们的专业技术顾问将在 1
        个工作日内与您联系，为您提供专属的自动化解决方案。
      </p>
    </div>

    <div class="grid grid-cols-1 gap-8 lg:grid-cols-5 lg:gap-12">
      <div class="flex flex-col gap-6 lg:col-span-2">
        <div
          class="overflow-hidden rounded-2xl bg-linear-to-br from-blue-600 to-indigo-700 p-8 text-white shadow-lg shadow-blue-500/20"
        >
          <h4 class="mb-6 text-xl font-bold tracking-wider">
            高效响应，专业服务
          </h4>

          <div class="space-y-6">
            <div class="flex items-start gap-4">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-white/10 backdrop-blur-xs"
              >
                <ZapIcon class="h-5 w-5 text-blue-200" />
              </div>
              <div>
                <h5 class="mb-1 font-bold text-blue-50">极速响应</h5>
                <p class="text-xs leading-relaxed text-blue-200">
                  收到留言后，我们会在24小时内安排专属工程师与您初步对接需求。
                </p>
              </div>
            </div>

            <div class="flex items-start gap-4">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-white/10 backdrop-blur-xs"
              >
                <ShieldCheckIcon class="h-5 w-5 text-blue-200" />
              </div>
              <div>
                <h5 class="mb-1 font-bold text-blue-50">隐私保护</h5>
                <p class="text-xs leading-relaxed text-blue-200">
                  您提交的所有商业信息和联系方式都将受到严格保密，绝不外泄。
                </p>
              </div>
            </div>

            <div class="flex items-start gap-4">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-white/10 backdrop-blur-xs"
              >
                <HeadphonesIcon class="h-5 w-5 text-blue-200" />
              </div>
              <div>
                <h5 class="mb-1 font-bold text-blue-50">全程跟踪</h5>
                <p class="text-xs leading-relaxed text-blue-200">
                  从需求沟通、方案定制到设备交付，提供一站式全周期技术支持。
                </p>
              </div>
            </div>
          </div>
        </div>

        <div
          class="rounded-2xl border border-slate-100 bg-slate-50 p-6 text-sm text-slate-500"
        >
          急需帮助？您可以直接拨打我们的技术热线：
          <span class="mt-1 block font-mono text-base font-bold text-blue-600"
            >0592-6818878</span
          >
        </div>
      </div>

      <div
        class="rounded-2xl border border-slate-100 bg-white p-1 shadow-sm lg:col-span-3"
      >
        <el-form
          ref="formRef"
          :model="formData"
          :rules="formRules"
          class="flex flex-col gap-6 p-4 sm:p-6"
          @submit.prevent="HandleSubmit"
        >
          <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
            <div class="space-y-2">
              <label class="block text-sm font-bold text-slate-700">
                您的姓名 <span class="text-red-500">*</span>
              </label>
              <el-form-item prop="name" class="mb-0">
                <el-input
                  v-model="formData.name"
                  placeholder="请输入您的称呼"
                  class="custom-el-input"
                  size="large"
                  :disabled="isSubmitting"
                >
                  <template #prefix>
                    <UserIcon class="h-5 w-5 transition-colors" />
                  </template>
                </el-input>
              </el-form-item>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-bold text-slate-700">
                联系电话 <span class="text-red-500">*</span>
              </label>
              <el-form-item prop="phone" class="mb-0">
                <el-input
                  v-model="formData.phone"
                  placeholder="请输入手机号或座机"
                  class="custom-el-input"
                  size="large"
                  :disabled="isSubmitting"
                >
                  <template #prefix>
                    <PhoneIcon class="h-5 w-5 transition-colors" />
                  </template>
                </el-input>
              </el-form-item>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-bold text-slate-700"
                >公司名称</label
              >
              <el-form-item prop="company" class="mb-0">
                <el-input
                  v-model="formData.company"
                  placeholder="请输入您的企业名称"
                  class="custom-el-input"
                  size="large"
                  :disabled="isSubmitting"
                >
                  <template #prefix>
                    <Building2Icon class="h-5 w-5 transition-colors" />
                  </template>
                </el-input>
              </el-form-item>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-bold text-slate-700"
                >电子邮箱</label
              >
              <el-form-item prop="email" class="mb-0">
                <el-input
                  v-model="formData.email"
                  placeholder="example@company.com"
                  class="custom-el-input"
                  size="large"
                  :disabled="isSubmitting"
                >
                  <template #prefix>
                    <MailIcon class="h-5 w-5 transition-colors" />
                  </template>
                </el-input>
              </el-form-item>
            </div>
          </div>

          <div class="space-y-2">
            <label class="block text-sm font-bold text-slate-700">
              留言内容 <span class="text-red-500">*</span>
            </label>
            <el-form-item prop="content" class="relative mb-0">
              <div
                class="pointer-events-none absolute top-3.5 left-4 z-10 flex"
              >
                <MessageSquareIcon class="h-5 w-5 text-slate-400" />
              </div>
              <el-input
                type="textarea"
                v-model="formData.content"
                :rows="5"
                placeholder="请简要描述您的自动化改造需求、设备定制意向或咨询的问题..."
                class="custom-el-textarea"
                :disabled="isSubmitting"
              />
            </el-form-item>
          </div>

          <div class="mt-2 flex items-center justify-between">
            <span class="text-xs text-slate-400"
              ><span class="text-red-500">*</span> 为必填项</span
            >
            <button
              type="submit"
              :disabled="isSubmitting"
              class="group flex w-full items-center justify-center gap-2 rounded-xl bg-linear-to-r from-blue-600 to-indigo-600 px-8 py-3.5 font-bold tracking-widest text-white shadow-lg shadow-blue-500/30 transition-all hover:from-blue-700 hover:to-indigo-700 hover:shadow-blue-500/50 active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-70 sm:w-auto"
            >
              <span>{{ isSubmitting ? "提交中..." : "立即提交" }}</span>
              <Loader2Icon v-if="isSubmitting" class="h-4 w-4 animate-spin" />
              <SendIcon
                v-else
                class="h-4 w-4 transition-transform duration-300 group-hover:translate-x-0.5 group-hover:-translate-y-0.5"
              />
            </button>
          </div>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { CreateSystemMessage } from "@/api/system"
import { ElMessage, type FormInstance, type FormRules } from "element-plus"
import {
  Building2Icon,
  HeadphonesIcon,
  Loader2Icon,
  MailIcon,
  MessageSquareIcon,
  PhoneIcon,
  SendIcon,
  ShieldCheckIcon,
  UserIcon,
  ZapIcon,
} from "lucide-vue-next"
import { reactive, ref } from "vue"

interface CreateSystemMessageReq {
  name: string
  phone: string
  company: string
  email: string
  content: string
}

const formRef = ref<FormInstance>()
const isSubmitting = ref(false)

const formData = reactive<CreateSystemMessageReq>({
  name: "",
  phone: "",
  company: "",
  email: "",
  content: "",
})

const formRules = reactive<FormRules<CreateSystemMessageReq>>({
  name: [
    { required: true, message: "请输入您的称呼", trigger: "blur" },
    { min: 2, max: 50, message: "长度在 2 到 50 个字符", trigger: "blur" },
  ],
  phone: [
    { required: true, message: "请输入联系电话", trigger: "blur" },
    {
      pattern: /^(?:(?:\+|00)86)?1[3-9]\d{9}$|^0\d{2,3}-?\d{7,8}$/,
      message: "请输入正确的手机号或座机号码",
      trigger: "blur",
    },
  ],
  email: [
    {
      type: "email",
      message: "请输入正确的邮箱地址",
      trigger: ["blur", "change"],
    },
  ],
  content: [
    { required: true, message: "请输入留言内容", trigger: "blur" },
    { min: 5, message: "留言内容至少 5 个字符", trigger: "blur" },
  ],
})

async function HandleSubmit() {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      isSubmitting.value = true
      try {
        await CreateSystemMessage({ ...formData })
        ElMessage.success({
          message: "留言提交成功！我们的技术顾问将尽快与您联系。",
          duration: 3000,
        })
        formRef.value?.resetFields()
      } catch (error) {
        console.error("提交留言失败:", error)
        ElMessage.error("提交失败，请稍后重试或直接拨打我们的联系电话。")
      } finally {
        isSubmitting.value = false
      }
    }
  })
}
</script>

<style scoped>
:deep(.custom-el-input .el-input__wrapper) {
  background-color: #f8fafc;
  border-radius: 0.75rem;
  box-shadow: 0 0 0 1px #e2e8f0 inset;
  padding-left: 0.75rem;
  padding-right: 1rem;
  transition: all 0.3s ease;
}

:deep(.custom-el-textarea .el-textarea__inner) {
  background-color: #f8fafc;
  border-radius: 0.75rem;
  box-shadow: 0 0 0 1px #e2e8f0 inset;
  padding-left: 2.75rem;
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
  padding-right: 1rem;
  transition: all 0.3s ease;
  font-family: inherit;
  resize: none;
}

:deep(.custom-el-input.is-focus .el-input__wrapper),
:deep(.custom-el-textarea .el-textarea__inner:focus) {
  background-color: #ffffff;
  box-shadow:
    0 0 0 1px #3b82f6 inset,
    0 0 0 4px rgba(59, 130, 246, 0.1) !important;
}

:deep(.custom-el-input .el-input__prefix) {
  color: #94a3b8;
  margin-right: 8px;
}
:deep(.custom-el-input.is-focus .el-input__prefix .lucide) {
  color: #3b82f6;
}

:deep(.el-form-item.is-error .el-input__wrapper),
:deep(.el-form-item.is-error .el-textarea__inner) {
  background-color: #fef2f2 !important;
  box-shadow: 0 0 0 1px #ef4444 inset !important;
}
:deep(.el-form-item__error) {
  padding-top: 4px;
}
</style>
