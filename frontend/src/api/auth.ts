import request from "@/utils/request"

import type { LoginParams, LoginResp, RegisterParams } from "@/types/auth"

export async function LoginApi(params: LoginParams) {
  try {
    const result = await request.post<LoginResp>("/public/login", params)
    if (!result || !result.success) throw new Error("登录失败")
    return result.data
  } catch (error) {
    throw error
  }
}

export async function RegisterApi(params: RegisterParams) {
  try {
    const result = await request.post<LoginResp>("/public/register", params)
    if (!result || !result.success) throw new Error("登录失败")
    return result.data
  } catch (error) {
    throw error
  }
}
