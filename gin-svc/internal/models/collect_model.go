package models

import "gorm.io/gorm"

type FavoriteModel struct {
	gorm.Model
	ID uint `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:收藏ID"`
}
