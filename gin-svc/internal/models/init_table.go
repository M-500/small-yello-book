package models

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&UserModel{},
		&SysRoleModel{},
		&SysPermissionModel{},
		&UserRoleModel{},
		&RolePermissionModel{},

		&NoteModel{},
		&TagModel{},
		&NoteTagModel{},
		&ImageModel{},

		&CommentModel{},      // 评论表
		&InteractiveModel{},  // 互动信息统计表
		&NotificationModel{}, // 通知表
		&BrowseLogModel{},    // 浏览记录表
		&CollectModel{},      // 收藏记录表
		&LikeModel{},         // 点赞记录表
	)
}
