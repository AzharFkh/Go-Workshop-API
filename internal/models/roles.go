package models

type Role struct {
	ID   uint
	Name string `gorm:"unique"` // admin, user
}
