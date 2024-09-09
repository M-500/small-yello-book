package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InteractiveDao interface {
	GetById(ctx context.Context, sourceId, bizType string) (models.InteractiveModel, error) // 通过ID获取资源的Social详情
	IncrLike(ctx context.Context, sourceId, bizType string) error                           // 点赞数+1
	DecrLike(ctx context.Context, sourceId, bizType string) error                           // 点赞数-1
	IncrViewLike(ctx context.Context, sourceId, bizType string) error                       // 浏览数+1
	DecrViewLike(ctx context.Context, sourceId, bizType string) error                       // 浏览数-1
	IncrCollection(ctx context.Context, sourceId, bizType string) error                     // 收藏数+1
	DecrCollection(ctx context.Context, sourceId, bizType string) error                     // 收藏数-1
	IncrComment(ctx context.Context, sourceId, bizType string) error                        // 评论数+1
	DecrComment(ctx context.Context, sourceId, bizType string) error                        // 评论数-1
	Insert(ctx context.Context, social models.InteractiveModel) error                       // 插入一条Social记录
	Update(ctx context.Context, social models.InteractiveModel) error                       // 更新一条Social记录
	Delete(ctx context.Context, sourceId, bizType string) error                             // 删除一条Social记录
}

func NewInteractiveDao(db *gorm.DB) InteractiveDao {
	return &gormInteractiveDao{db: db}
}

type gormInteractiveDao struct {
	db *gorm.DB
}

func (g *gormInteractiveDao) GetById(ctx context.Context, sourceId, bizType string) (models.InteractiveModel, error) {
	var social models.InteractiveModel
	err := g.db.WithContext(ctx).Where("source_gid = ? and biz_type = ?", sourceId, bizType).First(&social).Error
	if err != nil {
		return social, err
	}
	return social, nil
}

func (g *gormInteractiveDao) IncrLike(ctx context.Context, sourceId, bizType string) error {
	return g.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "source_gid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"like_cnt": gorm.Expr("like_cnt + 1"),
		}),
	}).Create(&models.InteractiveModel{
		SourceGID:  sourceId,
		ViewCnt:    0,
		LikeCnt:    1,
		ShareCnt:   0,
		CommentCnt: 0,
		CollectCnt: 0,
		BizType:    bizType,
	}).Error
}

func (g *gormInteractiveDao) DecrLike(ctx context.Context, sourceId, bizType string) error {
	return g.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "source_gid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"like_cnt": gorm.Expr("like_cnt - 1"),
		}),
	}).Error
}

func (g *gormInteractiveDao) IncrViewLike(ctx context.Context, sourceId, bizType string) error {
	return g.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "source_gid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"like_cnt": gorm.Expr("view_cnt + 1"),
		}),
	}).Create(&models.InteractiveModel{
		SourceGID:  sourceId,
		ViewCnt:    0,
		LikeCnt:    1,
		ShareCnt:   0,
		CommentCnt: 0,
		CollectCnt: 0,
		BizType:    bizType,
	}).Error
}

func (g *gormInteractiveDao) DecrViewLike(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (g *gormInteractiveDao) IncrCollection(ctx context.Context, sourceId, bizType string) error {
	return g.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "source_gid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"like_cnt": gorm.Expr("collect_cnt + 1"),
		}),
	}).Create(&models.InteractiveModel{
		SourceGID:  sourceId,
		ViewCnt:    0,
		LikeCnt:    1,
		ShareCnt:   0,
		CommentCnt: 0,
		CollectCnt: 0,
		BizType:    bizType,
	}).Error
}

func (g *gormInteractiveDao) DecrCollection(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (g *gormInteractiveDao) IncrComment(ctx context.Context, sourceId, bizType string) error {
	return g.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "source_gid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"like_cnt": gorm.Expr("comment_cnt + 1"),
		}),
	}).Create(&models.InteractiveModel{
		SourceGID:  sourceId,
		ViewCnt:    0,
		LikeCnt:    1,
		ShareCnt:   0,
		CommentCnt: 0,
		CollectCnt: 0,
		BizType:    bizType,
	}).Error
}

func (g *gormInteractiveDao) DecrComment(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}

func (g *gormInteractiveDao) Insert(ctx context.Context, social models.InteractiveModel) error {
	//TODO implement me
	panic("implement me")
}

func (g *gormInteractiveDao) Update(ctx context.Context, social models.InteractiveModel) error {
	//TODO implement me
	panic("implement me")
}

func (g *gormInteractiveDao) Delete(ctx context.Context, sourceId, bizType string) error {
	//TODO implement me
	panic("implement me")
}
