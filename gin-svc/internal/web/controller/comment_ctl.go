package controller

import (
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type CommentCtl struct {
}

func NewCommentCtl() *CommentCtl {
	return &CommentCtl{}
}

func (c *CommentCtl) AddCommentCtl(ctx *gin.Context, req types.AddCommentForm, uc jwt.UserClaims) (result ginx.JsonResult, err error) {

	return ginx.Success(), nil
}
