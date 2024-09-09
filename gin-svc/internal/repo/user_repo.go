package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/types"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../repo/mocks/mock_user_repo.go -package=mocks gin-svc/internal/repo UserRepoInterface
type UserRepoInterface interface {
	UpsertUser(ctx context.Context, user types.UpdateUserForm) error
	DeleteUser(ctx context.Context, user models.UserModel) error
	CreateUser(ctx context.Context, user models.UserModel) error
	FindByUserName(ctx context.Context, userName string) (models.UserModel, error)
	FindByEmail(ctx context.Context, email string) (models.UserModel, error)
	FindByUUID(ctx context.Context, uuid string) (domain.DUser, error)
}

func NewUserRepoInterface(userDao dao.UserDao) UserRepoInterface {
	return &userRepo{dao: userDao}
}

type userRepo struct {
	dao dao.UserDao
}

func (u *userRepo) FindByUUID(ctx context.Context, uuid string) (domain.DUser, error) {
	user, err := u.dao.GetByID(ctx, uuid)
	if err != nil {
		return domain.DUser{}, err
	}
	return u.userModelToDM(user), nil
}

func (u *userRepo) UpsertUser(ctx context.Context, user types.UpdateUserForm) error {
	userModel := models.UserModel{
		Model: gorm.Model{
			ID: uint(user.ID),
		},
		NickName: user.NickName,
		//IPAddr:        "",
		Avatar: user.Avatar,
		//Signature:     "",
		//FansCount:     0,
		//FollowerCount: 0,
		//LikeCntCount:  0,
		//Age:           0,
		//Male:          "",
		//BirthDay:      nil,
		//Addr:          "",
		//Profession:    "",
		//School:        "",
		//Level:         0,
	}
	return u.dao.Upsert(ctx, userModel, user.RoleIds)
}

func (u *userRepo) DeleteUser(ctx context.Context, user models.UserModel) error {
	return u.dao.Delete(ctx, int(user.ID))
}

func (u *userRepo) FindByUserName(ctx context.Context, userName string) (models.UserModel, error) {
	return u.dao.FindByUserName(ctx, userName)
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (models.UserModel, error) {
	return u.dao.FindByEmail(ctx, email)
}

func (u *userRepo) CreateUser(ctx context.Context, user models.UserModel) error {
	//return nil
	return u.dao.Upsert(ctx, user, []int{})
}

func (u *userRepo) userModelToDM(data models.UserModel) domain.DUser {
	return domain.DUser{
		GlobalNumber:  data.GlobalNumber,
		UserName:      data.UserName,
		NickName:      data.NickName,
		Avatar:        data.Avatar,
		Email:         data.Email,
		IPAddr:        data.IPAddr,
		FansCount:     data.FansCount,
		FollowerCount: data.FollowerCount,
		Password:      data.Password,
		Signature:     data.Signature,
	}
}
