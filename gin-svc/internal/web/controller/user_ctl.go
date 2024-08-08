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

// RegisterUser godoc
// @Summary RegisterUser
// @Description RegisterUser
// @Tags NA Model
// @Accept json
// @Produce json
// @Param user body types.UserForm true "user"
// @Success 200 {object} map[string]any
// @Router /api/v1/na/register [get]
func (ctl *userController) RegisterUser(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = ctl.userSvc.RegisterUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "创建/更新失败"), err
	}
	return ginx.Success(), nil
}

func (ctl *userController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/na/register", ginx.WrapJsonBody[types.UserForm](ctl.RegisterUser))
}
