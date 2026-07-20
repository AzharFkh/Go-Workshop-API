package repository

import (
	"go_bengkel/internal/models"

	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(dataService *models.DataService) error
	FindAll(userID uint, vehicleID uint) ([]models.DataService, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServicesRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) Create(dataService *models.DataService) error {
	return r.db.Create(dataService).Error
}

func (r *serviceRepository) FindAll(userID uint, vehicleID uint) ([]models.DataService, error) {
	var services []models.DataService

	err := r.db.Table("data_services").
		Select("data_services.*").
		Joins("JOIN vehicles ON vehicles.id = data_services.vehicle_id").
		Where("vehicles.user_id = ? AND data_services.vehicle_id = ?", userID, vehicleID).
		Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}
