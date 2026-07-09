package dto

import "go_bengkel/internal/models"

type VehicleResponse struct {
	UserID       uint   `json:"user_id" binding:"required"`
	PlateNum     string `json:"plate_num" binding:"required"`
	Range        int	`json:"range" binding:"required"`
	Vehicle_Type string `json:"vehicle_type" binding:"required"`
}

func ToVehicleResponse(vehicle models.Vehicle) VehicleResponse {

	return VehicleResponse{
		UserID:       vehicle.UserID,
		PlateNum:     vehicle.PlateNum,
		Range:        vehicle.Range,
		Vehicle_Type: vehicle.Vehicle_Type,
	}
}

func ToVehicleResponses(vehicle []models.Vehicle) []VehicleResponse {
	var responses []VehicleResponse

	for _, vehicle := range vehicle {
		responses = append(responses, ToVehicleResponse(vehicle))
	}

	return responses
}
