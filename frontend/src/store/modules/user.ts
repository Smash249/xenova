import { LoginApi, RegisterApi } from "@/api/auth"
import { defineStore } from "pinia"
import { computed, ref } from "vue"
import { useRouter } from "vue-router"

import type { LoginParams, RegisterParams } from "@/types/auth"
import type { User } from "@/types/user"

const userStore = defineStore(
  "user",
  () => {
    const router = useRouter()
    const user = ref<User | null>(null)
    const accessToken = ref<string>("")
    const refreshToken = ref<string>("")
    const isLogin = computed(() => {
      return user.value ? true : false
    })
    const email = computed(() => {
      return user.value?.email ?? ""
    })
    const userName = computed(() => {
      return user.value?.userName ?? ""
    })

    async function Login(params: LoginParams) {
      try {
        const result = await LoginApi(params)
        user.value = result.user
        accessToken.value = result.accessToken
        refreshToken.value = result.reFreshToken
      } catch (error) {
        throw error
      }
    }

    async function Register(params: RegisterParams) {
      try {
        const result = await RegisterApi(params)
        console.log(result)
      } catch (error) {
        throw error
      }
    }

    function Logout() {
      user.value = null
      accessToken.value = ""
      refreshToken.value = ""
      localStorage.removeItem("user")
      router.replace("/login")
    }

    return {
      isLogin,
      user,
      email,
      userName,
      accessToken,
      refreshToken,
      Login,
      Register,
      Logout,
    }
  },
  {
    persist: true,
  }
)

export default userStore
