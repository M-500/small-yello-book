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

func NewUserController(userSvc service.UserSvc) BaseController {
	return &userController{
		userSvc: userSvc,
	}
}

// UpdateUserInfo godoc
// @Summary Update user info
// @Description Update user info
// @Tags User Model
// @Accept json
// @Produce json
// @Param user body types.UserForm true "user"
// @Success 200 {object} map[string]any
// @Router /api/v1/users [put]
func (ctl *userController) UpdateUserInfo(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = ctl.userSvc.RegisterUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "更新失败"), err
	}
	return ginx.Success(), nil
}

// DeleteUserCtl godoc
// @Summary Delete user
// @Description Delete user
// @Tags User Model
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]any
// @Router /api/v1/users/{id} [delete]
func (ctl *userController) DeleteUserCtl(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = ctl.userSvc.RegisterUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "更新失败"), err
	}
	return ginx.Success(), nil
}

func (ctl *userController) RegisterRoute(group *gin.RouterGroup) {
	group.PUT("/users", ginx.WrapJsonBody[types.UserForm](ctl.UpdateUserInfo))
	group.DELETE("/users/:id", ginx.WrapJsonBody[types.UserForm](ctl.DeleteUserCtl))
}
