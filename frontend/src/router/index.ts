import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    name: "login",
    component: () => import("@/pages/auth/login.vue"),
  },
  {
    path: "/register",
    name: "register",
    component: () => import("@/pages/auth/register.vue"),
  },
  {
    path: "/",
    component: () => import("@/layout/index.vue"),
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: () => import("@/pages/dashboard/index.vue"),
      },
      {
        path: "product",
        name: "product",
        component: () => import("@/pages/product/index.vue"),
      },
      {
        path: "news",
        name: "news",
        component: () => import("@/pages/news/index.vue"),
      },
      {
        path: "download",
        name: "download",
        component: () => import("@/pages/download/index.vue"),
      },
      {
        path: "honor",
        name: "honor",
        component: () => import("@/pages/honor/index.vue"),
      },
      {
        path: "contact",
        name: "contact",
        component: () => import("@/pages/contact/index.vue"),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior(to) {
    if (to.hash) {
      return {
        el: to.hash,
        behavior: "smooth",
      }
    }
    return { top: 0 }
  },
  routes,
})

// 路由拦截器
router.beforeEach(() => {})

export default router
