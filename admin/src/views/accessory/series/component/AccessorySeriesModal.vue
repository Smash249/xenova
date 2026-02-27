<template>
  <NModal
    v-model:show="showModal"
    preset="card"
    :title="modalType === 'create' ? '新增配件系列' : '编辑配件系列'"
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

import { CreateAccessorySeriesApi, UpdateAccessorySeriesApi } from '@/api/accessory'

import type {
  CreateAccessorySeriesParams,
  AccessorySeries,
  UpdateAccessorySeriesParams,
} from '@/types/accessory'
import type { FormInst, FormRules } from 'naive-ui'

defineOptions({
  name: 'AccessorySeriesModal',
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

function Open(action: 'create' | 'update', row?: AccessorySeries) {
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
          await CreateAccessorySeries({
            name: modalForm.name,
            description: modalForm.description,
          })
          break
        case 'update':
          if (!modalForm.id) throw new Error('缺少系列ID')
          await UpdateAccessorySeries({
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

async function CreateAccessorySeries(params: CreateAccessorySeriesParams) {
  try {
    await CreateAccessorySeriesApi(params)
    message.success('新增成功')
    emit('success')
  } catch (error) {
    message.error('新增失败')
    console.error(error)
  }
}

async function UpdateAccessorySeries(params: UpdateAccessorySeriesParams) {
  try {
    await UpdateAccessorySeriesApi(params)
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
