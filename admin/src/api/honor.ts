import request from '@/utils/request'

import type {
  CreateCompanyHonorParams,
  CreateCompanyPatentParams,
  CreateLoveActivityParams,
  UpdateCompanyHonorParams,
  UpdateCompanyPatentParams,
  UpdateLoveActivityParams,
} from '@/types/honor'

export async function GetCompanyHonorListApi(
  page: number = 1,
  page_size: number = 10,
  title: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (title) searchParams.append('title', title)
    const result = await request.get('/public/company_honor' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取公司荣誉列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取公司荣誉列表请求失败')
  }
}

export async function CreateCompanyHonorApi(params: CreateCompanyHonorParams) {
  try {
    const result = await request.post('/admin/company_honor', params)
    if (!result.success) throw new Error('创建公司荣誉失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建公司荣誉请求失败')
  }
}

export async function UpdateCompanyHonorApi(params: UpdateCompanyHonorParams) {
  try {
    const result = await request.put(`/admin/company_honor`, params)
    if (!result.success) throw new Error('更新公司荣誉失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新公司荣誉请求失败')
  }
}

export async function DeleteCompanyHonorApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/company_honor`, { data: { id_list } })
    if (!result.success) throw new Error('删除公司荣誉失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除公司荣誉请求失败')
  }
}

export async function GetCompanyPatentListApi(
  page: number = 1,
  page_size: number = 10,
  title: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (title) searchParams.append('title', title)
    const result = await request.get('/public/company_patnet' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取公司专利列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取公司专利列表请求失败')
  }
}

export async function CreateCompanyPatentApi(params: CreateCompanyPatentParams) {
  try {
    const result = await request.post('/admin/company_patnet', params)
    if (!result.success) throw new Error('创建公司专利失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建公司专利请求失败')
  }
}

export async function UpdateCompanyPatentApi(params: UpdateCompanyPatentParams) {
  try {
    const result = await request.put(`/admin/company_patnet`, params)
    if (!result.success) throw new Error('更新公司专利失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新公司专利请求失败')
  }
}

export async function DeleteCompanyPatentApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/company_patnet`, { data: { id_list } })
    if (!result.success) throw new Error('删除公司专利失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除公司专利请求失败')
  }
}

export async function GetLoveActivityListApi(
  page: number = 1,
  page_size: number = 10,
  title: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (title) searchParams.append('title', title)
    const result = await request.get('/public/love_activity' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取爱心活动列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取爱心活动列表请求失败')
  }
}

export async function CreateLoveActivityApi(params: CreateLoveActivityParams) {
  try {
    const result = await request.post('/admin/love_activity', params)
    if (!result.success) throw new Error('创建爱心活动失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建爱心活动请求失败')
  }
}

export async function UpdateLoveActivityApi(params: UpdateLoveActivityParams) {
  try {
    const result = await request.put(`/admin/love_activity`, params)
    if (!result.success) throw new Error('更新爱心活动失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新爱心活动请求失败')
  }
}

export async function DeleteLoveActivityApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/love_activity`, { data: { id_list } })
    if (!result.success) throw new Error('删除爱心活动失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除爱心活动请求失败')
  }
}
