package dao

import (
	"context"
	"database/sql"

	"gin-svc/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_permissionDao_Insert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(t *testing.T) *sql.DB
		ctx  context.Context
		user models.SysPermissionModel

		wantErr error

		before func(t *testing.T)
		after  func(t *testing.T)
	}{
		//{
		//	name: "插入成功",
		//	mock: func(t *testing.T) *sql.DB {
		//		db, mock, err := sqlmock.New()
		//		assert.NoError(t, err)
		//		mockRes := sqlmock.NewResult(123, 1)
		//		// 这边要求传入的是 sql 的正则表达式
		//		mock.ExpectExec("INSERT INTO .*").
		//			WillReturnResult(mockRes)
		//		return db
		//	},
		//	ctx: context.Background(),
		//	user: models.SysPermissionModel{
		//		Title: "Tom",
		//	},
		//},
		//{
		//	name: "邮箱冲突",
		//	mock: func(t *testing.T) *sql.DB {
		//		db, mock, err := sqlmock.New()
		//		assert.NoError(t, err)
		//		// 这边要求传入的是 sql 的正则表达式
		//		mock.ExpectExec("INSERT INTO .*").
		//			WillReturnError(&dbDriver.MySQLError{Number: 1062})
		//		return db
		//	},
		//	ctx: context.Background(),
		//	user: User{
		//		Nickname: "Tom",
		//	},
		//	wantErr: ErrDuplicateEmail,
		//},
		//{
		//	name: "数据库错误",
		//	mock: func(t *testing.T) *sql.DB {
		//		db, mock, err := sqlmock.New()
		//		assert.NoError(t, err)
		//		// 这边要求传入的是 sql 的正则表达式
		//		mock.ExpectExec("INSERT INTO .*").
		//			WillReturnError(errors.New("数据库错误"))
		//		return db
		//	},
		//	ctx: context.Background(),
		//	user: User{
		//		Nickname: "Tom",
		//	},
		//	wantErr: errors.New("数据库错误"),
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sqlDB := tc.mock(t)
			db, err := gorm.Open(mysql.New(mysql.Config{
				Conn:                      sqlDB,
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				DisableAutomaticPing:   true,
				SkipDefaultTransaction: true,
			})
			assert.NoError(t, err)
			dao := NewPermissionDAO(db)
			err = dao.Insert(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
