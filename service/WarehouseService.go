package service

import (
	"errors"
	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type WarehouseService struct {
	Repo *repository.WarehouseRepository
}

func NewWarehouseService(repo *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{Repo: repo}
}

func (s *WarehouseService) CreateWarehouse(warehouse *models.Warehouse) error {
	return s.Repo.Create(warehouse)
}

func (s *WarehouseService) GetWarehouse(id uint64, clientCode string) (*models.Warehouse, error) {
	if clientCode == "" {
		return nil, errors.New("client_code is required")
	}
	return s.Repo.GetByID(id, clientCode)
}

func (s *WarehouseService) GetWarehouses(page, pageSize int, sortBy, order, clientCode string) ([]models.Warehouse, int64, error) {
	if clientCode == "" {
		return nil, 0, errors.New("client_code is required")
	}
	return s.Repo.GetAll(clientCode, page, pageSize, sortBy, order)
}

func (s *WarehouseService) UpdateWarehouse(warehouse *models.Warehouse) error {
	return s.Repo.Update(warehouse)
}

func (s *WarehouseService) DeleteWarehouse(id uint64, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}
	return s.Repo.Delete(id, clientCode)
}
