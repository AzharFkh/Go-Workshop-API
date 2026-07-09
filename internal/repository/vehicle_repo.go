package repository

import (
	"go_bengkel/internal/models"

	"gorm.io/gorm"
)


type VehicleRepository interface {
	Create(vehicle *models.Vehicle) error
	//FindAll
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

