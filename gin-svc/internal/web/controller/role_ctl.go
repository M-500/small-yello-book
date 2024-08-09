package controller

import (
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type roleController struct {
}

func NewRoleController() BaseController {
	return &roleController{}
}

func (r *roleController) RegisterRoute(group *gin.RouterGroup) {
	group.GET("/roles", ginx.WrapJsonBody[types.RolePageListReq](r.RolePageList))
}

// RolePageList godoc
// @Summary 获取权限列表
// @Description 获取权限列表
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param use body types.RolePageListReq true "RolePageListReq"
// @Success 200 {object} []domain.Role{}
// @Failure 500 {object} ginx.JsonResult{}
// @Router /api/v1/roles [get]
func (r *roleController) RolePageList(ctx *gin.Context, req types.RolePageListReq) (result ginx.JsonResult, err error) {
	return ginx.Success(), nil
}
