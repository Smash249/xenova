import type { User } from "./user"

interface LoginParams {
  email: string
  password: string
}

interface LoginResp {
  user: User
  accessToken: string
  reFreshToken: string
  success: boolean
}

interface RegisterParams {
  userName: string
  code: string
  email: string
  password: string
}

export type { LoginParams, RegisterParams, LoginResp }
