package dao

import (
	"context"
	"fmt"
	"gin-svc/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNewInteractiveDao(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want InteractiveDao
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewInteractiveDao(tt.args.db), "NewInteractiveDao(%v)", tt.args.db)
		})
	}
}

func Test_gormInteractiveDao_DecrCollection(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.DecrCollection(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("DecrCollection(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_DecrComment(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.DecrComment(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("DecrComment(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_DecrLike(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.DecrLike(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("DecrLike(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_DecrViewLike(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.DecrViewLike(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("DecrViewLike(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.Delete(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("Delete(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_GetById(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.InteractiveModel
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			got, err := g.GetById(tt.args.ctx, tt.args.sourceId, tt.args.bizType)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType)
		})
	}
}

func Test_gormInteractiveDao_IncrCollection(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.IncrCollection(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("IncrCollection(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_IncrComment(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.IncrComment(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("IncrComment(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_IncrLike(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.IncrLike(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("IncrLike(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_IncrViewLike(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		sourceId string
		bizType  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.IncrViewLike(tt.args.ctx, tt.args.sourceId, tt.args.bizType), fmt.Sprintf("IncrViewLike(%v, %v, %v)", tt.args.ctx, tt.args.sourceId, tt.args.bizType))
		})
	}
}

func Test_gormInteractiveDao_Insert(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		social models.InteractiveModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.Insert(tt.args.ctx, tt.args.social), fmt.Sprintf("Insert(%v, %v)", tt.args.ctx, tt.args.social))
		})
	}
}

func Test_gormInteractiveDao_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		social models.InteractiveModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gormInteractiveDao{
				db: tt.fields.db,
			}
			tt.wantErr(t, g.Update(tt.args.ctx, tt.args.social), fmt.Sprintf("Update(%v, %v)", tt.args.ctx, tt.args.social))
		})
	}
}
