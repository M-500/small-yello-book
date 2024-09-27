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
	CreateNote(ctx context.Context, note domain.DNote, uuid string) error                                         // 创建笔记
	GetNoteDetail(ctx context.Context, userId, noteId string) (domain.DNote, error)                               // 登陆用户获取笔记详情
	ListNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error)                 // 获取所有文章列表
	ListNoteByUserPublish(ctx context.Context, userId string, offset int, limit int) ([]domain.DNote, int, error) // 获取用户发布的文章列表
	FeedListNote(ctx context.Context, TagId int, offset int, limit int) ([]domain.DNote, error)
	UpdateNote(ctx context.Context, note domain.DNote) error // 更新笔记
	PassNote(ctx context.Context, uuid string) error
	DeleteNote(ctx context.Context, id int) error                                       // 删除笔记
	UpdatePermission(ctx context.Context, noteId int, userId int, permission int) error // 更新笔记权限
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

func (n *noteSvcImpl) ListNoteByUserPublish(ctx context.Context, userId string, offset int, limit int) ([]domain.DNote, int, error) {
	info, err1 := n.repo.FindAuthorInfo(ctx, userId)
	if err1 != nil {
		n.lg.Warn("find author info failed",
			ylog.String("err", err1.Error()),
			ylog.String("authorId", userId),
		)
		return nil, 0, err1
	}
	query, i, err := n.repo.FindNoteListByUser(ctx, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	res := make([]domain.DNote, 0, len(query))
	for _, model := range query {
		temp := n.toDomain(model)
		temp.AuthorInfo = &info // 组装用户信息
		// 获取点赞收藏评论数
		// 1. 查询文章的点赞数据
		data, err := n.interactiveRepo.GetById(ctx, model.UUID, "note")
		if err != nil {
			res = append(res, temp)
			continue
		}
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
	return res, int(i), nil
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
			// 记录最宽的图片宽度
			if note.CoverWidth < img.ImgWidth {
				note.CoverWidth = img.ImgWidth
				note.CoverHeight = img.ImgHeight
			}
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

func (n *noteSvcImpl) GetNoteDetail(ctx context.Context, userId, noteID string) (domain.DNote, error) {
	// 获取文章的基本信息
	detail, err := n.repo.NoteDetail(ctx, noteID)
	if err != nil {
		return domain.DNote{}, err
	}
	// 获取文章交互信息数据
	interactiveInfo, err := n.interactiveRepo.GetById(ctx, noteID, "note")
	if err != nil {
		n.lg.Warn("get interactive data failed", ylog.String("noteId", noteID))
	}
	detail.InteractiveInfo = &interactiveInfo

	// 文章阅读量 + 1
	_ = n.interactiveRepo.IncrReadCnt(ctx, noteID, "note")

	if utils.IsBlank(userId) {
		// 游客访问
		return detail, nil
	}

	// 判断用户是否点赞收藏

	// 将该文章添加到用户的浏览记录中

	return domain.DNote{}, err
}

func (n *noteSvcImpl) ListNote(ctx context.Context, status int, offset int, limit int) ([]domain.DNote, int, error) {
	list, total, err := n.repo.FindNoteList(ctx, status, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	res := make([]domain.DNote, 0, len(list))
	for _, model := range list {
		temp := n.toDomain(model)
		// 查询文章对应的作者信息
		info, err1 := n.repo.FindAuthorInfo(ctx, model.AuthorId)
		if err1 != nil {
			//n.lg.Warn("find author info failed",
			//	ylog.String("err", err.Error()),
			//	ylog.String("authorId", model.AuthorId),
			//)
			res = append(res, temp)
			continue
		}
		// 1. 查询文章的点赞数据
		data, err := n.interactiveRepo.GetById(ctx, model.UUID, "note")
		if err != nil {
			res = append(res, temp)
			continue
		}
		temp.AuthorInfo = &info
		temp.InteractiveInfo = &data
		res = append(res, temp)
	}
	return res, int(total), nil
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
