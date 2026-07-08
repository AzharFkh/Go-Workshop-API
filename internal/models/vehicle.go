package models

import "gorm.io/gorm"

type VehicleData struct {
	gorm.Model
	UserID uint `json:"user_id"`
	PlateNum string `json:"plate_num" gorm:"uniqueIndex"`
	Range int `json:"range"`
	Vehicle_Type string `json:"vehicle_type"`

	// one to many
	DataServices []DataService `json:"data_services" gorm:"foreignKey:VehicleID"`
}
