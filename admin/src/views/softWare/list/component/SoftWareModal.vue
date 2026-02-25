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
  NProgress,
  NTag,
  useMessage,
} from 'naive-ui'
import { ref, useTemplateRef, computed, nextTick } from 'vue'

import { CreateSoftWareApi, GetSoftWareSeriesListApi, UpdateSoftWareApi } from '@/api/softWare'
import { useUserStore } from '@/stores'

import type {
  CreateSoftWareParams,
  SoftWare,
  SoftWareSeries,
  UpdateSoftWareParams,
} from '@/types/softWare'
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

const DEFAULT_MODAL_FORM: ModelForm = {
  name: '',
  description: '',
  file_type: '',
  file_size: '',
  file_url: '',
  is_hot: false,
  series_id: null,
}

const uploadAction = import.meta.env.VITE_PROXY_PATH + '/private/upload'
const usser = useUserStore()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入软件名称', trigger: 'blur' }],
  series_id: [
    { required: true, type: 'number', message: '请选择所属系列', trigger: ['blur', 'change'] },
  ],
  file_url: [{ required: true, message: '请上传文件', trigger: 'change' }],
  description: [{ required: true, message: '请输入软件描述', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const seriesList = ref<SoftWareSeries[]>([])
const seriesLoading = ref(false)

const fileList = ref<UploadFileInfo[]>([])
const uploadProgress = ref(0)
const isUploading = ref(false)

const seriesOptions = computed(() =>
  seriesList.value.map((item) => ({
    label: item.name,
    value: item.id,
  })),
)

const uploadHeaders = computed(() => ({
  Authorization: `${usser.token}`,
}))

const modalForm = ref<ModelForm>({ ...DEFAULT_MODAL_FORM })

function FormatFileSize(bytes: number): string {
  if (bytes === 0) return '0 B'
  const units = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return `${(bytes / Math.pow(1024, i)).toFixed(2)} ${units[i]}`
}

function DetectFileType(fileName: string): string {
  const ext = fileName.split('.').pop()?.toLowerCase() || ''
  const knownTypes = [
    'exe',
    'msi',
    'dmg',
    'zip',
    'rar',
    'apk',
    'iso',
    'pdf',
    '7z',
    'tar',
    'gz',
    'deb',
    'pkg',
  ]
  return knownTypes.includes(ext) ? ext : ext || 'other'
}

function HandleBeforeUpload({ file }: { file: UploadFileInfo }) {
  if (file.file) {
    modalForm.value.file_size = FormatFileSize(file.file.size)
    modalForm.value.file_type = DetectFileType(file.file.name)
  }
  uploadProgress.value = 0
  isUploading.value = true
  return true
}

function HandleProgress({ event }: { file: UploadFileInfo; event?: ProgressEvent }) {
  if (event && event.lengthComputable) {
    uploadProgress.value = Math.round((event.loaded / event.total) * 100)
  }
}

function HandleFinish({ file, event }: { file: UploadFileInfo; event?: ProgressEvent }) {
  try {
    const response = JSON.parse((event?.target as XMLHttpRequest).response)
    modalForm.value.file_url = response.data.url
    file.status = 'finished'
    file.name = file.file?.name || file.name
    uploadProgress.value = 100
  } catch {
    file.status = 'error'
    message.error('上传解析失败')
  }

  nextTick().then(() => {
    isUploading.value = false
  })

  return file
}

function HandleError({ file }: { file: UploadFileInfo }) {
  message.error(`「${file.name}」上传失败`)
  isUploading.value = false
  uploadProgress.value = 0
  modalForm.value.file_type = ''
  modalForm.value.file_size = ''
  return file
}

function HandleRemove() {
  fileList.value = []
  modalForm.value.file_url = ''
  modalForm.value.file_type = ''
  modalForm.value.file_size = ''
  uploadProgress.value = 0
  return true
}

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

function Open(action: 'create' | 'update', row?: SoftWare) {
  modalType.value = action
  if (action === 'create') {
    modalForm.value = { ...DEFAULT_MODAL_FORM }
    fileList.value = []
  } else if (row) {
    modalForm.value = {
      id: row.id,
      name: row.name,
      description: row.description,
      file_type: row.file_type,
      file_size: row.file_size,
      file_url: row.file_url,
      is_hot: row.is_hot,
      series_id: row.series_id,
    }
    fileList.value = row.file_url
      ? [{ id: 'file', name: row.name, status: 'finished' as const }]
      : []
  }
  uploadProgress.value = 0
  isUploading.value = false
  showModal.value = true
  FetchSeriesList()
}

function HandleModalSubmit() {
  modalFormRef.value?.validate(async (errors) => {
    if (errors) return
    isLoading.value = true
    const form = modalForm.value
    try {
      switch (modalType.value) {
        case 'create':
          await CreateSoftWare({
            name: form.name,
            description: form.description,
            file_type: form.file_type,
            file_size: form.file_size,
            file_url: form.file_url,
            is_hot: form.is_hot,
            series_id: form.series_id!,
          })
          break
        case 'update':
          await UpdateSoftWare({
            id: form.id!,
            name: form.name,
            description: form.description,
            file_type: form.file_type,
            file_size: form.file_size,
            file_url: form.file_url,
            is_hot: form.is_hot,
            series_id: form.series_id!,
          })
          break
      }
      emit('success')
      showModal.value = false
    } catch (error) {
      message.error(modalType.value === 'create' ? '新增失败' : '更新失败')
      console.error(error)
    } finally {
      isLoading.value = false
    }
  })
}

async function CreateSoftWare(params: CreateSoftWareParams) {
  await CreateSoftWareApi(params)
  message.success('软件创建成功')
}

async function UpdateSoftWare(params: UpdateSoftWareParams) {
  await UpdateSoftWareApi(params)
  message.success('软件更新成功')
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
        </div>
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

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div class="mb-4">
        <div class="mb-3 flex items-center gap-2 text-base font-medium">文件管理</div>
        <NFormItem
          label="上传文件"
          path="file_url"
        >
          <div class="w-full">
            <NUpload
              v-model:file-list="fileList"
              :max="1"
              :action="uploadAction"
              :headers="uploadHeaders"
              :show-file-list="false"
              name="file"
              @before-upload="HandleBeforeUpload"
              @progress="HandleProgress"
              @finish="HandleFinish"
              @error="HandleError"
              @remove="HandleRemove"
            >
              <NButton
                type="primary"
                ghost
                :disabled="isUploading"
              >
                <template #icon>
                  <span class="iconify ph--upload-simple" />
                </template>
                {{ isUploading ? '上传中...' : '选择文件' }}
              </NButton>
            </NUpload>

            <div
              v-if="isUploading"
              class="mt-3"
            >
              <NProgress
                type="line"
                :percentage="uploadProgress"
                :height="8"
                indicator-placement="inside"
                processing
              />
            </div>

            <div
              v-if="modalForm.file_url && !isUploading"
              class="mt-3 flex items-center justify-between rounded-lg border border-gray-200 bg-gray-50 px-4 py-3 dark:border-gray-600 dark:bg-gray-800"
            >
              <div class="flex items-center gap-3">
                <span class="iconify text-2xl text-primary ph--file-arrow-down" />
                <div>
                  <div class="text-sm font-medium">
                    {{ fileList[0]?.name || modalForm.name }}
                  </div>
                  <div class="mt-0.5 flex items-center gap-2 text-xs text-gray-400">
                    <NTag
                      v-if="modalForm.file_type"
                      size="small"
                      :bordered="false"
                      type="info"
                    >
                      {{ modalForm.file_type.toUpperCase() }}
                    </NTag>
                    <span v-if="modalForm.file_size">{{ modalForm.file_size }}</span>
                  </div>
                </div>
              </div>
              <NButton
                text
                type="error"
                @click="HandleRemove"
              >
                <span class="iconify text-lg ph--trash" />
              </NButton>
            </div>
          </div>
        </NFormItem>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />
      <NFormItem
        path="description"
        label="软件描述"
      >
        <NInput
          v-model:value="modalForm.description"
          type="textarea"
          placeholder="请输入软件描述信息..."
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
            :disabled="isUploading"
          >
            {{ modalType === 'create' ? '立即创建' : '保存修改' }}
          </NButton>
        </NSpace>
      </div>
    </template>
  </NModal>
</template>
