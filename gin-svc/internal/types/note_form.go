package types

import (
	"gin-svc/internal/conf/cfg"
	"gin-svc/internal/domain"
	"gin-svc/pkg/utils"
	"strings"
)

type CreateNoteForm struct {
	NoteTitle   string    `form:"noteTitle" json:"noteTitle" binding:"required"`               // 标记
	ContentType int       `form:"contentType" json:"contentType" binding:"required,oneof=1 2"` // 文章类型，图文/视频
	NoteContent string    `form:"noteContent" json:"noteContent"  binding:"required"`          // 内容
	ImgList     []ImgItem `form:"imgList" json:"imgList" binding:"required"`                   // 图片列表
	//Address     string    `form:"address" json:"address"`
	Statement   string `form:"statement" json:"statement" binding:"required"` // 自主声明  1.原创 2.转载 3.其他
	PublishTime uint   `form:"publishTime" json:"publishTime"`                // 发布时间
	Private     bool   `json:"private" form:"private" binding:"required"`     // 是否私密
}

type ImgItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *CreateNoteForm) ToNoteDomain(cfg cfg.ServerCfg) domain.DNote {
	res := domain.DNote{
		ID:          utils.UUID(),
		NoteTitle:   c.NoteTitle,
		NoteContent: c.NoteContent,
		PublishTime: c.PublishTime,
		Statement:   c.Statement,
		ContentType: c.ContentType, // 设置文章类型
		Private:     c.Private,
	}
	if c.ContentType == 1 {
		for _, img := range c.ImgList {
			res.ImgList = append(res.ImgList, domain.ImageNote{
				ImgName:   img.Name,
				ImgUrl:    img.Url,
				ID:        utils.UUID(), // 默认生成UUID
				LocalPath: strings.Replace(img.Url, cfg.FileUploadHost, "", 1),
			})
		}
		return res
	}
	// 视频类型
	res.VideoInfo = &domain.VideoInfo{
		ID: utils.UUID(),
	}
	return res
}

var state uint

//const (
//	NoteStateNormal       = iota
//	NoteStateDeleted      // 删除
//	NoteStateDraft        // 草稿
//	NoteStateReviewed     // 审核通过
//	NoteStateReviewFailed // 审核未通过
//)

type QueryNoteForm struct {
	State int `form:"state" json:"state"` // 状态
	Page  int `form:"page" json:"page"`   // 页码
	Size  int `form:"size" json:"size"`   // 每页数量
}

type FeedNoteQueryForm struct {
	TagId int `json:"tag_id" binding:""`
	Page  int `form:"page" json:"page"` // 页码
	Size  int `form:"size" json:"size"` // 每页数量
}
