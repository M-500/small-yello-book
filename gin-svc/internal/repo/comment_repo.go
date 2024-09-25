package repo

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
)

type CommentRepo interface {
	CreateComment(ctx context.Context, data domain.DComment) error
	CommentListByResourceId(ctx context.Context, resourceId string) ([]*domain.DComment, error)
	CommentListByParentId(ctx context.Context, parentId int, deep int) ([]*domain.DComment, error)
}

func NewCommentRepo(dao dao.CommentDAO) CommentRepo {
	return &commentRepo{dao: dao}
}

type commentRepo struct {
	dao dao.CommentDAO
}

func (c *commentRepo) CommentListByResourceId(ctx context.Context, resourceId string) ([]*domain.DComment, error) {
	res, err := c.dao.CommentListBySourceId(ctx, resourceId, 0, 10)
	if err != nil {
		return nil, err
	}
	var comments []*domain.DComment
	for _, item := range res {
		comments = append(comments, c.toDomain(item))
	}
	return comments, nil
}

func (c *commentRepo) CommentListByParentId(ctx context.Context, parentId int, deep int) ([]*domain.DComment, error) {
	res, err := c.dao.CommentListByParentId(ctx, parentId, 0, 10)
	if err != nil {
		return nil, err
	}
	var comments []*domain.DComment
	for _, item := range res {
		comments = append(comments, c.toDomain(item))
	}
	return comments, nil
}

func (c *commentRepo) CreateComment(ctx context.Context, data domain.DComment) error {
	// TODO: 异步增加评论数

	// 插入评论记录
	return c.dao.InsertComment(ctx, c.toModel(data))
}

func (c *commentRepo) toModel(data domain.DComment) models.CommentModel {
	return models.CommentModel{
		UserUUId:     data.UserUUId,
		ParentId:     data.ParentId,
		CommentType:  data.CommentType,
		ResourceId:   data.ResourceId,
		ResourceType: data.ResourceType,
		Content:      data.Content,
		IsPinned:     data.IsPinned,
		MediaUrl:     data.MediaUrl,
	}
}
func (c *commentRepo) toDomain(item models.CommentModel) *domain.DComment {
	return &domain.DComment{
		ID:           item.ID,
		UserUUId:     item.UserUUId,
		ParentId:     item.ParentId,
		ResourceId:   item.ResourceId,
		ResourceType: item.ResourceType,
		CommentType:  item.CommentType,
		Content:      item.Content,
		IsPinned:     item.IsPinned,
		MediaUrl:     item.MediaUrl,
	}
}
