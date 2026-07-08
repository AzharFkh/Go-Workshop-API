package dto

import (
	"go_bengkel/internal/models"
)

type UserResponse struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func ToUserResponse(user models.User) UserResponse {

	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
