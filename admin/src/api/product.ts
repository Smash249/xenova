import request from '@/utils/request'

import type {
  CreateProductParams,
  CreateProductSeriesParams,
  UpdateProductParams,
  UpdateProductSeriesParams,
} from '@/types/product'

export async function GetProductSeriesListApi(page: number, page_size: number, name: string) {
  try {
    const searchParams = new URLSearchParams()
    if (page) searchParams.append('page', String(page))
    if (page_size) searchParams.append('page_size', String(page_size))
    if (name) searchParams.append('name', name)
    const query = searchParams.toString()
    const result = await request.get('/public/product_series?' + query)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function CreateProductSeriesApi(params: CreateProductSeriesParams) {
  try {
    const result = await request.post('/admin/product_series', params)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
export async function UpdateProductSeriesApi(parmas: UpdateProductSeriesParams) {
  try {
    const result = await request.put('/admin/product_series', parmas)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
export async function DeleteProductSeriesApi(id_list: number[]) {
  try {
    const result = await request.delete('/admin/product_series', {
      data: { id_list },
    })
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetProductListApi(page: number, pageSize: number, name: string) {
  const searchParams = new URLSearchParams()
  if (page) searchParams.append('page', String(page))
  if (name) searchParams.append('name', name)
  if (pageSize) searchParams.append('page_size', pageSize.toString())
  const query = searchParams.toString()
  try {
    const result = await request.get(`/public/products?${query}`)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function CreateProductApi(params: CreateProductParams) {
  try {
    const result = await request.post('/admin/products', params)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function UpdateProductApi(params: UpdateProductParams) {
  try {
    const result = await request.put('/admin/products', params)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}

export async function DeleteProductApi(id_list: number[]) {
  try {
    const result = await request.delete('/admin/products', {
      data: { id_list },
    })
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
