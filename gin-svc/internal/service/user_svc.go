package service

import (
	"context"
	"errors"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/types"
)

type UserSvc interface {
	RegisterUser(ctx context.Context, req types.UserForm) error
	DeleteUser(ctx context.Context, req types.UserForm) error
	PwdLogin(ctx context.Context, req types.LoginForm) error
	GetUserInfo(ctx context.Context, id int) (domain.DUser, error)
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

func (u *userSvcImpl) DeleteUser(ctx context.Context, req types.UserForm) error {
	return u.userRepo.DeleteUser(ctx, u.reqToModel(&req))
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

func (u *userSvcImpl) PwdLogin(ctx context.Context, req types.LoginForm) error {
	user, err := u.userRepo.FindByUserName(ctx, req.UserName)
	if errors.Is(err, dao.ErrRecordNotFound) {
		return errors.New("用户不存在")
	}
	if err != nil {
		return errors.New("查询用户失败")
	}
	if user.Password != req.Password {
		return errors.New("密码错误")
	}
	// 组合Jwt信息，返回Token
	return nil
}

func (u *userSvcImpl) GetUserInfo(ctx context.Context, id int) (domain.DUser, error) {
	return domain.DUser{}, nil
}
