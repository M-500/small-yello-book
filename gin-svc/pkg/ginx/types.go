package ginx

import (
	"github.com/golang-jwt/jwt/v5"
)

type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type UserClaims struct {
	Id  int64
	GId string
	jwt.RegisteredClaims
}
