interface CompanyHonor {
  id: number
  title: string
  date: string
  certNo: string
  issuer: string
  image: string
  description: string
}

interface CompanyPatent {
  id: number
  title: string
  type: string
  patentNo: string
  date: string
  inventor: string
  image: string
  summary: string
}

interface LoveActivity {
  id: number
  title: string
  date: string
  location: string
  participants: number
  cover: string
  summary: string
  content: string
}

export type { CompanyHonor, CompanyPatent, LoveActivity }
