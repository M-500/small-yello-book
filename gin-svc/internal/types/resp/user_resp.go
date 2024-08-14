package resp

import "gin-svc/internal/models"

// UserInfoResp 用户权限信息返回
type UserInfoResp struct {
	AllPerList []models.SysPermissionModel `json:"all_per_list"` // 系统所有权限列表
	Rights     []string                    `json:"rights"`       // 用户拥有的权限key
}
