package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
)

type Interactive interface {
	GetById(ctx context.Context, sourceId, bizType string) (domain.InteractiveInfo, error) // 通过ID获取资源的交互详情
	IncrLike(ctx context.Context, sourceId, bizType string) error                          // 点赞数+1
	DecrLike(ctx context.Context, sourceId, bizType string) error                          // 点赞数-1

	IncrReadCnt(ctx context.Context, sourceId, bizType string) error // 浏览数+1

	IncrCollection(ctx context.Context, sourceId, bizType string) error // 收藏数+1
	DecrCollection(ctx context.Context, sourceId, bizType string) error // 收藏数-1

	IncrComment(ctx context.Context, sourceId, bizType string) error // 评论数+1
	DecrComment(ctx context.Context, sourceId, bizType string) error // 评论数-1

	Insert(ctx context.Context, social domain.InteractiveInfo) error // 插入一条Social记录
	Update(ctx context.Context, social domain.InteractiveInfo) error // 更新一条Social记录
	Delete(ctx context.Context, sourceId, bizType string) error      // 删除一条Social记录

	GetIntrBySourceId(ctx context.Context, sourceId, bizType string) (domain.InteractiveInfo, error)
}

func NewInteractiveRepo(intrCache cache.InteractiveCache, dao dao.InteractiveDao) Interactive {
	return &interactive{
		cache: intrCache,
		dao:   dao,
	}
}

type interactive struct {
	cache cache.InteractiveCache
	dao   dao.InteractiveDao
}

func (s *interactive) toDomainInteractive(data models.InteractiveModel) domain.InteractiveInfo {
	return domain.InteractiveInfo{
		SourceGID:  data.SourceGID,
		ViewCnt:    data.ViewCnt,
		LikeCnt:    data.LikeCnt,
		ShareCnt:   data.ShareCnt,
		CommentCnt: data.CommentCnt,
		CollectCnt: data.CollectCnt,
		BizType:    data.BizType,
	}
}

func (s *interactive) GetIntrBySourceId(ctx context.Context, sourceId, bizType string) (domain.InteractiveInfo, error) {
	data, err := s.dao.GetById(ctx, sourceId, bizType)
	if err != nil {
		return domain.InteractiveInfo{}, err
	}
	return s.toDomainInteractive(data), nil
}

func (s *interactive) GetById(ctx context.Context, sourceId, bizType string) (domain.InteractiveInfo, error) {
	data, err := s.dao.GetById(ctx, sourceId, bizType)
	return s.toDomainInteractive(data), err
}

func (s *interactive) IncrLike(ctx context.Context, sourceId, bizType string) error {
	err := s.dao.IncrLike(ctx, sourceId, bizType)
	if err != nil {
		return err
	}
	go func() {
		//_ = s.cache.IncrLikeCnt(ctx, sourceId, bizType)
		// 若是采用协程更新缓存，必须使用 context.Background() 开启一个新的上下文，否则gin.Context的生命周期结束，协程大概率会被取消
		_ = s.cache.IncrLikeCnt(context.Background(), sourceId, bizType)
	}()
	return nil
}

func (s *interactive) DecrLike(ctx context.Context, sourceId, bizType string) error {
	err := s.dao.DecrLike(ctx, sourceId, bizType)
	if err != nil {
		return err
	}
	go func() {
		//_ = s.cache.DecrLikeCnt(ctx, sourceId, bizType)
		_ = s.cache.DecrLikeCnt(context.Background(), sourceId, bizType)
	}()
	return nil
}

func (s *interactive) IncrReadCnt(ctx context.Context, sourceId, bizType string) error {
	err := s.dao.IncrViewCnt(ctx, sourceId, bizType)
	if err != nil {
		return err
	}
	// 更新缓存，可能会数据不一致
	go func() {
		//_ = s.cache.IncrLikeCnt(ctx, sourceId, bizType)
		// 若是采用协程更新缓存，必须使用 context.Background() 开启一个新的上下文，否则gin.Context的生命周期结束，协程大概率会被取消
		_ = s.cache.AddViewCnt(context.Background(), sourceId, bizType)
	}()
	return nil
}

func (s *interactive) IncrCollection(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) DecrCollection(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) IncrComment(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) DecrComment(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) Insert(ctx context.Context, social domain.InteractiveInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) Update(ctx context.Context, social domain.InteractiveInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) Delete(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}
