package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
)

type NoteRepoInterface interface {
	FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error)
	CreateNote(ctx context.Context, note domain.DNote) error
}

type noteRepo struct {
	dao dao.NoteDaoInterface
}

func (n *noteRepo) FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noteRepo) CreateNote(ctx context.Context, note domain.DNote) error {
	err := n.dao.Insert(ctx, n.toModel(note))
	// 隐藏掉底层的错误
	return err
}

func (n *noteRepo) toModel(note domain.DNote) *models.NoteModel {
	return &models.NoteModel{}
}
