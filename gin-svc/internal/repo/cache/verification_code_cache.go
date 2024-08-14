package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	VerificationCodeCacheKey = "verification_code:"
	TimeExpire               = 5 * time.Minute
)

//go:generate mockgen -destination=./repo/mocks/mock_verification_code_cache.go -package=mocks gin-svc/internal/repo VerificationCodeCache
type VerificationCodeCache interface {
	// SetVerificationCode 设置验证码
	SetVerificationCode(ctx context.Context, email string, code string) error
	// GetVerificationCode 获取验证码
	GetVerificationCode(ctx context.Context, key string) (string, error)
	// DeleteVerificationCode 删除验证码
	DeleteVerificationCode(ctx context.Context, key string) error
}

type verificationCodeCacheImpl struct {
	cli redis.Cmdable
}

func NewRedisVerificationCodeCache(cli redis.Cmdable) VerificationCodeCache {
	return &verificationCodeCacheImpl{cli: cli}
}

func (v *verificationCodeCacheImpl) SetVerificationCode(ctx context.Context, email string, code string) error {
	return v.cli.Set(ctx, v.key(email), code, TimeExpire).Err()
}

func (v *verificationCodeCacheImpl) GetVerificationCode(ctx context.Context, email string) (string, error) {
	return v.cli.Get(ctx, v.key(email)).Result()
}

func (v *verificationCodeCacheImpl) DeleteVerificationCode(ctx context.Context, key string) error {
	return v.cli.Del(ctx, key).Err()
}

func (v *verificationCodeCacheImpl) key(email string) string {
	return VerificationCodeCacheKey + email
}
