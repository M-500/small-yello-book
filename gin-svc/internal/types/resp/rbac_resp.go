package resp

import "gin-svc/internal/models"

// RoleDetailResp 角色详情 包含角色基本信息和角色对应的权限列表
type RoleDetailResp struct {
	RoleInfo models.SysRoleModel         `json:"role_info"`
	PerList  []models.SysPermissionModel `json:"per_list"`
}

// RoleListResp 权限列表的翻页结果
type RoleListResp struct {
	List  []models.SysRoleModel `json:"list"`
	Total int                   `json:"total"`
}

// PermissionListResp 权限列表的翻页结果
type PermissionListResp struct {
	List  []models.SysPermissionModel `json:"list"`
	Total int                         `json:"total"`
}
