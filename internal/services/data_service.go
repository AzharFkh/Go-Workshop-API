package services

import (
	"errors"
	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
)

type DataService interface {
	CreateDataService(req dto.DataServiceRegister, userID uint, vehicleID uint) (*models.DataService, error)
	GetDataServices(userID uint, vehicleID uint) ([]models.DataService, error)
}

type dataService struct {
	serviceRepo repository.ServiceRepository
	vehicleRepo repository.VehicleRepository
}

func NewDataService(serviceRepo repository.ServiceRepository, vehicleRepo repository.VehicleRepository) DataService {
	return &dataService{serviceRepo, vehicleRepo}
}

func (s *dataService) CreateDataService(req dto.DataServiceRegister, userID uint, vehicleID uint) (*models.DataService, error) {

	isOwned, err := s.vehicleRepo.CheckOwnership(userID, vehicleID)

	if err != nil {
		return nil, err
	}
	if !isOwned {
		return nil, errors.New("vehicle is not owned by user")
	}
	dataService := models.DataService{
		VehicleID: vehicleID,
		PartName:  req.PartName,
		Amount:    req.Amount,
	}

	if err := s.serviceRepo.Create(&dataService); err != nil {
		return nil, err
	}

	return &dataService, nil

}

func (s *dataService) GetDataServices(userID uint, vehicleID uint) ([]models.DataService, error) {
	return s.serviceRepo.FindAll(userID, vehicleID)
}
