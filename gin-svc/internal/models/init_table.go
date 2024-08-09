package models

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&UserModel{},
		&SysRoleModel{},
		&SysPermissionModel{},
		&UserRoleModel{},
		&RolePermissionModel{},
	)
}
