package types

// RolePageListReq
// @Description:	获取权限列表
type RolePageListReq struct {
	Page     int    `json:"page" form:"page"`                             // 页码
	Size     int    `json:"size" form:"size"`                             // 页大小
	Keywords string `json:"keywords" form:"keywords" binding:"omitempty"` // 搜索关键词: 角色名，角色key，用于模糊搜索
	Status   bool   `json:"status" form:"status" binding:"omitempty"`     // 角色状态
}

type CreateRoleReq struct {
	RoleName string `json:"roleName" form:"roleName" binding:"required"` // 角色名
	RoleKey  string `json:"roleKey" form:"roleKey" binding:"required"`   // 角色key
	Status   bool   `json:"status" form:"status" binding:"required"`     // 角色状态
	Sort     int    `json:"sort" form:"sort" binding:"omitempty"`        // 排序
	PerIds   []int  `json:"perIds" form:"perIds" binding:"omitempty"`    // 权限ID列表
}

type UpdateRoleReq struct {
	ID       int    `json:"id" form:"id" binding:"required"`             // 角色ID
	RoleName string `json:"roleName" form:"roleName" binding:"required"` // 角色名
	Status   bool   `json:"status" form:"status" binding:"omitempty"`    // 角色状态
	Sort     int    `json:"sort" form:"sort" binding:"omitempty"`        // 排序
	PerIds   []int  `json:"perIds" form:"perIds" binding:"omitempty"`    // 权限ID列表
}
