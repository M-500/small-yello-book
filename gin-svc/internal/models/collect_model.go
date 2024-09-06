package models

import "gorm.io/gorm"

// CollectModel 收藏表  这里没有做收藏夹的概念，只是简单的收藏
type CollectModel struct {
	gorm.Model
	UserId     string `json:"userId" gorm:"column:user_id;type:varchar(36);comment:用户ID"`                  // 用户UUID
	SourceUUID string `json:"sourceUUID" gorm:"column:source_uuid;type:varchar(36);unique;comment:资源UUID"` // 资源UUID
	BizType    string `json:"bizType" gorm:"column:biz_type;type:varchar(36);comment:业务类型"`                // 业务类型
}
