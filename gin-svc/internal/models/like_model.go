package models

import "gorm.io/gorm"

type LikeModel struct {
	gorm.Model
	UserId     string `json:"userId" gorm:"column:user_id;type:varchar(36);comment:用户ID;uniqueIndex:idx_uer_source_type"`                  // 用户UUID
	SourceUUID string `json:"sourceUUID" gorm:"column:source_uuid;type:varchar(36);unique;comment:资源UUID;uniqueIndex:idx_uer_source_type"` // 资源UUID
	BizType    string `json:"bizType" gorm:"column:biz_type;type:varchar(36);comment:业务类型;uniqueIndex:idx_uer_source_type"`                // 业务类型
}

func (LikeModel) TableName() string {
	return "like_log"
}
