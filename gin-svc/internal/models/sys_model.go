package models

import "gorm.io/gorm"

type SysPermissionModel struct {
	gorm.Model
	Title  string `json:"title" gorm:"column:title;type:varchar(100);not null;default:'';comment:权限名称"`
	Action string `json:"action" gorm:"column:action;type:varchar(32);not null;default:'';comment:权限动作,对应HTTP请求方式"`
	PerKey string `json:"perKey" gorm:"unique;column:per_key;type:varchar(64);not null;default:'';comment:权限关键字(唯一索引)"`
	Mark   string `json:"mark" gorm:"column:mark;type:varchar(100);not null;default:'';comment:权限标识"`
	Path   string `json:"path" gorm:"column:path;type:varchar(255);not null;default:'';comment:权限路径"`
	Type   string `json:"type" gorm:"column:type;type:varchar(32);not null;default:'BUS';comment:权限类型 BUS/SYS"`
	Status bool   `json:"status" gorm:"column:status;type:tinyint(1);not null;default:1;comment:状态,0 禁用，1启用"`
}

func (SysPermissionModel) TableName() string {
	return "sys_permission"
}

type SysRoleModel struct {
	gorm.Model
	CreatorId int64  `gorm:"not null;type:int;column:creator_id;comment:创建人ID"`
	RoleName  string `gorm:"not null;type:varchar(32);column:role_name;comment:角色名"`
	RoleKey   string `gorm:"not null;unique;type:varchar(100);column:role_key;comment:角色字符串(唯一索引,不可更改)"`
	Status    bool   `gorm:"comment:角色状态"`
	RoleSort  int    `gorm:"comment:角色排序"`
	Remark    string `gorm:"type:varchar(100);comment:角色描述信息"`
}

func (SysRoleModel) TableName() string {
	return "sys_role"
}

type SysMenuModel struct {
}

func (SysMenuModel) TableName() string {
	return "sys_menu"
}

type UserRoleModel struct {
	UserID int `json:"userId" gorm:"column:user_id;type:int;comment:用户ID"`
	RoleID int `json:"roleId" gorm:"column:role_id;type:int;comment:角色ID"`
}

func (UserRoleModel) TableName() string {
	return "sys_user_role"
}

type RolePermissionModel struct {
	RoleID int `json:"roleId" gorm:"column:role_id;type:int;comment:角色ID"`
	PerID  int `json:"perId" gorm:"column:per_id;type:int;comment:权限ID"`
}

func (RolePermissionModel) TableName() string {
	return "sys_role_permission"
}
