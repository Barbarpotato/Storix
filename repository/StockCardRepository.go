package repository

import (
	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type StockCardRepository struct {
	DB *gorm.DB
}

func NewStockCardRepository(db *gorm.DB) *StockCardRepository {
	return &StockCardRepository{DB: db}
}

func (r *StockCardRepository) Create(stockCard *models.StockCard) error {
	return r.DB.Create(stockCard).Error
}

func (r *StockCardRepository) GetByID(id uint64) (*models.StockCard, error) {
	var stockCard models.StockCard
	if err := r.DB.First(&stockCard, id).Error; err != nil {
		return nil, err
	}
	return &stockCard, nil
}

func (r *StockCardRepository) GetAll() ([]models.StockCard, error) {
	var stockCards []models.StockCard
	if err := r.DB.Find(&stockCards).Error; err != nil {
		return nil, err
	}
	return stockCards, nil
}

func (r *StockCardRepository) Update(stockCard *models.StockCard) error {
	return r.DB.Save(stockCard).Error
}

func (r *StockCardRepository) Delete(id uint64) error {
	return r.DB.Delete(&models.StockCard{}, id).Error
}
