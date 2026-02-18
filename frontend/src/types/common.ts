type RequestResponse<T> = {
  data: T

  success: boolean
}

interface BaseResponse {
  createdAt: string
  updatedAt: string
}

interface Paginate {
  page_size: number
  page: number
  total_page: number
  total_count: number
}

export type { RequestResponse, BaseResponse, Paginate }
