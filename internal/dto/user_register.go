package dto

type UserRegister struct {
	Name     string `json:"name" binding:"required" example:"budi_admin"`
	Email    string `json:"email" binding:"required,email" example:"budi@email.com"`
	Password string `json:"password" binding:"required" example:"1234"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email" example:"budi@email.com"`
	Password string `json:"password" binding:"required" example:"1234"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
