package service

import (
	"context"
	"gin-svc/internal/types"
)

type UserSvc interface {
	UpsertUser(ctx context.Context, req types.UserForm) error
}

type userSvcImpl struct {
}

func (u *userSvcImpl) UpsertUser(ctx context.Context, req types.UserForm) error {
	//TODO implement me
	panic("implement me")
}
