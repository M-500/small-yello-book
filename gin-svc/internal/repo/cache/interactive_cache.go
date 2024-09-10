package cache

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	//go:embed lua/add_cnt.lua
	luaIncrCnt string
)

const (
	LikeKey    = "like_cnt"
	CommentKey = "comment_cnt"
	CollectKey = "collect_cnt"
	ViewKey    = "read_cnt"
)

type InteractiveCache interface {
	IncrLikeCnt(ctx context.Context, sourceId, bizType string) error
	DecrLikeCnt(ctx context.Context, sourceId, bizType string) error
	IncrCommentCnt(ctx context.Context, sourceId, bizType string) error
	DecrCommentCnt(ctx context.Context, sourceId, bizType string) error
	IncrCollectCnt(ctx context.Context, sourceId, bizType string) error
	DecrCollectCnt(ctx context.Context, sourceId, bizType string) error
	AddViewCnt(ctx context.Context, sourceId, bizType string) error
}

func NewInteractiveCache(cli redis.Cmdable) InteractiveCache {
	return &interactiveCache{cli: cli}
}

type interactiveCache struct {
	cli redis.Cmdable
}

func (s *interactiveCache) IncrLikeCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, LikeKey, 1).Err()
}

func (s *interactiveCache) DecrLikeCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, LikeKey, -1).Err()
}

func (s *interactiveCache) IncrCommentCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, CommentKey, 1).Err()
}

func (s *interactiveCache) DecrCommentCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, CommentKey, -1).Err()
}

func (s *interactiveCache) IncrCollectCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, CollectKey, 1).Err()
}

func (s *interactiveCache) DecrCollectCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, CollectKey, -1).Err()
}

func (s *interactiveCache) AddViewCnt(ctx context.Context, sourceId, bizType string) error {
	key := s.key(sourceId, bizType)
	// 不是特别需要处理 res
	//res, err := i.client.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Int()
	return s.cli.Eval(ctx, luaIncrCnt, []string{key}, ViewKey, 1).Err()
}

func (s *interactiveCache) DecrViewCnt(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactiveCache) key(sourceId, bizType string) string {
	return fmt.Sprintf("interactive:%s:%s", bizType, sourceId)
}
