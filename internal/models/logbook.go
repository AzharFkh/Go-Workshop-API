package models

import "gorm.io/gorm"

type DataService struct {
	gorm.Model

	VehicleID uint   `json:"vehicle_id"`
	PartName  string `json:"part_name"`
	Amount    int16  `json:"amount"`
}
