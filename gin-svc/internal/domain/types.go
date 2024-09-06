package domain

type NoteStatus int

const (
	NoteStateDraft     NoteStatus = iota
	NoteStatePublished            // 删除
	NoteStateRefuse               // 草稿
	NoteStateDeleted              // 审核通过
)
