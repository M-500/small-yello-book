package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
	"gorm.io/gorm"
	"strconv"
)

type NoteRepoInterface interface {
	FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error)
	FindNoteList(ctx context.Context, status int, page, size int) ([]models.NoteModel, error)
	FeedNoteList(ctx context.Context, tagID int, page, size int) ([]domain.DNote, error)
	FindAuthorInfo(ctx context.Context, authorID int) (*domain.Author, error)
	CreateNote(ctx context.Context, note domain.DNote) error
}

func NewNoteRepo(dao dao.NoteDaoInterface, userDao dao.UserDao) NoteRepoInterface {
	return &noteRepo{dao: dao, userDao: userDao}
}

type noteRepo struct {
	dao     dao.NoteDaoInterface
	userDao dao.UserDao
}

func (n *noteRepo) FindAuthorInfo(ctx context.Context, authorID int) (*domain.Author, error) {
	user, err := n.userDao.GetByID(ctx, authorID)
	if err != nil {
		return &domain.Author{}, err
	}
	return n.toDMAuthor(user), nil
}

func (n *noteRepo) FeedNoteList(ctx context.Context, tagID int, page, size int) ([]domain.DNote, error) {
	res := make([]domain.DNote, 0, size)
	list, err := n.dao.FeedNoteList(ctx, tagID, page, size)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		res = append(res, n.toDMNote(item))
	}
	return res, nil
}

func (n *noteRepo) FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n *noteRepo) FindNoteList(ctx context.Context, status int, page, size int) ([]models.NoteModel, error) {
	return n.dao.ListByStatus(ctx, status, page, size)
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

func (n *noteRepo) toDMNote(note models.NoteModel) domain.DNote {
	return domain.DNote{
		ID:          int(note.Model.ID),
		NoteTitle:   note.Title,
		NoteContent: note.Mark,
		Cover:       note.Cover,
		//Address:     note.Address,
		Statement: strconv.Itoa(note.Status),
		//PublishTime: note.Model.CreatedAt,
		//Private:     note.,
		ViewCnt:    note.ViewCnt,
		LikeCnt:    note.LikeCnt,
		ShareCnt:   note.ShareCnt,
		CommentCnt: note.CommentCnt,
		CollectCnt: note.CollectCnt,
		Status:     note.Status,
		AuthorId:   int(note.AuthorId),
		CreateTime: "",
		UpdateTime: "",
		ImgList:    nil,
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

func (n *noteRepo) toDMAuthor(user *models.UserModel) *domain.Author {
	return &domain.Author{
		ID:       int(user.ID),
		NickName: user.NickName,
		Avatar:   user.Avatar,
	}
}
