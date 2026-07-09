package models

import (
	"strings"

	"gorm.io/gorm"
)

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

func (v *Vehicle) BeforeCreate(tx *gorm.DB) error {
	v.PlateNum = normalizePlateNum(v.PlateNum)
	return nil
}

func (v *Vehicle) BeforeUpdate(tx *gorm.DB) error {
	v.PlateNum = normalizePlateNum(v.PlateNum)
	return nil
}

func normalizePlateNum(plate string) string {
	// Hapus semua spasi
	noSpace := strings.ReplaceAll(plate, " ", "")

	// Ubah menjadi huruf kapital semua
	return strings.ToUpper(noSpace)
}
