package repo

import (
	"context"
	"gin-svc/internal/models"
)

type TagRepoInterface interface {
	GetAllTag(ctx context.Context) ([]models.TagModel, error)
}

type tagRepo struct {
}

func (t *tagRepo) GetAllTag(ctx context.Context) ([]models.TagModel, error) {
	//TODO implement me
	panic("implement me")
}
