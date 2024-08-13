package repo

import (
	"context"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
)

type UserRepoInterface interface {
	UpsertUser(ctx context.Context, user *models.UserModel) error
	DeleteUser(ctx context.Context, user *models.UserModel) error
	CreateUser(ctx context.Context, user *models.UserModel) error
	FindByUserName(ctx context.Context, userName string) (*models.UserModel, error)
}

func NewUserRepoInterface(userDao dao.UserDao) UserRepoInterface {
	return &userRepo{dao: userDao}
}

type userRepo struct {
	dao dao.UserDao
}

func (u *userRepo) UpsertUser(ctx context.Context, user *models.UserModel) error {
	return u.dao.Upsert(ctx, user)
}

func (u *userRepo) DeleteUser(ctx context.Context, user *models.UserModel) error {
	return u.dao.Delete(ctx, int(user.ID))
}

func (u *userRepo) FindByUserName(ctx context.Context, userName string) (*models.UserModel, error) {
	return u.dao.FindByUserName(ctx, userName)
}

func (u *userRepo) CreateUser(ctx context.Context, user *models.UserModel) error {
	// 创建用户，是依赖唯一索引冲突 还是说
	return nil
}
