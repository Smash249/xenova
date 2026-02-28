<script setup lang="ts">
import { NButton, NForm, NFormItem, NInput, NModal, useMessage } from 'naive-ui'
import { reactive, ref, watch } from 'vue'

import { UpdateUserInfoApi, ChangePasswordApi } from '@/api/user'
import { useUserStore } from '@/stores'

const props = defineProps<{ show: boolean }>()
const emit = defineEmits<{ (e: 'update:show', val: boolean): void }>()

const message = useMessage()
const userStore = useUserStore()

const activeTab = ref('basic')

const isEditingName = ref(false)
const editName = ref('')
const isSavingName = ref(false)

const pwdForm = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: '',
})
const isSavingPwd = ref(false)

watch(
  () => props.show,
  (val) => {
    if (val) {
      editName.value = userStore.user?.userName || ''
      activeTab.value = 'basic'
      isEditingName.value = false
      pwdForm.oldPassword = ''
      pwdForm.newPassword = ''
      pwdForm.confirmPassword = ''
    }
  },
)

async function HandleSaveName() {
  if (!editName.value.trim()) {
    message.warning('用户名不能为空')
    return
  }
  isSavingName.value = true
  try {
    const result = await UpdateUserInfoApi(editName.value)
    userStore.UpdateUserInfo(result)
    message.success('修改用户名成功')
    isEditingName.value = false
  } catch (error) {
    message.error('修改用户名失败')
    console.error(error)
  } finally {
    isSavingName.value = false
  }
}

async function HandleSavePassword() {
  if (!pwdForm.oldPassword || !pwdForm.newPassword) {
    message.warning('请填写完整密码信息')
    return
  }
  if (pwdForm.newPassword !== pwdForm.confirmPassword) {
    message.warning('两次输入的新密码不一致')
    return
  }
  isSavingPwd.value = true
  try {
    await ChangePasswordApi({ oldPassword: pwdForm.oldPassword, newPassword: pwdForm.newPassword })
    message.success('密码修改成功')
    userStore.Logout()
  } catch (error) {
    message.error('密码修改失败')
    console.error(error)
  } finally {
    isSavingPwd.value = false
  }
}
</script>

<template>
  <NModal
    :show="show"
    @update:show="(val) => emit('update:show', val)"
    preset="card"
    title="个人中心"
    :bordered="false"
    style="width: 820px"
  >
    <div class="flex h-87.5">
      <div class="w-1/3 border-r border-neutral-100 pr-4 dark:border-neutral-800">
        <div
          class="mb-2 cursor-pointer rounded-lg px-4 py-3 transition-colors"
          :class="
            activeTab === 'basic'
              ? 'bg-blue-50 font-medium text-blue-600 dark:bg-blue-900/30 dark:text-blue-400'
              : 'text-neutral-600 hover:bg-neutral-50 dark:text-neutral-400 dark:hover:bg-neutral-800'
          "
          @click="activeTab = 'basic'"
        >
          基础信息
        </div>
        <div
          class="cursor-pointer rounded-lg px-4 py-3 transition-colors"
          :class="
            activeTab === 'security'
              ? 'bg-blue-50 font-medium text-blue-600 dark:bg-blue-900/30 dark:text-blue-400'
              : 'text-neutral-600 hover:bg-neutral-50 dark:text-neutral-400 dark:hover:bg-neutral-800'
          "
          @click="activeTab = 'security'"
        >
          账号安全
        </div>
      </div>

      <div class="w-2/3 pl-6">
        <div
          v-if="activeTab === 'basic'"
          class="space-y-8 pt-2"
        >
          <div>
            <span class="mb-2 block text-sm text-neutral-500">用户名</span>
            <div
              v-if="!isEditingName"
              class="flex items-center gap-3"
            >
              <span class="text-base font-medium dark:text-neutral-200">
                {{ userStore.user?.userName || '未登录用户' }}
              </span>
              <NButton
                size="tiny"
                quaternary
                type="info"
                @click="isEditingName = true"
              >
                修改
              </NButton>
            </div>
            <div
              v-else
              class="flex items-center gap-2"
            >
              <NInput
                v-model:value="editName"
                size="small"
                class="w-48"
                placeholder="请输入新用户名"
                @keyup.enter="HandleSaveName"
              />
              <NButton
                size="small"
                type="primary"
                :loading="isSavingName"
                @click="HandleSaveName"
              >
                保存
              </NButton>
              <NButton
                size="small"
                @click="isEditingName = false"
                >取消</NButton
              >
            </div>
          </div>

          <div>
            <span class="mb-1 block text-sm text-neutral-500">邮箱</span>
            <span class="text-base dark:text-neutral-200">
              {{ userStore.user?.email || '未绑定邮箱' }}
            </span>
          </div>
        </div>

        <div
          v-if="activeTab === 'security'"
          class="pt-2"
        >
          <h3 class="mb-6 text-base font-medium dark:text-neutral-200">修改密码</h3>
          <NForm
            ref="pwdFormRef"
            :model="pwdForm"
            label-placement="top"
          >
            <NFormItem
              label="当前密码"
              path="oldPassword"
            >
              <NInput
                v-model:value="pwdForm.oldPassword"
                type="password"
                show-password-on="click"
                placeholder="请输入当前密码"
              />
            </NFormItem>
            <NFormItem
              label="新密码"
              path="newPassword"
            >
              <NInput
                v-model:value="pwdForm.newPassword"
                type="password"
                show-password-on="click"
                placeholder="请输入新密码"
              />
            </NFormItem>
            <NFormItem
              label="确认新密码"
              path="confirmPassword"
            >
              <NInput
                v-model:value="pwdForm.confirmPassword"
                type="password"
                show-password-on="click"
                placeholder="请再次输入新密码"
              />
            </NFormItem>
            <div class="mt-2">
              <NButton
                type="primary"
                block
                :loading="isSavingPwd"
                @click="HandleSavePassword"
              >
                确认修改
              </NButton>
            </div>
          </NForm>
        </div>
      </div>
    </div>
  </NModal>
</template>
