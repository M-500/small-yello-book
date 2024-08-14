package service

import (
	"context"
	"errors"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/utils"
	"time"
)

type UserSvc interface {
	EmailLogin(ctx context.Context, req types.EmailLoginForm) (string, error)
	DeleteUser(ctx context.Context, req types.UserForm) error
	PwdLogin(ctx context.Context, req types.LoginForm) error
	GetUserInfo(ctx context.Context, id int) (domain.DUser, error)
}

func NewUserSvc(userRepo repo.UserRepoInterface,
	verCache cache.VerificationCodeCache,
	hdl jwt.Handler,
) UserSvc {
	return &userSvcImpl{
		userRepo: userRepo,
		verCache: verCache,
		Handler:  hdl,
	}
}

type userSvcImpl struct {
	jwt.Handler
	userRepo repo.UserRepoInterface
	verCache cache.VerificationCodeCache
}

func (u *userSvcImpl) EmailLogin(ctx context.Context, req types.EmailLoginForm) (string, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, dao.ErrRecordNotFound) {
		return "", err
	}
	// 校验验证码是否正确
	code, err := u.verCache.GetVerificationCode(ctx, req.Email)
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
			BirthDay:      time.Time{},
			Addr:          "",
			Profession:    "",
			School:        "",
			Level:         1,
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
	return tokenStr, nil
}

func (u *userSvcImpl) DeleteUser(ctx context.Context, req types.UserForm) error {
	return u.userRepo.DeleteUser(ctx, u.reqToModel(&req))
}

func (u *userSvcImpl) reqToModel(req *types.UserForm) *models.UserModel {
	return &models.UserModel{
		UserName:   req.UserName,
		NickName:   req.NickName,
		Avatar:     req.Avatar,
		Email:      req.Email,
		Password:   req.Password,
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
