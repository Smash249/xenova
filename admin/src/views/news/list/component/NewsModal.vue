<script setup lang="tsx">
import { MdEditor } from 'md-editor-v3'
import { NButton, NForm, NFormItem, NInput, NModal, NSelect, NSpace, useMessage } from 'naive-ui'
import { ref, useTemplateRef, computed } from 'vue'
import 'md-editor-v3/lib/style.css'

import { CreateNewsApi, GetNewsSeriesListApi, UpdateNewsApi } from '@/api/news'
import { useUserStore } from '@/stores'

import type { CreateNewsParams, News, NewsSeries, UpdateNewsParams } from '@/types/news'
import type { FormInst, FormRules } from 'naive-ui'

interface ModelForm {
  id?: number
  title: string
  content: string
  series_id: number | null
}

defineOptions({
  name: 'NewsModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const DEFAULT_MODAL_FORM: ModelForm = {
  title: '',
  content: '',
  series_id: null,
}

const uploadAction = import.meta.env.VITE_PROXY_PATH + '/private/upload'
const baseUrl = import.meta.env.VITE_SITE_BASE_API || ''
const usser = useUserStore()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  title: [{ required: true, message: '请输入新闻标题', trigger: 'blur' }],
  series_id: [
    { required: true, type: 'number', message: '请选择所属系列', trigger: ['blur', 'change'] },
  ],
  content: [{ required: true, message: '请输入新闻内容', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const seriesList = ref<NewsSeries[]>([])
const seriesLoading = ref(false)

const seriesOptions = computed(() =>
  seriesList.value.map((item) => ({
    label: item.name,
    value: item.id,
  })),
)

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

async function FetchSeriesList() {
  seriesLoading.value = true
  try {
    const result = await GetNewsSeriesListApi(1, 9999, '')
    seriesList.value = result.data
  } catch (error) {
    message.error('获取系列列表失败')
    console.error(error)
  } finally {
    seriesLoading.value = false
  }
}

function Open(action: 'create' | 'update', row?: News) {
  modalType.value = action
  if (action === 'create') {
    modalForm.value = { ...DEFAULT_MODAL_FORM }
  } else if (row) {
    modalForm.value = {
      id: row.id,
      title: row.title,
      content: row.content,
      series_id: row.series_id,
    }
  }
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
          await CreateNews({
            title: form.title,
            content: form.content,
            series_id: form.series_id!,
          })
          break
        case 'update':
          await UpdateNews({
            id: form.id!,
            title: form.title,
            content: form.content,
            series_id: form.series_id!,
          })
          break
      }
      emit('success')
      showModal.value = false
    } catch (error) {
      message.error(modalType.value === 'create' ? '发布失败' : '更新失败')
      console.error(error)
    } finally {
      isLoading.value = false
    }
  })
}

async function CreateNews(params: CreateNewsParams) {
  await CreateNewsApi(params)
  message.success('发布成功')
}

async function UpdateNews(params: UpdateNewsParams) {
  await UpdateNewsApi(params)
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
    :title="modalType === 'create' ? '发布新闻' : '编辑新闻'"
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
            label="新闻标题"
            path="title"
          >
            <NInput
              v-model:value="modalForm.title"
              placeholder="请输入新闻标题"
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
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <div>
        <div class="mb-3 flex items-center gap-2 text-base font-medium">新闻内容</div>
        <NFormItem
          path="content"
          :show-label="false"
        >
          <MdEditor
            v-model="modalForm.content"
            :style="{ height: '450px' }"
            placeholder="请输入新闻内容，支持 Markdown 格式..."
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
            {{ modalType === 'create' ? '立即发布' : '保存修改' }}
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
