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

func NewNoteSvcImpl(repo repo.NoteRepoInterface, lg ylog.Logger, intrRepo repo.Interactive) NoteService {
	return &noteSvcImpl{
		repo:            repo,
		interactiveRepo: intrRepo,
		lg:              lg}
}

type noteSvcImpl struct {
	repo            repo.NoteRepoInterface
	interactiveRepo repo.Interactive
	lg              ylog.Logger
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
		// 查询文章对应的作者信息
		info, err1 := n.repo.FindAuthorInfo(ctx, list[i].AuthorId)
		if err1 != nil {
			n.lg.Warn("find author info failed",
				ylog.String("err", err.Error()),
				ylog.String("authorId", list[i].AuthorId),
			)
			continue
		}
		list[i].AuthorInfo = &info

		// 查询文章对应的交互信息(点赞数，评论数等)
		data, err2 := n.interactiveRepo.GetById(ctx, list[i].ID, "note")
		if err2 != nil {
			n.lg.Warn("get interactive data failed", ylog.String("noteId", list[i].ID))
			// 没有交互信息很正常，所以不需要返回错误，直接跳过
			list[i].InteractiveInfo = &domain.InteractiveInfo{} // 给一个空的交互信息
			continue
		}
		//temp := n.toDomain(list[i])
		list[i].InteractiveInfo = &domain.InteractiveInfo{
			SourceGID:  data.SourceGID,
			ViewCnt:    data.ViewCnt,
			LikeCnt:    data.LikeCnt,
			ShareCnt:   data.ShareCnt,
			CommentCnt: data.CommentCnt,
			CollectCnt: data.CollectCnt,
			BizType:    data.BizType,
		}
	}
	return list, nil
}

func (n *noteSvcImpl) CreateNote(ctx context.Context, note domain.DNote, uuid string) error {

	if note.ContentType == 1 {
		// 保存图文笔记
		for _, img := range note.ImgList {
			absPath := img.LocalPath[1:]
			exist := utils.IsFileExist(absPath)
			if !exist {
				n.lg.Warn("image not exist", ylog.String("path", img.LocalPath))
				continue
			}
			md5, err := utils.FileMd5(absPath)
			if err != nil {
				continue
			}
			img.Size, _ = utils.GetFileSize(absPath)
			img.HashStr = md5
			img.ImgWidth, img.ImgHeight, _ = utils.GetImageSize(absPath)
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
	// 浏览数+1
	err := n.interactiveRepo.IncrReadCnt(ctx, id, "note")
	if err != nil {
		n.lg.Warn("incr view count failed", ylog.String("noteId", id))

	}
	return n.repo.NoteDetail(ctx, id)
}

func (n *noteSvcImpl) ListNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error) {
	list, err := n.repo.FindNoteList(ctx, status, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	res := make([]domain.DNote, 0, len(list))
	for _, model := range list {
		// 1. 查询文章的点赞数据
		data, err := n.interactiveRepo.GetById(ctx, model.UUID, "note")
		if err != nil {
			continue
		}
		temp := n.toDomain(model)
		temp.InteractiveInfo = &domain.InteractiveInfo{
			SourceGID:  data.SourceGID,
			ViewCnt:    data.ViewCnt,
			LikeCnt:    data.LikeCnt,
			ShareCnt:   data.ShareCnt,
			CommentCnt: data.CommentCnt,
			CollectCnt: data.CollectCnt,
			BizType:    data.BizType,
		}
		res = append(res, temp)
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
