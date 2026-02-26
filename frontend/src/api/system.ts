import request from "@/utils/request"

import type { CreateSystemMessageReq } from "@/types/system"

export async function GetSystemJobPositionList(page: number, pageSize: number) {
  try {
    const searchParams = new URLSearchParams()
    if (page) searchParams.append("page", String(page))
    if (pageSize) searchParams.append("page_size", String(pageSize))
    const result = await request.get(
      `/public/job_postions?${searchParams.toString()}`
    )
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function CreateSystemMessage(params: CreateSystemMessageReq) {
  try {
    const result = await request.post(`/public/messages`, params)
    if (!result || !result.success) throw new Error("请求出错")
    return result.data
  } catch (error) {
    throw error
  }
}
