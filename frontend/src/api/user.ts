import request from "@/utils/request"

import type { ChangePasswordParams, UpdateUserInfoParams } from "@/types/user"

async function GetUserInfoApi() {
  try {
    const result = await request.get("/private/userInfo")
    if (!result || !result.success) throw new Error("获取用户信息失败")
    return result.data
  } catch (error) {
    throw error
  }
}

async function UpdateUserInfoApi(params: UpdateUserInfoParams) {
  try {
    const result = await request.put("/private/userInfo", params)
    if (!result || !result.success) throw new Error("更新用户信息失败")
    return result.data
  } catch (error) {
    throw error
  }
}

async function ChangePasswordApi(params: ChangePasswordParams) {
  try {
    const result = await request.put("/private/changePassword", params)
    if (!result || !result.success) throw new Error("修改密码失败")
    return result.data
  } catch (error) {
    throw error
  }
}

export { GetUserInfoApi, UpdateUserInfoApi, ChangePasswordApi }
