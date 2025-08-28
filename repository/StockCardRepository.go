package repository

import (
	"errors"
	"strings"

	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type StockCardRepository struct {
	DB *gorm.DB
}

func NewStockCardRepository(db *gorm.DB) *StockCardRepository {
	return &StockCardRepository{DB: db}
}

func (r *StockCardRepository) Create(stockCard *models.StockCard, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}

	// Step 1: query the client
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return errors.New("client_code not found")
	}

	// Step 2: set client_id on stockCard
	stockCard.ClientID = client.ID

	// Step 3: create stockCard
	return r.DB.Create(stockCard).Error
}

func (r *StockCardRepository) GetByID(id uint64, clientCode string) (*models.StockCard, error) {
	if clientCode == "" {
		return nil, errors.New("client_code is required")
	}

	// Step 1: query the client
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return nil, errors.New("client_code not found")
	}

	// Step 2: query the stock card with client_id
	var stockCard models.StockCard
	if err := r.DB.Where("id = ? AND client_id = ?", id, client.ID).First(&stockCard).Error; err != nil {
		return nil, err
	}

	return &stockCard, nil
}

func (r *StockCardRepository) GetAll(clientCode string, page, pageSize int, sortBy, order string) ([]models.StockCard, int64, error) {
	var (
		stockCards []models.StockCard
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
	if err := r.DB.Model(&models.StockCard{}).
		Where("client_id = ?", client.ID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ---- allowed sort fields ----
	allowedSort := map[string]bool{
		"date":       true,
		"created_at": true,
	}
	if !allowedSort[sortBy] {
		sortBy = "date" // default sort field
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
		Find(&stockCards).Error; err != nil {
		return nil, 0, err
	}

	return stockCards, total, nil
}

func (r *StockCardRepository) Update(stockCard *models.StockCard) error {
	return r.DB.Save(stockCard).Error
}

func (r *StockCardRepository) Delete(id uint64, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}

	// validate client exists
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return errors.New("client_code not found")
	}

	// get the stock card for this client
	var stockCard models.StockCard
	if err := r.DB.Where("id = ? AND client_id = ?", id, client.ID).First(&stockCard).Error; err != nil {
		return errors.New("stock card not found for this client")
	}

	// delete the stock card
	if err := r.DB.Delete(&stockCard).Error; err != nil {
		return err
	}

	return nil
}
