<script setup lang="tsx">
import {
  NButton,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NModal,
  NSelect,
  NSpace,
  NUpload,
  useMessage,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed } from 'vue'

import { GetProductSeriesListApi } from '@/api/product'

import type { Product, ProductSeries } from '@/types/product'
import type { FormInst, FormRules, UploadFileInfo } from 'naive-ui'

interface ModelForm {
  id?: number
  name: string
  seriesId: number | null
  image: string
  description: string
  cover: string
  previews: string[]
  price: number
}

defineOptions({
  name: 'ProductModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入产品名称', trigger: 'blur' }],
  seriesId: [
    { required: true, type: 'number', message: '请选择所属系列', trigger: ['blur', 'change'] },
  ],
  price: [{ required: true, type: 'number', message: '请输入产品价格', trigger: 'blur' }],
  cover: [{ required: true, message: '请上传封面图', trigger: 'change' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const seriesList = ref<ProductSeries[]>([])
const seriesLoading = ref(false)

const seriesOptions = computed(() =>
  seriesList.value.map((item) => ({
    label: item.name,
    value: item.id,
  })),
)

const modalForm = reactive<ModelForm>({
  name: '',
  seriesId: null,
  image: '',
  description: '',
  cover: '',
  previews: [],
  price: 0,
})

const coverFileList = ref<UploadFileInfo[]>([])
const previewFileList = ref<UploadFileInfo[]>([])

async function FetchSeriesList() {
  seriesLoading.value = true
  try {
    const result = await GetProductSeriesListApi(1, 9999, '')
    seriesList.value = result.data
  } catch (error) {
    message.error('获取系列列表失败')
    console.error(error)
  } finally {
    seriesLoading.value = false
  }
}

function HandleCoverUploadChange(fileList: UploadFileInfo[]) {
  coverFileList.value = fileList.slice(-1)
  const file = coverFileList.value[0]
  if (file?.file) {
    modalForm.cover = URL.createObjectURL(file.file)
  } else if (fileList.length === 0) {
    modalForm.cover = ''
  }
}

function HandleCoverRemove() {
  coverFileList.value = []
  modalForm.cover = ''
  return true
}

function HandlePreviewUploadChange(fileList: UploadFileInfo[]) {
  previewFileList.value = fileList
  modalForm.previews = fileList
    .filter((f) => f.file || f.url)
    .map((f) => (f.file ? URL.createObjectURL(f.file) : f.url!))
}

function HandlePreviewRemove(options: { file: UploadFileInfo; fileList: UploadFileInfo[] }) {
  previewFileList.value = options.fileList
  modalForm.previews = options.fileList
    .filter((f) => f.file || f.url)
    .map((f) => (f.file ? URL.createObjectURL(f.file) : f.url!))
  return true
}

function Open(action: 'create' | 'update', row?: Product) {
  modalType.value = action
  if (action === 'create') {
    modalForm.id = undefined
    modalForm.name = ''
    modalForm.seriesId = null
    modalForm.image = ''
    modalForm.description = ''
    modalForm.cover = ''
    modalForm.previews = []
    modalForm.price = 0
    coverFileList.value = []
    previewFileList.value = []
  } else if (row) {
    modalForm.id = row.id
    modalForm.name = row.name
    modalForm.seriesId = Number(row.seriesId)
    modalForm.image = row.image
    modalForm.description = row.description
    modalForm.cover = row.cover
    modalForm.previews = [...row.previews]
    modalForm.price = row.price
    coverFileList.value = row.cover
      ? [{ id: 'cover', name: 'cover', status: 'finished', url: row.cover }]
      : []
    previewFileList.value = row.previews.map((url, index) => ({
      id: `preview-${index}`,
      name: `preview-${index}`,
      status: 'finished' as const,
      url,
    }))
  }
  showModal.value = true
  FetchSeriesList()
}

function HandleModalSubmit() {
  modalFormRef.value?.validate(async (errors) => {
    if (errors) return
    isLoading.value = true
    try {
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

defineExpose({
  Open,
})
</script>

<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增产品' : '编辑产品'"
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
            label="产品名称"
            path="name"
          >
            <NInput
              v-model:value="modalForm.name"
              placeholder="请输入产品名称"
            />
          </NFormItem>
          <NFormItem
            label="所属系列"
            path="seriesId"
          >
            <NSelect
              v-model:value="modalForm.seriesId"
              :options="seriesOptions"
              :loading="seriesLoading"
              placeholder="请选择所属系列"
              clearable
              filterable
            />
          </NFormItem>
          <NFormItem
            label="产品价格"
            path="price"
          >
            <NInputNumber
              v-model:value="modalForm.price"
              :min="0"
              :precision="2"
              placeholder="请输入产品价格"
              class="w-full"
            >
              <template #prefix>¥</template>
            </NInputNumber>
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div class="mb-4">
        <div class="mb-3 flex items-center gap-2 text-base font-medium">图片管理</div>
        <div class="grid grid-cols-2 gap-x-6 max-md:grid-cols-1">
          <NFormItem
            label="封面图"
            path="cover"
          >
            <NUpload
              :file-list="coverFileList"
              :max="1"
              list-type="image-card"
              accept="image/*"
              :default-upload="false"
              @update:file-list="HandleCoverUploadChange"
              @remove="HandleCoverRemove"
            />
          </NFormItem>
          <NFormItem
            label="预览图（最多10张）"
            path="previews"
          >
            <NUpload
              :file-list="previewFileList"
              :max="10"
              list-type="image-card"
              accept="image/*"
              :default-upload="false"
              @update:file-list="HandlePreviewUploadChange"
              @remove="HandlePreviewRemove"
            />
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div>
        <div class="mb-3 flex items-center gap-2 text-base font-medium">
          <span class="iconify text-lg text-primary ph--article" />
          产品描述
        </div>
        <NFormItem
          path="description"
          :show-label="false"
        >
          <NInput
            v-model:value="modalForm.description"
            type="textarea"
            placeholder="请输入产品描述信息..."
            :rows="4"
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
