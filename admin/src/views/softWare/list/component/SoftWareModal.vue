<script setup lang="tsx">
import {
  NButton,
  NForm,
  NFormItem,
  NInput,
  NModal,
  NSelect,
  NSpace,
  NSwitch,
  NUpload,
  useMessage,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed } from 'vue'

import { GetSoftWareSeriesListApi } from '@/api/softWare'

import type { SoftWare, SoftWareSeries } from '@/types/softWare'
import type { FormInst, FormRules, UploadFileInfo } from 'naive-ui'

interface ModelForm {
  id?: number
  name: string
  description: string
  file_type: string
  file_size: string
  file_url: string
  is_hot: boolean
  series_id: number | null
}

defineOptions({
  name: 'SoftWareModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入软件名称', trigger: 'blur' }],
  series_id: [
    { required: true, type: 'number', message: '请选择所属系列', trigger: ['blur', 'change'] },
  ],
  file_type: [{ required: true, message: '请输入文件类型', trigger: 'blur' }],
  file_url: [{ required: true, message: '请上传文件或输入下载地址', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const seriesList = ref<SoftWareSeries[]>([])
const seriesLoading = ref(false)

const seriesOptions = computed(() =>
  seriesList.value.map((item) => ({
    label: item.name,
    value: item.id,
  })),
)

const modalForm = reactive<ModelForm>({
  name: '',
  description: '',
  file_type: '',
  file_size: '',
  file_url: '',
  is_hot: false,
  series_id: null,
})

const fileList = ref<UploadFileInfo[]>([])

const fileTypeOptions = [
  { label: 'EXE', value: 'exe' },
  { label: 'MSI', value: 'msi' },
  { label: 'DMG', value: 'dmg' },
  { label: 'ZIP', value: 'zip' },
  { label: 'RAR', value: 'rar' },
  { label: 'APK', value: 'apk' },
  { label: 'ISO', value: 'iso' },
  { label: 'PDF', value: 'pdf' },
  { label: '其他', value: 'other' },
]

async function FetchSeriesList() {
  seriesLoading.value = true
  try {
    const result = await GetSoftWareSeriesListApi(1, 9999, '')
    seriesList.value = result.data
  } catch (error) {
    message.error('获取系列列表失败')
    console.error(error)
  } finally {
    seriesLoading.value = false
  }
}

function HandleFileUploadChange(options: { fileList: UploadFileInfo[] }) {
  fileList.value = options.fileList.slice(-1)
  const file = fileList.value[0]
  if (file?.file) {
    modalForm.file_url = URL.createObjectURL(file.file)
    modalForm.file_size = FormatFileSize(file.file.size)
    const ext = file.file.name.split('.').pop()?.toLowerCase() || ''
    const matched = fileTypeOptions.find((opt) => opt.value === ext)
    if (matched) modalForm.file_type = matched.value
  }
}

function HandleFileRemove() {
  fileList.value = []
  modalForm.file_url = ''
  modalForm.file_size = ''
  return true
}

function FormatFileSize(bytes: number): string {
  if (bytes === 0) return '0 B'
  const units = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return `${(bytes / Math.pow(1024, i)).toFixed(2)} ${units[i]}`
}

function Open(action: 'create' | 'update', row?: SoftWare) {
  modalType.value = action
  if (action === 'create') {
    modalForm.id = undefined
    modalForm.name = ''
    modalForm.description = ''
    modalForm.file_type = ''
    modalForm.file_size = ''
    modalForm.file_url = ''
    modalForm.is_hot = false
    modalForm.series_id = null
    fileList.value = []
  } else if (row) {
    modalForm.id = row.id
    modalForm.name = row.name
    modalForm.description = row.description
    modalForm.file_type = row.file_type
    modalForm.file_size = row.file_size
    modalForm.file_url = row.file_url
    modalForm.is_hot = row.is_hot
    modalForm.series_id = row.series_id
    fileList.value = row.file_url
      ? [{ id: 'file', name: row.name, status: 'finished', url: row.file_url }]
      : []
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
    :title="modalType === 'create' ? '新增下载' : '编辑下载'"
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
            label="软件名称"
            path="name"
          >
            <NInput
              v-model:value="modalForm.name"
              placeholder="请输入软件名称"
            />
          </NFormItem>
          <NFormItem
            label="所属系列"
            path="series_id"
          >
            <NSelect
              v-model:value="modalForm.series_id"
              :options="seriesOptions"
              :loading="seriesLoading"
              placeholder="请选择所属系列"
              clearable
              filterable
            />
          </NFormItem>
          <NFormItem
            label="文件类型"
            path="file_type"
          >
            <NSelect
              v-model:value="modalForm.file_type"
              :options="fileTypeOptions"
              placeholder="请选择文件类型"
              clearable
            />
          </NFormItem>
          <NFormItem
            label="热门推荐"
            path="is_hot"
          >
            <NSwitch v-model:value="modalForm.is_hot">
              <template #checked>是</template>
              <template #unchecked>否</template>
            </NSwitch>
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div class="mb-4">
        <div class="mb-3 flex items-center gap-2 text-base font-medium">文件管理</div>
        <div class="grid grid-cols-2 gap-x-6 max-md:grid-cols-1">
          <NFormItem
            label="上传文件"
            path="file_url"
          >
            <NUpload
              :file-list="fileList"
              :max="1"
              :default-upload="false"
              @change="HandleFileUploadChange"
              @remove="HandleFileRemove"
            >
              <NButton
                type="primary"
                ghost
              >
                <template #icon>
                  <span class="iconify ph--upload-simple" />
                </template>
                选择文件
              </NButton>
            </NUpload>
          </NFormItem>
          <NFormItem
            label="文件大小"
            path="file_size"
          >
            <NInput
              v-model:value="modalForm.file_size"
              placeholder="上传文件后自动填充"
              readonly
            />
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div>
        <div class="mb-3 flex items-center gap-2 text-base font-medium">
          <span class="iconify text-lg text-primary ph--article" />
          软件描述
        </div>
        <NFormItem
          path="description"
          :show-label="false"
        >
          <NInput
            v-model:value="modalForm.description"
            type="textarea"
            placeholder="请输入软件描述信息..."
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
