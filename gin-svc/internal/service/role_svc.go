package service

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/types"
)

type RoleServiceInterface interface {
	GetDetailByID(ctx context.Context, id int) (domain.RoleDetail, error)
	CreateRole(ctx context.Context, role types.CreateRoleReq) error
	DeleteRole(ctx context.Context, id int) error
	UpdateRole(ctx context.Context, role types.UpdateRoleReq) error
}

func NewRoleService() RoleServiceInterface {
	return nil
}
