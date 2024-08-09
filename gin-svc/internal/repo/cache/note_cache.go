package cache

import (
	"context"
	"gin-svc/internal/models"
)

type NoteCacheInterface interface {
	ListNoteByID(ctx context.Context, key string) ([]*models.NoteModel, error)
}
