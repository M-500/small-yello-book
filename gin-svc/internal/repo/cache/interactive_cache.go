package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	LikeKey    = "interactive:%s:like:%s"
	CommentKey = "interactive:%s:comment:%s"
	CollectKey = "interactive:%s:collect:%s"
	ViewKey    = "interactive:%s:view:%s"
)

type InteractiveCache interface {
	IncrLikeCnt(ctx context.Context, sourceId, bizType string) error
	DecrLikeCnt(ctx context.Context, sourceId, bizType string) error
	IncrCommentCnt(ctx context.Context, sourceId, bizType string) error
	DecrCommentCnt(ctx context.Context, sourceId, bizType string) error
	IncrCollectCnt(ctx context.Context, sourceId, bizType string) error
	DecrCollectCnt(ctx context.Context, sourceId, bizType string) error
	IncrViewCnt(ctx context.Context, sourceId, bizType string) error
	DecrViewCnt(ctx context.Context, sourceId, bizType string) error
}

func NewInteractiveCache(cli redis.Cmdable) InteractiveCache {
	return &interactiveCache{cli: cli}
}

type interactiveCache struct {
	cli redis.Cmdable
}

func (s *interactiveCache) IncrLikeCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) DecrLikeCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) IncrCommentCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) DecrCommentCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) IncrCollectCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) DecrCollectCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) IncrViewCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) DecrViewCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) likeKey(sourceId, bizType string) string {
	return fmt.Sprintf(LikeKey, bizType, sourceId)
}
func (s *interactiveCache) commentKey(sourceId, bizType string) string {
	return fmt.Sprintf(CommentKey, bizType, sourceId)
}
func (s *interactiveCache) collectKey(sourceId, bizType string) string {
	return fmt.Sprintf(CollectKey, bizType, sourceId)
}
func (s *interactiveCache) viewKey(sourceId, bizType string) string {
	return fmt.Sprintf(ViewKey, bizType, sourceId)
}
