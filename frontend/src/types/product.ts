import type { BaseResponse } from "./common"

interface Product {
  id: number
  name: string
  ProductSeries: string
  image: string
  description: string
}

interface ProductSeries extends BaseResponse {
  id: number
  name: string
}

export type { ProductSeries, Product }
