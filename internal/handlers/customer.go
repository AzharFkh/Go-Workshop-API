package handlers

import (
	"errors"
	"go_bengkel/internal/dto"
	"go_bengkel/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Melakukan autentikasi user dan mengembalikan JWT token.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserLogin	true	"Login Request"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Router			/users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}

// Create godoc
//
//	@Summary		Register new user
//	@Description	Membuat akun user baru.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserRegister	true	"Register Request"
//	@Success		201		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Failure		409		{object}	map[string]string
//	@Router			/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exist"})
		return
	}

	response := dto.ToUserResponse(*user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    response,
	})
}

// FindAll godoc
//
//	@Summary		Get all users
//	@Description	Mengambil seluruh data user.
//	@Tags			Users
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]string
//	@Router			/users [get]
//	@Security 		BearerAuth
func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := dto.ToUserResponses(users)

	c.JSON(http.StatusOK, gin.H{
		"users": response,
	})
}

// FindByID godoc
//
//	@Summary		Get user by ID
//	@Description	Mengambil data user berdasarkan ID.
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/users/{id} [get]
//	@Security 		BearerAuth
func (h *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	response := dto.ToUserResponse(*user)

	c.JSON(http.StatusOK, gin.H{
		"user": response,
	})
}

// Delete godoc
//
//	@Summary		Delete user
//	@Description	Menghapus user berdasarkan ID.
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	map[string]string
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users/{id} [delete]
//	@Security 		BearerAuth
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if err := h.userService.DeleteUserByID(id); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to delete user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been deleted",
	})
}
