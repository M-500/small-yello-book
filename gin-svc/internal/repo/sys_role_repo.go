package repo

import (
	"context"
	"errors"
	"fmt"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
	"golang.org/x/sync/errgroup"
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
	return r.perDao.CheckPermissionIds(ctx, perIds)
}

func (r *roleRepoImpl) GetRoleByID(ctx context.Context, id int) (*models.SysRoleModel, error) {
	// 从缓存中取
	info, err := r.cache.GetRoleInfo(ctx, id)
	if err == nil {
		return &info, nil
	}

	// 缓存miss 去DB中取
	res, err := r.roleDao.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 回写缓存
	err = r.cache.SetRoleInfo(ctx, res)
	if err != nil {
		// TODO: 记录日志，回写缓存失败
	}
	return res, nil
}

func (r *roleRepoImpl) FindPermissionListByRoleId(ctx context.Context, roleId int) ([]models.SysPermissionModel, error) {
	// 从缓存中查询角色对应的权限列表
	res, err := r.cache.GetPerListByRoleId(ctx, roleId)
	if err == nil {
		return res, err
	}
	// 缓存miss 去DB中读取
	res, err = r.perDao.ListByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}
	// 回写缓存
	err = r.cache.SetPerListByRoleId(ctx, roleId, res)
	if err != nil {
		// TODO: 记录日志，回写缓存失败
	}
	return res, nil
}

func (r *roleRepoImpl) ListRole(ctx context.Context, key string, page, pageSize int) ([]models.SysRoleModel, int64, error) {
	role, i, err := r.roleDao.ListRole(ctx, key, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	for _, item := range role {
		_ = r.cache.SetRoleInfo(ctx, &item)
	}
	return role, i, nil
}

func (r *roleRepoImpl) CreateRole(ctx context.Context, role *models.SysRoleModel) error {
	err := r.roleDao.Insert(ctx, role)
	if errors.Is(err, dao.ErrKeyDuplicate) {
		return errors.New("roleKey 已存在，创建角色失败")
	}
	return errors.New("创建角色失败")
}

func (r *roleRepoImpl) DeleteRole(ctx context.Context, id int) error {
	err2 := r.roleDao.DeleteByID(ctx, id)
	if err2 != nil {
		return err2
	}
	group := errgroup.Group{}
	group.Go(func() error {
		return r.cache.DelRoleInfo(context.Background(), id)
	})
	group.Go(func() error {
		return r.cache.DeletePerListByRoleId(context.Background(), id)
	})
	err := group.Wait()
	if err != nil {
		// TODO 记录日志
		fmt.Println("删除缓存失败啦！阿西！")
	}
	// 缓存删除失败，无伤大雅，不会影响业务
	return nil
}

func (r *roleRepoImpl) UpdateRole(ctx context.Context, role *models.SysRoleModel, perIds []int) error {
	err := r.roleDao.UpdateRole(ctx, role, perIds)
	if err != nil {
		return err
	}
	// 更新的时候只会删除缓存，并不加载缓存，尽可能保证数据
	return r.cache.DeletePerListByRoleId(ctx, int(role.ID))
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
