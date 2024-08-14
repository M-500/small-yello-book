package dao

import (
	"context"
	"database/sql"
	"fmt"
	"gin-svc/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// Test_userDaoImpl_Delete 测试删除用户
func Test_userDaoImpl_Delete(t *testing.T) {
	testCase := []struct {
		name string
		mock func(t *testing.T) *sql.DB
		ctx  context.Context
		uid  int

		wantErr error
	}{
		{
			name: "删除成功",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mock.ExpectExec("UPDATE sys_user SET delete_at = `1` WHERE id = ?").
					WithArgs(123).
					WillReturnResult(sqlmock.NewResult(0, 0))
				return db
			},
			ctx:     context.Background(),
			uid:     123, // 123这行记录肯定不存在
			wantErr: gorm.ErrRecordNotFound,
		},
	}
	for _, tc := range testCase {
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
			userDao := NewUserDao(db)
			err = userDao.Delete(tc.ctx, tc.uid) // 调用删除用户的方法
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func Test_userDaoImpl_FindByEmail(t *testing.T) {
	testCase := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
		mock   func(t *testing.T) *sql.DB
		ctx    context.Context
		email  string

		wantErr error
	}{
		{
			name: "查询成功",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				rows := sqlmock.NewRows([]string{"id", "email", "password", "ctime", "utime"}).
					AddRow(1, "18511112222@163.com", "123", "123", "123")
				mock.ExpectQuery("SELECT id, email, password, ctime, utime FROM sys_user WHERE email = ?").
					WithArgs("").WillReturnRows(rows)
				return db
			},
			ctx:     context.Background(),
			email:   "18511112222@163.com",
			wantErr: nil,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)

			//
		})
	}
}

func Test_userDaoImpl_FindByUserName(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		userName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.UserModel
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userDaoImpl{
				db: tt.fields.db,
			}
			got, err := u.FindByUserName(tt.args.ctx, tt.args.userName)
			if !tt.wantErr(t, err, fmt.Sprintf("FindByUserName(%v, %v)", tt.args.ctx, tt.args.userName)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindByUserName(%v, %v)", tt.args.ctx, tt.args.userName)
		})
	}
}

func Test_userDaoImpl_Upsert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(t *testing.T) *sql.DB
		ctx  context.Context
		user models.UserModel

		wantErr error
	}{
		{
			name: "插入成功",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mockRes := sqlmock.NewResult(123, 1)
				// 这边要求传入的是 sql 的正则表达式
				mock.ExpectExec("INSERT INTO .*").
					WillReturnResult(mockRes)
				return db
			},
			ctx: context.Background(),
			user: models.UserModel{
				Email: "1978992222@qq.com",
			},
			wantErr: nil,
		},
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
			dao := NewUserDao(db)
			err = dao.Upsert(tc.ctx, &tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
