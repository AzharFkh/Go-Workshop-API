package services

import (
	"errors"
	"go_bengkel/internal/dto"
	"go_bengkel/internal/models"
	"go_bengkel/internal/repository"
	"go_bengkel/internal/utils"

	"golang.org/x/crypto/bcrypt"
)


type UserService interface {
	CreateUser(req dto.UserRegister) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	DeleteUserByID(id int) error
	Login(req dto.UserLogin) (string, error)
}

type userService struct {
	repo repository.UserRepository
	roleRepo repository.RoleRepository // for RBAC
}

// constructor

func NewUserService(
	repo repository.UserRepository,
	roleRepo repository.RoleRepository,
) UserService {
	return &userService{
		repo: repo,
		roleRepo: roleRepo,
	}
}

// login

func (s *userService) Login(req dto.UserLogin) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(utils.TokenPayload{
		UserID: user.ID,
		Role: user.Role.Name,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

// create user

func (s *userService) CreateUser(req dto.UserRegister) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	// find default Role
	role, err := s.roleRepo.FindByName("user") // set default as user

	// masukan password ke req
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID: role.ID,
	}

	// kirim ke DB
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// get users

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

// get user by id

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.FindByID(id)
}

// delete user by id

func (s *userService) DeleteUserByID(id int) error {
	// mencari apakah user ada (tidak asal hapus)
	user, err := s.repo.FindByID(id)

	if err != nil {
		return err
	}

	return s.repo.Delete(user)
}
