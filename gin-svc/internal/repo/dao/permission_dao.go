package dao

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

// 确保实现接口，如果没有实现接口，编译期间会报错
var _ PermissionDAO = (*permissionDao)(nil)

type PermissionDAO interface {
	// CheckPermissionIds 检查多个权限ID是否存在
	CheckPermissionIds(ctx context.Context, perIds []int) error
	// Insert 创建权限
	Insert(ctx context.Context, data models.SysPermissionModel) error
	// ListByRoleId 根据角色ID获取权限列表
	ListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error)
	// GetByID 通过id获取权限信息
	GetByID(ctx context.Context, id int) (*models.SysPermissionModel, error)
}

func NewPermissionDAO(db *gorm.DB) PermissionDAO {
	return &permissionDao{
		db: db,
	}
}

type permissionDao struct {
	db *gorm.DB
}

func (p *permissionDao) CheckPermissionIds(ctx context.Context, perIds []int) error {
	var per []models.SysPermissionModel
	err := p.db.WithContext(ctx).Where("id in ? and deleted_at is null", perIds).Find(&per).Error
	if err != nil {
		return err
	}
	if len(per) != len(perIds) {
		return err
	}
	return nil
}

func (p *permissionDao) Insert(ctx context.Context, data models.SysPermissionModel) error {
	return p.db.WithContext(ctx).Create(&data).Error
}

func (p *permissionDao) ListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
	var res []models.SysPermissionModel
	err := p.db.WithContext(ctx).
		Joins("INNER JOIN sys_role_permission rp ON rp.permission_id = sys_permission.id").
		Where("rp.role_id = ?", roleId).
		Find(&res).Error
	return res, err
}

func (p *permissionDao) GetByID(ctx context.Context, id int) (*models.SysPermissionModel, error) {
	var res models.SysPermissionModel
	query := p.db.WithContext(ctx).Where("id = ?", id).First(&res)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &res, nil
}
