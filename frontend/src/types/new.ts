import type { BaseResponse } from "./common"

interface NewsSeries extends BaseResponse {
  id: number
  name: string
  description: string
}

interface News extends BaseResponse {
  id: number
  title: string
  content: string
  view_count: number
  series_id: number
}

interface NewsDetail extends BaseResponse {
  id: number
  series_id: number
  series_name: string
  title: string
  description: string
  content: string
  view_count: number
}

export type { NewsSeries, News, NewsDetail }
