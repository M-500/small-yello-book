package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
	"gorm.io/gorm"
)

type NoteRepoInterface interface {
	FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error)
	CreateNote(ctx context.Context, note domain.DNote) error
}

func NewNoteRepo(dao dao.NoteDaoInterface) NoteRepoInterface {
	return &noteRepo{dao: dao}
}

type noteRepo struct {
	dao dao.NoteDaoInterface
}

func (n *noteRepo) FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noteRepo) CreateNote(ctx context.Context, note domain.DNote) error {
	err := n.dao.SaveNoteWithTx(ctx, n.toModel(note), n.toImageModel(note))
	// 隐藏掉底层的错误
	return err
}

func (n *noteRepo) toModel(note domain.DNote) models.NoteModel {
	return models.NoteModel{
		Model:      gorm.Model{},
		Title:      note.NoteTitle,
		Cover:      note.ImgList[0].ImgUrl,
		Mark:       note.NoteContent,
		Type:       "",
		Status:     0,
		LikeCnt:    0,
		ViewCnt:    0,
		ShareCnt:   0,
		CommentCnt: 0,
		CollectCnt: 0,
		AuthorId:   0,
	}
}

func (n *noteRepo) toImageModel(data domain.DNote) []models.ImageModel {
	imgList := make([]models.ImageModel, 0, len(data.ImgList))
	for _, i2 := range data.ImgList {
		imgList = append(imgList, models.ImageModel{
			Model:   gorm.Model{},
			Width:   i2.ImgWidth,
			OldPath: i2.LocalPath,
			Hash:    i2.HashStr,
			Size:    i2.ImgHeight,
		})
	}
	return imgList
}
