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
          label="姓名"
          path="name"
        >
          <NInput
            v-model:value="queryForm.name"
            clearable
            placeholder="请输入姓名"
          />
        </NFormItem>
        <NFormItem
          label="电话"
          path="phone"
        >
          <NInput
            v-model:value="queryForm.phone"
            clearable
            placeholder="请输入电话"
          />
        </NFormItem>
        <div class="flex gap-2">
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
          :row-key="(row: SystemMessage) => row.id"
          :loading="isLoading"
          scroll-x="1500"
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
  NTooltip,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed, onMounted } from 'vue'

import { GetSystemMessageListApi, DeleteSystemMessagedApi } from '@/api/system'
import { ScrollContainer } from '@/components'

import type { SystemMessage } from '@/types/system'
import type { DataTableColumns, PaginationProps, DropdownProps, FormInst } from 'naive-ui'

defineOptions({
  name: 'SystemMessageAdmin',
})

const message = useMessage()

const formRef = useTemplateRef<FormInst>('formRef')

const showDropdown = ref(false)
const contextmenuId = ref<number | null>(null)

const tableData = ref<SystemMessage[]>([])
const checkedRowKeys = ref<Array<number | string>>([])
const isLoading = ref(false)

const queryForm = reactive({
  name: '',
  phone: '',
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
    GetMessageList()
  },
  onUpdatePageSize: (pageSize) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    GetMessageList()
  },
})

const dropdownOptions = reactive<DropdownProps>({
  x: 0,
  y: 0,
  options: [
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
      switch (v) {
        case 'delete':
          MutateDeleteData(contextmenuId.value)
          break
      }
    }
    showDropdown.value = false
  },
})

const hasChecked = computed(() => checkedRowKeys.value.length > 0)

const columns = computed<DataTableColumns<SystemMessage>>(() => {
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
      title: '姓名',
      width: 120,
      render: (row) => <span class='font-medium text-slate-700'>{row.name}</span>,
    },
    {
      key: 'phone',
      title: '电话',
      width: 140,
      render: (row) => <span class='font-mono text-slate-600'>{row.phone}</span>,
    },
    {
      key: 'email',
      title: '邮箱',
      width: 200,
      ellipsis: {
        tooltip: true,
      },
      render: (row) =>
        row.email ? (
          <span class='text-slate-600'>{row.email}</span>
        ) : (
          <NTag
            size='small'
            type='default'
          >
            暂无
          </NTag>
        ),
    },
    {
      key: 'company',
      title: '公司',
      width: 180,
      ellipsis: {
        tooltip: true,
      },
      render: (row) =>
        row.company ? (
          <span class='text-slate-600'>{row.company}</span>
        ) : (
          <NTag
            size='small'
            type='default'
          >
            暂无
          </NTag>
        ),
    },
    {
      key: 'content',
      title: '留言内容',
      width: 280,
      render: (row) => (
        <NTooltip
          placement='top'
          trigger='hover'
          disabled={!row.content || row.content.length <= 50}
        >
          {{
            default: () => <span class='max-w-80 whitespace-pre-wrap'>{row.content}</span>,
            trigger: () => <span class='line-clamp-2 text-sm text-slate-600'>{row.content}</span>,
          }}
        </NTooltip>
      ),
    },
    {
      key: 'created_at',
      title: '提交时间',
      width: 180,
      sorter: true,
      render: (row) => dayjs(row.created_at).format('YYYY-MM-DD HH:mm:ss'),
    },
    {
      key: 'actions',
      title: '操作',
      width: 100,
      align: 'center',
      fixed: 'right',
      render: (row) => CellActions(row),
    },
  ]
})

function CellActions(row: SystemMessage) {
  return (
    <div class='flex justify-center'>
      <NPopconfirm
        positiveText='确定'
        negativeText='取消'
        onPositiveClick={() => MutateDeleteData(row.id)}
      >
        {{
          default: () => '确认删除该留言吗？',
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
      GetMessageList()
    }
  })
}

function ResetForm() {
  queryForm.name = ''
  queryForm.phone = ''
  pagination.page = 1
  GetMessageList()
}

async function MutateDeleteData(id: number) {
  isLoading.value = true
  try {
    await DeleteSystemMessagedApi([id])
    message.success('删除成功')
    checkedRowKeys.value = checkedRowKeys.value.filter((key) => key !== id)
    GetMessageList()
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
    await DeleteSystemMessagedApi(checkedRowKeys.value as number[])
    message.success(`成功删除 ${checkedRowKeys.value.length} 条记录`)
    checkedRowKeys.value = []
    GetMessageList()
  } catch (error) {
    message.error('批量删除失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function GetMessageList() {
  isLoading.value = true
  try {
    const result = await GetSystemMessageListApi(
      pagination.page || 1,
      pagination.pageSize || 10,
      queryForm.name,
      queryForm.phone,
    )
    tableData.value = result.data
    if (result.paginate) pagination.itemCount = result.paginate.total_count
  } catch (error) {
    message.error('获取留言列表失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  GetMessageList()
})
</script>

<style scoped lang="scss"></style>
