package types

type AddCommentForm struct {
	Content    string `json:"content"`                        // 评论内容
	ResourceId string `json:"resource_id" binding:"required"` // 所属资源ID
	ParentId   uint   `json:"parent_id"`                      // 父评论ID 可以为空
	MediaUrl   string `json:"media_url"`                      // 媒体资源地址
	IsPinned   bool   `json:"is_pinned"`                      // 是否置顶
}
