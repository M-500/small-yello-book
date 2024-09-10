package domain

type IntrAction string

const (
	Like           IntrAction = "like"
	Cancel         IntrAction = "cancel_like"
	Collect        IntrAction = "collect"
	CancelCollect  IntrAction = "cancel_collect"
	Comment        IntrAction = "comment"
	Follower       IntrAction = "follower"
	CancelFollower IntrAction = "cancel_follower"
)

type DNotify struct {
	UpstreamUUID string     `json:"upstream_uuid"` // 上游用户UUID（操作者ID)
	ResourceId   string     `json:"resource_id"`   // 资源ID
	ResourceType string     `json:"resource_type"` // 资源类型
	OwnerId      string     `json:"owner_id"`      // 资源拥有者ID
	Action       IntrAction `json:"action"`        // 通知类型
	IsRead       bool       `json:"is_read"`       // 是否已读
	Content      string     `json:"content"`       // 通知内容
}
