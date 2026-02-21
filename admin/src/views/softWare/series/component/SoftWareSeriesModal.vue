<script setup lang="tsx">
import { NButton, NForm, NFormItem, NInput, NModal, NSpace, useMessage } from 'naive-ui'
import { reactive, ref, useTemplateRef } from 'vue'

import type { SoftWareSeries } from '@/types/softWare'
import type { FormInst, FormRules } from 'naive-ui'

interface ModelForm {
  id?: number
  name: string
  description: string
}

defineOptions({
  name: 'SoftWareSeriesModal',
})

const emit = defineEmits<{
  (e: 'success'): void
}>()

const message = useMessage()

const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入系列名称', trigger: 'blur' }],
}

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')
const isLoading = ref(false)

const modalForm = reactive<ModelForm>({
  name: '',
  description: '',
})

function Open(action: 'create' | 'update', row?: SoftWareSeries) {
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
    :title="modalType === 'create' ? '新增下载系列' : '编辑下载系列'"
    class="max-w-md"
    :segmented="{ content: true, footer: true }"
  >
    <NForm
      ref="modalFormRef"
      :model="modalForm"
      :rules="rules"
      label-placement="left"
      label-width="100"
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
