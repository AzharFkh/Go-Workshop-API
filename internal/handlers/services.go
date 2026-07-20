package handlers

import (
	"go_bengkel/internal/dto"
	"go_bengkel/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServicesHandler struct {
	dataService services.DataService
}

func NewServicesHandler(dataService services.DataService) *ServicesHandler {
	return &ServicesHandler{dataService}
}

func (h *ServicesHandler) Create(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID := userIDRaw.(uint)

	vehicleIDStr := c.Param("vehicle_id")
	vehicleID, err := strconv.Atoi(vehicleIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid vehicle id",
		})
		return
	}

	var req dto.DataServiceRegister

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dataService, err := h.dataService.CreateDataService(req, userID, uint(vehicleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "data service is created",
		"data":    dataService,
	})
}

func (h *ServicesHandler) FindAll(c *gin.Context) {

}
