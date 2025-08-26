package repository

import (
	"github.com/Barbarpotato/Storix/models"

	"gorm.io/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{DB: db}
}

func (r *ClientRepository) Create(client *models.Client) error {
	return r.DB.Create(client).Error
}

func (r *ClientRepository) GetByID(id uint64) (*models.Client, error) {
	var client models.Client
	if err := r.DB.First(&client, id).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) GetAll() ([]models.Client, error) {
	var clients []models.Client
	if err := r.DB.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *ClientRepository) Update(client *models.Client) error {
	return r.DB.Save(client).Error
}

func (r *ClientRepository) Delete(id uint64) error {
	return r.DB.Delete(&models.Client{}, id).Error
}
