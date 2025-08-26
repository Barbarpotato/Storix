package repository

import (
	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) Create(item *models.Item) error {
	return r.DB.Create(item).Error
}

func (r *ItemRepository) GetByID(id uint64) (*models.Item, error) {
	var item models.Item
	if err := r.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetAll() ([]models.Item, error) {
	var items []models.Item
	if err := r.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	return r.DB.Save(item).Error
}

func (r *ItemRepository) Delete(id uint64) error {
	return r.DB.Delete(&models.Item{}, id).Error
}
