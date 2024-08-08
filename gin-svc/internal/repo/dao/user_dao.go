package dao

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDao interface {
	Upsert(ctx context.Context, user *models.UserModel) error
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDaoImpl{db: db}
}

type userDaoImpl struct {
	db *gorm.DB
}

func (u *userDaoImpl) Upsert(ctx context.Context, user *models.UserModel) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return u.db.WithContext(ctx).Clauses(clause.OnConflict{
		// 当发生冲突时，更新所有字段
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"user_name":      user.UserName,
			"nick_name":      user.NickName,
			"avatar":         user.Avatar,
			"password":       user.Password,
			"phone":          user.Phone,
			"signature":      user.Signature,
			"ip_addr":        user.IPAddr,
			"fans_count":     user.FansCount,
			"follower_count": user.FollowerCount,
			"like_cnt_count": user.LikeCntCount,
			"age":            user.Age,
			"male":           user.Male,
			"birth_day":      user.BirthDay,
			"addr":           user.Addr,
			"profession":     user.Profession,
			"school":         user.School,
			"level":          user.Level,
		}),
	}).Create(user).Error
}
