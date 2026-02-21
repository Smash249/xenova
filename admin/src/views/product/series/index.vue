<script setup lang="tsx">
import dayjs from 'dayjs'
import {
  NButton,
  NDataTable,
  NCard,
  NForm,
  NFormItem,
  NInput,
  NPopconfirm,
  useMessage,
  NPagination,
  NDropdown,
  NModal,
  NSpace,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed, onMounted } from 'vue'

import { GetProductSeriesListApi, DeleteProductSeriesApi } from '@/api/product'
import { ScrollContainer } from '@/components'

import type {
  DataTableColumns,
  PaginationProps,
  FormRules,
  DropdownProps,
  FormInst,
} from 'naive-ui'

interface BaseResponse {
  created_at: string
  updated_at: string
}

interface ProductSeries extends BaseResponse {
  id: number
  name: string
  description: string
}

defineOptions({
  name: 'ProductSeriesAdmin',
})

const message = useMessage()

const formRef = useTemplateRef<FormInst>('formRef')
const modalFormRef = useTemplateRef<FormInst>('modalFormRef')

const rules: FormRules = {
  name: [{ required: true, message: '请输入系列名称', trigger: 'blur' }],
}

const showDropdown = ref(false)
const contextmenuId = ref<number | null>(null)

const productSeriesData = ref<ProductSeries[]>([])
const checkedRowKeys = ref<Array<number | string>>([])
const isLoading = ref(false)

const showModal = ref(false)
const modalType = ref<'create' | 'update'>('create')

const queryForm = reactive({
  name: '',
})

const modalForm = reactive<{ id?: number; name: string; description: string }>({
  name: '',
  description: '',
})

const pagination = reactive<PaginationProps>({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 50, 100],
  itemCount: 0,
  showQuickJumper: true,
  showQuickJumpDropdown: true,
  onUpdatePage: (page) => {
    pagination.page = page
    GetProductSeries()
  },
  onUpdatePageSize: (pageSize) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    GetProductSeries()
  },
})

const dropdownOptions = reactive<DropdownProps>({
  x: 0,
  y: 0,
  options: [
    {
      label: '编辑',
      key: 'edit',
    },
    {
      label: () => <span class='text-rose-500'>删除</span>,
      key: 'delete',
    },
  ],
  onClickoutside: () => {
    showDropdown.value = false
  },
  onSelect: (v) => {
    if (contextmenuId.value !== null) {
      const row = productSeriesData.value.find((item) => item.id === contextmenuId.value)
      switch (v) {
        case 'edit':
          if (row) CreateOrEditData('update', row)
          break
        case 'delete':
          MutateDeleteData(contextmenuId.value)
          break
      }
    }
    showDropdown.value = false
  },
})

const columns = computed<DataTableColumns<ProductSeries>>(() => {
  return [
    {
      type: 'selection',
      options: ['all', 'none'],
    },
    {
      key: 'id',
      title: 'ID',
      titleAlign: 'center',
      align: 'center',
      width: 80,
    },
    {
      key: 'name',
      title: '系列名称',
      width: 200,
    },
    {
      key: 'description',
      title: '描述',
      minWidth: 250,
    },
    {
      key: 'created_at',
      title: '创建时间',
      width: 180,
      render: (row) => dayjs(row.created_at).format('YYYY-MM-DD HH:mm:ss'),
    },
    {
      key: 'updated_at',
      title: '更新时间',
      width: 180,
      render: (row) => dayjs(row.updated_at).format('YYYY-MM-DD HH:mm:ss'),
    },
    {
      key: 'actions',
      title: '操作',
      width: 160,
      align: 'center',
      fixed: 'right',
      render: (row) => CellActions(row),
    },
  ]
})

function CellActions(row: ProductSeries) {
  return (
    <div class='flex justify-center gap-2'>
      <NButton
        secondary
        type='primary'
        size='small'
        onClick={() => CreateOrEditData('update', row)}
      >
        编辑
      </NButton>
      <NPopconfirm
        positiveText='确定'
        negativeText='取消'
        onPositiveClick={() => MutateDeleteData(row.id)}
      >
        {{
          default: () => '确认删除该产品系列吗？',
          trigger: () => (
            <NButton
              secondary
              type='error'
              size='small'
            >
              删除
            </NButton>
          ),
        }}
      </NPopconfirm>
    </div>
  )
}

function PaginationPrefix(info: { itemCount: number | undefined }) {
  const { itemCount } = info
  return (
    itemCount !== undefined && (
      <div>
        <span>总数&nbsp;</span>
        {itemCount}
        <span>&nbsp;条</span>
      </div>
    )
  )
}

function HandleQueryClick() {
  formRef.value?.validate((errors) => {
    if (!errors) {
      pagination.page = 1
      GetProductSeries()
    }
  })
}

function ResetForm() {
  queryForm.name = ''
  pagination.page = 1
  GetProductSeries()
}

function CreateOrEditData(action: 'create' | 'update', row?: ProductSeries) {
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
    if (!errors) {
      isLoading.value = true
      try {
        showModal.value = false
        GetProductSeries()
      } catch (error) {
        message.error(modalType.value === 'create' ? '新增失败' : '更新失败')
        console.error(error)
      } finally {
        isLoading.value = false
      }
    }
  })
}

async function MutateDeleteData(id: number) {
  isLoading.value = true
  try {
    await DeleteProductSeriesApi([id])
    message.success('删除成功')
    GetProductSeries()
  } catch (error) {
    message.error('删除失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function GetProductSeries() {
  isLoading.value = true
  try {
    const result = await GetProductSeriesListApi(
      pagination.page || 1,
      pagination.pageSize || 10,
      queryForm.name,
    )
    console.log('GetProductSeriesListApi result:', result.data)
    productSeriesData.value = result.data
    if (result.paginate) pagination.itemCount = result.paginate.total_count
  } catch (error) {
    message.error('获取列表失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  GetProductSeries()
})
</script>

<template>
  <ScrollContainer
    wrapper-class="flex flex-col gap-y-2"
    :scrollable="true"
  >
    <NCard
      class="flex-1"
      content-class="flex flex-col"
    >
      <div class="mb-2 flex justify-between gap-x-4 max-xl:mb-4 max-xl:flex-wrap">
        <NForm
          ref="formRef"
          :model="queryForm"
          inline
          label-placement="left"
          class="max-lg:w-full max-lg:flex-col"
        >
          <NFormItem
            label="系列名称"
            path="name"
          >
            <NInput
              v-model:value="queryForm.name"
              clearable
              placeholder="请输入系列名称"
            />
          </NFormItem>
        </NForm>
        <div class="flex gap-2">
          <NButton
            type="success"
            @click="CreateOrEditData('create')"
          >
            <template #icon>
              <span class="iconify ph--plus-circle" />
            </template>
            新增系列
          </NButton>
          <NButton
            type="info"
            @click="HandleQueryClick"
            :loading="isLoading"
            :disabled="isLoading"
          >
            <template #icon>
              <span class="iconify ph--magnifying-glass" />
            </template>
            查询
          </NButton>
          <NButton
            type="warning"
            @click="ResetForm"
          >
            <template #icon>
              <span class="iconify ph--arrow-clockwise" />
            </template>
            重置
          </NButton>
        </div>
      </div>

      <div class="flex flex-1 flex-col">
        <NDataTable
          class="flex-1"
          ref="dataTableRef"
          v-model:checked-row-keys="checkedRowKeys"
          :remote="true"
          :columns="columns"
          :data="productSeriesData"
          :row-key="(row) => row.id"
          :loading="isLoading"
        />

        <div class="mt-3 flex items-end justify-end max-xl:flex-col max-xl:gap-y-2">
          <NPagination
            v-bind="pagination"
            :prefix="PaginationPrefix"
            :disabled="isLoading"
          />
        </div>
      </div>
    </NCard>

    <NDropdown
      placement="bottom-start"
      trigger="manual"
      v-bind="dropdownOptions"
      :show="showDropdown"
    />

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
        label-width="80"
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
  </ScrollContainer>
</template>

<style scoped lang="scss"></style>
