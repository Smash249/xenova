interface User {
  id: number
  userName: string
  email: string
  phone: string
  updatedAt: string
  createdAt: string
}

type UpdateUserInfoParams = Pick<User, "userName" | "phone">

interface ChangePasswordParams {
  oldPassword: string
  newPassword: string
}
export type { User, UpdateUserInfoParams, ChangePasswordParams }
