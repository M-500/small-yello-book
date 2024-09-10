package repo

import (
	"context"
	"gin-svc/internal/domain"
)

type CommentRepo interface {
	CreateComment(ctx context.Context, userId string, data domain.DComment) error
}

type commentRepo struct {
}

func (c *commentRepo) CreateComment(ctx context.Context, userId string, data domain.DComment) error {
	//TODO implement me
	panic("implement me")
}
