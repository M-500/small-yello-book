package service

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/repo"
	"gin-svc/pkg/utils"
	"gin-svc/pkg/ylog"
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
	repo repo.NoteRepoInterface
	lg   ylog.Logger
}

func NewNoteSvcImpl(repo repo.NoteRepoInterface, lg ylog.Logger) NoteService {
	return &noteSvcImpl{repo: repo, lg: lg}
}

func (n *noteSvcImpl) CreateNote(ctx context.Context, note domain.DNote) error {
	// 1. 检查所有图片的链接是否存在
	for _, img := range note.ImgList {
		exist := utils.IsFileExist(img.LocalPath)
		if !exist {
			n.lg.Warn("image not exist", ylog.String("path", img.LocalPath))
			continue
		}
		md5, err := utils.FileMd5(img.LocalPath)
		if err != nil {
			continue
		}
		img.HashStr = md5
		//img., _ := utils.GetFileSize(img.LocalPath)
		img.ImgWidth, img.ImgHeight, _ = utils.GetImageSize(img.LocalPath)
	}
	// 1. 保存文章
	err := n.repo.CreateNote(ctx, note)
	if err != nil {
		n.lg.Warn("create note failed", ylog.String("err", err.Error()))
		return err
	}
	// 3. 将图片迁移到OSS，因为本地图片是临时图片，不会永久保存
	return err
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
