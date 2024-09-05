package controller

import (
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type SocialCtl struct {
}

// LikeCtl
// @Description: 点赞
// @receiver SocialCtl
func (s *SocialCtl) LikeCtl(ctx *gin.Context, req types.LikeForm) (result ginx.JsonResult, err error) {

	return ginx.Success(), nil
}

// CommentCtl
// @Description: 评论
// @receiver SocialCtl
func (s *SocialCtl) CommentCtl(ctx *gin.Context, req types.CommentForm) (result ginx.JsonResult, err error) {
	return ginx.Success(), nil
}

// CollectCtl
// @Description: 收藏
// @receiver SocialCtl
func (s *SocialCtl) CollectCtl(ctx *gin.Context, req types.CollectForm) (result ginx.JsonResult, err error) {
	return ginx.Success(), nil
}

// FollowCtl
// @Description: 关注
// @receiver SocialCtl
func (s *SocialCtl) FollowCtl(ctx *gin.Context, req types.FollowForm) (result ginx.JsonResult, err error) {
	return ginx.Success(), nil
}

func (s *SocialCtl) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/like", ginx.WrapJsonBody[types.LikeForm](s.LikeCtl))
	group.POST("/comment", ginx.WrapJsonBody[types.CommentForm](s.CommentCtl))
	group.POST("/collect", ginx.WrapJsonBody[types.CollectForm](s.CollectCtl))
	group.POST("/follow", ginx.WrapJsonBody[types.FollowForm](s.FollowCtl))
}
