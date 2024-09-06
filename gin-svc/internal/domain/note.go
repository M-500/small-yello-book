package domain

type DNote struct {
	ID          string      `json:"uuid"`
	NoteTitle   string      `json:"noteTitle"`
	NoteContent string      `json:"noteContent"`
	ContentType int         `json:"contentType"` // 文章类型，图文/视频
	Cover       string      `json:"cover"`
	Address     string      `json:"address"`
	Statement   string      `json:"statement"`
	PublishTime uint        `json:"publishTime"`
	Private     bool        `json:"private"`
	ViewCnt     int         `json:"viewCnt"`
	LikeCnt     int         `json:"likeCnt"`
	ShareCnt    int         `json:"shareCnt"`
	CommentCnt  int         `json:"commentCnt"`
	CollectCnt  int         `json:"collectCnt"`
	Status      NoteStatus  `json:"status"`
	AuthorId    int         `json:"authorId"`
	CreateTime  string      `json:"createTime"`
	UpdateTime  string      `json:"updateTime"`
	ImgList     []ImageNote `json:"imgList"`
	VideoInfo   *VideoInfo  `json:"videoInfo"`
	AuthorInfo  *Author     `json:"author"`
}

type Author struct {
	ID       string `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

type ImageNote struct {
	ID         string `json:"id"`
	NoteId     int    `json:"noteId"`
	ImgName    string `json:"imgName"`
	ImgUrl     string `json:"imgUrl"`
	ImgWidth   int    `json:"imgWidth"`
	ImgHeight  int    `json:"imgHeight"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	LocalPath  string `json:"localPath"`
	HashStr    string `json:"hashStr"`
}

type VideoInfo struct {
	ID        string `json:"id"` // 视频ID
	LocalPath string `json:"localPath"`
}
