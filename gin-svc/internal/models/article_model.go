package models

import "gorm.io/gorm"

type NoteModel struct {
	gorm.Model
	Title      string `json:"title" gorm:"column:title;type:varchar(255);comment:笔记主标题"`
	Mark       string `json:"mark" gorm:"column:mark;type:varchar(255);comment:笔记描述"`
	Type       string `json:"type" gorm:"column:type;type:varchar(8);comment:类型,图片，视频"`
	LikeCnt    int    `json:"like_cnt" gorm:"column:like_cnt;type:int;comment:点赞数"`
	CommentCnt int    `json:"comment_cnt" gorm:"column:comment_cnt;type:int;comment:评论数"`
	AuthorId   uint   `json:"author_id" gorm:"column:author_id;type:int;comment:作者ID"`
}

func (NoteModel) TableName() string {
	return "article"
}
