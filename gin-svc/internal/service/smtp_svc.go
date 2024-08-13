package service

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"strconv"
)

type SmtpServiceInterface interface {
	SendEmail(to []string, subject, body string) error
}

type smtp163Service struct {
	smtpService string // SMTP 服务器
	smtpPort    int    // SMTP 端口
	smtpKey     string // SMTP 密钥
	emailAddr   string // 发送邮件地址
}

func NewSMTPService() SmtpServiceInterface {
	return &smtp163Service{
		smtpService: "smtp.163.com",
		smtpPort:    25,
		smtpKey:     "PNNSCQKYTCCBCIIN",
		emailAddr:   "18574945291@163.com",
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
