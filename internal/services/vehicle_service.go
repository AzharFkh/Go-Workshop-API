package services

import (
	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
)

type VehicleService interface {
	CreateVehicle(req dto.VehicleRegister, userID uint) (*models.Vehicle, error)
	GetVehicle(userID uint) ([]models.Vehicle, error)
}

type vehicleService struct {
	repo repository.VehicleRepository
}

func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	return &vehicleService{repo: repo}
}

func (s *vehicleService) CreateVehicle(req dto.VehicleRegister, userID uint) (*models.Vehicle, error) {
	vehicle := models.Vehicle{
		UserID:      userID,
		PlateNum:    req.PlateNum,
		Range:       req.Range,
		VehicleType: req.VehicleType,
	}

	if err := s.repo.Create(&vehicle); err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (s *vehicleService) GetVehicle(userID uint) ([]models.Vehicle, error) {
	return s.repo.FindAll(userID)
}
