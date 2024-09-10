package models

import "gorm.io/gorm"

type NotificationModel struct {
	gorm.Model
	UpstreamUUID string `gorm:"type:varchar(32);index;not null;comment:上游用户UUID"`
	ResourceId   string `gorm:"type:varchar(32);index;not null;comment:资源ID"`
	ResourceType string `gorm:"type:varchar(32);index;not null;comment:资源类型"`
	OwnerId      string `gorm:"type:varchar(32);index;not null;comment:资源拥有者ID"`
	Type         string `gorm:"type:varchar(32);index;not null;comment:通知类型"`
	IsRead       bool   `gorm:"type:tinyint;index;not null;comment:是否已读"`
	Content      string `gorm:"type:varchar(255);not null;comment:通知内容"`
}

func (NotificationModel) TableName() string {
	return "notification"
}
