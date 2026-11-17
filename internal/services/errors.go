package services

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid email or password")
	ErrIncorrectOldPassword = errors.New("incorrect old password")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrVehicleNotOwned      = errors.New("vehicle is not owned by user")
	ErrVehicleNotFound      = errors.New("vehicle not found or unauthorized")
)
