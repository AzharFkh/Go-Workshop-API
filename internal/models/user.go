package models

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex:idx_email"`
	Password string    `json:"-"`
	RoleID   uint      `json:"role_id"`
	Role     Role      `json:"role" gorm:"foreignKey:RoleID"`
	Vehicles []Vehicle `json:"vehicles,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.Email = normalizeEmail(u.Email)
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) error {
	u.Email = normalizeEmail(u.Email)
	return nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
