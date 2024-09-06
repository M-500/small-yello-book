package models

import "gorm.io/gorm"

type SocialModel struct {
	gorm.Model
	SourceGID  string `json:"source_gid" gorm:"column:source_gid;type:varchar(36);unique;comment:资源的全局ID"` // 源ID
	ViewCnt    int    `json:"view_cnt" gorm:"column:view_cnt;type:int;comment:浏览数"`                        // 观看量
	LikeCnt    int    `json:"like_cnt" gorm:"column:like_cnt;type:int;comment:点赞数"`                        // 点赞量
	ShareCnt   int    `json:"share_cnt" gorm:"column:share_cnt;type:int;comment:分享数"`                      // 分享数
	CommentCnt int    `json:"comment_cnt" gorm:"column:comment_cnt;type:int;comment:评论数"`                  // 评论数
	BizType    string `json:"biz_type" gorm:"column:biz_type;type:varchar(36);comment:业务类型"`               // 业务类型
}

func (SocialModel) TableName() string {
	return "social"
}
