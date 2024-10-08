package service

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/repo/mocks"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	jwtMock "gin-svc/internal/web/middleware/jwt/mocks"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"

	"go.uber.org/mock/gomock" // 使用uber的mock包
)

func Test_userSvcImpl_EmailLogin(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler)

		// 预期的输入
		ctx context.Context
		req types.EmailLoginForm

		// 预期的输出
		tokenStr string // 预期的token
		wantErr  error  // 预期的错误
	}{
		{
			name: "登陆成功",
			mock: func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler) {
				userRepo := mocks.NewMockUserRepoInterface(ctrl)
				userRepo.EXPECT().FindByEmail(gomock.Any(), "19722221111@qq.com").
					Return(&models.UserModel{
						Email: "19722221111@qq.com",
					}, nil)
				codeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
				// 模拟查询成功
				codeRepo.EXPECT().GetVerificationCode(gomock.Any(), "19722221111@qq.com").
					Return("123456", nil)
				// 模拟删除成功
				codeRepo.EXPECT().DeleteVerificationCode(gomock.Any(), "19722221111@qq.com").
					Return(nil)
				handler := jwtMock.NewMockHandler(ctrl)
				handler.EXPECT().GenJWTToken(gomock.Any()).Return("xxx.weqweq.asijdiouqjwe", nil)
				return userRepo, codeRepo, handler
			},
			ctx: context.Background(),
			req: types.EmailLoginForm{
				Email:   "19722221111@qq.com",
				VerCode: "123456",
			},
			tokenStr: "xxx.weqweq.asijdiouqjwe",
			wantErr:  nil,
		},
		{
			name: "验证码不匹配",
			mock: func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler) {
				userRepo := mocks.NewMockUserRepoInterface(ctrl)
				userRepo.EXPECT().FindByEmail(gomock.Any(), "19722221111@qq.com").
					Return(&models.UserModel{
						Email: "19722221111@qq.com",
					}, nil)
				codeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
				// 模拟查询成功
				codeRepo.EXPECT().GetVerificationCode(gomock.Any(), "19722221111@qq.com").
					Return("214353", nil)
				// 模拟删除成功
				//codeRepo.EXPECT().DeleteVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return(nil)
				handler := jwtMock.NewMockHandler(ctrl)
				//handler.EXPECT().GenJWTToken(gomock.Any()).Return("xxx.weqweq.asijdiouqjwe", nil)
				return userRepo, codeRepo, handler
			},
			ctx: context.Background(),
			req: types.EmailLoginForm{
				Email:   "19722221111@qq.com",
				VerCode: "123456",
			},
			tokenStr: "",
			wantErr:  errors.New("验证码验证失败"),
		},
		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler) {
				userRepo := mocks.NewMockUserRepoInterface(ctrl)
				userRepo.EXPECT().FindByEmail(gomock.Any(), "19722221111@qq.com").
					Return(nil, dao.ErrRecordNotFound)
				codeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
				// 模拟查询成功
				//codeRepo.EXPECT().GetVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return("214353", nil)
				// 模拟删除成功
				//codeRepo.EXPECT().DeleteVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return(nil)
				handler := jwtMock.NewMockHandler(ctrl)
				//handler.EXPECT().GenJWTToken(gomock.Any()).Return("", nil)
				return userRepo, codeRepo, handler
			},
			ctx: context.Background(),
			req: types.EmailLoginForm{
				Email:   "19722221111@qq.com",
				VerCode: "123456",
			},
			tokenStr: "",
			wantErr:  errors.New("用户不存在"),
		},
		{
			name: "验证码过期",
			mock: func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler) {
				userRepo := mocks.NewMockUserRepoInterface(ctrl)
				userRepo.EXPECT().FindByEmail(gomock.Any(), "19722221111@qq.com").
					Return(&models.UserModel{
						Email: "19722221111@qq.com",
					}, nil)
				codeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
				// 模拟查询成功
				codeRepo.EXPECT().GetVerificationCode(gomock.Any(), "19722221111@qq.com").
					Return("", redis.Nil)
				// 模拟删除成功
				//codeRepo.EXPECT().DeleteVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return(nil)
				handler := jwtMock.NewMockHandler(ctrl)
				//handler.EXPECT().GenJWTToken(gomock.Any()).Return("xxx.weqweq.asijdiouqjwe", nil)
				return userRepo, codeRepo, handler
			},
			ctx: context.Background(),
			req: types.EmailLoginForm{
				Email:   "19722221111@qq.com",
				VerCode: "123456",
			},
			tokenStr: "",
			wantErr:  errors.New("验证码错误"),
		},
		{
			name: "MySQL查询错误",
			mock: func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler) {
				userRepo := mocks.NewMockUserRepoInterface(ctrl)
				userRepo.EXPECT().FindByEmail(gomock.Any(), "19722221111@qq.com").
					Return(nil, gorm.ErrInvalidDB)
				codeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
				// 模拟查询成功
				//codeRepo.EXPECT().GetVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return("123456", nil)
				// 模拟删除成功
				//codeRepo.EXPECT().DeleteVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return(nil)
				handler := jwtMock.NewMockHandler(ctrl)
				//handler.EXPECT().GenJWTToken(gomock.Any()).Return("xxx.weqweq.asijdiouqjwe", nil)
				return userRepo, codeRepo, handler
			},
			ctx: context.Background(),
			req: types.EmailLoginForm{
				Email:   "19722221111@qq.com",
				VerCode: "123456",
			},
			tokenStr: "",
			wantErr:  errors.New("系统错误，查询用户失败"),
		},
		{
			name: "Redis查询错误",
			mock: func(ctrl *gomock.Controller) (repo.UserRepoInterface, repo.VerificationCodeRepo, jwt.Handler) {
				userRepo := mocks.NewMockUserRepoInterface(ctrl)
				userRepo.EXPECT().FindByEmail(gomock.Any(), "19722221111@qq.com").
					Return(&models.UserModel{
						Email: "19722221111@qq.com",
					}, nil)
				codeRepo := mocks.NewMockVerificationCodeRepo(ctrl)
				// 模拟查询Redis报错
				codeRepo.EXPECT().GetVerificationCode(gomock.Any(), "19722221111@qq.com").
					Return("", redis.ErrClosed)
				// 模拟删除成功
				//codeRepo.EXPECT().DeleteVerificationCode(gomock.Any(), "19722221111@qq.com").
				//	Return(nil)
				handler := jwtMock.NewMockHandler(ctrl)
				//handler.EXPECT().GenJWTToken(gomock.Any()).Return("xxx.weqweq.asijdiouqjwe", nil)
				return userRepo, codeRepo, handler
			},
			ctx: context.Background(),
			req: types.EmailLoginForm{
				Email:   "19722221111@qq.com",
				VerCode: "123456",
			},
			tokenStr: "",
			wantErr:  errors.New("系统错误，查询验证码失败"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo, cacheRepo, hdl := tc.mock(ctrl)
			svc := &userSvcImpl{
				userRepo: userRepo,
				verRepo:  cacheRepo,
				Handler:  hdl,
			}

			tokenStr, err := svc.EmailLogin(tc.ctx, tc.req)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, tokenStr)
			}
		})
	}
}
