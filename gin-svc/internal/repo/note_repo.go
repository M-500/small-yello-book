package repo

import (
	"context"
	"errors"
	"gin-svc/internal/conf/cfg"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
	"gin-svc/pkg/utils"
	"gorm.io/gorm"
	"strconv"
)

type NoteRepoInterface interface {
	FindNoteListById(ctx context.Context, status int, page, size int, userId int) ([]models.NoteModel, error)
	FindNoteListByUser(ctx context.Context, userId string, page, size int) ([]models.NoteModel, int64, error)
	FindNoteList(ctx context.Context, status int, page, size int) ([]models.NoteModel, int64, error)
	FeedNoteList(ctx context.Context, tagID int, page, size int) ([]domain.DNote, error)
	FindAuthorInfo(ctx context.Context, authorID string) (domain.Author, error)
	CreateNote(ctx context.Context, note domain.DNote, uid string) error
	ChangeStatus(ctx context.Context, noteID string, status domain.NoteStatus) error
	NoteDetail(ctx context.Context, noteID string) (domain.DNote, error)
}

func NewNoteRepo(dao dao.NoteDaoInterface, userDao dao.UserDao, imgDao dao.ImageDao, cfg cfg.ServerCfg) NoteRepoInterface {
	return &noteRepo{
		dao:     dao,
		imgDao:  imgDao,
		cfg:     cfg,
		userDao: userDao}
}

type noteRepo struct {
	dao     dao.NoteDaoInterface
	imgDao  dao.ImageDao
	userDao dao.UserDao
	cfg     cfg.ServerCfg
}

func (n *noteRepo) FindNoteListByUser(ctx context.Context, userId string, page, size int) ([]models.NoteModel, int64, error) {
	return n.dao.ListPageByUserPublish(ctx, userId, page, size)
}

// NoteDetail 通过笔记ID获取笔记详情
func (n *noteRepo) NoteDetail(ctx context.Context, noteID string) (domain.DNote, error) {
	// 获取文章基本信息
	note, err := n.dao.FindByUUID(ctx, noteID)
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			return domain.DNote{}, errors.New("文章ID不存在")
		}
		return domain.DNote{}, errors.New("通过ID查询文章详情失败")
	}
	res := n.toDMNote(note)

	// 获取图片列表 / 视频信息
	imgList, err := n.imgDao.FindListByNoteId(ctx, noteID)
	if err != nil {
		return res, err
	}
	res.ImgList = make([]*domain.ImageNote, 0)
	for _, data := range imgList {
		res.ImgList = append(res.ImgList, n.toImageNote(data))
	}

	// 获取作者信息
	user, err := n.FindAuthorInfo(ctx, note.AuthorId)
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			return domain.DNote{}, errors.New("作者不存在")
		}
		return domain.DNote{}, errors.New("查询作者信息失败")
	}
	res.AuthorInfo = &user
	return res, nil
}

func (n *noteRepo) ChangeStatus(ctx context.Context, noteID string, status domain.NoteStatus) error {
	return n.dao.ChangeStatus(ctx, noteID, int(status))
}

func (n *noteRepo) FindAuthorInfo(ctx context.Context, authorID string) (domain.Author, error) {
	user, err := n.userDao.GetByID(ctx, authorID)
	if err != nil {
		return domain.Author{}, err
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

func (n *noteRepo) FindNoteList(ctx context.Context, status int, page, size int) ([]models.NoteModel, int64, error) {
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
		CoverWidth:  note.CoverWidth,
		CoverHeight: note.CoverHeight,
		Mark:        note.NoteContent,
		ContentType: note.ContentType,
		Status:      0,
	}
}

func (n *noteRepo) toDMNote(note models.NoteModel) domain.DNote {
	dNote := domain.DNote{
		ID:          note.UUID,
		NoteTitle:   note.Title,
		NoteContent: note.Mark,
		Cover:       note.Cover,
		CoverHeight: note.CoverHeight,
		CoverWidth:  note.CoverWidth,
		ContentType: note.ContentType,
		//Address:     note.Address,
		Statement: strconv.Itoa(note.Status),
		//PublishTime: note.Model.CreatedAt,
		//Private:     note.,
		Status:     domain.NoteStatus(note.Status),
		AuthorId:   note.AuthorId,
		CreateTime: note.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateTime: note.UpdatedAt.Format("2006-01-02 15:04:05"),
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
			Height:  i2.ImgHeight,
			OldPath: i2.LocalPath,
			Hash:    i2.HashStr,
			Size:    i2.Size,
		})
	}
	return imgList
}

func (n *noteRepo) toImageNote(data models.ImageModel) *domain.ImageNote {
	return &domain.ImageNote{
		ID:        data.UUID,
		NoteId:    data.NoteId,
		ImgWidth:  data.Width,
		LocalPath: n.cfg.FileUploadHost + data.OldPath,
		HashStr:   data.Hash,
		ImgHeight: data.Height,
		Size:      data.Size,
	}
}

func (n *noteRepo) toDMAuthor(user models.UserModel) domain.Author {
	return domain.Author{
		ID:       user.GlobalNumber,
		NickName: user.NickName,
		Avatar:   user.Avatar,
	}
}
