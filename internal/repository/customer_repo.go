package repository

import (
	"errors"
	"go_bengkel/internal/models"

	"gorm.io/gorm"
)


type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id int) (*models.User, error)
	Delete(user *models.User) error
	FindByEmail(email string) (*models.User, error)
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

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id int) (*models.User, error){
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