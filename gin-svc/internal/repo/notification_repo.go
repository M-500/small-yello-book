package repo

import (
	"context"
	"gin-svc/internal/models"
	"gin-svc/internal/repo/dao"
)

type NotifyRepo interface {
	GetNotifyNum(ctx context.Context, userID string) (int64, error)
	NotifyList(ctx context.Context, userID string, action string, offset, limit int) ([]*models.NotificationModel, error)
	LikeNotify(ctx context.Context, userID, authId, sourceId, sourceType string) error
	CancelLikeNotify(ctx context.Context, userID, authId, sourceId string) error
}

func NewNotifyRepo(dao dao.NotificationDao) NotifyRepo {
	return &notifyRepo{notifyDao: dao}
}

type notifyRepo struct {
	notifyDao dao.NotificationDao
}

// NotifyList
//
//	@Description: 通过用户ID和类型获取通知列表
//	@receiver n
//	@param ctx
//	@param userID
//	@param action
//	@param offset
//	@param limit
//	@return []*models.NotificationModel
//	@return error
func (n *notifyRepo) NotifyList(ctx context.Context, userID string, action string, offset, limit int) ([]*models.NotificationModel, error) {
	//TODO implement me
	panic("implement me")
}

// GetNotifyNum
//
//	@Description: 获取未读通知数量
//	@receiver n
//	@param ctx
//	@param userID
//	@return int64
//	@return error
func (n *notifyRepo) GetNotifyNum(ctx context.Context, userID string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

// CancelLikeNotify
//
//	@Description: 取消点赞 通知到用户
//	@receiver n
//	@param ctx
//	@param userID
//	@param authId
//	@param sourceId
//	@return error
func (n *notifyRepo) CancelLikeNotify(ctx context.Context, userID, authId, sourceId string) error {
	return n.notifyDao.Insert(ctx, models.NotificationModel{
		UpstreamUUID: userID,
		OwnerId:      authId,
		ResourceId:   sourceId,
		IsRead:       false,
		Type:         "CANCEL_LIKE",
	})
}

// LikeNotify
//
//	@Description: 点赞 通知到用户
//	@receiver n
//	@param ctx
//	@param userID  谁
//	@param authId  谁的资源
//	@param sourceId 资源ID
//	@return error
func (n *notifyRepo) LikeNotify(ctx context.Context, userID, authId, sourceId, sourceType string) error {
	return n.notifyDao.Insert(ctx, models.NotificationModel{
		UpstreamUUID: userID,
		OwnerId:      authId,
		ResourceId:   sourceId,
		ResourceType: sourceType,
		IsRead:       false,
		Type:         "LIKE",
	})
}
