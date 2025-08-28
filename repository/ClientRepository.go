package repository

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/Barbarpotato/Storix/models"
	"gorm.io/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{DB: db}
}

func (r *ClientRepository) Create(client *models.Client) error {
	// generate random 16-char hex code
	bytes := make([]byte, 8) // 8 bytes = 16 hex chars
	if _, err := rand.Read(bytes); err != nil {
		return err
	}
	client.Code = hex.EncodeToString(bytes)

	// save to DB
	if err := r.DB.Create(client).Error; err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) GetByID(id uint64) (*models.Client, error) {
	var client models.Client
	if err := r.DB.First(&client, id).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) GetAll(page, pageSize int, sortBy, order string) ([]models.Client, int64, error) {
	var (
		clients []models.Client
		total   int64
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
	if err := r.DB.Model(&models.Client{}).Count(&total).Error; err != nil {
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
		Find(&clients).Error; err != nil {
		return nil, 0, err
	}

	return clients, total, nil
}

func (r *ClientRepository) Update(client *models.Client) error {

	return errors.New("update operation is not permitted on Client")

	// return r.DB.Save(client).Error
}

func (r *ClientRepository) Delete(id uint64) error {

	return errors.New("delete operation is not permitted on Client")

	// return r.DB.Delete(&models.Client{}, id).Error
}
