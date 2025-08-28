package repository

import (
	"strings"

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

func (r *UnitRepository) GetAll(page, pageSize int, sortBy, order string) ([]models.Unit, int64, error) {
	var (
		units []models.Unit
		total int64
	)

	// ---- pagination defaults ----
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10 // default perPage
	}
	if pageSize > 100 {
		pageSize = 100 // max perPage
	}

	// ---- count total ----
	if err := r.DB.Model(&models.Unit{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ---- allowed sort fields ----
	allowedSort := map[string]bool{
		"name":       true,
		"code":       true,
		"created_at": true,
	}
	if !allowedSort[sortBy] {
		sortBy = "name" // default sort field
	}

	// ---- validate order ----
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	// ---- query with pagination + sorting ----
	offset := (page - 1) * pageSize
	if err := r.DB.
		Order(sortBy + " " + order).
		Limit(pageSize).
		Offset(offset).
		Find(&units).Error; err != nil {
		return nil, 0, err
	}

	return units, total, nil
}

func (r *UnitRepository) Update(unit *models.Unit) error {
	return r.DB.Save(unit).Error
}

func (r *UnitRepository) Delete(id uint64) error {
	return r.DB.Delete(&models.Unit{}, id).Error
}
