package service

import (
	"context"
	"gin-svc/internal/domain"
)

type NoteService interface {
	// 创建笔记
	CreateNote(ctx context.Context, note domain.DNote) error

	// 获取笔记详情
	GetNoteDetail(ctx context.Context, id int) (domain.DNote, error)

	// 获取某个用户的笔记列表
	ListUserNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error)

	// 更新笔记
	UpdateNote(ctx context.Context, note domain.DNote) error

	// 删除笔记
	DeleteNote(ctx context.Context, id int) error

	// 更新笔记权限
	UpdatePermission(ctx context.Context, noteId int, userId int, permission int) error
}

type noteSvcImpl struct {
}

func (n *noteSvcImpl) CreateNote(ctx context.Context, note domain.DNote) error {
	//TODO implement me
	panic("implement me")
}

func (n *noteSvcImpl) GetNoteDetail(ctx context.Context, id int) (domain.DNote, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noteSvcImpl) ListUserNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noteSvcImpl) UpdateNote(ctx context.Context, note domain.DNote) error {
	//TODO implement me
	panic("implement me")
}

func (n *noteSvcImpl) DeleteNote(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (n *noteSvcImpl) UpdatePermission(ctx context.Context, noteId int, userId int, permission int) error {
	return nil
}
