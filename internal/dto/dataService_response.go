package dto

import "go_bengkel/internal/models"

type DataServiceResponse struct {
	ID        uint   `json:"id"`
	VehicleID uint   `json:"vehicle_id"`
	PartName  string `json:"part_name" binding:"required"`
	Amount    int16  `json:"amount" binding:"required"`
}

func ToDataServiceResponse(dataService models.DataService) DataServiceResponse {
	return DataServiceResponse{
		ID:        dataService.ID,
		VehicleID: dataService.VehicleID,
		PartName:  dataService.PartName,
		Amount:    dataService.Amount,
	}
}

func ToDataServiceResponses(dataServices []models.DataService) []DataServiceResponse {
	responses := []DataServiceResponse{}

	for _, dataService := range dataServices {
		responses = append(responses, ToDataServiceResponse(dataService))
	}
	return responses
}
