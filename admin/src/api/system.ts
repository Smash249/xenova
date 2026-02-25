import request from '@/utils/request'

export async function GetSystemUserListApi(
  page: number = 1,
  page_size: number = 10,
  email: string = '',
) {
  try {
    const searchParams = new URLSearchParams({
      page: page.toString(),
      page_size: page_size.toString(),
    })
    if (email) searchParams.append('email', email)
    const result = await request.get('/admin/users' + '?' + searchParams.toString())
    if (!result.success) throw new Error('登录失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('登录请求失败')
  }
}

export async function BanedUserByIdApi(userId: number) {
  try {
    const result = await request.post(`/admin/users/${userId}`)
    if (!result.success) throw new Error('封禁用户失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('封禁用户请求失败')
  }
}
