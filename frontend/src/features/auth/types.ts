export interface User {
  id: number
  organization_id: number
  email: string
  first_name: string
  roles: string | string[]
}

export interface DataAuth {
  user: User
  token: string
}

export interface SessionResponse {
  data: DataAuth
  success: boolean
  message?: string
}

export interface LoginPayload {
  email: string
  password: string
  remember?: boolean
}
