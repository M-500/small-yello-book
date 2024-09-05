package dao

import (
	"context"
	"gin-svc/internal/models"
	"gorm.io/gorm"
)

type NoteDaoInterface interface {
	FindById(ctx context.Context, id int) (models.NoteModel, error)
	ListByKey(ctx context.Context, keyword string, page, size int) ([]models.NoteModel, error)
	ListByStatus(ctx context.Context, status int, page, size int) ([]models.NoteModel, error)
	Insert(ctx context.Context, note models.NoteModel) error
	SaveNoteWithTx(ctx context.Context, note models.NoteModel, imgList []models.ImageModel) error
	Update(ctx context.Context, note *models.NoteModel) error
	Delete(ctx context.Context, id int) error
	HardDelete(ctx context.Context, id int) error

	FeedNoteList(ctx context.Context, tagID int, page, size int) ([]models.NoteModel, error)
}

func NewNoteDao(db *gorm.DB) NoteDaoInterface {
	return NewNoteDaoImpl(db)
}

type noteDaoImpl struct {
	db *gorm.DB
}

func (n *noteDaoImpl) FeedNoteList(ctx context.Context, tagID int, page, size int) ([]models.NoteModel, error) {
	// TODO: 这里有问题，其实应该从大数据平台/推荐算法平台获取数据
	notes := make([]models.NoteModel, 0, size)
	query := n.db.WithContext(ctx).Model(&models.NoteModel{}).Where("status = ?", 1).Limit(size).Offset((page - 1) * size).Find(&notes)
	if query.Error != nil {
		return nil, query.Error
	}
	return notes, nil
}

func NewNoteDaoImpl(db *gorm.DB) NoteDaoInterface {
	return &noteDaoImpl{db: db}
}

func (n *noteDaoImpl) SaveNoteWithTx(ctx context.Context, note models.NoteModel, imgList []models.ImageModel) error {
	return n.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, img := range imgList {
			err := tx.Model(&models.ImageModel{}).Create(&img).Error
			if err != nil {
				return err
			}
		}
		return tx.Model(&models.NoteModel{}).Create(&note).Error
	})
}

func (n *noteDaoImpl) FindById(ctx context.Context, id int) (models.NoteModel, error) {
	note := models.NoteModel{}
	query := n.db.WithContext(ctx).Model(&models.NoteModel{}).Where("id = ?", id).First(&note)
	if query.Error != nil {
		return note, query.Error
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

func (n *noteDaoImpl) Insert(ctx context.Context, note models.NoteModel) error {
	query := n.db.WithContext(ctx).Create(note)
	if err := query.Error; err != nil {
		// 主键冲突？ 没主键冲突吧
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
