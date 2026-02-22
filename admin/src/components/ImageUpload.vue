<script setup lang="ts">
import { NProgress, NUpload, useMessage } from 'naive-ui'
import { computed, nextTick, ref, watch } from 'vue'

import { useUserStore } from '@/stores'

import type { UploadFileInfo } from 'naive-ui'

interface Props {
  max?: number
  action?: string
}

const usser = useUserStore()
const baseUrl = import.meta.env.VITE_SITE_BASE_API || ''

const props = withDefaults(defineProps<Props>(), {
  max: 1,
  action: import.meta.env.VITE_PROXY_PATH + '/private/upload',
})

const modelValue = defineModel<string | string[]>({
  required: true,
})

const message = useMessage()

const fileList = ref<UploadFileInfo[]>([])
const isMultiple = computed(() => props.max > 1)
const uploadProgress = ref<Record<string, number>>({})
const isUploading = computed(() => fileList.value.some((f) => f.status === 'uploading'))

const uploadHeaders = computed(() => {
  return {
    Authorization: `${usser.token}`,
  }
})

function WithPrefix(url: string): string {
  if (!url) return ''
  if (url.startsWith('http://') || url.startsWith('https://') || url.startsWith('blob:')) {
    return url
  }
  return baseUrl + url
}

function StripPrefix(url: string): string {
  if (!url) return ''
  if (baseUrl && url.startsWith(baseUrl)) {
    return url.slice(baseUrl.length)
  }
  return url
}

watch(
  modelValue,
  (val) => {
    if (isMultiple.value) {
      const urls = (val as string[]) || []
      const currentUrls = fileList.value.map((f) => StripPrefix(f.url || '')).filter(Boolean)
      if (JSON.stringify(urls) !== JSON.stringify(currentUrls)) {
        fileList.value = urls.map((url, index) => ({
          id: `file-${index}`,
          name: `file-${index}`,
          status: 'finished' as const,
          url: WithPrefix(url),
        }))
      }
    } else {
      const url = (val as string) || ''
      const currentUrl = StripPrefix(fileList.value[0]?.url || '')
      if (url !== currentUrl) {
        fileList.value = url
          ? [{ id: 'file-0', name: 'file-0', status: 'finished' as const, url: WithPrefix(url) }]
          : []
      }
    }
  },
  { immediate: true },
)

function SyncModelValue() {
  const urls = fileList.value
    .filter((f) => f.status === 'finished' && f.url)
    .map((f) => StripPrefix(f.url!))

  if (isMultiple.value) {
    modelValue.value = urls
  } else {
    modelValue.value = urls[0] || ''
  }
}

function HandleProgress({ file, event }: { file: UploadFileInfo; event?: ProgressEvent }) {
  if (event && event.lengthComputable) {
    uploadProgress.value[file.id] = Math.round((event.loaded / event.total) * 100)
  }
  return file
}

function HandleBeforeUpload({ file }: { file: UploadFileInfo }) {
  uploadProgress.value[file.id] = 0
  return true
}

function HandleFinish({ file, event }: { file: UploadFileInfo; event?: ProgressEvent }) {
  try {
    const response = JSON.parse((event?.target as XMLHttpRequest).response)
    const rawUrl = response.data.url
    file.url = WithPrefix(rawUrl)
    file.status = 'finished'
  } catch {
    file.status = 'error'
    message.error('上传解析失败')
  }
  uploadProgress.value[file.id] = 100

  nextTick().then(() => {
    SyncModelValue()
    setTimeout(() => {
      delete uploadProgress.value[file.id]
    }, 500)
  })

  return file
}

function HandleError({ file }: { file: UploadFileInfo }) {
  message.error(`「${file.name}」上传失败`)
  delete uploadProgress.value[file.id]
  return file
}

function HandleRemove({
  file,
  fileList: newFileList,
}: {
  file: UploadFileInfo
  fileList: UploadFileInfo[]
}) {
  delete uploadProgress.value[file.id]
  fileList.value = newFileList
  SyncModelValue()
  return true
}
</script>

<template>
  <div class="w-full">
    <NUpload
      v-model:file-list="fileList"
      :max="props.max"
      :multiple="isMultiple"
      list-type="image-card"
      accept="image/*"
      :action="props.action"
      :headers="uploadHeaders"
      name="file"
      @before-upload="HandleBeforeUpload"
      @progress="HandleProgress"
      @finish="HandleFinish"
      @error="HandleError"
      @remove="HandleRemove"
    />
    <div
      v-if="isUploading"
      class="mt-2 space-y-1"
    >
      <div
        v-for="file in fileList.filter((f) => f.status === 'uploading')"
        :key="file.id"
        class="flex items-center gap-2 text-sm text-gray-500"
      >
        <span class="w-24 truncate">{{ file.name }}</span>
        <NProgress
          type="line"
          :percentage="uploadProgress[file.id] || 0"
          :show-indicator="false"
          :height="6"
          class="flex-1"
        />
        <span class="w-10 text-right text-xs">{{ uploadProgress[file.id] || 0 }}%</span>
      </div>
    </div>
  </div>
</template>
