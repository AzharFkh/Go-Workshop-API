package models

import (
	"strings"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model

	UserID       uint          `json:"user_id"`
	PlateNum     string        `json:"plate_num" gorm:"uniqueIndex"`
	Range        int           `json:"range"`
	VehicleType  string        `json:"vehicle_type"`
	DataServices []DataService `json:"data_services,omitempty" gorm:"foreignKey:VehicleID"`
}

func (v *Vehicle) BeforeCreate(_ *gorm.DB) error {
	v.PlateNum = normalizePlateNum(v.PlateNum)
	return nil
}

func (v *Vehicle) BeforeUpdate(_ *gorm.DB) error {
	v.PlateNum = normalizePlateNum(v.PlateNum)
	return nil
}

func normalizePlateNum(plate string) string {
	noSpace := strings.ReplaceAll(plate, " ", "")
	return strings.ToUpper(noSpace)
}
