import request from '@/utils/request'

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

export async function CreateProductSeriesApi() {
  try {
    const result = await request.post('/admin/product_series')
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
export async function UpdateProductSeriesApi() {
  try {
    const result = await request.put('/admin/product_series')
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

export async function GetProductListApi(page: number, name: string, series_id: number | null) {
  const searchParams = new URLSearchParams()
  if (page) searchParams.append('page', String(page))
  if (name) searchParams.append('name', name)
  if (series_id) searchParams.append('series_id', series_id.toString())
  const query = searchParams.toString()
  try {
    const result = await request.get(`/public/products?${query}`)
    if (!result || !result.success) throw new Error('请求出错')
    return result.data
  } catch (error) {
    throw error
  }
}
