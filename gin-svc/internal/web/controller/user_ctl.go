package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userSvc service.UserSvc
}

func NewUserController() BaseController {
	return &userController{}
}

func (ctl *userController) Upsert(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = ctl.userSvc.UpsertUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "创建/更新失败"), err
	}
	return ginx.Success(), nil
}

func (ctl *userController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/users", ginx.WrapJsonBody[types.UserForm](ctl.Upsert))
}
