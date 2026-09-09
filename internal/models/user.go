package models

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"uniqueIndex:idx_email"`
	Password string `json:"password" binding:"required"`

	RoleID uint `json:"role_id"`
	Role   Role `json:"role" gorm:"foreignKey:RoleID"`

	// one to many
	Vehicle []Vehicle `json:"vehicle" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Email = normalizeEmail(u.Email)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.Email = normalizeEmail(u.Email)
	return nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(
		strings.TrimSpace(email),
	)
}
