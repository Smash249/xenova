import type { BaseResponse } from "./common"

interface SoftWareSeries extends BaseResponse {
  id: number
  name: string
  description: string
}

interface SoftWare extends BaseResponse {
  id: number
  name: string
  description: string
  file_type: string
  file_size: string
  file_url: string
  is_hot: boolean
  series_id: number
}

export type { SoftWareSeries, SoftWare }
