package service

import (
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

func (s *WarehouseService) GetWarehouse(id uint64) (*models.Warehouse, error) {
	return s.Repo.GetByID(id)
}

func (s *WarehouseService) GetWarehouses() ([]models.Warehouse, error) {
	return s.Repo.GetAll()
}

func (s *WarehouseService) UpdateWarehouse(warehouse *models.Warehouse) error {
	return s.Repo.Update(warehouse)
}

func (s *WarehouseService) DeleteWarehouse(id uint64) error {
	return s.Repo.Delete(id)
}
