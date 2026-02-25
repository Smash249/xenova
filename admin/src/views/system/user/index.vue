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
  NTag,
} from 'naive-ui'
import { reactive, ref, useTemplateRef, computed, onMounted } from 'vue'

import { GetSystemUserListApi, BanedUserByIdApi } from '@/api/system'
import { ScrollContainer } from '@/components'

import type { UserInfo } from '@/types/user'
import type { DataTableColumns, PaginationProps, FormInst } from 'naive-ui'

defineOptions({
  name: 'UserAdmin',
})

const message = useMessage()

const formRef = useTemplateRef<FormInst>('formRef')

const userData = ref<UserInfo[]>([])
const checkedRowKeys = ref<Array<number | string>>([])
const isLoading = ref(false)

const queryForm = reactive({
  email: '',
  userName: '',
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
    GetUserList()
  },
  onUpdatePageSize: (pageSize) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    GetUserList()
  },
})

const hasChecked = computed(() => checkedRowKeys.value.length > 0)

const columns = computed<DataTableColumns<UserInfo>>(() => {
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
      key: 'userName',
      title: '用户名',
      width: 150,
      ellipsis: {
        tooltip: true,
      },
    },
    {
      key: 'email',
      title: '邮箱',
      width: 200,
      ellipsis: {
        tooltip: true,
      },
    },
    {
      key: 'role',
      title: '角色',
      width: 120,
      align: 'center',
      render: (row) => {
        const roleMap: Record<string, { label: string; type: 'success' | 'info' | 'warning' }> = {
          admin: { label: '管理员', type: 'success' },
          user: { label: '普通用户', type: 'info' },
        }
        const role = roleMap[row.role] || { label: row.role, type: 'warning' as const }
        return (
          <NTag
            type={role.type}
            size='small'
            round
          >
            {role.label}
          </NTag>
        )
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
      width: 120,
      align: 'center',
      fixed: 'right',
      render: (row) => CellActions(row),
    },
  ]
})

function CellActions(row: UserInfo) {
  return (
    <div class='flex justify-center gap-2'>
      <NPopconfirm
        positiveText='确定'
        negativeText='取消'
        onPositiveClick={() => MutateBanUser(row.id)}
      >
        {{
          default: () => `确认封禁用户「${row.userName}」吗？`,
          trigger: () => (
            <NButton
              secondary
              type='error'
              size='small'
            >
              封禁
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
      GetUserList()
    }
  })
}

function ResetForm() {
  queryForm.email = ''
  pagination.page = 1
  GetUserList()
}

async function MutateBanUser(id: number) {
  isLoading.value = true
  try {
    await BanedUserByIdApi(id)
    message.success('封禁成功')
    GetUserList()
  } catch (error) {
    message.error('封禁失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function MutateBatchBanUsers() {
  if (!hasChecked.value) return
  isLoading.value = true
  try {
    const promises = (checkedRowKeys.value as number[]).map((id) => BanedUserByIdApi(id))
    await Promise.all(promises)
    message.success(`成功封禁 ${checkedRowKeys.value.length} 个用户`)
    checkedRowKeys.value = []
    GetUserList()
  } catch (error) {
    message.error('批量封禁失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

async function GetUserList() {
  isLoading.value = true
  try {
    const result = await GetSystemUserListApi(
      pagination.page || 1,
      pagination.pageSize || 10,
      queryForm.email,
    )
    console.log('GetSystemUserListApi result:', result.data)
    userData.value = result.data
    if (result.paginate) pagination.itemCount = result.paginate.total_count
  } catch (error) {
    message.error('获取用户列表失败')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  GetUserList()
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
          label="邮箱"
          path="email"
        >
          <NInput
            v-model:value="queryForm.email"
            clearable
            placeholder="请输入用户邮箱"
          />
        </NFormItem>
        <NFormItem
          label="用户名"
          path="userName"
        >
          <NInput
            v-model:value="queryForm.userName"
            clearable
            placeholder="请输入用户名"
          />
        </NFormItem>
        <div class="flex gap-2">
          <NPopconfirm
            positive-text="确定"
            negative-text="取消"
            @positive-click="MutateBatchBanUsers"
          >
            <template #default> 确认封禁选中的 {{ checkedRowKeys.length }} 个用户吗？ </template>
            <template #trigger>
              <NButton
                type="error"
                :disabled="!hasChecked"
              >
                <template #icon>
                  <span class="iconify ph--prohibit" />
                </template>
                批量封禁 ({{ checkedRowKeys.length }})
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
          :data="userData"
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
  </ScrollContainer>
</template>

<style scoped lang="scss"></style>
