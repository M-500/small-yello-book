package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-svc/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type RoleCache interface {
	GetPerListByRoleId(ctx context.Context, id int) ([]models.SysPermissionModel, error)
	DeletePerListByRoleId(ctx context.Context, id int) error
	SetPerListByRoleId(ctx context.Context, id int, data []models.SysPermissionModel) error

	SetRoleInfo(ctx context.Context, data *models.SysRoleModel) error
	GetRoleInfo(ctx context.Context, id int) (models.SysRoleModel, error)
	DelRoleInfo(ctx context.Context, id int) error

	//SetRolePerList(ctx context.Context, roleId int, data []models.SysPermissionModel) error
	//GetRolePerList(ctx context.Context, roleId int) ([]models.SysPermissionModel, error)
	//DelRolePerList(ctx context.Context, roleId int) error
}

type roleCacheImpl struct {
	cli     redis.Cmdable
	timeout time.Duration
}

func (r *roleCacheImpl) key(roleId int) string {
	return fmt.Sprintf("role:per:list:%d", roleId)
}

func (r *roleCacheImpl) roleInfoKey(roleId int) string {
	return fmt.Sprintf("role:info:%d", roleId)
}

func (r *roleCacheImpl) GetPerListByRoleId(ctx context.Context, id int) ([]models.SysPermissionModel, error) {
	bytes, err := r.cli.Get(ctx, r.key(id)).Bytes()
	if err != nil {
		return nil, err
	}
	var data []models.SysPermissionModel
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		// TODO: 记录日志
		return nil, err
	}
	return data, nil
}

func (r *roleCacheImpl) DeletePerListByRoleId(ctx context.Context, id int) error {
	return r.cli.Del(ctx, r.key(id)).Err()
}

func (r *roleCacheImpl) SetPerListByRoleId(ctx context.Context, id int, data []models.SysPermissionModel) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.cli.Set(ctx, r.key(id), bytes, r.timeout).Err()
}

func (r *roleCacheImpl) SetRoleInfo(ctx context.Context, data *models.SysRoleModel) error {
	if data == nil {
		return errors.New("无法设置空的角色信息到缓存中")
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		// TODO: 记录日志
		return err
	}
	return r.cli.Set(ctx, r.roleInfoKey(int(data.ID)), bytes, r.timeout).Err()
}

func (r *roleCacheImpl) GetRoleInfo(ctx context.Context, id int) (models.SysRoleModel, error) {
	bytes, err := r.cli.Get(ctx, r.roleInfoKey(id)).Bytes()
	if err != nil {
		// TODO 记录日志
		return models.SysRoleModel{}, err
	}
	var data models.SysRoleModel
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		// TODO: 记录日志
		return models.SysRoleModel{}, err
	}
	return data, nil
}

func (r *roleCacheImpl) DelRoleInfo(ctx context.Context, id int) error {
	return r.cli.Del(ctx, r.roleInfoKey(id)).Err()
}

func NewRoleCache(cli redis.Cmdable) RoleCache {
	return &roleCacheImpl{
		timeout: time.Minute * 30,
		cli:     cli}
}
