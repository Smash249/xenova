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
  seriesId: string
  image: string
  description: string
  cover: string
  previews: string[]
  price: number
}

type CreateProductParams = Omit<Product, 'id'>

type UpdateProductParams = Product

export type {
  ProductSeries,
  CreateProductSeriesParams,
  UpdateProductSeriesParams,
  Product,
  CreateProductParams,
  UpdateProductParams,
}
