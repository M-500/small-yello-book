package repo

import (
	"context"
	"gin-svc/internal/models"
)

type SysPermissionRepoInterface interface {
	// GetSysPermissionByID 根据ID获取SysPermission
	GetSysPermissionByID(ctx context.Context, id uint) (*models.SysPermissionModel, error)
}
