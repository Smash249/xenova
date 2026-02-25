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
  NImage,
  NTag,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed, onMounted } from 'vue'

import { GetCompanyPatentListApi, DeleteCompanyPatentApi } from '@/api/honor'
import { ScrollContainer } from '@/components'

import CompanyPatentModal from './component/CompanyPatentModal.vue'

import type { CompanyPatent } from '@/types/honor'
import type { DataTableColumns, PaginationProps, DropdownProps, FormInst } from 'naive-ui'

defineOptions({
  name: 'CompanyPatentAdmin',
})

const message = useMessage()

const formRef = useTemplateRef<FormInst>('formRef')
const modalRef = ref<InstanceType<typeof CompanyPatentModal> | null>(null)

const showDropdown = ref(false)
const contextmenuId = ref<number | null>(null)

const patentData = ref<CompanyPatent[]>([])
const checkedRowKeys = ref<Array<number | string>>([])
const isLoading = ref(false)

const queryForm = reactive({
  title: '',
})

const patentTypeColorMap: Record<string, 'success' | 'info' | 'warning'> = {
  发明专利: 'success',
  实用新型专利: 'info',
  外观设计专利: 'warning',
}

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
    GetCompanyPatentList()
  },
  onUpdatePageSize: (pageSize) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    GetCompanyPatentList()
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
      const row = patentData.value.find((item) => item.id === contextmenuId.value)
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

const columns = computed<DataTableColumns<CompanyPatent>>(() => {
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
      key: 'title',
      title: '专利名称',
      minWidth: 200,
      ellipsis: {
        tooltip: true,
      },
    },
    {
      key: 'image',
      title: '专利图片',
      width: 100,
      align: 'center',
      render: (row) =>
        row.image ? (
          <NImage
            src={import.meta.env.VITE_SITE_BASE_API + row.image}
            width={60}
            height={60}
            object-fit='cover'
            class='rounded'
            lazy
          />
        ) : (
          <span class='text-gray-400'>暂无</span>
        ),
    },
    {
      key: 'type',
      title: '专利类型',
      width: 130,
      align: 'center',
      render: (row) => (
        <NTag
          type={patentTypeColorMap[row.type] || 'default'}
          size='small'
          round
        >
          {row.type}
        </NTag>
      ),
    },
    {
      key: 'patent_no',
      title: '专利号',
      width: 180,
      align: 'center',
      render: (row) => (
        <NTag
          type='info'
          size='small'
          bordered={false}
        >
          {row.patent_no}
        </NTag>
      ),
    },
    {
      key: 'auth_date',
      title: '授权日期',
      width: 130,
      align: 'center',
      render: (row) =>
        row.auth_date ? (
          <NTag
            type='success'
            size='small'
            round
          >
            {dayjs(row.auth_date).format('YYYY-MM-DD')}
          </NTag>
        ) : (
          <span class='text-gray-400'>—</span>
        ),
    },
    {
      key: 'inventor',
      title: '发明人',
      width: 150,
      ellipsis: {
        tooltip: true,
      },
    },
    {
      key: 'summary',
      title: '摘要',
      minWidth: 200,
      ellipsis: {
        tooltip: true,
      },
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

function CellActions(row: CompanyPatent) {
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
          default: () => '确认删除该专利吗？',
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
      GetCompanyPatentList()
    }
  })
}

function ResetForm() {
  queryForm.title = ''
  pagination.page = 1
  GetCompanyPatentList()
}

async function MutateDeleteData(id: number) {
  isLoading.value = true
  try {
    await DeleteCompanyPatentApi([id])
    message.success('删除成功')
    checkedRowKeys.value = checkedRowKeys.value.filter((key) => key !== id)
    GetCompanyPatentList()
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
    await DeleteCompanyPatentApi(checkedRowKeys.value as number[])
    message.success(`成功删除 ${checkedRowKeys.value.length} 条记录`)
    checkedRowKeys.value = []
    GetCompanyPatentList()
  } catch (error) {
    message.error('批量删除失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function GetCompanyPatentList() {
  isLoading.value = true
  try {
    const result = await GetCompanyPatentListApi(
      pagination.page || 1,
      pagination.pageSize || 10,
      queryForm.title,
    )
    console.log('GetCompanyPatentListApi result:', result.data)
    patentData.value = result.data
    if (result.paginate) pagination.itemCount = result.paginate.total_count
  } catch (error) {
    message.error('获取列表失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  GetCompanyPatentList()
})
</script>

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
          label="专利名称"
          path="title"
        >
          <NInput
            v-model:value="queryForm.title"
            clearable
            placeholder="请输入专利名称"
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
            新增专利
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
      </NForm>

      <div class="flex min-h-0 flex-1 flex-col">
        <NDataTable
          class="flex-1"
          ref="dataTableRef"
          v-model:checked-row-keys="checkedRowKeys"
          :remote="true"
          :columns="columns"
          :data="patentData"
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

    <CompanyPatentModal
      ref="modalRef"
      @success="GetCompanyPatentList"
    />
  </ScrollContainer>
</template>

<style scoped lang="scss"></style>
