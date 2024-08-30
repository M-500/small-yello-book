package domain

type DNote struct {
	ID          int    `json:"id"`
	NoteTitle   string `json:"noteTitle"`
	NoteContent string `json:"noteContent"`
	Cover       string `json:"cover"`
	Address     string `json:"address"`
	Statement   string `json:"statement"`
	PublishTime uint   `json:"publishTime"`
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
	ImgList     []ImageNote
}

type ImageNote struct {
	ID         int    `json:"id"`
	NoteId     int    `json:"noteId"`
	ImgName    string `json:"imgName"`
	ImgUrl     string `json:"imgUrl"`
	ImgWidth   int    `json:"imgWidth"`
	ImgHeight  int    `json:"imgHeight"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	LocalPath  string `json:"localPath"`
}
