import request from "@/utils/request"

export async function GetNewsSeriesApi() {
  try {
    const result = await request.get("/public/journalism_series")
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetNewsListApi(
  page: number,
  title: string,
  series_id: number | null
) {
  const searchParams = new URLSearchParams()
  if (page) searchParams.append("page", String(page))
  if (title) searchParams.append("name", title)
  if (series_id) searchParams.append("series_id", series_id.toString())
  const query = searchParams.toString()
  try {
    const result = await request.get(`/public/journalisms?${query}`)
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetNewsDetailApi(newsId: number) {
  try {
    const result = await request.get(`/private/journalisms/${newsId}`)
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}
