package services

import (
	"errors"
	"strings"

	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
	"go_bengkel/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(req dto.UserRegister) (*models.User, error)
	GetUsers(page, limit int) (utils.PaginatedResult[models.User], error)
	GetUserByID(id uint) (*models.User, error)
	DeleteUserByID(id uint) error
	Login(req dto.UserLogin) (string, error)
	ChangePassword(id uint, req dto.ChangePasswordRequest) error
}

type userService struct {
	repo     repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(repo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &userService{
		repo:     repo,
		roleRepo: roleRepo,
	}
}

func (s *userService) Login(req dto.UserLogin) (string, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(utils.TokenPayload{
		UserID: user.ID,
		Role:   user.Role.Name,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUser(req dto.UserRegister) (*models.User, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepo.FindByName("user")
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     req.Name,
		Email:    email,
		Password: string(hashedPassword),
		RoleID:   role.ID,
	}

	if err := s.repo.Create(&user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrEmailAlreadyExists
		}
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUsers(page, limit int) (utils.PaginatedResult[models.User], error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	return s.repo.FindAll(page, limit)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) DeleteUserByID(id uint) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(user)
}

func (s *userService) ChangePassword(id uint, req dto.ChangePasswordRequest) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return ErrIncorrectOldPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.Update(user)
}
