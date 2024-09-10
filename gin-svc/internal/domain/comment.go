package domain

type DComment struct {
	ID           uint        `json:"id"`
	UserUUId     string      `json:"userUUId"`
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
