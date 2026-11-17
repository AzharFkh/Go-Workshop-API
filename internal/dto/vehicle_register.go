package dto

type VehicleRegister struct {
	PlateNum    string `json:"plate_num" binding:"required" example:"B1234XYZ"`
	Range       int    `json:"range" binding:"required" example:"15000"`
	VehicleType string `json:"vehicle_type" binding:"required" example:"motor"`
}
