package repo

import (
	"context"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
)

type RoleRepo interface {
	// GetRoleByID 通过角色ID获取角色信息
	GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error)

	ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error)

	CreateRole(ctx context.Context, role *models.SysRoleModel) error

	DeleteRole(ctx context.Context, id int) error

	UpdateRole(ctx context.Context, role *models.SysRoleModel) error
	// FindPermissionListByRoleId 查询某个角色 ID 对应的权限列表
	FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error)
}

type roleRepoImpl struct {
	perDao  dao.PermissionDAO
	roleDao dao.RoleDAO
	cache   cache.RoleCache
}

func NewRoleRepo(pDao dao.PermissionDAO, rDao dao.RoleDAO, cache cache.RoleCache) RoleRepo {
	return &roleRepoImpl{
		perDao:  pDao,
		roleDao: rDao,
		cache:   cache,
	}
}

func (r *roleRepoImpl) GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error) {
	return r.roleDao.GetRoleByID(ctx, id)
}

func (r *roleRepoImpl) ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error) {
	return r.roleDao.ListRole(ctx, key, page, pageSize)
}

func (r *roleRepoImpl) CreateRole(ctx context.Context, role *models.SysRoleModel) error {
	return r.roleDao.Insert(ctx, role)
}

func (r *roleRepoImpl) DeleteRole(ctx context.Context, id int) error {
	go func() {
		_ = r.cache.DeletePerListByRoleId(context.Background(), id)
		// TODO: 如果缓存删除数据失败，要记录日志
	}()
	return r.roleDao.DeleteByID(ctx, id)
}

func (r *roleRepoImpl) UpdateRole(ctx context.Context, role *models.SysRoleModel) error {

	err := r.roleDao.UpdateRole(ctx, role)
	if err != nil {
		return err
	}
	go func() {
		// TODO: 删除缓存要记录日志
		_ = r.cache.DeletePerListByRoleId(context.Background(), int(role.ID))
	}()
	return nil
}

func (r *roleRepoImpl) FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
	res, err := r.cache.GetPerListByRoleId(ctx, roleId)
	if err == nil {
		// 如果从缓存中读取到了数据就直接返回
		return res, nil
	}
	res, err = r.roleDao.FindPermissionListByRoleId(ctx, roleId)
	go func() {
		_ = r.cache.SetPerListByRoleId(ctx, roleId, res)
		// TODO: 异步设置缓存，如果出错要记录一下日志
	}()
	return res, err
}
