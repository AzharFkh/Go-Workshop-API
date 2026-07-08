package services

import (
	"errors"
	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
	"go_bengkel/internal/utils"
)

// revisi :

type UserService interface{
	CreateUser(req dto.UserRegister) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	DeleteUserByID(id int) error
	Login(req dto.UserLogin) (string, error)
}

type userService struct{
	repo repository.UserRepository
}

// constructor 

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Login(req dto.UserLogin) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}


func (s *userService) CreateUser(req dto.UserRegister) (*models.User, error) {
	user := models.User {
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) DeleteUserByID(id int) error {
	// mencari apakah user ada (tidak asal hapus)
	user, err := s.repo.FindByID(id)

	if err != nil{
		return err
	}

	return s.repo.Delete(user)
}