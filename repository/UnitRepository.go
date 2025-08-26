package repository

import (
	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type UnitRepository struct {
	DB *gorm.DB
}

func NewUnitRepository(db *gorm.DB) *UnitRepository {
	return &UnitRepository{DB: db}
}

func (r *UnitRepository) Create(unit *models.Unit) error {
	return r.DB.Create(unit).Error
}

func (r *UnitRepository) GetByID(id uint64) (*models.Unit, error) {
	var unit models.Unit
	if err := r.DB.First(&unit, id).Error; err != nil {
		return nil, err
	}
	return &unit, nil
}

func (r *UnitRepository) GetAll() ([]models.Unit, error) {
	var units []models.Unit
	if err := r.DB.Find(&units).Error; err != nil {
		return nil, err
	}
	return units, nil
}

func (r *UnitRepository) Update(unit *models.Unit) error {
	return r.DB.Save(unit).Error
}

func (r *UnitRepository) Delete(id uint64) error {
	return r.DB.Delete(&models.Unit{}, id).Error
}
