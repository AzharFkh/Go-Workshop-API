package repository

import (
	"go_bengkel/internal/models"

	"gorm.io/gorm"
)

type VehicleRepository interface {
	Create(vehicle *models.Vehicle) error
	FindAll(userID uint) ([]models.Vehicle, error)
	CheckOwnership(userID uint, vehicleID uint) (bool, error)
	// tambah query lainnya
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
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

func (r *vehicleRepository) CheckOwnership(userID uint, vehicleID uint) (bool, error) {
	var count int64

	err := r.db.Model(&models.Vehicle{}).
		Where("id=? AND user_id = ?", vehicleID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
