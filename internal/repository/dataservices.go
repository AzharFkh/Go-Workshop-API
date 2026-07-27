package repository

import (
	"errors"
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
	var ErrVehicleNotFound = errors.New("vehicle not found or unauthorized")
	var vehicleCount int64

	err := r.db.Table("vehicles").
		Where("id = ? AND user_id = ?", vehicleID, userID).
		Count(&vehicleCount).Error

	if err != nil {
		return nil, err
	}

	if vehicleCount == 0 {
		return nil, ErrVehicleNotFound
	}

	err = r.db.Table("data_services").
			Where("vehicle_id = ?", vehicleID).
			Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}
