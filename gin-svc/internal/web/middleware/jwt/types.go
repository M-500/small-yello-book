package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -destination=./mocks/mock_handler.go -package=mocks gin-svc/internal/web/middleware/jwt Handler
type Handler interface {
	ClearToken(ctx *gin.Context) error
	ExtractToken(ctx *gin.Context) string
	GenJWTToken(uid int64) (string, error)
	SetLoginToken(ctx *gin.Context, uid int64) error
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	Uid int64
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid int64
	//UserAgent string
}

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
var RCJWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgA")
