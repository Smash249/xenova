import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw[] = [
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
        path: "about",
        name: "about",
        component: () => import("@/pages/about/index.vue"),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_BASEPATH || ""),
  routes,
})

export default router
