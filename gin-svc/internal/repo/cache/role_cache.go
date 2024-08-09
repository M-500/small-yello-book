package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-svc/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type RoleCache interface {
	GetPerListByRoleId(ctx context.Context, id int) ([]models.SysPermissionModel, error)
	DeletePerListByRoleId(ctx context.Context, id int) error
	SetPerListByRoleId(ctx context.Context, id int, data []models.SysPermissionModel) error
}

type roleCacheImpl struct {
	cli     redis.Cmdable
	timeout time.Duration
}

func (r *roleCacheImpl) key(roleId int) string {
	return fmt.Sprintf("role:per:list:%d", roleId)
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

func NewRoleCache(cli redis.Cmdable) RoleCache {
	return &roleCacheImpl{
		timeout: time.Minute * 30,
		cli:     cli}
}
