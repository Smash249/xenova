import { useStorage } from '@vueuse/core'
import { acceptHMRUpdate, defineStore, storeToRefs } from 'pinia'
import { computed } from 'vue'

import { LoginApi } from '@/api/user'
import router from '@/router'
import { resolveMenu, resolveRoute } from '@/router/helper'

import { pinia } from '.'

import type { MenuMixedOptions } from '@/router/interface'
import type { UserInfo } from '@/types/user'

const DetaultUserInfo: UserInfo = {
  id: -1,
  userName: '',
  email: '',
  role: '',
  created_at: '',
  updated_at: '',
}

export const useUserStore = defineStore('userStore', () => {
  const user = useStorage<UserInfo>('user', DetaultUserInfo)

  const token = useStorage<string | null>('token', null)

  const refreshToken = useStorage<string | null>('refreshToken', null)

  const menu = useStorage<MenuMixedOptions[]>('menu', [
    {
      path: 'dashboard',
      name: 'dashboard',
      icon: 'icon-[lucide--layout-dashboard]',
      label: '仪表板',
      meta: {
        componentName: 'Dashboard',
        pinned: true,
        showTab: true,
      },
      component: 'dashboard/index',
    },
    {
      path: 'product',
      name: 'product',
      icon: 'icon-[lucide--package]',
      label: '产品管理',
      meta: {
        componentName: 'Product',
        pinned: true,
        showTab: true,
      },
      children: [
        {
          path: 'series',
          name: 'productSeries',
          icon: 'icon-[lucide--shapes]',
          label: '产品系列',
          meta: {
            componentName: 'ProductSeries',
            pinned: true,
            showTab: true,
          },
          component: 'product/series/index',
        },
        {
          path: 'list',
          name: 'productList',
          icon: 'icon-[lucide--list]',
          label: '产品列表',
          meta: {
            componentName: 'ProductList',
            pinned: true,
            showTab: true,
          },
          component: 'product/list/index',
        },
      ],
    },
    {
      path: 'news',
      name: 'news',
      icon: 'icon-[lucide--newspaper]',
      label: '新闻管理',
      meta: {
        componentName: 'News',
        pinned: true,
        showTab: true,
      },
      component: 'news/index',
      children: [
        {
          path: 'series',
          name: 'newsSeries',
          icon: 'icon-[lucide--shapes]',
          label: '新闻系列',
          meta: {
            componentName: 'NewsSeries',
            pinned: true,
            showTab: true,
          },
          component: 'news/series/index',
        },
        {
          path: 'list',
          name: 'newsList',
          icon: 'icon-[lucide--list]',
          label: '新闻列表',
          meta: {
            componentName: 'NewsList',
            pinned: true,
            showTab: true,
          },
          component: 'news/list/index',
        },
      ],
    },
    {
      path: 'softWare',
      name: 'softWare',
      icon: 'icon-[lucide--code]',
      label: '软件管理',
      meta: {
        componentName: 'Software',
        pinned: true,
        showTab: true,
      },
      component: 'softWare/index',
      children: [
        {
          path: 'series',
          name: 'softwareSeries',
          icon: 'icon-[lucide--shapes]',
          label: '软件系列',
          meta: {
            componentName: 'SoftwareSeries',
            pinned: true,
            showTab: true,
          },
          component: 'softWare/series/index',
        },
        {
          path: 'list',
          name: 'softwareList',
          icon: 'icon-[lucide--list]',
          label: '软件列表',
          meta: {
            componentName: 'SoftwareList',
            pinned: true,
            showTab: true,
          },
          component: 'softWare/list/index',
        },
      ],
    },
    {
      path: 'honor',
      name: 'honor',
      icon: 'icon-[lucide--award]',
      label: '荣誉管理',
      meta: {
        componentName: 'Honor',
        pinned: true,
        showTab: true,
      },
      children: [
        {
          path: 'companyHonor',
          name: 'companyHonor',
          icon: 'icon-[lucide--trophy]',
          label: '公司荣誉',
          meta: {
            componentName: 'CompanyHonor',
            pinned: true,
            showTab: true,
          },
          component: 'honor/companyHonor/index',
        },
        {
          path: 'companyPatent',
          name: 'companyPatent',
          icon: 'icon-[lucide--file-badge]',
          label: '公司专利',
          meta: {
            componentName: 'CompanyPatent',
            pinned: true,
            showTab: true,
          },
          component: 'honor/companyPatent/index',
        },
        {
          path: 'loveActivity',
          name: 'loveActivity',
          icon: 'icon-[lucide--heart-handshake]',
          label: '公益活动',
          meta: {
            componentName: 'LoveActivity',
            pinned: true,
            showTab: true,
          },
          component: 'honor/loveActivity/index',
        },
      ],
    },
  ])

  const userMenu = computed(() => {
    return resolveMenu(menu.value)
  })

  const userRoute = computed(() => {
    return resolveRoute(menu.value)
  })

  function cleanup(redirectPath?: string) {
    router.replace({
      name: 'login',
      ...(redirectPath ? { query: { r: redirectPath } } : {}),
    })
    token.value = null
    user.value = null
    if (router.hasRoute('layout')) {
      router.removeRoute('layout')
    }
  }

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
