package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/types/resp"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RoleController struct {
	svc service.RoleSvc
}

func NewRoleController(svc service.RoleSvc) *RoleController {
	return &RoleController{
		svc: svc,
	}
}

// RolePageList godoc
// @Security ApiKeyAuth
// @Summary 获取角色列表
// @Description 获取角色列表
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param use body types.RolePageListReq true "RolePageListReq"
// @Success 200 {object} []domain.Role{}
// @Failure 500 {object} ginx.JsonResult{}
// @Router /api/v1/roles [get]
func (r *RoleController) RolePageList(ctx *gin.Context, req types.RolePageListReq) (result ginx.JsonResult, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 || req.Size > 50 {
		req.Page = 10
	}
	role, total, err := r.svc.PageListRole(ctx.Request.Context(), req.Keywords, req.Page, req.Size)
	if err != nil {
		return ginx.Error(10009, err.Error()), err
	}
	return ginx.SuccessJson(resp.RoleListResp{
		List:  role,
		Total: int(total),
	}), nil
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
func (r *RoleController) CreateRoleCtl(ctx *gin.Context, req types.CreateRoleReq) (result ginx.JsonResult, err error) {
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
func (r *RoleController) DetailRoleCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
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
func (r *RoleController) DeleteRoleCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
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
func (r *RoleController) UpdateRoleCtl(ctx *gin.Context, req types.UpdateRoleReq) (result ginx.JsonResult, err error) {
	err = r.svc.UpdateRole(ctx.Request.Context(), req)
	if err != nil {
		return ginx.Error(10013, err.Error()), err
	}
	return ginx.Success(), nil
}
