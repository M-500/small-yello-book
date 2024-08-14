package types

type EmailForm struct {
	Email    string `json:"email"`     // 邮箱
	TypeCode int    `json:"type_code"` // 用来表示发送邮件的类型 1:注册 2:找回密码
}
