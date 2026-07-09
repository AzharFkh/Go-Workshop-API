package services

import (
	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
)

// masih kurang 3 fitur lagi yaa

type VehicleService interface {
	CreateVehicle(req dto.VehicleRegister, userID uint) (*models.Vehicle, error)
	// GetVehicle(userID uint) ([]models.Vehicle, error)
	// GetVehicleByID(id uint, userID uint) (*models.Vehicle, error)
	// DeleteVehicleByID(id uint, userID uint) error
}

type vehicleService struct {
	repo repository.VehicleRepository
}

func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	return &vehicleService{repo}
}

func (s *vehicleService) CreateVehicle(req dto.VehicleRegister, userID uint) (*models.Vehicle, error) {

	// mapping DTO to model
	vehicle := models.Vehicle{
		UserID:       userID,
		PlateNum:     req.PlateNum,
		Range:        req.Range,
		Vehicle_Type: req.VehicleType,
	}

	if err := s.repo.Create(&vehicle); err != nil {
		return nil, err
	}

	return &vehicle, nil
}
