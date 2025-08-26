package service

import (
	"github.com/Barbarpotato/Storix/models"
	"github.com/Barbarpotato/Storix/repository"
)

type ClientService struct {
	Repo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{Repo: repo}
}

func (s *ClientService) CreateClient(client *models.Client) error {
	return s.Repo.Create(client)
}

func (s *ClientService) GetClient(id uint64) (*models.Client, error) {
	return s.Repo.GetByID(id)
}

func (s *ClientService) GetClients() ([]models.Client, error) {
	return s.Repo.GetAll()
}

func (s *ClientService) UpdateClient(client *models.Client) error {
	return s.Repo.Update(client)
}

func (s *ClientService) DeleteClient(id uint64) error {
	return s.Repo.Delete(id)
}
