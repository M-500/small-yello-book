package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	UserUUId     string `gorm:"index:idx_comment_user_uuid;column:user_uuid;type:varchar(32);not null" json:"userId" form:"userId"`         // 用户编号
	ParentId     uint   `gorm:"index:idx_comment_parent_id;column:parent_id;default null;comment:父ID，可以为空" json:"parentId" form:"parentId"` // 父评论ID
	CommentType  string `gorm:"not null;type:varchar(16);default:'text';comment:评论类型" json:"commentType" form:"commentType"`                // 评论类型  三种类型 text media  mixed
	ResourceType string `gorm:"not null;type:varchar(32);column:评论资源类型;column:resource_type"`                                               // 被评论实体类型
	ResourceId   string `gorm:"index:idx_comment_resource_id;not null;column:resource_id" json:"entityId" form:"entityId"`                  // 被评论实体编号 上级ID
	Content      string `gorm:"type:text;" json:"content" form:"content"`                                                                   // 内容
	IsPinned     bool   `gorm:"type:tinyint;not null;default 0;comment:是否置顶" json:"isPinned" form:"isPinned"`                               // 是否置顶
	MediaUrl     string `gorm:"size:1024" json:"mediaUrl" form:"mediaUrl"`                                                                  // 媒体资源地址
	UserAgent    string `gorm:"size:1024" json:"userAgent" form:"userAgent"`                                                                // UserAgent
	Ip           string `gorm:"size:128" json:"ip" form:"ip"`                                                                               // IP
	IpLocation   string `gorm:"size:64" json:"ipLocation" form:"ipLocation"`                                                                // IP属地
	Status       int    `gorm:"type:int(11);index:idx_comment_status" json:"status" form:"status"`                                          // 状态：0：待审核、1：审核通过、2：审核失败、3：已发布 4:下架
}

func (CommentModel) TableName() string {
	return "comment"
}
