export interface emialForm {
  email: string
  type_code: number
}

export interface loginForm {
  email: string
  ver_code: string
}

export interface loginResponse {
  token: string
}
