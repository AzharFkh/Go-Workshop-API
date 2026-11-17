package repository

import (
	"errors"

	"go_bengkel/internal/models"

	"gorm.io/gorm"
)

var ErrVehicleNotFound = errors.New("vehicle not found or unauthorized")

type ServiceRepository interface {
	Create(dataService *models.DataService) error
	FindAll(userID, vehicleID uint) ([]models.DataService, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServicesRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) Create(dataService *models.DataService) error {
	return r.db.Create(dataService).Error
}

func (r *serviceRepository) FindAll(userID, vehicleID uint) ([]models.DataService, error) {
	var vehicleCount int64
	err := r.db.Model(&models.Vehicle{}).
		Where("id = ? AND user_id = ?", vehicleID, userID).
		Count(&vehicleCount).Error
	if err != nil {
		return nil, err
	}
	if vehicleCount == 0 {
		return nil, ErrVehicleNotFound
	}

	var services []models.DataService
	if err := r.db.Where("vehicle_id = ?", vehicleID).Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}
