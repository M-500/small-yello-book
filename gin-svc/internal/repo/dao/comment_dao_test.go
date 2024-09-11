package dao

import (
	"context"
	"gin-svc/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_gormCommentDAO_CommentListByParentId(t *testing.T) {
	tests := []struct {
		name     string                     // 测试用例的名字
		mock     func(mock sqlmock.Sqlmock) // 使用SQL Mock
		ctx      context.Context            // 预期的输入参数
		parentId int                        // 预期的输入参数
		page     int                        // 预期的输入参数
		pageSize int                        // 预期的输入参数
		want     []models.CommentModel      // 预期的返回值
		wantErr  bool                       // 预期的返回值
	}{
		{},
	}
	for _, tc := range tests {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()
		// 这里使用gorm的mysql驱动，传入刚才mock出来的sql.DB
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{
			DisableAutomaticPing:   true,
			SkipDefaultTransaction: true,
		})
		assert.NoError(t, err)
		tc.mock(mock) // 调用mock函数，传入mock对象
		dao := NewCommentDAO(gormDB)
		id, err := dao.CommentListByParentId(tc.ctx, tc.parentId, tc.page, tc.pageSize)
		assert.Equal(t, tc.wantErr, err)
		assert.Equal(t, len(tc.want), len(id))
	}
}
