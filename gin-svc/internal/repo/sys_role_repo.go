package repo

import (
	"context"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
)

//go:generate mockgen -destination=../repo/mocks/mock_role_repo.go -package=mocks -source=./sys_role_repo.go RoleRepo
type RoleRepo interface {
	// CheckPermissionIds 批量检查权限 ID 是否存在
	CheckPermissionIds(ctx context.Context, perIds []int) error
	// GetRoleByID 通过角色ID获取角色信息
	GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error)
	// ListRole 获取角色列表分页信息
	ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error)
	// CreateRole 创建角色
	CreateRole(ctx context.Context, role *models.SysRoleModel) error
	// DeleteRole 删除角色
	DeleteRole(ctx context.Context, id int) error
	// UpdateRole 更新角色信息
	UpdateRole(ctx context.Context, role *models.SysRoleModel, perIds []int) error
	// FindPermissionListByRoleId 查询某个角色 ID 对应的权限列表
	FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error)
	// FindAllPermissionsByUserID 通过用户ID查询用户的所有权限
	FindAllPermissionsByUserID(ctx context.Context, uid int) ([]models.SysPermissionModel, error)
	// FindAllRolesByUserID 通过用户ID查询用户的所有角色
	FindAllRolesByUserID(ctx context.Context, uid int) ([]models.SysRoleModel, error)
}

type roleRepoImpl struct {
	perDao  dao.PermissionDAO
	roleDao dao.RoleDAO
	cache   cache.RoleCache
}

func (r *roleRepoImpl) CheckPermissionIds(ctx context.Context, perIds []int) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepoImpl) GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error) {
	res, err := r.roleDao.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *roleRepoImpl) ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepoImpl) CreateRole(ctx context.Context, role *models.SysRoleModel) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepoImpl) DeleteRole(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepoImpl) UpdateRole(ctx context.Context, role *models.SysRoleModel, perIds []int) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepoImpl) FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
	return r.perDao.ListByRoleId(ctx, roleId)
}

func (r *roleRepoImpl) FindAllPermissionsByUserID(ctx context.Context, uid int) ([]models.SysPermissionModel, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepoImpl) FindAllRolesByUserID(ctx context.Context, uid int) ([]models.SysRoleModel, error) {
	//TODO implement me
	panic("implement me")
}

func NewRoleRepo(pDao dao.PermissionDAO, rDao dao.RoleDAO, cache cache.RoleCache) RoleRepo {
	return &roleRepoImpl{
		perDao:  pDao,
		roleDao: rDao,
		cache:   cache,
	}
}

//
//func (r *roleRepoImpl) CheckPermissionIds(ctx context.Context, perIds []int) error {
//	return r.perDao.CheckPermissionIds(ctx, perIds)
//}
//func (r *roleRepoImpl) GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error) {
//	return r.roleDao.GetRoleByID(ctx, id)
//}
//
//func (r *roleRepoImpl) ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error) {
//	return r.roleDao.ListRole(ctx, key, page, pageSize)
//}
//
//func (r *roleRepoImpl) CreateRole(ctx context.Context, role *models.SysRoleModel) error {
//	return r.roleDao.Insert(ctx, role)
//}
//
//func (r *roleRepoImpl) DeleteRole(ctx context.Context, id int) error {
//	go func() {
//		_ = r.cache.DeletePerListByRoleId(context.Background(), id)
//		// TODO: 如果缓存删除数据失败，要记录日志
//	}()
//	return r.roleDao.DeleteByID(ctx, id)
//}
//
//func (r *roleRepoImpl) UpdateRole(ctx context.Context, role *models.SysRoleModel, perIds []int) error {
//
//	err := r.roleDao.UpdateRole(ctx, role, perIds)
//	if err != nil {
//		return err
//	}
//	go func() {
//		// TODO: 删除缓存要记录日志
//		_ = r.cache.DeletePerListByRoleId(context.Background(), int(role.ID))
//	}()
//	return nil
//}
//
//func (r *roleRepoImpl) FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
//	res, err := r.cache.GetPerListByRoleId(ctx, roleId)
//	if err == nil {
//		// 如果从缓存中读取到了数据就直接返回
//		return res, nil
//	}
//	res, err = r.roleDao.FindPermissionListByRoleId(ctx, roleId)
//	go func() {
//		_ = r.cache.SetPerListByRoleId(context.Background(), roleId, res)
//		// TODO: 异步设置缓存，如果出错要记录一下日志
//	}()
//	return res, err
//}
