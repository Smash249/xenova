import axios from "axios"
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios"
import { ElMessage } from "element-plus"

import type { RequestResponse } from "@/types/common"

let axiosCache: Request | null = null

class Request {
  private instance: AxiosInstance
  private constructor(config?: AxiosRequestConfig) {
    this.instance = axios.create({
      baseURL: config?.baseURL || "",
      timeout: config?.timeout || 10000,
      ...config,
    })
    this.setupInterceptors()
  }

  private handelErrorCode(code: number) {
    switch (code) {
      case 401:
        localStorage.removeItem("user")

        ElMessage.error({
          message: "登录失效，请重新登录",
        })
        window.location.href = "/login"
        break
      case 403:
        console.error("403 Forbidden")
        break
      case 404:
        console.error("404 Not Found")
        break
      case 500:
        console.error("500 Internal Server Error")
        break
      default:
        console.error("Unknown Error")
        break
    }
  }

  private setupInterceptors() {
    // 请求拦截器
    this.instance.interceptors.request.use(
      (config) => {
        const userStore = localStorage.getItem("user")
        if (userStore) {
          try {
            const user = JSON.parse(userStore)
            if (user.accessToken) {
              config.headers.Authorization = `${user.accessToken}`
            }
          } catch (e) {
            console.error("Failed to parse user store:", e)
          }
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      }
    )

    // 响应拦截器
    this.instance.interceptors.response.use(
      (response: AxiosResponse) => {
        return response.data
      },
      (error) => {
        const status = error.response?.status
        this.handelErrorCode(status)
        return Promise.reject(error)
      }
    )
  }

  public static getInstance(config?: AxiosRequestConfig): Request {
    if (!axiosCache) axiosCache = new Request(config)
    return axiosCache
  }

  public get<T = any>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<RequestResponse<T>> {
    return this.instance.get(url, config)
  }

  public post<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<RequestResponse<T>> {
    return this.instance.post(url, data, config)
  }

  public put<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<RequestResponse<T>> {
    return this.instance.put(url, data, config)
  }

  public delete<T = any>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<RequestResponse<T>> {
    return this.instance.delete(url, config)
  }
}

export default Request.getInstance({
  baseURL: import.meta.env.VITE_PROXY_PATH,
})
