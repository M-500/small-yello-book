package dao

import (
	"context"
	"gin-svc/internal/models"
)

type LikeLogDAO interface {
	Insert(ctx context.Context, data models.LikeModel) error // 新增一条点赞记录
	Delete(ctx context.Context, data models.LikeModel) error // 删除一条点赞记录
	PageListLikeLog(ctx context.Context, userID, resourceID, bizType string, page, size int) ([]models.LikeModel, int64, error)
}
