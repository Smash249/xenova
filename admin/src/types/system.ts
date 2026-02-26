import type { BaseResponse } from './common'

interface SystemMessage extends BaseResponse {
  id: number
  name: string
  phone: string
  content: string
  company: string
  email: string
}

interface UpdateSystemMessageReq {
  id: string
  name: string
  phone: string
  content: string
  company: string
  email: string
}

interface SystemJobPosition extends BaseResponse {
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
  status: number // 1: 招聘中, 0: 已停止/下线
  sort: number
}

interface CreateJobPositionReq {
  title: string
  department: string
  location: string
  headcount: string
  experience: string
  education: string
  salaryRange: string
  responsibilities: string
  requirements: string
}

interface UpdateJobPositionReq {
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

export type {
  SystemMessage,
  UpdateSystemMessageReq,
  SystemJobPosition,
  CreateJobPositionReq,
  UpdateJobPositionReq,
}
