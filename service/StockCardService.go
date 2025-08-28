package service

import (
	"errors"
	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type StockCardService struct {
	Repo *repository.StockCardRepository
}

func NewStockCardService(repo *repository.StockCardRepository) *StockCardService {
	return &StockCardService{Repo: repo}
}

func (s *StockCardService) CreateStockCard(stockCard *models.StockCard, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}
	return s.Repo.Create(stockCard, clientCode)
}

func (s *StockCardService) GetStockCard(id uint64, clientCode string) (*models.StockCard, error) {
	if clientCode == "" {
		return nil, errors.New("client_code is required")
	}
	return s.Repo.GetByID(id, clientCode)
}

func (s *StockCardService) GetStockCards(page, pageSize int, sortBy, order, clientCode string) ([]models.StockCard, int64, error) {
	if clientCode == "" {
		return nil, 0, errors.New("client_code is required")
	}
	return s.Repo.GetAll(clientCode, page, pageSize, sortBy, order)
}

func (s *StockCardService) UpdateStockCard(stockCard *models.StockCard) error {
	return s.Repo.Update(stockCard)
}

func (s *StockCardService) DeleteStockCard(id uint64, clientCode string) error {
	if clientCode == "" {
		return errors.New("client_code is required")
	}
	return s.Repo.Delete(id, clientCode)
}
