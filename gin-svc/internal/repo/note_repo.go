package repo

import (
	"context"
	"gin-svc/internal/models"
)

type NoteRepoInterface interface {
	FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error)
}

type noteRepo struct {
}

func (n *noteRepo) FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error) {
	//TODO implement me
	panic("implement me")
}
