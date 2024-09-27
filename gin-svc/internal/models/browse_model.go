package models

import "gorm.io/gorm"

type BrowseLogModel struct {
	gorm.Model
	UserId string `json:"user_id" gorm:"column:user_id;type:varchar(32);comment:用户UUID"` // 用户UUID
	NoteId string `json:"note_id" gorm:"column:note_id;type:varchar(32);comment:笔记ID"`   // 笔记ID
}

func (BrowseLogModel) TableName() string {
	return "browse_logs"
}
