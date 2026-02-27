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
  useMessage,
} from 'naive-ui'
import { ref, useTemplateRef, computed } from 'vue'

import { CreateAccessoryApi, GetAccessorySeriesListApi, UpdateAccessoryApi } from '@/api/accessory'
import ImageUpload from '@/components/ImageUpload.vue'

import type {
  CreateAccessoryParams,
  Accessory,
  AccessorySeries,
  UpdateAccessoryParams,
} from '@/types/accessory'
import type { FormInst, FormRules } from 'naive-ui'

interface ModelForm {
  id?: number
  name: string
  series_id: number
  image: string
  description: string
  cover: string
  previews: string[]
  price: number
}

defineOptions({
  name: 'AccessoryModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const DEFAULT_MODAL_FORM: ModelForm = {
  name: '',
  series_id: 0,
  image: '',
  description: '',
  cover: '',
  previews: [],
  price: 0,
}

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入配件名称', trigger: 'blur' }],
  series_id: [
    { required: true, type: 'number', message: '请选择所属系列', trigger: ['blur', 'change'] },
  ],
  price: [{ required: true, type: 'number', message: '请输入配件价格', trigger: 'blur' }],
  cover: [{ required: true, message: '请上传封面图', trigger: 'change' }],
  description: [{ required: true, message: '请输入配件描述', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const seriesList = ref<AccessorySeries[]>([])
const seriesLoading = ref(false)

const seriesOptions = computed(() =>
  seriesList.value.map((item) => ({
    label: item.name,
    value: item.id,
  })),
)

const modalForm = ref<ModelForm>({ ...DEFAULT_MODAL_FORM, previews: [] })

async function FetchSeriesList() {
  seriesLoading.value = true
  try {
    const result = await GetAccessorySeriesListApi(1, 9999, '')
    seriesList.value = result.data
  } catch (error) {
    message.error('获取系列列表失败')
    console.error(error)
  } finally {
    seriesLoading.value = false
  }
}

function Open(action: 'create' | 'update', row?: Accessory) {
  modalType.value = action
  if (action === 'create') {
    modalForm.value = { ...DEFAULT_MODAL_FORM, previews: [] }
  } else if (row) {
    modalForm.value = {
      id: row.id,
      name: row.name,
      series_id: row.series_id,
      image: row.image,
      description: row.description,
      cover: row.cover,
      previews: [...row.previews],
      price: row.price,
    }
  }
  showModal.value = true
  FetchSeriesList()
}

function HandleModalSubmit() {
  modalFormRef.value?.validate(async (errors) => {
    if (errors) return
    const form = modalForm.value
    try {
      switch (modalType.value) {
        case 'create':
          await CreateAccessory({
            name: form.name,
            series_id: form.series_id,
            image: form.image,
            description: form.description,
            cover: form.cover,
            previews: form.previews,
            price: form.price,
          })
          break
        case 'update':
          await UpdateAccessory({
            id: form.id!,
            name: form.name,
            series_id: form.series_id,
            image: form.image,
            description: form.description,
            cover: form.cover,
            previews: form.previews,
            price: form.price,
          })
          break
      }
    } catch (error) {
      message.error(modalType.value === 'create' ? '新增失败' : '更新失败')
      console.error(error)
    }
  })
}

async function CreateAccessory(params: CreateAccessoryParams) {
  try {
    await CreateAccessoryApi(params)
    message.success('配件创建成功')
    showModal.value = false
    emit('success')
  } catch (error) {
    message.error('配件创建失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function UpdateAccessory(params: UpdateAccessoryParams) {
  try {
    await UpdateAccessoryApi(params)
    message.success('配件更新成功')
    showModal.value = false
    emit('success')
  } catch (error) {
    message.error('配件更新失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

defineExpose({
  Open,
})
</script>

<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增配件' : '编辑配件'"
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
            label="配件名称"
            path="name"
          >
            <NInput
              v-model:value="modalForm.name"
              placeholder="请输入配件名称"
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
            label="配件价格"
            path="price"
          >
            <NInputNumber
              v-model:value="modalForm.price"
              :min="0"
              :precision="2"
              placeholder="请输入配件价格"
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
            <ImageUpload
              v-model="modalForm.cover"
              :max="1"
            />
          </NFormItem>
          <NFormItem
            label="预览图（最多10张）"
            path="previews"
          >
            <ImageUpload
              v-model="modalForm.previews"
              :max="10"
            />
          </NFormItem>
        </div>
      </div>

      <div class="mb-4 border-t border-gray-200 dark:border-gray-700" />

      <NFormItem
        path="description"
        label="配件描述"
      >
        <NInput
          v-model:value="modalForm.description"
          type="textarea"
          placeholder="请输入配件描述信息..."
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
