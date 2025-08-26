package repository

import (
	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type WarehouseRepository struct {
	DB *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{DB: db}
}

func (r *WarehouseRepository) Create(warehouse *models.Warehouse) error {
	return r.DB.Create(warehouse).Error
}

func (r *WarehouseRepository) GetByID(id uint64) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	if err := r.DB.First(&warehouse, id).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *WarehouseRepository) GetAll() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	if err := r.DB.Find(&warehouses).Error; err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (r *WarehouseRepository) Update(warehouse *models.Warehouse) error {
	return r.DB.Save(warehouse).Error
}

func (r *WarehouseRepository) Delete(id uint64) error {
	return r.DB.Delete(&models.Warehouse{}, id).Error
}
