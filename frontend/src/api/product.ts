import request from "@/utils/request"

export async function GetProductSeriesApi() {
  try {
    const result = await request.get("/public/product_series")
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetProductListApi(
  page: number,
  name: string,
  series_id: number | null
) {
  const searchParams = new URLSearchParams()
  if (page) searchParams.append("page", String(page))
  if (name) searchParams.append("name", name)
  if (series_id) searchParams.append("series_id", series_id.toString())
  const query = searchParams.toString()
  try {
    const result = await request.get(`/public/products?${query}`)
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}
