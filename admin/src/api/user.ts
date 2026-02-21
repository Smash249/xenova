import request from '@/utils/request'

export async function LoginApi(data: { email: string; password: string }) {
  try {
    const result = await request.post('/public/login', data)
    if (!result.success) throw new Error('登录失败')
    return result.data
  } catch (error) {
    throw error instanceof Error ? error : new Error('登录请求失败')
  }
}
