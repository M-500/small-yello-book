package domain

type DNote struct {
	ID              string           `json:"uuid"`
	NoteTitle       string           `json:"noteTitle"`
	NoteContent     string           `json:"noteContent"`
	ContentType     int              `json:"contentType"` // 文章类型，图文/视频
	Cover           string           `json:"cover"`       // 封面
	CoverWidth      int              `json:"coverWidth"`  // 封面宽度
	CoverHeight     int              `json:"coverHeight"` // 封面高度
	Address         string           `json:"address"`     // 地址
	Statement       string           `json:"statement"`
	IPLocation      string           `json:"IPLocation"` // IP归属地
	PublishTime     uint             `json:"publishTime"`
	Private         bool             `json:"private"`
	Status          NoteStatus       `json:"status"`
	AuthorId        string           `json:"authorId"`
	CreateTime      string           `json:"createTime"`
	UpdateTime      string           `json:"updateTime"`
	ImgList         []*ImageNote     `json:"imgList"`         // 图片信息
	VideoInfo       *VideoInfo       `json:"videoInfo"`       // 视频信息
	AuthorInfo      *Author          `json:"author"`          // 作者信息
	InteractiveInfo *InteractiveInfo `json:"interactiveInfo"` // 交互信息
}

type InteractiveInfo struct {
	SourceGID  string `json:"sourceGID"`
	ViewCnt    int    `json:"viewCnt"`
	LikeCnt    int    `json:"likeCnt"`
	ShareCnt   int    `json:"shareCnt"`
	CommentCnt int    `json:"commentCnt"`
	CollectCnt int    `json:"collectCnt"`
	BizType    string `json:"bizType"`
}

type Author struct {
	ID       string `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

type ImageNote struct {
	ID         string `json:"id"`
	NoteId     string `json:"noteId"`
	ImgName    string `json:"imgName"`
	ImgUrl     string `json:"imgUrl"`
	ImgWidth   int    `json:"imgWidth"`
	ImgHeight  int    `json:"imgHeight"`
	Size       int64  `json:"size"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	LocalPath  string `json:"localPath"`
	Url        string `json:"url"`
	HashStr    string `json:"hashStr"`
}

type VideoInfo struct {
	ID        string `json:"id"` // 视频ID
	LocalPath string `json:"localPath"`
}
