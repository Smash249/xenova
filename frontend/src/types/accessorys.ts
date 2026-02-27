import type { BaseResponse } from "./common"

interface AccessorySeries extends BaseResponse {
  id: number
  name: string
}

interface Accessory extends BaseResponse {
  id: number
  name: string
  seriesId: string
  image: string
  description: string
  cover: string
  previews: string[]
  price: number
}
interface AccessoryDetail extends BaseResponse {
  id: number
  name: string
  description: string
  cover: string
  previews: string[] | null
  price: number
  series_id: number
  series_name: string
}

export type { AccessorySeries, Accessory, AccessoryDetail }
