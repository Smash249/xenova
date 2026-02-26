<template>
  <ScrollContainer
    wrapper-class="flex flex-col gap-y-2"
    :scrollable="true"
  >
    <NCard
      class="h-full flex-1"
      content-class="flex flex-col h-full"
    >
      <NForm
        ref="formRef"
        :model="queryForm"
        inline
        label-placement="left"
        class="max-lg:w-full max-lg:flex-col"
      >
        <NFormItem
          label="职位名称"
          path="title"
        >
          <NInput
            v-model:value="queryForm.title"
            clearable
            placeholder="请输入职位名称"
          />
        </NFormItem>
        <NFormItem
          label="部门"
          path="department"
        >
          <NInput
            v-model:value="queryForm.department"
            clearable
            placeholder="请输入部门"
          />
        </NFormItem>
        <div class="flex gap-2">
          <NButton
            type="success"
            @click="modalRef?.Open('create')"
          >
            <template #icon>
              <span class="iconify ph--plus-circle" />
            </template>
            新增职位
          </NButton>
          <NPopconfirm
            positive-text="确定"
            negative-text="取消"
            @positive-click="MutateBatchDeleteData"
          >
            <template #default> 确认删除选中的 {{ checkedRowKeys.length }} 条记录吗？ </template>
            <template #trigger>
              <NButton
                type="error"
                :disabled="!hasChecked"
              >
                <template #icon>
                  <span class="iconify ph--trash" />
                </template>
                批量删除 ({{ checkedRowKeys.length }})
              </NButton>
            </template>
          </NPopconfirm>
          <NButton
            type="info"
            :loading="isLoading"
            :disabled="isLoading"
            @click="HandleQueryClick"
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
      </NForm>

      <div class="flex min-h-0 flex-1 flex-col">
        <NDataTable
          class="flex-1"
          v-model:checked-row-keys="checkedRowKeys"
          :remote="true"
          :columns="columns"
          :data="tableData"
          :row-key="(row: SystemJobPosition) => row.id"
          :loading="isLoading"
          scroll-x="1600"
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

    <JobPositionModal
      ref="modalRef"
      @success="GetJobPositionList"
    />
  </ScrollContainer>
</template>

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
  NTag,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed, onMounted } from 'vue'

import { GetSystemJobPositionListApi, DeleteSystemJobPositionApi } from '@/api/system'
import { ScrollContainer } from '@/components'

import JobPositionModal from './component/JobPositionModal.vue'

import type { SystemJobPosition } from '@/types/system'
import type { DataTableColumns, PaginationProps, DropdownProps, FormInst } from 'naive-ui'

defineOptions({
  name: 'JobPositionAdmin',
})

const message = useMessage()

const formRef = useTemplateRef<FormInst>('formRef')
const modalRef = ref<InstanceType<typeof JobPositionModal> | null>(null)

const showDropdown = ref(false)
const contextmenuId = ref<number | null>(null)

const tableData = ref<SystemJobPosition[]>([])
const checkedRowKeys = ref<Array<number | string>>([])
const isLoading = ref(false)

const queryForm = reactive({
  title: '',
  department: '',
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
    GetJobPositionList()
  },
  onUpdatePageSize: (pageSize) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    GetJobPositionList()
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
      const row = tableData.value.find((item) => item.id === contextmenuId.value)
      switch (v) {
        case 'edit':
          if (row) modalRef.value?.Open('update', row)
          break
        case 'delete':
          MutateDeleteData(contextmenuId.value)
          break
      }
    }
    showDropdown.value = false
  },
})

const hasChecked = computed(() => checkedRowKeys.value.length > 0)

const columns = computed<DataTableColumns<SystemJobPosition>>(() => {
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
      width: 70,
    },
    {
      key: 'title',
      title: '职位名称',
      width: 150,
      render: (row) => <span class='font-medium text-slate-700'>{row.title}</span>,
    },
    {
      key: 'department',
      title: '部门',
      width: 120,
      render: (row) => (
        <NTag
          size='small'
          type='info'
          bordered={false}
        >
          {row.department}
        </NTag>
      ),
    },
    {
      key: 'location',
      title: '工作地点',
      width: 120,
    },
    {
      key: 'headcount',
      title: '招聘人数',
      width: 100,
      align: 'center',
    },
    {
      key: 'salaryRange',
      title: '薪资范围',
      width: 130,
      render: (row) => (
        <NTag
          size='small'
          type='warning'
          bordered={false}
        >
          {row.salaryRange}
        </NTag>
      ),
    },
    {
      key: 'experience',
      title: '经验要求',
      width: 110,
    },
    {
      key: 'education',
      title: '学历要求',
      width: 100,
    },
    {
      key: 'status',
      title: '状态',
      width: 100,
      align: 'center',
      render: (row) => (
        <NTag
          size='small'
          type={row.status === 1 ? 'success' : 'error'}
          bordered={false}
        >
          {row.status === 1 ? '招聘中' : '已停止'}
        </NTag>
      ),
    },
    {
      key: 'sort',
      title: '排序',
      width: 80,
      align: 'center',
    },
    {
      key: 'created_at',
      title: '创建时间',
      width: 180,
      render: (row) => dayjs(row.created_at).format('YYYY-MM-DD HH:mm:ss'),
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

function CellActions(row: SystemJobPosition) {
  return (
    <div class='flex justify-center gap-2'>
      <NButton
        secondary
        type='primary'
        size='small'
        onClick={() => modalRef.value?.Open('update', row)}
      >
        编辑
      </NButton>
      <NPopconfirm
        positiveText='确定'
        negativeText='取消'
        onPositiveClick={() => MutateDeleteData(row.id)}
      >
        {{
          default: () => '确认删除该职位吗？',
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
      GetJobPositionList()
    }
  })
}

function ResetForm() {
  queryForm.title = ''
  queryForm.department = ''
  pagination.page = 1
  GetJobPositionList()
}

async function MutateDeleteData(id: number) {
  isLoading.value = true
  try {
    await DeleteSystemJobPositionApi([id])
    message.success('删除成功')
    checkedRowKeys.value = checkedRowKeys.value.filter((key) => key !== id)
    GetJobPositionList()
  } catch (error) {
    message.error('删除失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function MutateBatchDeleteData() {
  if (!hasChecked.value) return
  isLoading.value = true
  try {
    await DeleteSystemJobPositionApi(checkedRowKeys.value as number[])
    message.success(`成功删除 ${checkedRowKeys.value.length} 条记录`)
    checkedRowKeys.value = []
    GetJobPositionList()
  } catch (error) {
    message.error('批量删除失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function GetJobPositionList() {
  isLoading.value = true
  try {
    const result = await GetSystemJobPositionListApi(
      pagination.page || 1,
      pagination.pageSize || 10,
      queryForm.title,
      queryForm.department,
    )
    tableData.value = result.data
    if (result.paginate) pagination.itemCount = result.paginate.total_count
  } catch (error) {
    message.error('获取职位列表失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  GetJobPositionList()
})
</script>

<style scoped lang="scss"></style>
