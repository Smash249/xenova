import request from '@/utils/request'

import type {
  CreateNewsParams,
  CreateNewsSeriesParams,
  UpdateNewsParams,
  UpdateNewsSeriesParams,
} from '@/types/news'

export async function GetNewsSeriesListApi(
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
    const result = await request.get('/public/journalism_series' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取新闻系列列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取新闻系列列表请求失败')
  }
}
export async function CreateNewsSeriesApi(params: CreateNewsSeriesParams) {
  try {
    const result = await request.post('/admin/journalism_series', params)
    if (!result.success) throw new Error('创建新闻系列失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建新闻系列请求失败')
  }
}

export async function UpdateNewsSeriesApi(params: UpdateNewsSeriesParams) {
  try {
    const result = await request.put(`/admin/journalism_series`, params)
    if (!result.success) throw new Error('更新新闻系列失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新新闻系列请求失败')
  }
}

export async function DeleteNewsSeriesApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/journalism_series`, { data: { id_list } })
    if (!result.success) throw new Error('删除新闻系列失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除新闻系列请求失败')
  }
}

export async function GetNewsListApi(
  page: number = 1,
  page_size: number = 10,
  title: string = '',
  series_id: number | null = null,
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (title) searchParams.append('title', title)
    if (series_id !== null) searchParams.append('series_id', series_id.toString())
    const result = await request.get('/public/journalisms' + '?' + searchParams.toString())
    if (!result.success) throw new Error('获取新闻列表失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取新闻列表请求失败')
  }
}

export async function CreateNewsApi(params: CreateNewsParams) {
  try {
    const result = await request.post('/admin/journalisms', params)
    if (!result.success) throw new Error('创建新闻失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('创建新闻请求失败')
  }
}

export async function UpdateNewsApi(params: UpdateNewsParams) {
  try {
    const result = await request.put(`/admin/journalisms`, params)
    if (!result.success) throw new Error('更新新闻失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('更新新闻请求失败')
  }
}

export async function DeleteNewsApi(id_list: number[]) {
  try {
    const result = await request.delete(`/admin/journalisms`, { data: { id_list } })
    if (!result.success) throw new Error('删除新闻失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('删除新闻请求失败')
  }
}
