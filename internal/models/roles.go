package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"unique"` // admin, user
}
