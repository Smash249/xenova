import request from '@/utils/request'

import type {
  CreateJobPositionReq,
  UpdateJobPositionReq,
  UpdateSystemMessageReq,
} from '@/types/system'

export async function GetSystemUserListApi(
  page: number = 1,
  page_size: number = 10,
  email: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (email) searchParams.append('email', email)
    const result = await request.get('/admin/users' + '?' + searchParams.toString())
    if (!result.success) throw new Error('登录失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('登录请求失败')
  }
}

export async function BanedUserByIdApi(userId: number) {
  try {
    const result = await request.post(`/admin/users/${userId}`)
    if (!result.success) throw new Error('封禁用户失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('封禁用户请求失败')
  }
}

export async function GetSystemMessageListApi(
  page: number = 1,
  page_size: number = 10,
  name: string = '',
  phone: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (name) searchParams.append('name', name)
    if (phone) searchParams.append('phone', phone)
    const result = await request.get('/admin/messages' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取菜单列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取菜单列表请求失败')
  }
}

export async function UpdateSystemMessageApi(params: UpdateSystemMessageReq) {
  try {
    const result = await request.put('/admin/messages', params)
    if (!result.success) throw new Error('更新菜单失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新菜单请求失败')
  }
}

export async function DeleteSystemMessagedApi(id_list: number[]) {
  try {
    const result = await request.delete('/admin/messages/', {
      data: { id_list },
    })
    if (!result.success) throw new Error('删除菜单失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除菜单请求失败')
  }
}

export async function GetSystemJobPositionListApi(
  page: number = 1,
  page_size: number = 10,
  title: string = '',
  department: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (title) searchParams.append('title', title)
    if (department) searchParams.append('department', department)
    const result = await request.get('/public/job_positions' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取职位列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取职位列表请求失败')
  }
}

export async function CreateSystemJobPositionApi(params: CreateJobPositionReq) {
  try {
    const result = await request.post('/admin/job_positions', params)
    if (!result.success) throw new Error('创建职位失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建职位请求失败')
  }
}

export async function UpdateSystemJobPositionApi(params: UpdateJobPositionReq) {
  try {
    const result = await request.put('/admin/job_positions', params)
    if (!result.success) throw new Error('更新职位失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新职位请求失败')
  }
}

export async function DeleteSystemJobPositionApi(id_list: number[]) {
  try {
    const result = await request.delete('/admin/job_positions/', {
      data: { id_list },
    })
    if (!result.success) throw new Error('删除职位失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除职位请求失败')
  }
}
