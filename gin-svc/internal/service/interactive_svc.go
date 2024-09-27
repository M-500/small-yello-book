package service

import (
	"context"
	"errors"
	"gin-svc/internal/domain"
	"gin-svc/internal/repo"
)

type InteractiveSvc interface {
	IncrLikeCnt(ctx context.Context, data domain.DNotify) error
	DecrLikeCnt(ctx context.Context, data domain.DNotify) error
	IncrViewCnt(ctx context.Context, data domain.DNotify) error
	DecrViewCnt(ctx context.Context, data domain.DNotify) error
	IncrCollectCnt(ctx context.Context, data domain.DNotify) error
	DecrCollectCnt(ctx context.Context, data domain.DNotify) error
	IncrCommentCnt(ctx context.Context, data domain.DNotify) error
	DecrCommentCnt(ctx context.Context, data domain.DNotify) error
}

func NewInteractiveSvc(repo repo.Interactive,
	notifyRepo repo.NotifyRepo, userRepo repo.UserRepoInterface) InteractiveSvc {
	return &interactiveSvc{
		repo:       repo,
		notifyRepo: notifyRepo,
		userRepo:   userRepo,
	}
}

type interactiveSvc struct {
	repo       repo.Interactive
	notifyRepo repo.NotifyRepo
	userRepo   repo.UserRepoInterface
}

func (i *interactiveSvc) IncrCommentCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.IncrComment(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) DecrCommentCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.DecrComment(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) IncrLikeCnt(ctx context.Context, data domain.DNotify) error {
	// 1. 查看资源所属ID
	author, err2 := i.userRepo.FindByUUID(ctx, data.OwnerId)
	if err2 != nil {
		return err2
	}
	if data.OwnerId != author.GlobalNumber {
		return errors.New("资源所属者不匹配")
	}

	// 2. 通知资源所属者 什么人 对你的什么资源进行的点赞
	err := i.notifyRepo.LikeNotify(ctx, data.UpstreamUUID, data.OwnerId, data.ResourceId, data.ResourceType)
	if err != nil {
		// TODO: 记录日志
	}
	// 3. 新增一条点赞记录

	// 将资源对应的点赞值+1
	return i.repo.IncrLike(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) DecrLikeCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.DecrLike(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) IncrViewCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.IncrReadCnt(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) DecrViewCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.DecrCollection(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) IncrCollectCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.IncrCollection(ctx, data.ResourceId, data.ResourceType)
}

func (i *interactiveSvc) DecrCollectCnt(ctx context.Context, data domain.DNotify) error {
	return i.repo.DecrCollection(ctx, data.ResourceId, data.ResourceType)
}
