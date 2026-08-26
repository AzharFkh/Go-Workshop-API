package repository

import (
	"errors"
	"go_bengkel/internal/models"
	"go_bengkel/internal/utils"

	"gorm.io/gorm"
)


type UserRepository interface {
	Create(user *models.User) error
	FindAll(page int, limit int) (utils.PaginatedResult[models.User], error)
	FindByID(id uint) (*models.User, error)
	Delete(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll(page int, limit int) (utils.PaginatedResult[models.User], error) {
	var users []models.User
	var totalItems int64
	var result utils.PaginatedResult[models.User]

	offset := (page-1)*limit

	// counting total item
	if err := r.db.Model(&models.User{}).Count(&totalItems).Error; err != nil{
		return result, err
	}

	if totalItems == 0{
		return result, nil
	}

	// querry with limit
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil{
		return result, err
	}

	totalPages := (totalItems + int64(limit) - 1) / int64(limit)

	// insert data
	result.Data = users
	result.Page = page
	result.Limit = limit
	result.TotalPages = totalPages
	result.TotalItems = totalItems

	return result, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error){
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, gorm.ErrRecordNotFound
		}
	}
	return &user, nil
}

func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Preload("Role").First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}