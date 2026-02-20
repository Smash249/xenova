import request from "@/utils/request"

export async function GetCompanyHonorApi(page: number) {
  try {
    const searchParams = new URLSearchParams()
    if (page) searchParams.append("page", String(page))
    const result = await request.get(
      `/public/company_honor?${searchParams.toString()}`
    )
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetLoveActivityListApi(page: number) {
  try {
    const searchParams = new URLSearchParams()
    if (page) searchParams.append("page", String(page))

    const result = await request.get(
      `/public/love_activity?${searchParams.toString()}`
    )
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function GetCompanyPatentApi(page: number) {
  try {
    const searchParams = new URLSearchParams()
    if (page) searchParams.append("page", String(page))
    const result = await request.get(
      `/public/company_patnet?${searchParams.toString()}`
    )
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}
