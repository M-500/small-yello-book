package service

import (
	"context"
	"errors"
	"fmt"
	"gin-svc/internal/repo/cache"
	"gin-svc/pkg/utils"
	"github.com/jordan-wright/email"
	"github.com/redis/go-redis/v9"
	"net/smtp"
	"strconv"
)

type SmtpServiceInterface interface {
	SendEmail(to []string, subject, body string) error
	LoginEmailSend(ctx context.Context, email string) error
}

type smtp163Service struct {
	smtpService string // SMTP 服务器
	smtpPort    int    // SMTP 端口
	smtpKey     string // SMTP 密钥
	emailAddr   string // 发送邮件地址
	cache       cache.VerificationCodeCache
}

func NewSMTPService(cach cache.VerificationCodeCache) SmtpServiceInterface {
	return &smtp163Service{
		smtpService: "smtp.163.com",
		smtpPort:    25,
		smtpKey:     "PNNSCQKYTCCBCIIN",
		emailAddr:   "18574945291@163.com",
		cache:       cach,
	}
}

func (s *smtp163Service) SendEmail(to []string, subject, body string) error {
	em := email.NewEmail()
	em.From = s.emailAddr // 发件人
	em.To = to            // 收件人(目标邮箱)
	em.Subject = "小黄书 ybook" + subject
	em.Text = []byte(body)
	// 调用接口发送邮件
	// 此处端口号不一定为25使用对应邮箱时需要具体更换
	return em.Send(s.smtpService+":"+strconv.Itoa(s.smtpPort), smtp.PlainAuth("", s.emailAddr, s.smtpKey, s.smtpService))
}

func (s *smtp163Service) LoginEmailSend(ctx context.Context, to string) error {
	// 生成6位随机数字
	verificationCode := utils.GenerateRandomNumberStr()
	fmt.Println("发送的验证码为:", verificationCode)
	// 查询redis是否存在
	code, err := s.cache.GetVerificationCode(ctx, to)
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	if len(code) > 0 {
		// 重复发送
		return errors.New("重复发送，请稍后再试")
	}
	err = s.cache.SetVerificationCode(ctx, to, verificationCode)
	if err != nil {
		return err
	}
	// 发送邮件
	//go func() {
	//	// 应该要写入kafka中，以便后续处理
	//	err = s.SendEmail([]string{to}, "登陆验证码", verificationCode)
	//}()
	err = s.SendEmail([]string{to}, "登陆验证码", verificationCode)
	return err
}
