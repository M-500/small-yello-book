package repo

import (
	"context"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
)

type Interactive interface {
	GetById(ctx context.Context, sourceId, bizType string) (models.InteractiveModel, error) // 通过ID获取资源的Social详情

	IncrLike(ctx context.Context, sourceId, bizType string) error // 点赞数+1
	DecrLike(ctx context.Context, sourceId, bizType string) error // 点赞数-1

	IncrViewLike(ctx context.Context, sourceId, bizType string) error // 浏览数+1
	DecrViewLike(ctx context.Context, sourceId, bizType string) error // 浏览数-1

	IncrCollection(ctx context.Context, sourceId, bizType string) error // 收藏数+1
	DecrCollection(ctx context.Context, sourceId, bizType string) error // 收藏数-1

	IncrComment(ctx context.Context, sourceId, bizType string) error // 评论数+1
	DecrComment(ctx context.Context, sourceId, bizType string) error // 评论数-1

	Insert(ctx context.Context, social models.InteractiveModel) error // 插入一条Social记录
	Update(ctx context.Context, social models.InteractiveModel) error // 更新一条Social记录
	Delete(ctx context.Context, sourceId, bizType string) error       // 删除一条Social记录
}

func NewSchoolRepo(socialCache cache.InteractiveCache, socialDao dao.InteractiveDao) Interactive {
	return &interactive{
		socialCache: socialCache,
		socialDao:   socialDao,
	}
}

type interactive struct {
	socialCache cache.InteractiveCache
	socialDao   dao.InteractiveDao
}

func (s *interactive) GetById(ctx context.Context, sourceId, bizType string) (models.InteractiveModel, error) {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) IncrLike(ctx context.Context, sourceId, bizType string) error {
	err := s.socialDao.IncrLike(ctx, sourceId, bizType)
	if err != nil {
		return err
	}
	go func() {
		//_ = s.socialCache.IncrLikeCnt(ctx, sourceId, bizType)
		// 若是采用协程更新缓存，必须使用 context.Background() 开启一个新的上下文，否则gin.Context的生命周期结束，协程大概率会被取消
		_ = s.socialCache.IncrLikeCnt(context.Background(), sourceId, bizType)
	}()
	return nil
}

func (s *interactive) DecrLike(ctx context.Context, sourceId, bizType string) error {
	err := s.socialDao.DecrLike(ctx, sourceId, bizType)
	if err != nil {
		return err
	}
	go func() {
		//_ = s.socialCache.DecrLikeCnt(ctx, sourceId, bizType)
		_ = s.socialCache.DecrLikeCnt(context.Background(), sourceId, bizType)
	}()
	return nil
}

func (s *interactive) IncrViewLike(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) DecrViewLike(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
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

func (s *interactive) Insert(ctx context.Context, social models.InteractiveModel) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) Update(ctx context.Context, social models.InteractiveModel) error {
	//TODO implement me
	panic("implement me")
}

func (s *interactive) Delete(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}
