package dto

import (
	"go_bengkel/internal/models"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserResponse(user models.User) UserResponse {

	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserResponses(users []models.User) []UserResponse {
	var responses []UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}

	return responses
}
