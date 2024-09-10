package controller

import (
	"gin-svc/internal/domain"
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func NewInteractiveController(intrSvc service.InteractiveSvc) *InteractiveCtl {
	return &InteractiveCtl{
		intrSvc: intrSvc,
	}
}

type InteractiveCtl struct {
	intrSvc service.InteractiveSvc
}

// LikeCtl
//
//	@Description: 点赞
//	@receiver c
//	@param ctx
//	@return result
//	@return err
func (i *InteractiveCtl) LikeCtl(ctx *gin.Context, data types.InteractiveForm, uc jwt.UserClaims) (result ginx.JsonResult, err error) {
	notify := domain.DNotify{
		UpstreamUUID: uc.Uid,
		ResourceId:   data.ResourceId,
		ResourceType: data.ResourceType,
		OwnerId:      data.OwnerId,
		Action:       domain.Like,
		IsRead:       false,
	}
	err = i.intrSvc.IncrLikeCnt(ctx, notify)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.Success(), nil
}

func (i *InteractiveCtl) CollectCtl(ctx *gin.Context, uc jwt.UserClaims) (result ginx.JsonResult, err error) {
	return ginx.Success(), nil
}
