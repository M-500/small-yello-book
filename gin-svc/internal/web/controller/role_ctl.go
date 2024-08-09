package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
	"strconv"
)

type roleController struct {
	svc service.RoleServiceInterface
}

func NewRoleController(svc service.RoleServiceInterface) BaseController {
	return &roleController{
		svc: svc,
	}
}

func (r *roleController) RegisterRoute(group *gin.RouterGroup) {
	group.GET("/roles", ginx.WrapJsonBody[types.RolePageListReq](r.RolePageList))
	group.GET("/roles/:id", ginx.WrapResponse(r.DetailRoleCtl))
	group.POST("/roles", ginx.WrapJsonBody[types.CreateRoleReq](r.CreateRoleCtl))
	group.DELETE("/roles/:id", ginx.WrapResponse(r.DeleteRoleCtl))
	group.PUT("/roles", ginx.WrapJsonBody[types.UpdateRoleReq](r.UpdateRoleCtl))
}

// RolePageList godoc
// @Summary 获取角色列表
// @Description 获取角色列表
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

// CreateRoleCtl godoc
// @Summary 创建角色
// @Description 创建角色
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param input body types.CreateRoleReq true "CreateRoleReq"
// @Success 200 {object} ginx.JsonResult{}
// @Router /api/v1/roles [post]
func (r *roleController) CreateRoleCtl(ctx *gin.Context, req types.CreateRoleReq) (result ginx.JsonResult, err error) {
	err = r.svc.CreateRole(ctx.Request.Context(), req)
	if err != nil {
		return ginx.Error(10010, err.Error()), err
	}
	return ginx.Success(), nil
}

// DetailRoleCtl godoc
// @Summary 查询角色详情
// @Description 查询角色详情
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} domain.RoleDetail
// @Router /api/v1/roles/{id} [get]
func (r *roleController) DetailRoleCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ginx.JsonResult{Msg: "id 不能为空"}, err
	}
	detail, err := r.svc.GetDetailByID(ctx.Request.Context(), id)
	if err != nil {
		return ginx.Error(10011, err.Error()), err
	}
	return ginx.SuccessJson(detail), nil
}

// DeleteRoleCtl godoc
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} ginx.JsonResult{}
// @Router /api/v1/roles/{id} [delete]
func (r *roleController) DeleteRoleCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ginx.JsonResult{Msg: "id 不能为空"}, err
	}
	err = r.svc.DeleteRole(ctx.Request.Context(), id)
	if err != nil {
		return ginx.Error(10012, err.Error()), err
	}
	return ginx.Success(), nil
}

// UpdateRoleCtl godoc
// @Summary 更新角色
// @Description 更新角色
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param input body types.UpdateRoleReq true "UpdateRoleReq"
// @Success 200 {object} ginx.JsonResult{}
// @Router /api/v1/roles [put]
func (r *roleController) UpdateRoleCtl(ctx *gin.Context, req types.UpdateRoleReq) (result ginx.JsonResult, err error) {
	err = r.svc.UpdateRole(ctx.Request.Context(), req)
	if err != nil {
		return ginx.Error(10013, err.Error()), err
	}
	return ginx.Success(), nil
}
