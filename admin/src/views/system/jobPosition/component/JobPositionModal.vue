<script setup lang="tsx">
import {
  NModal,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NButton,
  NSelect,
  useMessage,
  NGrid,
  NGi,
} from 'naive-ui'
import { reactive, ref, computed } from 'vue'

import { CreateSystemJobPositionApi, UpdateSystemJobPositionApi } from '@/api/system'

import type { SystemJobPosition } from '@/types/system'
import type { FormInst, FormRules } from 'naive-ui'

type ModalMode = 'create' | 'update'

const emit = defineEmits<{
  success: []
}>()

const message = useMessage()

const formRef = ref<FormInst | null>(null)
const visible = ref(false)
const mode = ref<ModalMode>('create')
const isSubmitting = ref(false)
const editId = ref<number | null>(null)

const formModel = reactive({
  title: '',
  department: '',
  location: '',
  headcount: '',
  experience: '',
  education: '',
  salaryRange: '',
  responsibilities: '',
  requirements: '',
  status: 1,
  sort: 0,
})

const modalTitle = computed(() => (mode.value === 'create' ? '新增职位' : '编辑职位'))

const statusOptions = [
  { label: '招聘中', value: 1 },
  { label: '已停止', value: 0 },
]

const rules: FormRules = {
  title: { required: true, message: '请输入职位名称', trigger: 'blur' },
  department: { required: true, message: '请输入部门', trigger: 'blur' },
  location: { required: true, message: '请输入工作地点', trigger: 'blur' },
  headcount: { required: true, message: '请输入招聘人数', trigger: 'blur' },
  experience: { required: true, message: '请输入经验要求', trigger: 'blur' },
  education: { required: true, message: '请输入学历要求', trigger: 'blur' },
  salaryRange: { required: true, message: '请输入薪资范围', trigger: 'blur' },
  responsibilities: { required: true, message: '请输入岗位职责', trigger: 'blur' },
  requirements: { required: true, message: '请输入任职要求', trigger: 'blur' },
}

function ResetForm() {
  formModel.title = ''
  formModel.department = ''
  formModel.location = ''
  formModel.headcount = ''
  formModel.experience = ''
  formModel.education = ''
  formModel.salaryRange = ''
  formModel.responsibilities = ''
  formModel.requirements = ''
  formModel.status = 1
  formModel.sort = 0
  editId.value = null
}

function Open(m: ModalMode, row?: SystemJobPosition) {
  mode.value = m
  ResetForm()

  if (m === 'update' && row) {
    editId.value = row.id
    formModel.title = row.title
    formModel.department = row.department
    formModel.location = row.location
    formModel.headcount = row.headcount
    formModel.experience = row.experience
    formModel.education = row.education
    formModel.salaryRange = row.salaryRange
    formModel.responsibilities = row.responsibilities
    formModel.requirements = row.requirements
    formModel.status = row.status
    formModel.sort = row.sort
  }

  visible.value = true
}

async function HandleSubmit() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }

  isSubmitting.value = true
  try {
    if (mode.value === 'create') {
      await CreateSystemJobPositionApi({
        title: formModel.title,
        department: formModel.department,
        location: formModel.location,
        headcount: formModel.headcount,
        experience: formModel.experience,
        education: formModel.education,
        salaryRange: formModel.salaryRange,
        responsibilities: formModel.responsibilities,
        requirements: formModel.requirements,
      })
      message.success('创建成功')
    } else {
      await UpdateSystemJobPositionApi({
        id: editId.value!,
        title: formModel.title,
        department: formModel.department,
        location: formModel.location,
        headcount: formModel.headcount,
        experience: formModel.experience,
        education: formModel.education,
        salaryRange: formModel.salaryRange,
        responsibilities: formModel.responsibilities,
        requirements: formModel.requirements,
        status: formModel.status,
        sort: formModel.sort,
      })
      message.success('更新成功')
    }
    visible.value = false
    emit('success')
  } catch (error) {
    message.error(mode.value === 'create' ? '创建失败' : '更新失败')
    console.error(error)
  } finally {
    isSubmitting.value = false
  }
}

defineExpose({ Open })
</script>

<template>
  <NModal
    v-model:show="visible"
    preset="card"
    :title="modalTitle"
    :bordered="false"
    :mask-closable="false"
    style="width: 720px"
    :segmented="{ content: true, action: true }"
  >
    <NForm
      ref="formRef"
      :model="formModel"
      :rules="rules"
      label-placement="left"
      label-width="auto"
      require-mark-placement="right-hanging"
    >
      <NGrid
        :cols="2"
        :x-gap="16"
      >
        <NGi>
          <NFormItem
            label="职位名称"
            path="title"
          >
            <NInput
              v-model:value="formModel.title"
              placeholder="请输入职位名称"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi>
          <NFormItem
            label="部门"
            path="department"
          >
            <NInput
              v-model:value="formModel.department"
              placeholder="请输入部门"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi>
          <NFormItem
            label="工作地点"
            path="location"
          >
            <NInput
              v-model:value="formModel.location"
              placeholder="请输入工作地点"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi>
          <NFormItem
            label="招聘人数"
            path="headcount"
          >
            <NInput
              v-model:value="formModel.headcount"
              placeholder="例如：2人、若干"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi>
          <NFormItem
            label="经验要求"
            path="experience"
          >
            <NInput
              v-model:value="formModel.experience"
              placeholder="例如：3-5年"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi>
          <NFormItem
            label="学历要求"
            path="education"
          >
            <NInput
              v-model:value="formModel.education"
              placeholder="例如：本科及以上"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi>
          <NFormItem
            label="薪资范围"
            path="salaryRange"
          >
            <NInput
              v-model:value="formModel.salaryRange"
              placeholder="例如：15K-25K"
              clearable
            />
          </NFormItem>
        </NGi>
        <NGi v-if="mode === 'update'">
          <NFormItem label="状态">
            <NSelect
              v-model:value="formModel.status"
              :options="statusOptions"
              placeholder="请选择状态"
            />
          </NFormItem>
        </NGi>
        <NGi v-if="mode === 'update'">
          <NFormItem label="排序">
            <NInputNumber
              v-model:value="formModel.sort"
              :min="0"
              placeholder="排序值"
              class="w-full"
            />
          </NFormItem>
        </NGi>
      </NGrid>

      <NFormItem
        label="岗位职责"
        path="responsibilities"
      >
        <NInput
          v-model:value="formModel.responsibilities"
          type="textarea"
          placeholder="请输入岗位职责，每条一行"
          :rows="4"
        />
      </NFormItem>
      <NFormItem
        label="任职要求"
        path="requirements"
      >
        <NInput
          v-model:value="formModel.requirements"
          type="textarea"
          placeholder="请输入任职要求，每条一行"
          :rows="4"
        />
      </NFormItem>
    </NForm>

    <template #action>
      <div class="flex justify-end gap-3">
        <NButton @click="visible = false">取消</NButton>
        <NButton
          type="primary"
          :loading="isSubmitting"
          :disabled="isSubmitting"
          @click="HandleSubmit"
        >
          {{ mode === 'create' ? '创建' : '保存' }}
        </NButton>
      </div>
    </template>
  </NModal>
</template>

<style scoped lang="scss"></style>
