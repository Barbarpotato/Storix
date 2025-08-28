package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type ItemService struct {
	Repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) CreateItem(item *models.Item, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}
	return s.Repo.Create(item, clientCode)
}

func (s *ItemService) GetItem(id uint64, clientCode string) (*models.Item, error) {
	if clientCode == "" {
		return nil, errors.New("client_code is required")
	}
	return s.Repo.GetByID(id, clientCode)
}

func (s *ItemService) GetItems(page, pageSize int, sortBy, order, clientCode string) ([]models.Item, int64, error) {
	if clientCode == "" {
		return nil, 0, errors.New("client_code is required")
	}
	return s.Repo.GetAll(clientCode, page, pageSize, sortBy, order)
}

func (s *ItemService) SetActive(ctx context.Context, id uint64, clientCode string) error {
	if err := s.Repo.SetActive(ctx, id, clientCode); err != nil {
		return fmt.Errorf("service: failed to activate item: %w", err)
	}
	return nil
}

func (s *ItemService) SetInactive(ctx context.Context, id uint64, clientCode string) error {
	if err := s.Repo.SetInactive(ctx, id, clientCode); err != nil {
		return fmt.Errorf("service: failed to deactivate item: %w", err)
	}
	return nil
}

func (s *ItemService) UpdateItem(item *models.Item, clientCode string) error {
	return s.Repo.Update(item, clientCode)
}

func (s *ItemService) DeleteItem(id uint64, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}
	return s.Repo.Delete(id, clientCode)
}
