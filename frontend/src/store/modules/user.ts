import { LoginApi } from "@/api/auth"
import { UpdateUserInfoApi } from "@/api/user"
import { defineStore } from "pinia"
import { computed, ref } from "vue"
import { useRouter } from "vue-router"

import type { LoginParams } from "@/types/auth"
import type { UpdateUserInfoParams, User } from "@/types/user"

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

    async function UpdateUserInfo(params: UpdateUserInfoParams) {
      try {
        const result = await UpdateUserInfoApi(params)
        user.value = result.data
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

      Logout,
      UpdateUserInfo,
    }
  },
  {
    persist: true,
  }
)

export default userStore
