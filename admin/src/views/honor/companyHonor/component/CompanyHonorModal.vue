<script setup lang="tsx">
import {
  NButton,
  NDatePicker,
  NForm,
  NFormItem,
  NInput,
  NModal,
  NSpace,
  useMessage,
} from 'naive-ui'
import { reactive, ref, useTemplateRef } from 'vue'

import { CreateCompanyHonorApi, UpdateCompanyHonorApi } from '@/api/honor'
import ImageUpload from '@/components/ImageUpload.vue'

import type {
  CompanyHonor,
  CreateCompanyHonorParams,
  UpdateCompanyHonorParams,
} from '@/types/honor'
import type { FormInst, FormRules, UploadFileInfo } from 'naive-ui'

interface ModelForm {
  id?: number
  title: string
  issue_date: number | null
  cert_no: string
  issuer: string
  image: string
  description: string
}

defineOptions({
  name: 'CompanyHonorModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  title: [{ required: true, message: '请输入荣誉名称', trigger: 'blur' }],
  date: [
    { required: true, type: 'number', message: '请选择获得日期', trigger: ['blur', 'change'] },
  ],
  issuer: [{ required: true, message: '请输入颁发机构', trigger: 'blur' }],
  issue_date: [
    { required: true, type: 'number', message: '请选择获得日期', trigger: ['blur', 'change'] },
  ],
  cert_no: [{ required: true, message: '请输入证书编号', trigger: 'blur' }],
  image: [{ required: true, message: '请上传荣誉图片', trigger: 'change' }],
  description: [{ required: true, message: '请输入荣誉描述', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const modalForm = reactive<ModelForm>({
  title: '',
  issue_date: null,
  cert_no: '',
  issuer: '',
  image: '',
  description: '',
})

const imageFileList = ref<UploadFileInfo[]>([])

function Open(action: 'create' | 'update', row?: CompanyHonor) {
  modalType.value = action
  if (action === 'create') {
    modalForm.id = undefined
    modalForm.title = ''
    modalForm.issue_date = null
    modalForm.cert_no = ''
    modalForm.issuer = ''
    modalForm.image = ''
    modalForm.description = ''
    imageFileList.value = []
  } else if (row) {
    modalForm.id = row.id
    modalForm.title = row.title
    modalForm.issue_date = row.issue_date ? new Date(row.issue_date).getTime() : null
    modalForm.cert_no = row.cert_no
    modalForm.issuer = row.issuer
    modalForm.image = row.image
    modalForm.description = row.description
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
    const issue_date = modalForm.issue_date ? new Date(modalForm.issue_date).toISOString() : null

    try {
      switch (modalType.value) {
        case 'create':
          await CreateCompanyHonor({
            title: modalForm.title,
            issue_date: issue_date ? issue_date.toString() : '',
            cert_no: modalForm.cert_no,
            issuer: modalForm.issuer,
            image: modalForm.image,
            description: modalForm.description,
          })
          break
        case 'update':
          await UpdateCompanyHonor({
            id: modalForm.id!,
            title: modalForm.title,
            issue_date: issue_date ? issue_date.toString() : '',
            cert_no: modalForm.cert_no,
            issuer: modalForm.issuer,
            image: modalForm.image,
            description: modalForm.description,
          })
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

async function CreateCompanyHonor(params: CreateCompanyHonorParams) {
  await CreateCompanyHonorApi(params)
  message.success('新增成功')
}

async function UpdateCompanyHonor(params: UpdateCompanyHonorParams) {
  await UpdateCompanyHonorApi(params)
  message.success('更新成功')
}

defineExpose({
  Open,
})
</script>

<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增荣誉' : '编辑荣誉'"
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
            label="荣誉名称"
            path="title"
          >
            <NInput
              v-model:value="modalForm.title"
              placeholder="请输入荣誉名称"
            />
          </NFormItem>
          <NFormItem
            label="颁发机构"
            path="issuer"
          >
            <NInput
              v-model:value="modalForm.issuer"
              placeholder="请输入颁发机构"
            />
          </NFormItem>
          <NFormItem
            label="获得日期"
            path="issue_date"
          >
            <NDatePicker
              v-model:value="modalForm.issue_date"
              type="date"
              placeholder="请选择获得日期"
              class="w-full"
              clearable
            />
          </NFormItem>
          <NFormItem
            label="证书编号"
            path="cert_no"
          >
            <NInput
              v-model:value="modalForm.cert_no"
              placeholder="请输入证书编号（选填）"
            />
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div class="mb-4">
        <NFormItem
          path="image"
          label="荣誉图片"
        >
          <ImageUpload
            v-model="modalForm.image"
            :max="1"
          />
        </NFormItem>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <NFormItem
        path="description"
        label="荣誉描述"
      >
        <NInput
          v-model:value="modalForm.description"
          type="textarea"
          placeholder="请输入荣誉描述信息..."
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
