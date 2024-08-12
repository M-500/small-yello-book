package dao

import (
	"context"
	"errors"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type NoteDaoInterface interface {
	FindById(ctx context.Context, id int) (*models.NoteModel, error)
	ListByKey(ctx context.Context, keyword string, page, size int) ([]models.NoteModel, error)
	ListByStatus(ctx context.Context, status int, page, size int) ([]models.NoteModel, error)
	Insert(ctx context.Context, note *models.NoteModel) error
	Update(ctx context.Context, note *models.NoteModel) error
	Delete(ctx context.Context, id int) error
	HardDelete(ctx context.Context, id int) error
}

type noteDaoImpl struct {
	db *gorm.DB
}

func (n *noteDaoImpl) FindById(ctx context.Context, id int) (*models.NoteModel, error) {
	note := &models.NoteModel{}
	query := n.db.WithContext(ctx).Model(&models.NoteModel{}).Where("id = ?", id).First(note)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return note, nil
}

func (n *noteDaoImpl) ListByKey(ctx context.Context, keyword string, page, size int) ([]models.NoteModel, error) {
	var notes []models.NoteModel
	query := n.db.WithContext(ctx).Model(&models.NoteModel{}).Where("title like ?", "%"+keyword+"%").
		Limit(size).Offset((page - 1) * size).Find(&notes)
	if err := query.Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *noteDaoImpl) ListByStatus(ctx context.Context, status int, page, size int) ([]models.NoteModel, error) {
	var notes []models.NoteModel
	query := n.db.WithContext(ctx).Model(&models.NoteModel{}).
		Where("status = ?", status).Limit(size).Offset((page - 1) * size).
		Find(&notes)
	if err := query.Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *noteDaoImpl) Insert(ctx context.Context, note *models.NoteModel) error {
	query := n.db.WithContext(ctx).Create(note)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (n *noteDaoImpl) Update(ctx context.Context, note *models.NoteModel) error {
	query := n.db.WithContext(ctx).Model(&models.NoteModel{}).Where("id = ?", note.ID).Updates(note)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (n *noteDaoImpl) Delete(ctx context.Context, id int) error {
	return n.db.WithContext(ctx).Delete(&models.NoteModel{}, id).Error
}

func (n *noteDaoImpl) HardDelete(ctx context.Context, id int) error {
	return n.db.WithContext(ctx).Unscoped().Delete(&models.NoteModel{}, id).Error
}
