package domain

type DNote struct {
	ID          int `json:"id"`
	NoteTitle   string
	NoteContent string
	Cover       string
	Address     string
	PublishTime string
	Private     bool
	ViewCnt     int
	LikeCnt     int
	ShareCnt    int
	CommentCnt  int
	CollectCnt  int
	Status      int
	AuthorId    int
	CreateTime  string
	UpdateTime  string
}
