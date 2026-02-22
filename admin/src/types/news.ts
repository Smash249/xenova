import type { BaseResponse } from './common'

interface NewsSeries extends BaseResponse {
  id: number
  name: string
  description: string
}
type CreateNewsSeriesParams = Pick<NewsSeries, 'name' | 'description'>

type UpdateNewsSeriesParams = CreateNewsSeriesParams & { id: number }

interface News extends BaseResponse {
  id: number
  title: string
  content: string
  view_count: number
  series_id: number
}

type CreateNewsParams = Pick<News, 'title' | 'content' | 'series_id'>

type UpdateNewsParams = CreateNewsParams & { id: number }

export type {
  NewsSeries,
  CreateNewsSeriesParams,
  UpdateNewsSeriesParams,
  News,
  CreateNewsParams,
  UpdateNewsParams,
}
