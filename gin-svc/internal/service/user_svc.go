package service

import (
	"context"
	"errors"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/utils"
)

type UserSvc interface {
	EmailLogin(ctx context.Context, req types.EmailLoginForm) (string, error)
	DeleteUser(ctx context.Context, req types.UserForm) error
	GetUserInfo(ctx context.Context, id int) (domain.DUser, error)
}

func NewUserSvc(userRepo repo.UserRepoInterface,
	verRepo repo.VerificationCodeRepo,
	hdl jwt.Handler,
) UserSvc {
	return &userSvcImpl{
		userRepo: userRepo,
		verRepo:  verRepo,
		Handler:  hdl,
	}
}

type userSvcImpl struct {
	jwt.Handler
	userRepo repo.UserRepoInterface
	verRepo  repo.VerificationCodeRepo
}

func (u *userSvcImpl) EmailLogin(ctx context.Context, req types.EmailLoginForm) (string, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, dao.ErrRecordNotFound) {
		return "", err
	}
	// 校验验证码是否正确
	code, err := u.verRepo.GetVerificationCode(ctx, req.Email)
	if err != nil {
		return "", errors.New("验证码错误")
	}
	if code != req.VerCode {
		return "", errors.New("验证码验证失败")
	}
	if user == nil {
		// 记录日志
		user = &models.UserModel{
			GlobalNumber:  utils.UUID(),
			Email:         req.Email,
			NickName:      utils.RandNickName(),
			IPAddr:        "",
			Avatar:        "",
			Password:      "",
			Signature:     "这家伙很懒，什么都没留下~",
			FansCount:     0,
			FollowerCount: 0,
			LikeCntCount:  0,
			Age:           0,
			Male:          "",
			//BirthDay:      &time.Time{},
			Addr:       "",
			Profession: "",
			School:     "",
			Level:      1,
		}
		// 初始化用户逻辑
		err = u.userRepo.CreateUser(ctx, user)
		if err != nil {
			return "", err
		}
	}
	// 处理登录逻辑
	tokenStr, err := u.GenJWTToken(int64(user.ID))
	if err != nil {
		return "", err
	}

	// 删除验证码
	_ = u.verRepo.DeleteVerificationCode(ctx, req.Email)
	return tokenStr, nil
}

func (u *userSvcImpl) DeleteUser(ctx context.Context, req types.UserForm) error {
	return u.userRepo.DeleteUser(ctx, u.reqToModel(&req))
}

func (u *userSvcImpl) reqToModel(req *types.UserForm) *models.UserModel {
	return &models.UserModel{
		UserName:  req.UserName,
		NickName:  req.NickName,
		Avatar:    req.Avatar,
		Email:     req.Email,
		Password:  req.Password,
		Signature: req.Signature,
		Male:      req.Male,
		//BirthDay:   &req.BirthDay,
		Addr:       req.Addr,
		Profession: req.Profession,
		School:     req.School,
	}
}

func (u *userSvcImpl) GetUserInfo(ctx context.Context, id int) (domain.DUser, error) {
	return domain.DUser{}, nil
}
