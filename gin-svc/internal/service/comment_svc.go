package service

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/repo"
	"gin-svc/internal/types"
)

type CommentSvc interface {
	CreateComment(ctx context.Context, userId string, data types.AddCommentForm) error
	CommentList(ctx context.Context, resourceId string) ([]*domain.DComment, error)
}

func NewCommentSvc(commRepo repo.CommentRepo) CommentSvc {
	return &commSvc{commRepo: commRepo}
}

type commSvc struct {
	commRepo repo.CommentRepo
}

func (c *commSvc) CommentList(ctx context.Context, resourceId string) ([]*domain.DComment, error) {
	listLay1, err := c.commRepo.CommentListByResourceId(ctx, resourceId)
	if err != nil {
		return nil, err
	}
	for _, item := range listLay1 {
		res, err := c.commRepo.CommentListByParentId(ctx, int(item.ID), 3)
		if err != nil {
			continue
		}
		item.Sub = res
	}
	return listLay1, nil
}

func (c *commSvc) CreateComment(ctx context.Context, userId string, data types.AddCommentForm) error {
	dmComm := c.toDomainComment(data)
	dmComm.UserUUId = userId
	//dmComm.Ip = ctx.Value("ip").(string)

	return c.commRepo.CreateComment(ctx, dmComm)
}

func (c *commSvc) toDomainComment(data types.AddCommentForm) domain.DComment {
	return domain.DComment{
		ParentId:   data.ParentId,
		ResourceId: data.ResourceId,
		Content:    data.Content,
		IsPinned:   data.IsPinned,
		MediaUrl:   data.MediaUrl,
	}
}
