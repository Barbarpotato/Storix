package service

import (
	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type UnitService struct {
	Repo *repository.UnitRepository
}

func NewUnitService(repo *repository.UnitRepository) *UnitService {
	return &UnitService{Repo: repo}
}

func (s *UnitService) CreateUnit(unit *models.Unit) error {
	return s.Repo.Create(unit)
}

func (s *UnitService) GetUnit(id uint64) (*models.Unit, error) {
	return s.Repo.GetByID(id)
}

func (s *UnitService) GetUnits(page, pageSize int, sortBy, order string) ([]models.Unit, int64, error) {
	return s.Repo.GetAll(page, pageSize, sortBy, order)
}

func (s *UnitService) UpdateUnit(unit *models.Unit) error {
	return s.Repo.Update(unit)
}

func (s *UnitService) DeleteUnit(id uint64) error {
	return s.Repo.Delete(id)
}
