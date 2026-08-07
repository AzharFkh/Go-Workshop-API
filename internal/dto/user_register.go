package dto

import (
	"go_bengkel/internal/models"
)


type UserRegister struct {
	Name     string `json:"name" binding:"required" example:"budi_admin"`
	Email    string `json:"email" binding:"required,email" example:"budi@email.com"`
	Password string `json:"password" binding:"required" example:"1234"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" example:"budi@email.com"`
	Password string `json:"password" binding:"required" example:"1234"`
}

func ToUserRegister(user models.User) UserRegister {

	return UserRegister{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
