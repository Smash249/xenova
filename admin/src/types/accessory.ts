import type { BaseResponse } from './common'

interface AccessorySeries extends BaseResponse {
  id: number
  name: string
  description: string
}

type CreateAccessorySeriesParams = Pick<AccessorySeries, 'name' | 'description'>

type UpdateAccessorySeriesParams = Pick<AccessorySeries, 'id' | 'name' | 'description'>

interface Accessory extends BaseResponse {
  id: number
  name: string
  series_id: number
  image: string
  description: string
  cover: string
  previews: string[]
  price: number
}

type CreateAccessoryParams = Omit<Accessory, 'id' | 'created_at' | 'updated_at'>

type UpdateAccessoryParams = CreateAccessoryParams & { id: number }

export type {
  AccessorySeries,
  CreateAccessorySeriesParams,
  UpdateAccessorySeriesParams,
  Accessory,
  CreateAccessoryParams,
  UpdateAccessoryParams,
}
