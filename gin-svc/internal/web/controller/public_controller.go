package controller

import (
	"errors"
	"fmt"
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/types/resp"
	"gin-svc/pkg/ginx"
	"gin-svc/pkg/utils"
	"github.com/gin-gonic/gin"
)

type publicController struct {
	userSvc service.UserSvc
	smtpSvc service.SmtpServiceInterface
}

func NewPublicController(userSvc service.UserSvc, smtp service.SmtpServiceInterface) BaseController {
	return &publicController{
		userSvc: userSvc,
		smtpSvc: smtp,
	}
}

// EmailLoginCtl godoc
// @Summary 邮箱登录/注册
// @Description 邮箱登录/注册
// @Tags 公共模块
// @Accept json
// @Produce json
// @Param user body types.EmailLoginForm true "user"
// @Success 200 {object} resp.LoginResp
// @Failure 10011 {string} string "登陆失败"
// @Router /api/v1/na/login [post]
func (p *publicController) EmailLoginCtl(ctx *gin.Context, req types.EmailLoginForm) (result ginx.JsonResult, err error) {
	token, err := p.userSvc.EmailLogin(ctx, req)
	if err != nil {
		fmt.Println(err)
		return ginx.Error(10011, "登陆失败"), err
	}
	res := resp.LoginResp{Token: token}
	return ginx.JsonResult{Data: res}, nil
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

// EmailSendCtl godoc
// @Summary 发送邮件
// @Description 发送邮件
// @Tags 公共模块
// @Accept json
// @Produce json
// @Param user body types.EmailForm true "user"
// @Success 200 {object} map[string]any
// @Router /api/v1/na/email/send [post]
func (p *publicController) EmailSendCtl(ctx *gin.Context, req types.EmailForm) (result ginx.JsonResult, err error) {
	// 校验邮箱格式是否合法
	email := utils.IsValidEmail(req.Email)
	if !email {
		return ginx.Error(10011, "邮箱格式不合法"), errors.New("邮箱格式不合法")
	}
	err = p.smtpSvc.LoginEmailSend(ctx, req.Email)
	return ginx.JsonResult{
		Code: 500,
		Msg:  err.Error(),
	}, err
}

func (p *publicController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/login", ginx.WrapJsonBody[types.EmailLoginForm](p.EmailLoginCtl))
	//group.POST("/login", ginx.WrapJsonBody[types.LoginForm](p.LoginCtl))
	group.POST("/email/send", ginx.WrapJsonBody[types.EmailForm](p.EmailSendCtl))
}
