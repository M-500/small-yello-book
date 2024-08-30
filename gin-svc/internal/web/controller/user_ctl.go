package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc service.UserSvc
}

func NewUserController(userSvc service.UserSvc) *UserController {
	return &UserController{
		userSvc: userSvc,
	}
}

// UpdateUserInfo godoc
// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param user body types.UserForm true "user"
// @Success 200 {object} map[string]any
// @Router /api/v1/users [put]
func (u *UserController) UpdateUserInfo(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = u.userSvc.DeleteUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "更新失败"), err
	}
	return ginx.Success(), nil
}

func (u *UserController) GetUserInfo(ctx *gin.Context, claim jwt.UserClaims) (result ginx.JsonResult, err error) {
	//id := claim.Uid

	return ginx.JsonResult{}, nil
}

// DeleteUserCtl godoc
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]any
// @Router /api/v1/users/{id} [delete]
func (u *UserController) DeleteUserCtl(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = u.userSvc.DeleteUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "更新失败"), err
	}
	return ginx.Success(), nil
}

func (u *UserController) RegisterRoute(group *gin.RouterGroup) {
	group.PUT("/users", ginx.WrapJsonBody[types.UserForm](u.UpdateUserInfo))
	group.DELETE("/users/:id", ginx.WrapJsonBody[types.UserForm](u.DeleteUserCtl))
}
