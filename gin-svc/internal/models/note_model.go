package models

import "gorm.io/gorm"

type NoteModel struct {
	gorm.Model
	Title      string `json:"title" gorm:"column:title;type:varchar(255);comment:笔记主标题"`
	Cover      string `json:"cover" gorm:"column:cover;type:varchar(255);comment:封面图片URL"`
	Mark       string `json:"mark" gorm:"column:mark;type:varchar(255);comment:笔记描述"`
	Type       string `json:"type" gorm:"column:type;type:varchar(8);comment:类型,图片，视频"`
	Status     int    `json:"status" gorm:"column:status;type:int;comment:状态,0 草稿/审核中，1 已发布，2 未通过，3 已删除"` // 笔记状态
	LikeCnt    int    `json:"like_cnt" gorm:"column:like_cnt;type:int;comment:点赞数"`                       // 点赞量
	ViewCnt    int    `json:"view_cnt" gorm:"column:view_cnt;type:int;comment:浏览数"`                       // 观看量
	ShareCnt   int    `json:"share_cnt" gorm:"column:share_cnt;type:int;comment:分享数"`                     // 分享数
	CommentCnt int    `json:"comment_cnt" gorm:"column:comment_cnt;type:int;comment:评论数"`                 // 评论数
	CollectCnt int    `json:"collect_cnt" gorm:"column:collect_cnt;type:int;comment:收藏数"`                 // 收藏数
	AuthorId   uint   `json:"author_id" gorm:"column:author_id;type:int;comment:作者ID"`
}

func (NoteModel) TableName() string {
	return "notes"
}
