package dto

import "go_bengkel/internal/models"

type VehicleResponse struct {
	ID          uint   `json:"vehicle_id"`
	UserID      uint   `json:"user_id"`
	PlateNum    string `json:"plate_num"`
	Range       int    `json:"range"`
	VehicleType string `json:"vehicle_type"`
}

func ToVehicleResponse(vehicle models.Vehicle) VehicleResponse {
	return VehicleResponse{
		ID:          vehicle.ID,
		UserID:      vehicle.UserID,
		PlateNum:    vehicle.PlateNum,
		Range:       vehicle.Range,
		VehicleType: vehicle.VehicleType,
	}
}

func ToVehicleResponses(vehicles []models.Vehicle) []VehicleResponse {
	responses := make([]VehicleResponse, 0, len(vehicles))
	for _, vehicle := range vehicles {
		responses = append(responses, ToVehicleResponse(vehicle))
	}
	return responses
}
