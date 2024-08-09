package dao

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type RoleDAO interface {
	// GetRoleByID 通过角色ID获取角色信息
	GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error)

	ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error)

	Insert(ctx context.Context, role *models.SysRoleModel) error

	DeleteByID(ctx context.Context, id int) error

	UpdateRole(ctx context.Context, role *models.SysRoleModel) error
	// FindPermissionListByRoleId 查询某个角色 ID 对应的权限列表
	FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error)
}

func NewRoleDAO(db *gorm.DB) RoleDAO {
	return &roleDaoImpl{db: db}
}

type roleDaoImpl struct {
	db *gorm.DB
}

func (r *roleDaoImpl) GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error) {
	var res models.SysRoleModel
	query := r.db.WithContext(ctx).Model(&models.SysRoleModel{}).Where("id = ?", id).First(&res)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &res, nil
}

func (r *roleDaoImpl) ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error) {
	var (
		total int64
		res   []models.SysRoleModel
	)
	err := r.db.WithContext(ctx).Model(&models.SysRoleModel{}).
		Where("role_key like ?", "%"+key+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.WithContext(ctx).Model(&models.SysRoleModel{}).Where("role_key like ?", "%"+key+"%").
		Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (r *roleDaoImpl) Insert(ctx context.Context, role *models.SysRoleModel) error {
	return r.db.WithContext(ctx).Model(&models.SysRoleModel{}).Create(role).Error
}

func (r *roleDaoImpl) DeleteByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleDaoImpl) UpdateRole(ctx context.Context, role *models.SysRoleModel) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleDaoImpl) FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
	var perList []models.SysPermissionModel
	err := r.db.WithContext(ctx).Table("sys_permission").
		Select("sys_permission.*").
		Joins("INNER JOIN sys_role_permission ON sys_permission.id = sys_role_permission.permission_id").
		Where("sys_role_permission.role_id = ?", roleId).Find(&perList).Error
	return perList, err
}
