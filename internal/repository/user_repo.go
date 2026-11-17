package repository

import (
	"go_bengkel/internal/models"
	"go_bengkel/internal/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll(page, limit int) (utils.PaginatedResult[models.User], error)
	FindByID(id uint) (*models.User, error)
	Delete(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll(page, limit int) (utils.PaginatedResult[models.User], error) {
	result := utils.PaginatedResult[models.User]{
		Page:  page,
		Limit: limit,
	}

	var totalItems int64
	if err := r.db.Model(&models.User{}).Count(&totalItems).Error; err != nil {
		return result, err
	}

	result.TotalItems = totalItems
	if totalItems == 0 {
		return result, nil
	}

	offset := (page - 1) * limit
	var users []models.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return result, err
	}

	result.Data = users
	result.TotalPages = (totalItems + int64(limit) - 1) / int64(limit)

	return result, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}
