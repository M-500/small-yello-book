package types

import "time"

type UserForm struct {
	UserName   string    `json:"user_name"`
	NickName   string    `json:"nick_name"`
	Avatar     string    `json:"avatar"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Signature  string    `json:"signature"`
	Male       string    `json:"male"`
	BirthDay   time.Time `json:"birthDay"`
	Addr       string    `json:"addr"`
	Profession string    `json:"profession"`
	School     string    `json:"school"`
}

type UpdateUserForm struct {
	ID       int    `json:"id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	RoleIds  []int  `json:"roleIds"`
}

type EmailLoginForm struct {
	Email   string `json:"email" binding:"required,email"`
	VerCode string `json:"ver_code" binding:"required"`
}
