package domain

type DComment struct {
	ID           uint        `json:"id"`
	UserUUId     string      `json:"userUUId"`
	UserAvatar   string      `json:"userAvatar"`
	UserNickName string      `json:"userNickName"`
	IsAuthor     bool        `json:"isAuthor"` // 是否是作者
	ParentId     uint        `json:"parentId"`
	CommentType  string      `json:"commentType"`
	ResourceType string      `json:"resourceType"`
	ResourceId   string      `json:"resourceId"`
	Content      string      `json:"content"`
	IsPinned     bool        `json:"isPinned"`
	MediaUrl     string      `json:"mediaUrl"`
	UserAgent    string      `json:"userAgent"`
	Ip           string      `json:"ip"`
	IpLocation   string      `json:"ipLocation"`
	Status       int         `json:"status"`
	Sub          []*DComment `json:"sub"`
}
