package dao

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type RoleDAO interface {
	// GetRoleByID 通过角色ID获取角色信息
	GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error)

	ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error)

	Insert(ctx context.Context, role *models.SysRoleModel) error

	DeleteByID(ctx context.Context, id int) error

	UpdateRole(ctx context.Context, role *models.SysRoleModel, perIds []int) error
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
		// 记录日志，不要把错误抛出给上层
		return nil, errors.New("DB查询角色信息失败")
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
	query := r.db.WithContext(ctx).Model(&models.SysRoleModel{}).Create(role)
	var me *mysql.MySQLError
	if errors.As(query.Error, &me) {
		const uniqueIndexErrNo uint16 = 1062
		if me.Number == uniqueIndexErrNo {
			return ErrKeyDuplicate
		}
	}
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *roleDaoImpl) DeleteByID(ctx context.Context, id int) error {
	// TODO：删除角色，是不是要删除对应的角色权限关联表 是吧
	return r.db.WithContext(ctx).Model(&models.SysRoleModel{}).Delete("id = ?", id).Error
}

func (r *roleDaoImpl) UpdateRole(ctx context.Context, role *models.SysRoleModel, perIds []int) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&models.SysRoleModel{}).Where("id=?", role.ID).Updates(role).Error
		if err != nil {
			return err
		}
		err = tx.Model(&models.RolePermissionModel{}).Where("role_id=?", role.ID).Delete("role_id=?", role.ID).Error
		if err != nil {
			return err
		}
		for _, id := range perIds {
			err = tx.Model(&models.RolePermissionModel{}).Create(&models.RolePermissionModel{
				RoleID: int(role.ID),
				PerID:  id,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *roleDaoImpl) FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
	var perList []models.SysPermissionModel
	err := r.db.WithContext(ctx).Table("sys_permission").
		Select("sys_permission.*").
		Joins("LEFT JOIN sys_role_permission ON sys_permission.id = sys_role_permission.per_id").
		Where("sys_role_permission.role_id = ?", roleId).Find(&perList).Error
	return perList, err
}
