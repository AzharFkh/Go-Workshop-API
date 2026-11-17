package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"go_bengkel/internal/dto"
	"go_bengkel/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
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
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Router			/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		409		{object}	dto.ErrorResponse
//	@Router			/auth/register [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		if errors.Is(err, services.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    dto.ToUserResponse(*user),
	})
}

// FindAll godoc
//
//	@Summary		Get all users
//	@Description	Mengambil seluruh data user.
//	@Tags			Users
//	@Produce		json
//	@Param			page	query		int	false	"page number (default: 1)"
//	@Param			limit	query		int	false	"total data per page (default: 5)"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/users [get]
//	@Security		BearerAuth
func (h *UserHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	result, err := h.userService.GetUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userDTOs := dto.ToUserResponses(result.Data)
	c.JSON(http.StatusOK, dto.ToPaginatedResponse(userDTOs, result))
}

// FindByID godoc
//
//	@Summary		Get user by ID
//	@Description	Mengambil data user berdasarkan ID.
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (h *UserHandler) FindByID(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": dto.ToUserResponse(*user),
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
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
func (h *UserHandler) Delete(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	if err := h.userService.DeleteUserByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been deleted",
	})
}

// ChangePassword godoc
//
//	@Summary		Change User Password
//	@Description	Mengubah password untuk user yang sedang login.
//	@Tags			Users
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.ChangePasswordRequest	true	"Data Update Password"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Router			/users/me/password [patch]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.ChangePassword(userID, req); err != nil {
		if errors.Is(err, services.ErrIncorrectOldPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password changed successfully",
	})
}
