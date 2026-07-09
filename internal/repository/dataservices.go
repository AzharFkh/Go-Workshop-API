package repository

import (
	"go_bengkel/internal/models"

	"gorm.io/gorm"
)


type ServiceRepository interface {
	GetServicesByUserAndVehicle(userID uint, vehicleID uint) ([]models.DataService, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServicesRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) GetServicesByUserAndVehicle(userID uint, VehicleID uint) ([]models.DataService, error) {
	var services []models.DataService

	err := r.db.Table("data_services").
		Select("data_services.*"). 
		Joins("JOIN vehicle_data ON vehicle_data.id = data_services.vehicle_id"). 
		Where("vehicle_data.user_id = ? AND data_services.vehicle_id = ?", userID, VehicleID). 
		Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}