<template>
  <el-drawer v-model="visible" direction="rtl" size="300px">
    <div class="flex h-full flex-col">
      <div class="mb-6 flex items-center gap-4 border-b border-gray-100 pb-6">
        <el-avatar
          :size="50"
          class="bg-blue-100 text-xl font-bold text-blue-600"
        >
          {{ avatarLetter }}
        </el-avatar>
        <div class="flex flex-col">
          <span class="text-lg font-bold text-gray-800">{{
            user.userName || "用户"
          }}</span>
          <span class="text-sm text-gray-500">{{ user.email }}</span>
        </div>
      </div>

      <div class="flex-1 space-y-2">
        <div
          v-for="(item, index) in menuItems"
          :key="index"
          class="flex cursor-pointer items-center rounded-lg py-2 text-gray-600 transition-colors hover:bg-gray-50 hover:text-blue-600"
          @click="HandleMenuClick(item)"
        >
          <i :class="item.icon" class="mr-3 text-lg"></i>
          <span class="font-medium">{{ item.title }}</span>
        </div>
      </div>

      <div class="border-t border-gray-100 pt-4">
        <button
          class="flex w-full items-center justify-center rounded-lg bg-red-50 px-4 py-3 text-red-600 transition-colors hover:bg-red-100"
          @click="HandleLogout"
        >
          <span class="font-bold">退出登录</span>
        </button>
      </div>
    </div>
  </el-drawer>

  <el-dialog
    v-model="dialogVisible"
    width="680px"
    :show-close="true"
    :close-on-click-modal="false"
    class="user-dialog"
  >
    <div class="flex h-105">
      <div class="w-40 border-r border-gray-100 pr-4">
        <div
          v-for="(tab, index) in dialogTabs"
          :key="index"
          class="mb-1 flex cursor-pointer items-center rounded-lg px-3 py-2.5 text-sm transition-colors"
          :class="
            activeTab === tab.key
              ? 'bg-blue-50 font-bold text-blue-600'
              : 'text-gray-600 hover:bg-gray-50 hover:text-blue-600'
          "
          @click="activeTab = tab.key"
        >
          <i :class="tab.icon" class="mr-2"></i>
          <span>{{ tab.title }}</span>
        </div>
      </div>

      <div class="flex-1 pl-6">
        <div v-if="activeTab === 'profile'">
          <h3 class="mb-6 text-lg font-bold text-gray-800">用户信息</h3>
          <el-form
            ref="profileFormRef"
            :model="profileForm"
            :rules="profileRules"
            label-width="80px"
            label-position="left"
          >
            <el-form-item label="用户名" prop="userName">
              <el-input
                v-model="profileForm.userName"
                placeholder="请输入用户名"
              />
            </el-form-item>
            <el-form-item label="手机号" prop="phone">
              <el-input
                v-model="profileForm.phone"
                placeholder="请输入手机号"
              />
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                :loading="profileLoading"
                @click="HandleSaveProfile"
              >
                保存修改
              </el-button>
            </el-form-item>
          </el-form>
        </div>

        <div v-if="activeTab === 'security'">
          <h3 class="mb-6 text-lg font-bold text-gray-800">账户安全</h3>
          <el-form
            ref="passwordFormRef"
            :model="passwordForm"
            :rules="passwordRules"
            label-width="80px"
            label-position="left"
          >
            <el-form-item label="旧密码" prop="oldPassword">
              <el-input
                v-model="passwordForm.oldPassword"
                type="password"
                show-password
                placeholder="请输入旧密码"
              />
            </el-form-item>
            <el-form-item label="新密码" prop="newPassword">
              <el-input
                v-model="passwordForm.newPassword"
                type="password"
                show-password
                placeholder="请输入新密码"
              />
            </el-form-item>
            <el-form-item label="确认密码" prop="confirmPassword">
              <el-input
                v-model="passwordForm.confirmPassword"
                type="password"
                show-password
                placeholder="请再次输入新密码"
              />
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                :loading="passwordLoading"
                @click="HandleChangePassword"
              >
                修改密码
              </el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>
    </div>
  </el-dialog>
</template>

<script setup>
import {
  ChangePasswordApi,
  GetUserInfoApi,
  UpdateUserInfoApi,
} from "@/api/user"
import userStore from "@/store/modules/user"
import { ElMessage } from "element-plus"
import { computed, nextTick, reactive, ref } from "vue"

const menuItems = [
  { title: "个人中心", icon: "el-icon-user", action: "profile" },
  { title: "更改密码", icon: "el-icon-setting", action: "settings" },
]

const dialogTabs = [
  { title: "用户信息", icon: "el-icon-user", key: "profile" },
  { title: "账户安全", icon: "el-icon-lock", key: "security" },
]

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(["update:modelValue"])
const user = userStore()

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
})

const avatarLetter = computed(() => {
  return user.email ? user.email.charAt(0).toUpperCase() : ""
})

const dialogVisible = ref(false)
const activeTab = ref("profile")
const profileLoading = ref(false)
const passwordLoading = ref(false)
const profileFormRef = ref(null)
const passwordFormRef = ref(null)

const profileForm = reactive({
  userName: "",
  phone: "",
})

const passwordForm = reactive({
  oldPassword: "",
  newPassword: "",
  confirmPassword: "",
})

const profileRules = {
  userName: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  email: [
    { required: true, message: "请输入邮箱", trigger: "blur" },
    { type: "email", message: "请输入正确的邮箱格式", trigger: "blur" },
  ],
  phone: [{ required: true, message: "请输入手机号", trigger: "blur" }],
}

const ValidateConfirmPassword = (rule, value, callback) => {
  if (value !== passwordForm.newPassword) {
    callback(new Error("两次输入的密码不一致"))
  } else {
    callback()
  }
}

const passwordRules = {
  oldPassword: [{ required: true, message: "请输入旧密码", trigger: "blur" }],
  newPassword: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 6, message: "密码长度不能少于6位", trigger: "blur" },
  ],
  confirmPassword: [
    { required: true, message: "请再次输入新密码", trigger: "blur" },
    { validator: ValidateConfirmPassword, trigger: "blur" },
  ],
}

async function FetchUserInfo() {
  try {
    const data = await GetUserInfoApi()
    profileForm.userName = data.userName || ""
    profileForm.email = data.email || ""
    profileForm.phone = data.phone || ""
  } catch (error) {
    ElMessage.error("获取用户信息失败")
  }
}

function HandleMenuClick(item) {
  visible.value = false
  if (item.action === "profile") {
    activeTab.value = "profile"
  } else if (item.action === "settings") {
    activeTab.value = "security"
  }
  dialogVisible.value = true
  nextTick(() => {
    FetchUserInfo()
    if (passwordFormRef.value) {
      passwordFormRef.value.resetFields()
    }
  })
}

async function HandleSaveProfile() {
  try {
    await profileFormRef.value.validate()
  } catch {
    return
  }
  profileLoading.value = true
  try {
    await UpdateUserInfoApi({
      userName: profileForm.userName,
      phone: profileForm.phone,
    })
    ElMessage.success("用户信息更新成功")
    user.userName = profileForm.userName
    user.email = profileForm.email
  } catch (error) {
    ElMessage.error(error.message || "更新用户信息失败")
  } finally {
    profileLoading.value = false
  }
}

async function HandleChangePassword() {
  try {
    await passwordFormRef.value.validate()
  } catch {
    return
  }
  passwordLoading.value = true
  try {
    await ChangePasswordApi({
      oldPassword: passwordForm.oldPassword,
      newPassword: passwordForm.newPassword,
    })
    ElMessage.success("密码修改成功，请重新登录")
    dialogVisible.value = false
    user.Logout()
  } catch (error) {
    ElMessage.error(error.message || "修改密码失败")
  } finally {
    passwordLoading.value = false
  }
}

function HandleLogout() {
  visible.value = false
  user.Logout()
}
</script>

<style scoped>
:deep(.el-drawer__body) {
  padding: 0;
}

:deep(.user-dialog .el-dialog__header) {
  display: none;
}

:deep(.user-dialog .el-dialog__body) {
  padding: 24px;
}
</style>
