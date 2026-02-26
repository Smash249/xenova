import type { BaseResponse } from "./common"

interface JobPosition extends BaseResponse {
  id: number
  title: string
  department: string
  location: string
  headcount: string
  experience: string
  education: string
  salaryRange: string
  responsibilities: string
  requirements: string
  status: number
  sort: number
}

interface CreateSystemMessageReq {
  name: string
  content: string
  phone: string
  email: string
  company: string
}

export type { JobPosition, CreateSystemMessageReq }
