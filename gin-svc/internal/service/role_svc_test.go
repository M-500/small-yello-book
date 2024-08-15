package service

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock" // 使用uber的mock包
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_roleSvcImpl_GetDetailByID(t *testing.T) {
	nowDate := time.Now()
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) repo.RoleRepo

		ctx context.Context
		id  int

		wantRes domain.RoleDetail
		wantErr error
	}{
		{
			name: "查询成功",
			mock: func(ctrl *gomock.Controller) repo.RoleRepo {
				roleRepo := mocks.NewMockRoleRepo(ctrl)
				roleRepo.EXPECT().GetRoleByID(gomock.Any(), 1).
					Return(&models.SysRoleModel{
						Model:    gorm.Model{ID: 1, CreatedAt: nowDate},
						RoleName: "admin",
						RoleKey:  "admin",
					}, nil)
				roleRepo.EXPECT().FindPermissionListByRoleId(gomock.Any(), 1).
					Return([]models.SysPermissionModel{
						{
							Model:  gorm.Model{ID: 1},
							Title:  "查看",
							PerKey: "query:view",
							Path:   "/view",
							Action: "GET",
						},
						{
							Model:  gorm.Model{ID: 2},
							Title:  "删除",
							PerKey: "del:view",
							Path:   "/view",
							Action: "DELETE",
						},
						{
							Model:  gorm.Model{ID: 3},
							Title:  "更新",
							PerKey: "update:view",
							Path:   "/view",
							Action: "PUT",
						},
					}, nil)
				return roleRepo
			},
			ctx: context.Background(),
			id:  1,

			wantRes: domain.RoleDetail{
				RoleBase: domain.Role{
					Id:         1,
					RoleName:   "admin",
					RoleKey:    "admin",
					CreateTime: uint(nowDate.Unix()),
				},
				PerList: []domain.Permission{
					{
						Id:     1,
						Title:  "查看",
						PerKey: "query:view",
						Path:   "/view",
						Action: "GET",
					},
					{
						Id:     2,
						Title:  "删除",
						PerKey: "del:view",
						Path:   "/view",
						Action: "DELETE",
					},
					{
						Id:     3,
						Title:  "更新",
						PerKey: "update:view",
						Path:   "/view",
						Action: "PUT",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 固定写法 照抄就好了
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			roleRepoMock := tc.mock(ctrl)
			svc := &roleSvcImpl{
				roleRepo: roleRepoMock,
			}
			id, err := svc.GetDetailByID(tc.ctx, tc.id)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.Equal(t, tc.wantRes, id)
			}
		})
	}
}
