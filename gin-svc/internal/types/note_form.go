package types

type CreateNoteArticleForm struct {
	NoteTitle   string `form:"noteTitle" json:"noteTitle" binding:"required"`
	NoteContent string `form:"noteContent" json:"noteContent"  binding:"required"`
	Cover       string `form:"cover" json:"cover"  binding:"required"`
	Address     string `form:"address" json:"address"`
	PublishTime string `form:"publishTime" json:"publishTime"`
	Private     bool   `json:"private" form:"private"`
}
