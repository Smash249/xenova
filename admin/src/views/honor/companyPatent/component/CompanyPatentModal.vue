<script setup lang="tsx">
import {
  NButton,
  NDatePicker,
  NForm,
  NFormItem,
  NInput,
  NModal,
  NSelect,
  NSpace,
  useMessage,
} from 'naive-ui'
import { reactive, ref, useTemplateRef } from 'vue'

import { CreateCompanyPatentApi, UpdateCompanyPatentApi } from '@/api/honor'
import ImageUpload from '@/components/ImageUpload.vue'

import type {
  CompanyPatent,
  CreateCompanyPatentParams,
  UpdateCompanyPatentParams,
} from '@/types/honor'
import type { FormInst, FormRules, UploadFileInfo } from 'naive-ui'

interface ModelForm {
  id?: number
  title: string
  type: string
  patent_no: string
  auth_date: number | null
  inventor: string
  image: string
  summary: string
}

defineOptions({
  name: 'CompanyPatentModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  title: [{ required: true, message: '请输入专利名称', trigger: 'blur' }],
  type: [{ required: true, message: '请选择专利类型', trigger: ['blur', 'change'] }],
  patent_no: [{ required: true, message: '请输入专利号', trigger: 'blur' }],
  auth_date: [
    { required: true, type: 'number', message: '请选择授权日期', trigger: ['blur', 'change'] },
  ],
  inventor: [{ required: true, message: '请输入发明人', trigger: 'blur' }],
  summary: [{ required: true, message: '请输入专利摘要', trigger: 'blur' }],
  image: [{ required: true, message: '请上传专利图片', trigger: 'change' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const patentTypeOptions = [
  { label: '发明专利', value: '发明专利' },
  { label: '实用新型专利', value: '实用新型专利' },
  { label: '外观设计专利', value: '外观设计专利' },
]

const modalForm = reactive<ModelForm>({
  title: '',
  type: '发明专利',
  patent_no: '',
  auth_date: null,
  inventor: '',
  image: '',
  summary: '',
})

const imageFileList = ref<UploadFileInfo[]>([])

function Open(action: 'create' | 'update', row?: CompanyPatent) {
  modalType.value = action
  if (action === 'create') {
    modalForm.id = undefined
    modalForm.title = ''
    modalForm.type = '发明专利'
    modalForm.patent_no = ''
    modalForm.auth_date = null
    modalForm.inventor = ''
    modalForm.image = ''
    modalForm.summary = ''
    imageFileList.value = []
  } else if (row) {
    modalForm.id = row.id
    modalForm.title = row.title
    modalForm.type = row.type
    modalForm.patent_no = row.patent_no
    modalForm.auth_date = row.auth_date ? new Date(row.auth_date).getTime() : null
    modalForm.inventor = row.inventor
    modalForm.image = row.image
    modalForm.summary = row.summary
    imageFileList.value = row.image
      ? [{ id: 'image', name: 'image', status: 'finished', url: row.image }]
      : []
  }
  showModal.value = true
}

function HandleModalSubmit() {
  modalFormRef.value?.validate(async (errors) => {
    if (errors) return
    isLoading.value = true
    try {
      const auth_date = modalForm.auth_date ? new Date(modalForm.auth_date).toISOString() : ''
      switch (modalType.value) {
        case 'create':
          await CreateCompanyPatent({
            title: modalForm.title,
            type: modalForm.type,
            patent_no: modalForm.patent_no,
            auth_date,
            inventor: modalForm.inventor,
            image: modalForm.image,
            summary: modalForm.summary,
          })
          break
        case 'update':
          await UpdateCompanyPatent({
            id: modalForm.id!,
            title: modalForm.title,
            type: modalForm.type,
            patent_no: modalForm.patent_no,
            auth_date,
            inventor: modalForm.inventor,
            image: modalForm.image,
            summary: modalForm.summary,
          })
          break
      }
      showModal.value = false
      emit('success')
    } catch (error) {
      message.error(modalType.value === 'create' ? '新增失败' : '更新失败')
      console.error(error)
    } finally {
      isLoading.value = false
    }
  })
}

async function CreateCompanyPatent(params: CreateCompanyPatentParams) {
  await CreateCompanyPatentApi(params)
  message.success('专利创建成功')
}

async function UpdateCompanyPatent(params: UpdateCompanyPatentParams) {
  await UpdateCompanyPatentApi(params)
  message.success('专利更新成功')
}

defineExpose({
  Open,
})
</script>

<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增专利' : '编辑专利'"
    class="max-w-3xl"
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
            label="专利名称"
            path="title"
          >
            <NInput
              v-model:value="modalForm.title"
              placeholder="请输入专利名称"
            />
          </NFormItem>
          <NFormItem
            label="专利类型"
            path="type"
          >
            <NSelect
              v-model:value="modalForm.type"
              :options="patentTypeOptions"
              placeholder="请选择专利类型"
              clearable
            />
          </NFormItem>
          <NFormItem
            label="专利号"
            path="patent_no"
          >
            <NInput
              v-model:value="modalForm.patent_no"
              placeholder="请输入专利号"
            />
          </NFormItem>
          <NFormItem
            label="授权日期"
            path="auth_date"
          >
            <NDatePicker
              v-model:value="modalForm.auth_date"
              type="date"
              placeholder="请选择授权日期"
              class="w-full"
              clearable
            />
          </NFormItem>
          <NFormItem
            label="发明人"
            path="inventor"
            class="col-span-2 max-md:col-span-1"
          >
            <NInput
              v-model:value="modalForm.inventor"
              placeholder="请输入发明人，多人用逗号分隔"
            />
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div class="mb-4">
        <NFormItem
          path="image"
          label="专利图片"
        >
          <ImageUpload
            v-model="modalForm.image"
            :max="1"
          />
        </NFormItem>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <NFormItem
        path="summary"
        label="专利摘要"
      >
        <NInput
          v-model:value="modalForm.summary"
          type="textarea"
          placeholder="请输入专利摘要信息..."
          :rows="4"
        />
      </NFormItem>
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
