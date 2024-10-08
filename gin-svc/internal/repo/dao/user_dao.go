package dao

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -source=./user_dao.go -package=daomocks -destination=./mocks/user_dao.mock.go UserDao
type UserDao interface {
	Upsert(ctx context.Context, user models.UserModel, rid []int) error
	Delete(ctx context.Context, id int) error
	FindByUserName(ctx context.Context, userName string) (models.UserModel, error)
	FindByEmail(ctx context.Context, email string) (models.UserModel, error)
	GetByID(ctx context.Context, id string) (models.UserModel, error)
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDaoImpl{db: db}
}

type userDaoImpl struct {
	db *gorm.DB
}

func (u *userDaoImpl) GetByID(ctx context.Context, id string) (models.UserModel, error) {
	var user models.UserModel
	err := u.db.WithContext(ctx).Model(&models.UserModel{}).Where("global_number = ?", id).First(&user).Error
	if err != nil {
		return models.UserModel{}, err
	}
	return user, nil
}

func (u *userDaoImpl) Delete(ctx context.Context, id int) error {
	return u.db.WithContext(ctx).Delete(&models.UserModel{}, id).Error
}

func (u *userDaoImpl) Upsert(ctx context.Context, user models.UserModel, rid []int) error {
	//if user == nil {
	//	return errors.New("user is nil")
	//}
	return u.db.WithContext(ctx).Clauses(clause.OnConflict{
		// 当发生冲突时，更新所有字段
		Columns: []clause.Column{{Name: "email"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"user_name":      user.UserName,
			"nick_name":      user.NickName,
			"avatar":         user.Avatar,
			"password":       user.Password,
			"email":          user.Email,
			"signature":      user.Signature,
			"ip_addr":        user.IPAddr,
			"fans_count":     user.FansCount,
			"follower_count": user.FollowerCount,
			"like_cnt_count": user.LikeCntCount,
			"age":            user.Age,
			"male":           user.Male,
			//"birth_day":      user.BirthDay,
			"addr":       user.Addr,
			"profession": user.Profession,
			"school":     user.School,
			"level":      user.Level,
		}),
	}).Create(&user).Error
}

func (u *userDaoImpl) FindByUserName(ctx context.Context, userName string) (models.UserModel, error) {
	user := models.UserModel{}
	query := u.db.WithContext(ctx).Where("user_name = ?", userName).First(&user)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserModel{}, ErrRecordNotFound
		}
		return models.UserModel{}, err
	}
	return user, nil
}

func (u *userDaoImpl) FindByEmail(ctx context.Context, email string) (models.UserModel, error) {
	user := models.UserModel{}
	query := u.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserModel{}, ErrRecordNotFound
		}
		return models.UserModel{}, err
	}
	return user, nil
}
