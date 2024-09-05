package types

type LikeForm struct {
	ResourceUUID string `json:"resource_uuid" binding:"required"` // 被点赞的资源ID
	BizType      int    `json:"biz_type" binding:"required"`      // 业务类型
	Action       int    `json:"action" binding:"required"`        // 1.点赞 2.取消点赞
}

type CommentForm struct {
	ResourceUUID string `json:"resource_uuid" binding:"required"` // 被评论的资源ID
	BizType      int    `json:"biz_type" binding:"required"`      // 业务类型
	Content      string `json:"content" binding:"required"`       // 评论内容
}

type CollectForm struct {
	ResourceUUID string `json:"resource_uuid" binding:"required"` // 被收藏的资源ID
	BizType      int    `json:"biz_type" binding:"required"`      // 业务类型
	Action       int    `json:"action" binding:"required"`        // 1.收藏 2.取消收藏
}

type FollowForm struct {
	FollowUUID string `json:"follow_uuid" binding:"required"` // 被关注的用户ID
	Action     int    `json:"action" binding:"required"`      // 1.关注 2.取消关注
}
