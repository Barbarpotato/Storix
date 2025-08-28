package repository

import (
	"errors"
	"strings"

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

func (r *WarehouseRepository) GetByID(id uint64, clientCode string) (*models.Warehouse, error) {
	if clientCode == "" {
		return nil, errors.New("client_code is required")
	}

	// Step 1: query the client
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return nil, errors.New("client_code not found")
	}

	// Step 2: query the warehouse with client_id
	var warehouse models.Warehouse
	if err := r.DB.Where("id = ? AND client_id = ?", id, client.ID).First(&warehouse).Error; err != nil {
		return nil, err
	}

	return &warehouse, nil
}

func (r *WarehouseRepository) GetAll(clientCode string, page, pageSize int, sortBy, order string) ([]models.Warehouse, int64, error) {
	var (
		warehouses []models.Warehouse
		total      int64
	)

	// ---- client_code validation ----
	if clientCode == "" {
		return nil, 0, errors.New("client_code is required")
	}

	// query client table
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return nil, 0, errors.New("client_code not found")
	}

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

	// ---- count total for this client ----
	if err := r.DB.Model(&models.Warehouse{}).
		Where("client_id = ?", client.ID).
		Count(&total).Error; err != nil {
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
		Where("client_id = ?", client.ID).
		Order(sortBy + " " + order).
		Limit(pageSize).
		Offset(offset).
		Find(&warehouses).Error; err != nil {
		return nil, 0, err
	}

	return warehouses, total, nil
}

func (r *WarehouseRepository) Update(warehouse *models.Warehouse) error {
	return r.DB.Save(warehouse).Error
}

func (r *WarehouseRepository) Delete(id uint64, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}

	// validate client exists
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return errors.New("client_code not found")
	}

	// get the warehouse for this client
	var warehouse models.Warehouse
	if err := r.DB.Where("id = ? AND client_id = ?", id, client.ID).First(&warehouse).Error; err != nil {
		return errors.New("warehouse not found for this client")
	}

	// delete the warehouse
	if err := r.DB.Delete(&warehouse).Error; err != nil {
		return err
	}

	return nil
}
