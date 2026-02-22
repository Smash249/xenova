import type { BaseResponse } from './common'

interface SoftWareSeries extends BaseResponse {
  id: number
  name: string
  description: string
}

type CreateSoftWareSeriesParams = Pick<SoftWareSeries, 'name' | 'description'>

type UpdateSoftWareSeriesParams = Pick<SoftWareSeries, 'id' | 'name' | 'description'>

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

type CreateSoftWareParams = Omit<SoftWare, 'id' | 'created_at' | 'updated_at'>

type UpdateSoftWareParams = CreateSoftWareParams & Pick<SoftWare, 'id'>

export type {
  SoftWareSeries,
  CreateSoftWareSeriesParams,
  UpdateSoftWareSeriesParams,
  SoftWare,
  CreateSoftWareParams,
  UpdateSoftWareParams,
}
