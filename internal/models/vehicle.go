package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	// Foreign Key ke tabel users
	UserID       uint   `json:"user_id"`
	PlateNum     string `json:"plate_num" gorm:"uniqueIndex"`
	Range        int    `json:"range"`
	Vehicle_Type string `json:"vehicle_type"`

	// one to many
	DataServices []DataService `json:"data_services" gorm:"foreignKey:VehicleID"`
}
