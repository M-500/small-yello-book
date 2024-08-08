package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type publicController struct {
	userSvc service.UserSvc
}

func NewPublicController(userSvc service.UserSvc) BaseController {
	return &publicController{
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
// @Router /api/v1/na/register [put]
func (p *publicController) RegisterUser(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = p.userSvc.RegisterUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "创建/更新失败"), err
	}
	return ginx.Success(), nil
}

func (p *publicController) RegisterRoute(group *gin.RouterGroup) {
	group.PUT("/register", ginx.WrapJsonBody[types.UserForm](p.RegisterUser))
}
