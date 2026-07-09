package dto

import "go_bengkel/internal/models"


type VehicleRegister struct {
	PlateNum    string `json:"plate_num" binding:"required"`
	Range       int    `json:"range" binding:"required"` 
	VehicleType string `json:"vehicle_type" binding:"required"`
}

func ToVehicleRegister(vehicle models.Vehicle) VehicleRegister {
	return VehicleRegister{
		PlateNum: vehicle.PlateNum,
		Range: vehicle.Range,
		VehicleType: vehicle.Vehicle_Type,
	}
}