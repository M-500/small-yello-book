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
// @Summary 注册用户
// @Description 注册用户
// @Tags 公共模块
// @Accept json
// @Produce json
// @Param user body types.UserForm true "user"
// @Success 200 {object} map[string]any
// @Router /api/v1/na/register [post]
func (p *publicController) RegisterUser(ctx *gin.Context, req types.UserForm) (result ginx.JsonResult, err error) {
	err = p.userSvc.RegisterUser(ctx, req)
	if err != nil {
		return ginx.Error(10011, "创建/更新失败"), err
	}
	return ginx.Success(), nil
}

// LoginCtl godoc
// @Summary 用户密码登录
// @Description 用户密码登录
// @Tags 公共模块
// @Accept json
// @Produce json
// @Param user body types.LoginForm true "user"
// @Success 200 {object} map[string]any
// @Router /api/v1/na/login [post]
func (p *publicController) LoginCtl(ctx *gin.Context, req types.LoginForm) (result ginx.JsonResult, err error) {
	err = p.userSvc.PwdLogin(ctx, req)
	if err != nil {
		return ginx.Error(10011, "登陆失败"), err
	}
	return ginx.Success(), nil
}

func (p *publicController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/register", ginx.WrapJsonBody[types.UserForm](p.RegisterUser))
	group.POST("/login", ginx.WrapJsonBody[types.LoginForm](p.LoginCtl))
}
