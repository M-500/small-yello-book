package service

import (
	"context"
	"gin-svc/internal/repo"
)

type InteractiveSvc interface {
	IncrLikeCnt(ctx context.Context, uid string, bizType string) error
	DecrLikeCnt(ctx context.Context, uid string, bizType string) error
	IncrViewCnt(ctx context.Context, uid string, bizType string) error
	DecrViewCnt(ctx context.Context, uid string, bizType string) error
	IncrCollectCnt(ctx context.Context, uid string, bizType string) error
	DecrCollectCnt(ctx context.Context, uid string, bizType string) error
	IncrCommentCnt(ctx context.Context, uid string, bizType string) error
	DecrCommentCnt(ctx context.Context, uid string, bizType string) error
}

type interactiveSvc struct {
	repo repo.Interactive
}

func (i *interactiveSvc) IncrCommentCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.IncrComment(ctx, uid, bizType)
}

func (i *interactiveSvc) DecrCommentCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.DecrComment(ctx, uid, bizType)
}

func (i *interactiveSvc) IncrLikeCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.IncrLike(ctx, uid, bizType)
}

func (i *interactiveSvc) DecrLikeCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.DecrLike(ctx, uid, bizType)
}

func (i *interactiveSvc) IncrViewCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.IncrReadCnt(ctx, uid, bizType)
}

func (i *interactiveSvc) DecrViewCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.DecrCollection(ctx, uid, bizType)
}

func (i *interactiveSvc) IncrCollectCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.IncrCollection(ctx, uid, bizType)
}

func (i *interactiveSvc) DecrCollectCnt(ctx context.Context, uid string, bizType string) error {
	return i.repo.DecrCollection(ctx, uid, bizType)
}
