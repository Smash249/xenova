import axios from "axios"
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios"

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

  private setupInterceptors() {
    // 请求拦截器
    this.instance.interceptors.request.use(
      (config) => {
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
        const message = error.response?.data?.message || error.message
        console.error("Request error:", message)
        return Promise.reject(error)
      }
    )
  }

  public static getInstance(config?: AxiosRequestConfig): Request {
    if (!axiosCache) axiosCache = new Request(config)
    return axiosCache
  }

  public get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.get(url, config)
  }

  public post<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<T> {
    return this.instance.post(url, data, config)
  }

  public put<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<T> {
    return this.instance.put(url, data, config)
  }

  public delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.delete(url, config)
  }
}

export default Request.getInstance()
