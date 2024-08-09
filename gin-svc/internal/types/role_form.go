package types

// RolePageListReq
// @Description:	获取权限列表
type RolePageListReq struct {
	Page     int    `json:"page" form:"page"`                             // 页码
	Size     int    `json:"size" form:"size"`                             // 页大小
	Keywords string `json:"keywords" form:"keywords" binding:"omitempty"` // 搜索关键词: 角色名，角色key，用于模糊搜索
	Status   bool   `json:"status" form:"status" binding:"omitempty"`     // 角色状态
}
