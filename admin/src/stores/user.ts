import { useStorage } from '@vueuse/core'
import { acceptHMRUpdate, defineStore, storeToRefs } from 'pinia'
import { computed } from 'vue'

import { LoginApi } from '@/api/user'
import router from '@/router'
import { resolveMenu, resolveRoute } from '@/router/helper'

import { pinia } from '.'

import type { MenuMixedOptions } from '@/router/interface'
import type { UserInfo } from '@/types/user'

export const useUserStore = defineStore('userStore', () => {
  const user = useStorage<UserInfo | null>('user', null)

  const token = useStorage<string | null>('token', null)

  const refreshToken = useStorage<string | null>('refreshToken', null)

  const menu = useStorage<MenuMixedOptions[]>('menu', [
    {
      path: 'dashboard',
      name: 'dashboard',
      icon: 'icon-[mage--dashboard-chart]',
      label: '仪表板',
      meta: {
        componentName: 'Dashboard',
        pinned: true,
        showTab: true,
      },
      component: 'dashboard/index',
    },
  ])

  async function Login(data: Parameters<typeof LoginApi>[0]) {
    try {
      const result = await LoginApi(data)
      console.log('登录成功:', result)
      token.value = result.accessToken
      refreshToken.value = result.refreshToken
      user.value = result.user
    } catch (error) {
      console.error('登录失败:', error)
      throw new Error('登录失败')
    }
  }

  function cleanup(redirectPath?: string) {
    router.replace({
      name: 'signIn',
      ...(redirectPath ? { query: { r: redirectPath } } : {}),
    })
    token.value = null
    user.value = null
    if (router.hasRoute('layout')) {
      router.removeRoute('layout')
    }
  }

  const userMenu = computed(() => {
    return resolveMenu(menu.value)
  })
  const userRoute = computed(() => {
    return resolveRoute(menu.value)
  })

  return {
    user,
    token,
    userMenu,
    userRoute,
    Login,
    cleanup,
  }
})

export function toRefsUserStore() {
  return {
    ...storeToRefs(useUserStore(pinia)),
  }
}

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
}
