package dto

// ErrorResponse untuk API
type ErrorResponse struct {
	Error string `json:"error" example:"keterangan error yang terjadi, misal: 'invalid credentials'"`
}
