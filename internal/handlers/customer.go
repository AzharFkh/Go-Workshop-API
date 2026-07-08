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

func NewUserHandler(userService services.UserService) *UserHandler{
	return &UserHandler{userService}
}

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
		"token": token,
	})
}

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

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user": user,
	})
}


func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil{
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

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
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
