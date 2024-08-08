package service

import (
	"context"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/types"
)

type UserSvc interface {
	RegisterUser(ctx context.Context, req types.UserForm) error
}

func NewUserSvc(userRepo repo.UserRepoInterface) UserSvc {
	return &userSvcImpl{userRepo: userRepo}
}

type userSvcImpl struct {
	userRepo repo.UserRepoInterface
}

func (u *userSvcImpl) RegisterUser(ctx context.Context, req types.UserForm) error {
	return u.userRepo.UpsertUser(ctx, u.reqToModel(&req))
}

func (u *userSvcImpl) reqToModel(req *types.UserForm) *models.UserModel {
	return &models.UserModel{
		UserName:   req.UserName,
		NickName:   req.NickName,
		Avatar:     req.Avatar,
		Password:   req.Password,
		Phone:      req.Phone,
		Signature:  req.Signature,
		Male:       req.Male,
		BirthDay:   req.BirthDay,
		Addr:       req.Addr,
		Profession: req.Profession,
		School:     req.School,
	}
}
