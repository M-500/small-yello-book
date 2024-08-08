package types

import "time"

type UserForm struct {
	UserName   string    `json:"user_name"`
	NickName   string    `json:"nick_name"`
	Avatar     string    `json:"avatar"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Signature  string    `json:"signature"`
	Male       string    `json:"male"`
	BirthDay   time.Time `json:"birthDay"`
	Addr       string    `json:"addr"`
	Profession string    `json:"profession"`
	School     string    `json:"school"`
}
