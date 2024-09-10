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

type PublicController struct {
	userSvc service.UserSvc
	smtpSvc service.SmtpServiceInterface
}

func NewPublicController(userSvc service.UserSvc, smtp service.SmtpServiceInterface) *PublicController {
	return &PublicController{
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
func (p *PublicController) EmailLoginCtl(ctx *gin.Context, req types.EmailLoginForm) (result ginx.JsonResult, err error) {
	token, err := p.userSvc.EmailLogin(ctx, req)
	if err != nil {
		//fmt.Println(err)
		return ginx.Error(10011, err.Error()), err
	}
	res := resp.LoginResp{Token: token}
	return ginx.JsonResult{Data: res}, nil
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
func (p *PublicController) EmailSendCtl(ctx *gin.Context, req types.EmailForm) (result ginx.JsonResult, err error) {
	// 校验邮箱格式是否合法
	email := utils.IsValidEmail(req.Email)
	if !email {
		return ginx.Error(10011, "邮箱格式不合法"), errors.New("邮箱格式不合法")
	}
	err = p.smtpSvc.LoginEmailSend(ctx, req.Email)
	if err != nil {
		fmt.Println(err)
		return ginx.Error(10011, err.Error()), err
	}
	return ginx.Success(), err
}
