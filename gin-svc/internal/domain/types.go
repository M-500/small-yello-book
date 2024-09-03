package domain

var NoteStates uint

const (
	// 草稿
	NoteStateDraft uint = iota
	// 审核中
	NoteStateReviewing
	// 审核通过
	NoteStateReviewed
	// 审核未通过
	NoteStateReviewFailed
	// 已发布
	NoteStatePublished
	// 已删除
	NoteStateDeleted
)
