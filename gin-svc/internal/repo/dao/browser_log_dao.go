package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type BrowserLogDAO interface {
	Insert(ctx context.Context, data models.BrowseLogModel) error // 新增一条浏览记录
	FindByUserID(ctx context.Context, userID int64, page, size int) ([]models.BrowseLogModel, int64, error)
}

type gormBrowserLogDAO struct {
	db *gorm.DB
}

func (g *gormBrowserLogDAO) Insert(ctx context.Context, data models.BrowseLogModel) error {
	return g.db.WithContext(ctx).Model(&models.BrowseLogModel{}).Create(&data).Error
}

func (g *gormBrowserLogDAO) FindByUserID(ctx context.Context, userID int64, page, size int) ([]models.BrowseLogModel, int64, error) {
	var (
		total int64
		res   []models.BrowseLogModel
	)
	err := g.db.WithContext(ctx).Model(&models.BrowseLogModel{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = g.db.WithContext(ctx).Model(&models.BrowseLogModel{}).Where("user_id = ?", userID).Offset((page - 1) * size).Limit(size).Find(&res).Error
	return res, total, err
}
