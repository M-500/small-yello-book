package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type LikeLogDAO interface {
	Insert(ctx context.Context, data models.LikeModel) error // 新增一条点赞记录
	Delete(ctx context.Context, data models.LikeModel) error // 删除一条点赞记录
	Exists(ctx context.Context, userID, resourceID, bizType string) (bool, error)
	PageListLikeLog(ctx context.Context, userID, resourceID, bizType string, page, size int) ([]models.LikeModel, int64, error)
}

type likeLogDAO struct {
	db *gorm.DB
}

func (l *likeLogDAO) Exists(ctx context.Context, userID, resourceID, bizType string) (bool, error) {
	var count int64
	err := l.db.WithContext(ctx).Model(&models.LikeModel{}).Where("user_id = ? AND source_uuid = ? AND biz_type = ?", userID, resourceID, bizType).Count(&count).Error
	return count > 0, err
}

func (l *likeLogDAO) Insert(ctx context.Context, data models.LikeModel) error {
	return l.db.WithContext(ctx).Model(&models.LikeModel{}).Create(&data).Error
}

func (l *likeLogDAO) Delete(ctx context.Context, data models.LikeModel) error {
	// TODO:
	return l.db.WithContext(ctx).Where("user_id = ? AND source_uuid = ? AND biz_type = ?", data.UserId, data.SourceUUID, data.BizType).Delete(&models.LikeModel{}).Error
}

func (l *likeLogDAO) PageListLikeLog(ctx context.Context, userID, resourceID, bizType string, page, size int) ([]models.LikeModel, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewLikeLogDAO(db *gorm.DB) LikeLogDAO {
	return &likeLogDAO{db: db}
}
