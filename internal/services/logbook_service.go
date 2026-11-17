package services

import (
	"errors"

	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
)

type DataService interface {
	CreateDataService(req dto.DataServiceRegister, userID, vehicleID uint) (*models.DataService, error)
	GetDataServices(userID, vehicleID uint) ([]models.DataService, error)
}

type dataService struct {
	serviceRepo repository.ServiceRepository
	vehicleRepo repository.VehicleRepository
}

func NewDataService(serviceRepo repository.ServiceRepository, vehicleRepo repository.VehicleRepository) DataService {
	return &dataService{
		serviceRepo: serviceRepo,
		vehicleRepo: vehicleRepo,
	}
}

func (s *dataService) CreateDataService(req dto.DataServiceRegister, userID, vehicleID uint) (*models.DataService, error) {
	isOwned, err := s.vehicleRepo.CheckOwnership(userID, vehicleID)
	if err != nil {
		return nil, err
	}
	if !isOwned {
		return nil, ErrVehicleNotOwned
	}

	record := models.DataService{
		VehicleID: vehicleID,
		PartName:  req.PartName,
		Amount:    req.Amount,
	}

	if err := s.serviceRepo.Create(&record); err != nil {
		return nil, err
	}

	return &record, nil
}

func (s *dataService) GetDataServices(userID, vehicleID uint) ([]models.DataService, error) {
	records, err := s.serviceRepo.FindAll(userID, vehicleID)
	if err != nil {
		if errors.Is(err, repository.ErrVehicleNotFound) {
			return nil, ErrVehicleNotFound
		}
		return nil, err
	}
	return records, nil
}
