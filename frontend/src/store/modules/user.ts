import { defineStore } from "pinia"
import { ref } from "vue"
import request from "@/utils/request"

interface User {
  ID: number
  UserName: string
  Email: string
}

interface LoginResponse {
  user: User
  accessToken: string
  reFreshToken: string
}

const userStore = defineStore(
  "user",
  () => {
    const user = ref<User | null>(null)
    const accessToken = ref<string>("")
    const refreshToken = ref<string>("")

    const login = async (email: string, password: string) => {
      try {
        const res = await request.post<LoginResponse>("/api/login", {
          email,
          password,
        })
        user.value = res.user
        accessToken.value = res.accessToken
        refreshToken.value = res.reFreshToken
        return res
      } catch (error) {
        throw error
      }
    }

    const register = async (
      UserName: string,
      email: string,
      code: string,
      password: string
    ) => {
      try {
        await request.post("/api/register", {
          UserName,
          email,
          code,
          password,
        })
      } catch (error) {
        throw error
      }
    }

    const logout = () => {
      user.value = null
      accessToken.value = ""
      refreshToken.value = ""
    }

    return {
      user,
      accessToken,
      refreshToken,
      login,
      register,
      logout,
    }
  },
  {
    persist: true,
  }
)

export default userStore
