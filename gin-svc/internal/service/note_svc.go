package service

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/pkg/utils"
	"gin-svc/pkg/ylog"
	"strconv"
)

type NoteService interface {
	// 创建笔记
	CreateNote(ctx context.Context, note domain.DNote, uuid string) error

	// 获取笔记详情
	GetNoteDetail(ctx context.Context, noteId string) (domain.DNote, error)

	// 获取所有文章列表
	ListNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error)

	FeedListNote(ctx context.Context, TagId int, offset int, limit int) ([]domain.DNote, error)

	// 更新笔记
	UpdateNote(ctx context.Context, note domain.DNote) error

	PassNote(ctx context.Context, uuid string) error

	// 删除笔记
	DeleteNote(ctx context.Context, id int) error

	// 更新笔记权限
	UpdatePermission(ctx context.Context, noteId int, userId int, permission int) error
}

type noteSvcImpl struct {
	repo repo.NoteRepoInterface
	lg   ylog.Logger
}

func (n *noteSvcImpl) PassNote(ctx context.Context, uuid string) error {
	return n.repo.ChangeStatus(ctx, uuid, domain.NoteStatePublished)
}

func (n *noteSvcImpl) FeedListNote(ctx context.Context, TagId int, offset int, limit int) ([]domain.DNote, error) {
	list, err := n.repo.FeedNoteList(ctx, TagId, offset, limit)
	if err != nil {
		return nil, err
	}
	for i := range list {
		info, err := n.repo.FindAuthorInfo(ctx, list[i].AuthorId)
		if err != nil {
			continue
		}
		list[i].AuthorInfo = info
	}
	return list, nil
}

func NewNoteSvcImpl(repo repo.NoteRepoInterface, lg ylog.Logger) NoteService {
	return &noteSvcImpl{repo: repo, lg: lg}
}

func (n *noteSvcImpl) CreateNote(ctx context.Context, note domain.DNote, uuid string) error {

	if note.ContentType == 1 {
		// 保存图文笔记
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
		err := n.repo.CreateNote(ctx, note, uuid)
		if err != nil {
			n.lg.Warn("create note failed", ylog.String("err", err.Error()))
			return err
		}
		// TODO 3. 将图片迁移到OSS，因为本地图片是临时图片，不会永久保存
		return nil
	}
	// 保存视频笔记

	return nil
}

func (n *noteSvcImpl) GetNoteDetail(ctx context.Context, id string) (domain.DNote, error) {
	return n.repo.NoteDetail(ctx, id)
}

func (n *noteSvcImpl) ListNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error) {
	list, err := n.repo.FindNoteList(ctx, status, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	res := make([]domain.DNote, 0, len(list))
	for _, model := range list {
		res = append(res, n.toDomain(model))
	}
	return res, len(list), nil
}

func (n *noteSvcImpl) toDomain(note models.NoteModel) domain.DNote {
	return domain.DNote{
		ID:          note.UUID,
		NoteTitle:   note.Title,
		NoteContent: note.Mark,
		Cover:       note.Cover,
		//Address:     note.Address,
		Statement: strconv.Itoa(note.Status),
		//PublishTime: note.Model.CreatedAt,
		//Private:     note.,
		Status:     domain.NoteStatus(note.Status),
		AuthorId:   note.AuthorId,
		CreateTime: "",
		UpdateTime: "",
		ImgList:    nil,
	}
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
