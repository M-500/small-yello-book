package dao

import (
	"context"
	"gin-svc/internal/conf/cfg"
	"gin-svc/internal/ioc"
	"gin-svc/internal/models"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func Test_permissionDao_Insert(t *testing.T) {
	d := cfg.DatabaseCfg{
		DSN:         "root:root@tcp(127.0.0.1:13317)/y_book?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConn: 10,
		MaxOpenConn: 10,
	}
	//redis := ioc.SetUpRedis(&config.Redis)
	db := ioc.SetUpDB(&d)
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data models.SysPermissionModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				data: models.SysPermissionModel{
					Title:  "创建角色",
					Action: http.MethodPost,
					PerKey: "create:role",
					Mark:   "",
					Path:   "/api/v1/roles",
					Type:   "SYS",
					Status: true,
				},
			},
		},
		{
			name: "",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				data: models.SysPermissionModel{
					Title:  "查看角色列表",
					Action: http.MethodGet,
					PerKey: "query:role:list",
					Mark:   "",
					Path:   "/api/v1/roles",
					Type:   "SYS",
					Status: true,
				},
			},
		},
		{
			name: "",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				data: models.SysPermissionModel{
					Title:  "查看角色详情",
					Action: http.MethodGet,
					PerKey: "query:role:detail",
					Mark:   "",
					Path:   "/api/v1/roles/:id",
					Type:   "SYS",
					Status: true,
				},
			},
		},
		{
			name: "",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				data: models.SysPermissionModel{
					Title:  "删除角色",
					Action: http.MethodDelete,
					PerKey: "delete:role",
					Mark:   "",
					Path:   "/api/v1/roles/:id",
					Type:   "SYS",
					Status: true,
				},
			},
		},
		{
			name: "",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				data: models.SysPermissionModel{
					Title:  "更新角色信息",
					Action: http.MethodPut,
					PerKey: "update:role:detail",
					Mark:   "",
					Path:   "/api/v1/roles",
					Type:   "SYS",
					Status: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &permissionDao{
				db: tt.fields.db,
			}
			if err := p.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
