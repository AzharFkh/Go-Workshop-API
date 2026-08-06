package dto

import "go_bengkel/internal/models"


type VehicleRegister struct {
	PlateNum    string `json:"plate_num" binding:"required" example:"B1234XYZ"`
	Range       int    `json:"range" binding:"required" example:"15000"` 
	VehicleType string `json:"vehicle_type" binding:"required" example:"motor"`
}

func ToVehicleRegister(vehicle models.Vehicle) VehicleRegister {
	return VehicleRegister{
		PlateNum: vehicle.PlateNum,
		Range: vehicle.Range,
		VehicleType: vehicle.Vehicle_Type,
	}
}