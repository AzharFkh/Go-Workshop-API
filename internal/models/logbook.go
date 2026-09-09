package models

import "gorm.io/gorm"


type DataService struct {
	gorm.Model

	// Foreign Key ke tabel vehicle_data
	VehicleID uint   `json:"vehicle_id"`
	PartName  string `json:"part_name"`
	Amount    int16  `json:"amount"`
}

