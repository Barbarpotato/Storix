package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) Create(item *models.Item, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}

	// query the client
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return errors.New("client_code not found")
	}

	// Validate UnitID
	var unit models.Unit
	if err := r.DB.Where("id = ?", item.UnitID).Take(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("invalid unit_id %d: record not found", item.UnitID)
		}
		return fmt.Errorf("failed to validate unit_id %d: %w", item.UnitID, err)
	}

	// set client_id on item
	item.ClientID = client.ID

	// Inject unit info into item
	item.UnitCode = unit.Code
	item.UnitName = unit.Name

	// insert into DB with error guard
	if err := r.DB.Create(item).Error; err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetByID(id uint64, clientCode string) (*models.Item, error) {
	if clientCode == "" {
		return nil, errors.New("client_code is required")
	}

	// Step 1: query the client
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return nil, errors.New("client_code not found")
	}

	// Step 2: query the item with client_id
	var item models.Item
	if err := r.DB.Where("id = ? AND client_id = ?", id, client.ID).First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) GetAll(clientCode string, page, pageSize int, sortBy, order string) ([]models.Item, int64, error) {
	var (
		items []models.Item
		total int64
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
	if err := r.DB.Model(&models.Item{}).
		Where("client_id = ?", client.ID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ---- allowed sort fields ----
	allowedSort := map[string]bool{
		"name":       true,
		"number":     true,
		"created_at": true,
	}
	if !allowedSort[sortBy] {
		sortBy = "name" // default sort
	}

	// ---- validate order ----
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	// ---- query items with client_id, pagination + sorting ----
	offset := (page - 1) * pageSize
	if err := r.DB.
		Where("client_id = ?", client.ID).
		Order(sortBy + " " + order).
		Limit(pageSize).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
func (r *ItemRepository) SetActive(ctx context.Context, id uint64, clientCode string) error {
	// Step 1: Get client by code
	var client models.Client
	if err := r.DB.WithContext(ctx).
		Where("code = ?", clientCode).
		Take(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("client with code %s not found", clientCode)
		}
		return fmt.Errorf("failed to query client: %w", err)
	}

	// Step 2: Get item by id + client_id
	var item models.Item
	if err := r.DB.WithContext(ctx).
		Where("id = ? AND client_id = ?", id, client.ID).
		Take(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("item with id %d for client %s not found", id, clientCode)
		}
		return fmt.Errorf("failed to query item: %w", err)
	}

	if item.IsActive {
		return fmt.Errorf("item with id %d is already active", id)
	}

	item.IsActive = true

	// Generate Number if draft
	if item.Number == "DRAFT" {
		now := time.Now()
		item.Number = fmt.Sprintf("ITM-%s-%05d", now.Format("20060102"), item.ID)
	}

	if err := r.DB.WithContext(ctx).Save(&item).Error; err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

func (r *ItemRepository) SetInactive(ctx context.Context, id uint64, clientCode string) error {
	// Step 1: Get client by code
	var client models.Client
	if err := r.DB.WithContext(ctx).
		Where("code = ?", clientCode).
		Take(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("client with code %s not found", clientCode)
		}
		return fmt.Errorf("failed to query client: %w", err)
	}

	// Step 2: Get item by id + client_id
	var item models.Item
	if err := r.DB.WithContext(ctx).
		Where("id = ? AND client_id = ?", id, client.ID).
		Take(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("item with id %d for client %s not found", id, clientCode)
		}
		return fmt.Errorf("failed to query item: %w", err)
	}

	if !item.IsActive {
		// already inactive, nothing to do
		return nil
	}

	item.IsActive = false

	if err := r.DB.WithContext(ctx).Save(&item).Error; err != nil {
		return fmt.Errorf("failed to deactivate item: %w", err)
	}

	return nil
}

func (r *ItemRepository) Update(item *models.Item, clientCode string) error {
	if item == nil {
		return errors.New("item is required")
	}
	if clientCode == "" {
		return errors.New("client_code is required")
	}

	// validate client exists
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return errors.New("client_code not found")
	}

	// update only name and description
	return r.DB.Model(&models.Item{}).
		Where("id = ? AND client_id = ?", item.ID, client.ID).
		Updates(map[string]interface{}{
			"name":        item.Name,
			"description": item.Description,
		}).Error
}

func (r *ItemRepository) Delete(id uint64, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}

	// validate client exists
	var client models.Client
	if err := r.DB.Where("code = ?", clientCode).First(&client).Error; err != nil {
		return errors.New("client_code not found")
	}

	// get the item for this client
	var item models.Item
	if err := r.DB.Where("id = ? AND client_id = ?", id, client.ID).First(&item).Error; err != nil {
		return errors.New("item not found for this client")
	}

	// cannot delete if item is active
	if item.IsActive {
		return errors.New("cannot delete active item; set inactive first")
	}

	// check if item exists in stockcard
	var count int64
	if err := r.DB.Model(&models.StockCard{}).
		Where("item_id = ?", item.ID).
		Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check stockcard: %w", err)
	}

	if count > 0 {
		return errors.New("cannot delete item; it is referenced in stockcard")
	}

	// delete the item
	if err := r.DB.Delete(&item).Error; err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
