package service

import (
	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type StockCardService struct {
	Repo *repository.StockCardRepository
}

func NewStockCardService(repo *repository.StockCardRepository) *StockCardService {
	return &StockCardService{Repo: repo}
}

func (s *StockCardService) CreateStockCard(stockCard *models.StockCard) error {
	return s.Repo.Create(stockCard)
}

func (s *StockCardService) GetStockCard(id uint64) (*models.StockCard, error) {
	return s.Repo.GetByID(id)
}

func (s *StockCardService) GetStockCards() ([]models.StockCard, error) {
	return s.Repo.GetAll()
}

func (s *StockCardService) UpdateStockCard(stockCard *models.StockCard) error {
	return s.Repo.Update(stockCard)
}

func (s *StockCardService) DeleteStockCard(id uint64) error {
	return s.Repo.Delete(id)
}
