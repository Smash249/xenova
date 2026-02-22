interface CompanyHonor {
  id: number
  title: string
  issue_date: string
  cert_no: string
  issuer: string
  image: string
  description: string
}

interface CompanyPatent {
  id: number
  title: string
  type: string
  patent_no: string
  auth_date: string
  inventor: string
  image: string
  summary: string
}

interface LoveActivity {
  id: number
  title: string
  activity_date: string
  location: string
  participants: number
  cover: string
  summary: string
  content: string
}

export type { CompanyHonor, CompanyPatent, LoveActivity }
