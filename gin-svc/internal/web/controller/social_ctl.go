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
