package models

import "gorm.io/gorm"

type InteractiveModel struct {
	gorm.Model
	//UserUUID   string `json:"user_uuid" gorm:"column:user_uuid;type:varchar(36);comment:用户UUID"`           // 用户UUID
	//OwnUserID  string `json:"own_user_id" gorm:"column:own_user_id;type:varchar(36);comment:拥有者用户ID"`      // 拥有者用户ID
	SourceGID  string `json:"source_gid" gorm:"column:source_gid;type:varchar(36);unique;comment:资源的全局ID"` // 源ID
	ViewCnt    int    `json:"view_cnt" gorm:"column:view_cnt;type:int;comment:浏览数"`                        // 观看量
	LikeCnt    int    `json:"like_cnt" gorm:"column:like_cnt;type:int;comment:点赞数"`                        // 点赞量
	ShareCnt   int    `json:"share_cnt" gorm:"column:share_cnt;type:int;comment:分享数"`                      // 分享数
	CommentCnt int    `json:"comment_cnt" gorm:"column:comment_cnt;type:int;comment:评论数"`                  // 评论数
	CollectCnt int    `json:"collect_cnt" gorm:"column:collect_cnt;type:int;comment:收藏数"`                  // 收藏数
	BizType    string `json:"biz_type" gorm:"column:biz_type;type:varchar(36);comment:业务类型"`               // 业务类型
}

func (InteractiveModel) TableName() string {
	return "interactive"
}
