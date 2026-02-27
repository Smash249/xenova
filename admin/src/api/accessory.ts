import request from '@/utils/request'

import type {
  CreateAccessoryParams,
  CreateAccessorySeriesParams,
  UpdateAccessoryParams,
  UpdateAccessorySeriesParams,
} from '@/types/accessory'

export async function GetAccessorySeriesListApi(page: number, page_size: number, name: string) {
  try {
    const searchParams = new URLSearchParams()
    if (page) searchParams.append('page', String(page))
    if (page_size) searchParams.append('page_size', String(page_size))
    if (name) searchParams.append('name', name)
    const query = searchParams.toString()
    const result = await request.get('/public/accessory_series?' + query)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function CreateAccessorySeriesApi(params: CreateAccessorySeriesParams) {
  try {
    const result = await request.post('/admin/accessory_series', params)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
export async function UpdateAccessorySeriesApi(parmas: UpdateAccessorySeriesParams) {
  try {
    const result = await request.put('/admin/accessory_series', parmas)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
export async function DeleteAccessorySeriesApi(id_list: number[]) {
  try {
    const result = await request.delete('/admin/accessory_series', {
      data: { id_list },
    })
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetAccessoryListApi(page: number, pageSize: number, name: string) {
  const searchParams = new URLSearchParams()
  if (page) searchParams.append('page', String(page))
  if (name) searchParams.append('name', name)
  if (pageSize) searchParams.append('page_size', pageSize.toString())
  const query = searchParams.toString()
  try {
    const result = await request.get(`/public/accessorys?${query}`)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function CreateAccessoryApi(params: CreateAccessoryParams) {
  try {
    const result = await request.post('/admin/accessorys', params)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function UpdateAccessoryApi(params: UpdateAccessoryParams) {
  try {
    const result = await request.put('/admin/accessorys', params)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function DeleteAccessoryApi(id_list: number[]) {
  try {
    const result = await request.delete('/admin/accessorys', {
      data: { id_list },
    })
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
