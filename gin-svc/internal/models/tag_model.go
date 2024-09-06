package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	UUID        string `json:"uuid" gorm:"type:char(36);not null;unique"`    // UUID
	Name        string `json:"name" gorm:"type:varchar(50);not null;unique"` // 标签名称
	Description string `json:"description" gorm:"type:varchar(255)"`
	Status      bool   `json:"status" gorm:"type:tinyint(1);default:1"`
}

func (TagModel) TableName() string {
	return "tags"
}
