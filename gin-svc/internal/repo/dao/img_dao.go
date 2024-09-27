package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type ImageDao interface {
	FindListByNoteId(ctx context.Context, noteId string) ([]models.ImageModel, error)
}

type imgDao struct {
	db *gorm.DB
}

func NewImageDao(db *gorm.DB) ImageDao {
	return &imgDao{db: db}
}

func (i *imgDao) FindListByNoteId(ctx context.Context, noteId string) ([]models.ImageModel, error) {
	var res []models.ImageModel
	err := i.db.WithContext(ctx).Model(&models.ImageModel{}).Where("note_id = ?", noteId).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
