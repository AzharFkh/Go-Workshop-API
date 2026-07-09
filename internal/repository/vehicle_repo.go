package repository

import (
	"go_bengkel/internal/models"

	"gorm.io/gorm"
)


type VehicleRepository interface {
	Create(vehicle *models.Vehicle) error
	FindAll(userID uint) ([]models.Vehicle, error)
	// tambah query lainnya
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository{
	return &vehicleRepository{db}
}

func (r *vehicleRepository) Create(vehicle *models.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *vehicleRepository) FindAll(userID uint) ([]models.Vehicle, error) {
	
	var vehicles []models.Vehicle

	if err := r.db.Where("user_id = ?", userID).Find(&vehicles).Error; err != nil {
		return nil, err
	}

	return vehicles, nil 
}

