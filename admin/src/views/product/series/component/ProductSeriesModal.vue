<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增产品系列' : '编辑产品系列'"
    class="max-w-md"
  >
    <NForm
      ref="modalFormRef"
      :model="modalForm"
      :rules="rules"
      label-placement="left"
      label-width="100px"
    >
      <NFormItem
        label="系列名称"
        path="name"
      >
        <NInput
          v-model:value="modalForm.name"
          placeholder="请输入系列名称"
        />
      </NFormItem>
      <NFormItem
        label="描述"
        path="description"
      >
        <NInput
          v-model:value="modalForm.description"
          type="textarea"
          placeholder="请输入系列描述"
          :rows="3"
        />
      </NFormItem>
      <div class="mt-4 flex justify-end">
        <NSpace>
          <NButton @click="showModal = false">取消</NButton>
          <NButton
            type="primary"
            @click="HandleModalSubmit"
            :loading="isLoading"
            >确认</NButton
          >
        </NSpace>
      </div>
    </NForm>
  </NModal>
</template>

<script setup lang="tsx">
import { NButton, NForm, NFormItem, NInput, NModal, NSpace, useMessage } from 'naive-ui'
import { reactive, ref, useTemplateRef } from 'vue'

import { CreateProductSeriesApi, UpdateProductSeriesApi } from '@/api/product'

import type {
  CreateProductSeriesParams,
  ProductSeries,
  UpdateProductSeriesParams,
} from '@/types/product'
import type { FormInst, FormRules } from 'naive-ui'

defineOptions({
  name: 'ProductSeriesModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入系列名称', trigger: 'blur' }],
  description: [{ required: true, message: '请输入系列描述', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const modalForm = reactive<{ id?: number; name: string; description: string }>({
  name: '',
  description: '',
})

function Open(action: 'create' | 'update', row?: ProductSeries) {
  modalType.value = action
  if (action === 'create') {
    modalForm.id = undefined
    modalForm.name = ''
    modalForm.description = ''
  } else if (row) {
    modalForm.id = row.id
    modalForm.name = row.name
    modalForm.description = row.description
  }
  showModal.value = true
}

function HandleModalSubmit() {
  modalFormRef.value?.validate(async (errors) => {
    if (errors) return
    isLoading.value = true
    try {
      switch (modalType.value) {
        case 'create':
          await CreateProductSeries({
            name: modalForm.name,
            description: modalForm.description,
          })
          break
        case 'update':
          if (!modalForm.id) throw new Error('缺少系列ID')
          await UpdateProductSeries({
            id: modalForm.id,
            name: modalForm.name,
            description: modalForm.description,
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

async function CreateProductSeries(params: CreateProductSeriesParams) {
  try {
    await CreateProductSeriesApi(params)
    message.success('新增成功')
    emit('success')
  } catch (error) {
    message.error('新增失败')
    console.error(error)
  }
}

async function UpdateProductSeries(params: UpdateProductSeriesParams) {
  try {
    await UpdateProductSeriesApi(params)
    message.success('更新成功')
    emit('success')
  } catch (error) {
    message.error('更新失败')
    console.error(error)
  }
}

defineExpose({
  Open,
})
</script>
