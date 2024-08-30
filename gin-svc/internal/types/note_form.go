package types

import "gin-svc/internal/domain"

type CreateNoteForm struct {
	NoteTitle   string    `form:"noteTitle" json:"noteTitle" binding:"required"`      // 标记
	NoteContent string    `form:"noteContent" json:"noteContent"  binding:"required"` // 内容
	ImgList     []ImgItem `form:"imgList" json:"imgList" binding:"required"`          // 图片列表
	//Address     string    `form:"address" json:"address"`
	Statement   string `form:"statement" json:"statement"`     // 自主声明  1.原创 2.转载 3.其他
	PublishTime uint   `form:"publishTime" json:"publishTime"` // 发布时间
	Private     bool   `json:"private" form:"private"`         // 是否私密
}

type ImgItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *CreateNoteForm) ToNoteDomain() domain.DNote {
	res := domain.DNote{
		NoteTitle:   c.NoteTitle,
		NoteContent: c.NoteContent,
		PublishTime: c.PublishTime,
		Statement:   c.Statement,
		Private:     c.Private,
	}
	for _, img := range c.ImgList {
		res.ImgList = append(res.ImgList, domain.ImageNote{
			ImgName: img.Name,
			ImgUrl:  img.Url,
		})
	}
	return res
}
