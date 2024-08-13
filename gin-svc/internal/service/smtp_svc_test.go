package service

import "testing"

func Test_smtp163Service_SendEmail(t *testing.T) {
	type fields struct {
		smtpService string
		smtpPort    int
		smtpKey     string
		emailAddr   string
	}
	type args struct {
		to      []string
		subject string
		body    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				smtpService: "smtp.163.com",
				smtpPort:    25,
				smtpKey:     "PNNSCQKYCIIN",
				emailAddr:   "18574945291@163.com",
			},
			args: args{
				to:      []string{"1978992154@qq.com"},
				subject: "登陆验证码",
				body:    "1234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &smtp163Service{
				smtpService: tt.fields.smtpService,
				smtpPort:    tt.fields.smtpPort,
				smtpKey:     tt.fields.smtpKey,
				emailAddr:   tt.fields.emailAddr,
			}
			if err := s.SendEmail(tt.args.to, tt.args.subject, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
