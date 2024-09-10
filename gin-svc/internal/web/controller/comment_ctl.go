package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
	"gin-svc/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CommentCtl struct {
	commentSvc service.CommentSvc
}

func NewCommentCtl(commSvc service.CommentSvc) *CommentCtl {
	return &CommentCtl{
		commentSvc: commSvc,
	}
}

func (c *CommentCtl) AddCommentCtl(ctx *gin.Context, req types.AddCommentForm, uc jwt.UserClaims) (result ginx.JsonResult, err error) {
	err = c.commentSvc.CreateComment(ctx, uc.Uid, req)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.Success(), nil
}

func (c *CommentCtl) CommentListCtl(ctx *gin.Context) (ginx.JsonResult, error) {
	resourceId := ctx.Param("resource_id")
	if utils.IsBlank(resourceId) {
		return ginx.Error(400, "resource_id is required"), nil
	}
	comments, err := c.commentSvc.CommentList(ctx, resourceId)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.SuccessJson(comments), nil
}
