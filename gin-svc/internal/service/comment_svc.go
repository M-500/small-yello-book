package service

import (
	"context"
	"gin-svc/internal/types"
)

type CommentSvc interface {
	CreateComment(ctx context.Context, data types.AddCommentForm) error
}
