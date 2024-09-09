package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	UserUUId   string `gorm:"index:idx_comment_user_uuid;not null" json:"userId" form:"userId"`           // 用户编号
	EntityType string `gorm:"index:idx_comment_entity_type;not null" json:"entityType" form:"entityType"` // 被评论实体类型
	EntityId   string `gorm:"index:idx_comment_entity_id;not null" json:"entityId" form:"entityId"`       // 被评论实体编号 上级ID
	Content    string `gorm:"type:text;not null" json:"content" form:"content"`                           // 内容
	ImgUrl     string `gorm:"type:text;not null" json:"imgUrl" form:"imgUrl"`                             // 图片的URL地址
	UserAgent  string `gorm:"size:1024" json:"userAgent" form:"userAgent"`                                // UserAgent
	Ip         string `gorm:"size:128" json:"ip" form:"ip"`                                               // IP
	IpLocation string `gorm:"size:64" json:"ipLocation" form:"ipLocation"`                                // IP属地
	Status     int    `gorm:"type:int(11);index:idx_comment_status" json:"status" form:"status"`          // 状态：0：待审核、1：审核通过、2：审核失败、3：已发布 4:下架
	CreateTime int64  `json:"createTime" form:"createTime"`                                               // 创建时间
}

func (CommentModel) TableName() string {
	return "comment"
}
