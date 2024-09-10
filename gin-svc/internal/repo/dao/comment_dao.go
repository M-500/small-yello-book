package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type CommentDAO interface {
	InsertComment(ctx context.Context, data models.CommentModel) error
	DeleteComment(ctx context.Context, id int) error
	GetCommentById(ctx context.Context, id int) (models.CommentModel, error)
	UpdateComment(ctx context.Context, id int, data models.CommentModel) error
	CommentListBySourceId(ctx context.Context, sourceId string, page, pageSize int) ([]models.CommentModel, error)
	CommentListByParentId(ctx context.Context, parentId int, page, pageSize int) ([]models.CommentModel, error)
}

func NewCommentDAO(db *gorm.DB) CommentDAO {
	return &gormCommentDAO{db: db}
}

type gormCommentDAO struct {
	db *gorm.DB
}

func (g *gormCommentDAO) GetCommentById(ctx context.Context, id int) (models.CommentModel, error) {
	var comment models.CommentModel
	err := g.db.WithContext(ctx).Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (g *gormCommentDAO) DeleteComment(ctx context.Context, id int) error {
	return g.db.WithContext(ctx).Where("id = ?", id).Delete(&models.CommentModel{}).Error
}

func (g *gormCommentDAO) UpdateComment(ctx context.Context, id int, data models.CommentModel) error {
	//TODO implement me
	panic("implement me")
}

// CommentListBySourceId
//
//	@Description: 获取文章的一级评论
//	@receiver g
//	@param ctx
//	@param sourceId
//	@param page
//	@param pageSize
//	@return []models.CommentModel
//	@return error
func (g *gormCommentDAO) CommentListBySourceId(ctx context.Context, sourceId string, page, pageSize int) ([]models.CommentModel, error) {
	var res []models.CommentModel
	err := g.db.WithContext(ctx).Where("resource_id = ? and resource_type = null", sourceId).Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	return res, err
}

// CommentListByParentId
//
//	@Description: 获取评论的子评论
//	@receiver g
//	@param ctx
//	@param parentId
//	@param page
//	@param pageSize
//	@return []models.CommentModel
//	@return error
func (g *gormCommentDAO) CommentListByParentId(ctx context.Context, parentId int, page, pageSize int) ([]models.CommentModel, error) {
	var res []models.CommentModel
	err := g.db.WithContext(ctx).Where("parent_id = ?", parentId).
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *gormCommentDAO) InsertComment(ctx context.Context, data models.CommentModel) error {
	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 评论记录+1
		err := g.db.Model(&models.CommentModel{}).Create(&data).Error
		if err != nil {
			return err
		}
		// 2. 资源的评论数+1
		err = g.db.Model(&models.InteractiveModel{}).Where("source_id = ? and biz_type = ?", data.ResourceId, data.ResourceType).
			Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return err
	})

}
