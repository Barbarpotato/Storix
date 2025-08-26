package service

import (
	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type ItemService struct {
	Repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) CreateItem(item *models.Item) error {
	return s.Repo.Create(item)
}

func (s *ItemService) GetItem(id uint64) (*models.Item, error) {
	return s.Repo.GetByID(id)
}

func (s *ItemService) GetItems() ([]models.Item, error) {
	return s.Repo.GetAll()
}

func (s *ItemService) UpdateItem(item *models.Item) error {
	return s.Repo.Update(item)
}

func (s *ItemService) DeleteItem(id uint64) error {
	return s.Repo.Delete(id)
}
