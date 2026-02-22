import type { BaseResponse } from './common'

interface CompanyHonor extends BaseResponse {
  id: number
  title: string
  issue_date: string
  cert_no: string
  issuer: string
  image: string
  description: string
}

type CreateCompanyHonorParams = Omit<CompanyHonor, 'id' | 'created_at' | 'updated_at'>

type UpdateCompanyHonorParams = CreateCompanyHonorParams & { id: number }

interface CompanyPatent extends BaseResponse {
  id: number
  title: string
  type: string
  patent_no: string
  auth_date: string
  inventor: string
  image: string
  summary: string
}

type CreateCompanyPatentParams = Omit<CompanyPatent, 'id' | 'created_at' | 'updated_at'>

type UpdateCompanyPatentParams = CreateCompanyPatentParams & { id: number }

interface LoveActivity extends BaseResponse {
  id: number
  title: string
  activity_date: string
  location: string
  participants: number
  cover: string
  summary: string
  content: string
}

type CreateLoveActivityParams = Omit<LoveActivity, 'id' | 'created_at' | 'updated_at'>

type UpdateLoveActivityParams = CreateLoveActivityParams & { id: number }

export type {
  CompanyHonor,
  CreateCompanyHonorParams,
  UpdateCompanyHonorParams,
  CompanyPatent,
  CreateCompanyPatentParams,
  UpdateCompanyPatentParams,
  LoveActivity,
  CreateLoveActivityParams,
  UpdateLoveActivityParams,
}
