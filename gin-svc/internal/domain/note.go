package domain

type DNote struct {
	ID          int    `json:"id"`
	NoteTitle   string `json:"noteTitle"`
	NoteContent string `json:"noteContent"`
	Cover       string `json:"cover"`
	Address     string `json:"address"`
	PublishTime string `json:"publishTime"`
	Private     bool   `json:"private"`
	ViewCnt     int    `json:"viewCnt"`
	LikeCnt     int    `json:"likeCnt"`
	ShareCnt    int    `json:"shareCnt"`
	CommentCnt  int    `json:"commentCnt"`
	CollectCnt  int    `json:"collectCnt"`
	Status      int    `json:"status"`
	AuthorId    int    `json:"authorId"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
