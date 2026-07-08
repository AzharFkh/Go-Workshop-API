package models

import "gorm.io/gorm"

// add gorm for easier write to db

// data here
type DataService struct {
	gorm.Model

	VehicleID uint   `json:"vehicle_id"`
	PartName  string `json:"part_name"`
	Amount    int16  `json:"amount"`
}

// read data and show from db

// write data to db

// patch update and delete
