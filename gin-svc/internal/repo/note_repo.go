package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
	"gin-svc/pkg/utils"
	"gorm.io/gorm"
	"strconv"
)

type NoteRepoInterface interface {
	FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error)
	FindNoteList(ctx context.Context, status int, page, size int) ([]models.NoteModel, error)
	FeedNoteList(ctx context.Context, tagID int, page, size int) ([]domain.DNote, error)
	FindAuthorInfo(ctx context.Context, authorID string) (*domain.Author, error)
	CreateNote(ctx context.Context, note domain.DNote, uid string) error
	ChangeStatus(ctx context.Context, noteID string, status domain.NoteStatus) error
	NoteDetail(ctx context.Context, noteID string) (domain.DNote, error)
}

func NewNoteRepo(dao dao.NoteDaoInterface, userDao dao.UserDao, imgDao dao.ImageDao) NoteRepoInterface {
	return &noteRepo{
		dao:     dao,
		imgDao:  imgDao,
		userDao: userDao}
}

type noteRepo struct {
	dao     dao.NoteDaoInterface
	imgDao  dao.ImageDao
	userDao dao.UserDao
}

func (n *noteRepo) NoteDetail(ctx context.Context, noteID string) (domain.DNote, error) {
	note, err := n.dao.FindByUUID(ctx, noteID)
	if err != nil {
		return domain.DNote{}, err
	}
	res := n.toDMNote(note)
	// 获取图片列表 / 视频信息
	res.ImgList = make([]domain.ImageNote, 0)
	imgList, err := n.imgDao.FindListByNoteId(ctx, noteID)
	if err != nil {
		return res, err
	}
	imgs := make([]domain.ImageNote, 0, len(imgList))
	for _, data := range imgList {
		//tmp := n.toImageNote(data)
		//tmp.Url = "http://127.0.0.1:8122" + tmp.LocalPath
		imgs = append(imgs, n.toImageNote(data))
	}
	res.ImgList = imgs
	// 获取用户信息
	return res, nil
}

func (n *noteRepo) ChangeStatus(ctx context.Context, noteID string, status domain.NoteStatus) error {
	return n.dao.ChangeStatus(ctx, noteID, int(status))
}

func (n *noteRepo) FindAuthorInfo(ctx context.Context, authorID string) (*domain.Author, error) {
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

func (n *noteRepo) CreateNote(ctx context.Context, note domain.DNote, uid string) error {
	note.Status = domain.NoteStateDraft // 限制死状态
	res := n.toModel(note)
	res.AuthorId = uid
	err := n.dao.SaveNoteWithTx(ctx, res, n.toImageModel(note))
	// 隐藏掉底层的错误
	return err
}

func (n *noteRepo) toModel(note domain.DNote) models.NoteModel {
	return models.NoteModel{
		Model:       gorm.Model{},
		UUID:        note.ID,
		Title:       note.NoteTitle,
		Cover:       note.ImgList[0].ImgUrl,
		Mark:        note.NoteContent,
		ContentType: note.ContentType,
		Status:      0,
		LikeCnt:     0,
		ViewCnt:     0,
		ShareCnt:    0,
		CommentCnt:  0,
		CollectCnt:  0,
	}
}

func (n *noteRepo) toDMNote(note models.NoteModel) domain.DNote {
	dNote := domain.DNote{
		ID:          note.UUID,
		NoteTitle:   note.Title,
		NoteContent: note.Mark,
		Cover:       note.Cover,
		ContentType: note.ContentType,
		//Address:     note.Address,
		Statement: strconv.Itoa(note.Status),
		//PublishTime: note.Model.CreatedAt,
		//Private:     note.,
		ViewCnt:    note.ViewCnt,
		LikeCnt:    note.LikeCnt,
		ShareCnt:   note.ShareCnt,
		CommentCnt: note.CommentCnt,
		CollectCnt: note.CollectCnt,
		Status:     domain.NoteStatus(note.Status),
		AuthorId:   note.AuthorId,
		CreateTime: "",
		UpdateTime: "",
	}
	if utils.IsBlank(dNote.ID) {
		dNote.ID = utils.UUID()
	}
	return dNote
}

func (n *noteRepo) toImageModel(data domain.DNote) []models.ImageModel {
	imgList := make([]models.ImageModel, 0, len(data.ImgList))
	for _, i2 := range data.ImgList {
		imgList = append(imgList, models.ImageModel{
			UUID:    i2.ID,
			NoteId:  data.ID,
			Model:   gorm.Model{},
			Width:   i2.ImgWidth,
			OldPath: i2.LocalPath,
			Hash:    i2.HashStr,
			Size:    i2.ImgHeight,
		})
	}
	return imgList
}

func (n *noteRepo) toImageNote(data models.ImageModel) domain.ImageNote {
	return domain.ImageNote{
		ID:        data.UUID,
		NoteId:    data.NoteId,
		ImgWidth:  data.Width,
		LocalPath: "http://127.0.0.1:8122" + data.OldPath,
		HashStr:   data.Hash,
		ImgHeight: data.Size,
	}
}

func (n *noteRepo) toDMAuthor(user *models.UserModel) *domain.Author {
	return &domain.Author{
		ID:       user.GlobalNumber,
		NickName: user.NickName,
		Avatar:   user.Avatar,
	}
}
