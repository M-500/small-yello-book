package repo

import (
	"context"
	"gin-svc/internal/repo/cache"
)

//go:generate mockgen -destination=../repo/mocks/mock_verification_code_repo.go -package=mocks gin-svc/internal/repo VerificationCodeRepo
type VerificationCodeRepo interface {
	// SetVerificationCode 设置验证码
	SetVerificationCode(ctx context.Context, email string, code string) error
	// GetVerificationCode 获取验证码
	GetVerificationCode(ctx context.Context, key string) (string, error)
	// DeleteVerificationCode 删除验证码
	DeleteVerificationCode(ctx context.Context, key string) error
}

type verificationCodeRepo struct {
	cache cache.VerificationCodeCache
}

func NewVerificationCodeRepo(cachePar cache.VerificationCodeCache) VerificationCodeRepo {
	return &verificationCodeRepo{
		cache: cachePar,
	}
}

func (v *verificationCodeRepo) SetVerificationCode(ctx context.Context, email string, code string) error {
	return v.cache.SetVerificationCode(ctx, email, code)
}

func (v *verificationCodeRepo) GetVerificationCode(ctx context.Context, key string) (string, error) {
	return v.cache.GetVerificationCode(ctx, key)
}

func (v *verificationCodeRepo) DeleteVerificationCode(ctx context.Context, key string) error {
	return v.cache.DeleteVerificationCode(ctx, key)
}
