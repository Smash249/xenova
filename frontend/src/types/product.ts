import type { BaseResponse } from "./common"

interface ProductSeries extends BaseResponse {
  id: number
  name: string
}

interface Product extends BaseResponse {
  id: number
  name: string
  seriesId: string
  image: string
  description: string
  cover: string
  previews: string[]
  price: number
}
interface ProductDetail extends BaseResponse {
  id: number
  name: string
  description: string
  cover: string
  previews: string[] | null
  price: number
  series_id: number
  series_name: string
}

export type { ProductSeries, Product, ProductDetail }
