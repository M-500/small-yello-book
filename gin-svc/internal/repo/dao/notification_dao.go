package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type NotificationDao interface {
	Insert(ctx context.Context, notification models.NotificationModel) error
}

func NewNotificationDao(db *gorm.DB) NotificationDao {
	return &notificationDAO{db: db}
}

type notificationDAO struct {
	db *gorm.DB
}

func (n *notificationDAO) Insert(ctx context.Context, notification models.NotificationModel) error {
	return n.db.WithContext(ctx).Create(&notification).Error
}
