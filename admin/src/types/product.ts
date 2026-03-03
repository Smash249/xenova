import type { BaseResponse } from './common'

interface ProductSeries extends BaseResponse {
  id: number
  name: string
  description: string
}

type CreateProductSeriesParams = Pick<ProductSeries, 'name' | 'description'>

type UpdateProductSeriesParams = Pick<ProductSeries, 'id' | 'name' | 'description'>

interface Product extends BaseResponse {
  id: number
  name: string
  series_id: number
  image: string
  description: string
  cover: string
  previews: string[]
  is_hot: boolean
}

type CreateProductParams = Omit<Product, 'id' | 'created_at' | 'updated_at'>

type UpdateProductParams = CreateProductParams & { id: number }

type UpdateProductHotParams = { is_hot: boolean } & { id: number }

export type {
  ProductSeries,
  CreateProductSeriesParams,
  UpdateProductSeriesParams,
  UpdateProductHotParams,
  Product,
  CreateProductParams,
  UpdateProductParams,
}
