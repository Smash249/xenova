<script setup lang="tsx">
import { MdEditor } from 'md-editor-v3'
import {
  NButton,
  NDatePicker,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NModal,
  NSpace,
  useMessage,
} from 'naive-ui'
import { ref, useTemplateRef } from 'vue'
import 'md-editor-v3/lib/style.css'

import { CreateLoveActivityApi, UpdateLoveActivityApi } from '@/api/honor'
import ImageUpload from '@/components/ImageUpload.vue'
import { useUserStore } from '@/stores'

import type {
  CreateLoveActivityParams,
  LoveActivity,
  UpdateLoveActivityParams,
} from '@/types/honor'
import type { FormInst, FormRules } from 'naive-ui'

interface ModelForm {
  id?: number
  title: string
  activity_date: number | null
  location: string
  participants: number
  cover: string
  summary: string
  content: string
}

defineOptions({
  name: 'LoveActivityModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const DEFAULT_MODAL_FORM: ModelForm = {
  title: '',
  activity_date: null,
  location: '',
  participants: 0,
  cover: '',
  summary: '',
  content: '',
}

const uploadAction = import.meta.env.VITE_PROXY_PATH + '/private/upload'
const baseUrl = import.meta.env.VITE_SITE_BASE_API || ''
const usser = useUserStore()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  title: [{ required: true, message: '请输入活动标题', trigger: 'blur' }],
  activity_date: [
    { required: true, type: 'number', message: '请选择活动日期', trigger: ['blur', 'change'] },
  ],
  location: [{ required: true, message: '请输入活动地点', trigger: 'blur' }],
  participants: [{ required: true, type: 'number', message: '请输入参与人数', trigger: 'blur' }],
  content: [{ required: true, message: '请输入活动内容', trigger: 'blur' }],
  summary: [{ required: true, max: 200, message: '活动摘要不能超过200字', trigger: 'blur' }],
  cover: [{ required: true, message: '请上传活动封面', trigger: 'change' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const modalForm = ref<ModelForm>({ ...DEFAULT_MODAL_FORM })

function WithPrefix(url: string): string {
  if (!url) return ''
  if (url.startsWith('http://') || url.startsWith('https://') || url.startsWith('blob:')) {
    return url
  }
  return baseUrl + url
}

async function HandleUploadImg(files: File[], callback: (urls: string[]) => void) {
  const urls: string[] = []

  for (const file of files) {
    try {
      const formData = new FormData()
      formData.append('file', file)

      const response = await fetch(uploadAction, {
        method: 'POST',
        headers: {
          Authorization: `${usser.token}`,
        },
        body: formData,
      })

      if (!response.ok) {
        message.error(`「${file.name}」上传失败`)
        continue
      }

      const result = await response.json()
      urls.push(WithPrefix(result.data.url))
    } catch (error) {
      message.error(`「${file.name}」上传失败`)
      console.error(error)
    }
  }

  callback(urls)
}

function Open(action: 'create' | 'update', row?: LoveActivity) {
  modalType.value = action
  if (action === 'create') {
    modalForm.value = { ...DEFAULT_MODAL_FORM }
  } else if (row) {
    modalForm.value = {
      id: row.id,
      title: row.title,
      activity_date: row.activity_date ? new Date(row.activity_date).getTime() : null,
      location: row.location,
      participants: row.participants,
      cover: row.cover,
      summary: row.summary,
      content: row.content,
    }
  }
  showModal.value = true
}

function HandleModalSubmit() {
  modalFormRef.value?.validate(async (errors) => {
    if (errors) return
    isLoading.value = true
    const form = modalForm.value
    try {
      const activity_date = form.activity_date ? new Date(form.activity_date).toISOString() : ''
      switch (modalType.value) {
        case 'create':
          await CreateLoveActivity({
            title: form.title,
            activity_date,
            location: form.location,
            participants: form.participants,
            cover: form.cover,
            summary: form.summary,
            content: form.content,
          })
          break
        case 'update':
          await UpdateLoveActivity({
            id: form.id!,
            title: form.title,
            activity_date,
            location: form.location,
            participants: form.participants,
            cover: form.cover,
            summary: form.summary,
            content: form.content,
          })
          break
      }
    } catch (error) {
      message.error(modalType.value === 'create' ? '新增失败' : '更新失败')
      console.error(error)
    } finally {
      isLoading.value = false
    }
  })
}

async function CreateLoveActivity(params: CreateLoveActivityParams) {
  await CreateLoveActivityApi(params)
  message.success('新增成功')
  showModal.value = false
  emit('success')
}

async function UpdateLoveActivity(params: UpdateLoveActivityParams) {
  await UpdateLoveActivityApi(params)
  message.success('更新成功')
  showModal.value = false
  emit('success')
}

defineExpose({
  Open,
})
</script>

<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增爱心活动' : '编辑爱心活动'"
    class="max-w-5xl"
    :segmented="{ content: true, footer: true }"
  >
    <NForm
      ref="modalFormRef"
      :model="modalForm"
      :rules="rules"
      label-placement="top"
    >
      <div class="mb-4">
        <div class="mb-3 flex items-center gap-2 text-base font-medium">基本信息</div>
        <div class="grid grid-cols-2 gap-x-6 max-md:grid-cols-1">
          <NFormItem
            label="活动标题"
            path="title"
          >
            <NInput
              v-model:value="modalForm.title"
              placeholder="请输入活动标题"
            />
          </NFormItem>
          <NFormItem
            label="活动地点"
            path="location"
          >
            <NInput
              v-model:value="modalForm.location"
              placeholder="请输入活动地点"
            />
          </NFormItem>
          <NFormItem
            label="活动日期"
            path="activity_date"
          >
            <NDatePicker
              v-model:value="modalForm.activity_date"
              type="date"
              placeholder="请选择活动日期"
              class="w-full"
              clearable
            />
          </NFormItem>
          <NFormItem
            label="参与人数"
            path="participants"
          >
            <NInputNumber
              v-model:value="modalForm.participants"
              :min="0"
              placeholder="请输入参与人数"
              class="w-full"
            >
              <template #suffix>人</template>
            </NInputNumber>
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div class="mb-4">
        <div class="mb-3 flex items-center gap-2 text-base font-medium">封面与摘要</div>
        <div class="grid grid-cols-2 gap-x-6 max-md:grid-cols-1">
          <NFormItem
            label="活动封面"
            path="cover"
          >
            <ImageUpload
              v-model="modalForm.cover"
              :max="1"
            />
          </NFormItem>
          <NFormItem
            label="活动摘要"
            path="summary"
          >
            <NInput
              v-model:value="modalForm.summary"
              type="textarea"
              placeholder="请输入活动摘要"
              :rows="4"
            />
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div>
        <div class="mb-3 flex items-center gap-2 text-base font-medium">活动内容</div>
        <NFormItem
          path="content"
          :show-label="false"
        >
          <MdEditor
            v-model="modalForm.content"
            :style="{ height: '400px' }"
            placeholder="请输入活动内容，支持 Markdown 格式..."
            :on-upload-img="HandleUploadImg"
          />
        </NFormItem>
      </div>
    </NForm>

    <template #footer>
      <div class="flex justify-end">
        <NSpace>
          <NButton @click="showModal = false">取消</NButton>
          <NButton
            type="primary"
            @click="HandleModalSubmit"
            :loading="isLoading"
          >
            {{ modalType === 'create' ? '立即创建' : '保存修改' }}
          </NButton>
        </NSpace>
      </div>
    </template>
  </NModal>
</template>

<style scoped lang="scss">
:deep(.md-editor) {
  width: 100%;
}
</style>
