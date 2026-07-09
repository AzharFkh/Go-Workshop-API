package services

import (
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
)

type DataService interface{
	GetVehicleServices(userID uint, vehicleID uint) ([]models.DataService, error)
}

type dataService struct{
	repo repository.ServiceRepository
}

func NewDataService(repo repository.ServiceRepository) DataService {
	return &dataService{repo}
}

func (s *dataService) GetVehicleServices(userID uint, vehicleID uint) ([]models.DataService, error) {
	return s.repo.GetServicesByUserAndVehicle(userID, vehicleID)
}