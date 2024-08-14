package service

import (
	"context"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"testing"
)

func Test_userSvcImpl_EmailLogin(t *testing.T) {
	type fields struct {
		Handler  jwt.Handler
		userRepo repo.UserRepoInterface
		verCache cache.VerificationCodeCache
	}
	type args struct {
		ctx context.Context
		req types.EmailLoginForm
	}
	tests := []struct {
		name string
		// 预期输入 依赖等
		fields fields
		args   args
		// 预期输出
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				Handler:  nil,
				userRepo: nil,
				verCache: nil,
			},
			args:    args{},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userSvcImpl{
				Handler:  tt.fields.Handler,
				userRepo: tt.fields.userRepo,
				verCache: tt.fields.verCache,
			}
			got, err := u.EmailLogin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EmailLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
