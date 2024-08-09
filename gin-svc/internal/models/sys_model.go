package models

import "gorm.io/gorm"

type SysPermissionModel struct {
}

func (SysPermissionModel) TableName() string {
	return "sys_permission"
}

type SysRoleModel struct {
	gorm.Model
	RoleName string `json:"roleName" gorm:"type:varchar(128);not null;comment:角色名称"`
	RoleKey  string `json:"roleKey" gorm:"type:varchar(128);not null;comment:角色标识"`
	Sort     int    `json:"sort" gorm:"type:int;not null;comment:排序"`
}

func (SysRoleModel) TableName() string {
	return "sys_role"
}

type SysMenuModel struct {
}

func (SysMenuModel) TableName() string {
	return "sys_menu"
}
