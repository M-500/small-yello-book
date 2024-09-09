package models

import "gorm.io/gorm"

type NoteModel struct {
	gorm.Model
	UUID        string `json:"uuid" gorm:"column:uuid;type:varchar(36);unique;comment:笔记UUID"`                           // 笔记UUID
	ContentType int    `json:"content_type" gorm:"column:content_type;type:int;not null;default: 1;comment:内容类型,1图文,视频"` // 内容类型
	Title       string `json:"title" gorm:"column:title;type:varchar(255);comment:笔记主标题"`                                // 笔记主标题
	Cover       string `json:"cover" gorm:"column:cover;type:varchar(255);comment:封面图片URL"`                              // 封面图片
	Mark        string `json:"mark" gorm:"column:mark;type:text;comment:笔记描述"`                                           // 笔记描述,内容
	Status      int    `json:"status" gorm:"column:status;type:int;comment:状态,0 草稿/审核中，1 已发布，2 未通过，3 已删除"`               // 笔记状态
	AuthorId    string `json:"author_id" gorm:"column:author_id;type:varchar(36);comment:作者UUID"`
	UserAgent   string `gorm:"size:1024" json:"userAgent" form:"userAgent"` // UserAgent
	Ip          string `gorm:"size:128" json:"ip" form:"ip"`                // IP
	IpLocation  string `gorm:"size:64" json:"ipLocation" form:"ipLocation"` // IP属地
}

func (NoteModel) TableName() string {
	return "notes"
}

// 笔记类型关联表
type NoteTagModel struct {
	NoteId uint `json:"note_id" gorm:"column:note_id;type:int;comment:笔记ID"`
	TagId  uint `json:"tag_id" gorm:"column:tag_id;type:int;comment:标签ID"`
}

func (NoteTagModel) TableName() string {
	return "note_tag"
}
