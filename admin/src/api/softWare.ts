import request from '@/utils/request'

import type { CreateSoftWareParams, UpdateSoftWareParams } from '@/types/softWare'

export async function GetSoftWareSeriesListApi(
  page: number = 1,
  page_size: number = 10,
  name: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (name) searchParams.append('name', name)
    const result = await request.get('/public/software_series' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取软件系列列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取软件系列列表请求失败')
  }
}

export async function CreateSoftWareSeriesApi(params: CreateSoftWareParams) {
  try {
    const result = await request.post('/admin/software_series', params)
    if (!result.success) throw new Error('创建软件系列失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建软件系列请求失败')
  }
}

export async function UpdateSoftWareSeriesApi(params: UpdateSoftWareParams) {
  try {
    const result = await request.put(`/admin/software_series`, params)
    if (!result.success) throw new Error('更新软件系列失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新软件系列请求失败')
  }
}

export async function DeleteSoftWareSeriesApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/software_series`, { data: { id_list } })
    if (!result.success) throw new Error('删除软件系列失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除软件系列请求失败')
  }
}

export async function GetSoftWareListApi(
  page: number = 1,
  page_size: number = 10,
  name: string = '',
  series_id?: number,
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (name) searchParams.append('name', name)
    if (series_id) searchParams.append('series_id', series_id.toString())
    const result = await request.get('/public/softwares' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取软件列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取软件列表请求失败')
  }
}

export async function CreateSoftWareApi(params: CreateSoftWareParams) {
  try {
    const result = await request.post('/admin/softwares', params)
    if (!result.success) throw new Error('创建软件失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建软件请求失败')
  }
}

export async function UpdateSoftWareApi(params: UpdateSoftWareParams) {
  try {
    const result = await request.put(`/admin/softwares`, params)
    if (!result.success) throw new Error('更新软件失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新软件请求失败')
  }
}

export async function DeleteSoftWareApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/softwares`, { data: { id_list } })
    if (!result.success) throw new Error('删除软件失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除软件请求失败')
  }
}
