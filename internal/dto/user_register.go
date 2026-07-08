package dto

import (
	"go_bengkel/internal/models"
)


type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ToUserRegister(user models.User) UserRegister {

	return UserRegister{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
