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

func (h *ServicesHandler) GetID(c *gin.Context) (uint, uint, bool) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return 0, 0, false
	}

	vehicleID, err := strconv.Atoi(c.Param("vehicle_id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid vehicle_id",
		})
		return 0, 0, false
	}
	return userIDRaw.(uint), uint(vehicleID), true
}

func (h *ServicesHandler) Create(c *gin.Context) {
	userID, vehicleID, exists := h.GetID(c)

	if !exists {
		return
	}
	var req dto.DataServiceRegister

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	servicesData, err := h.dataService.CreateDataService(req, userID, vehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "data service is created",
		"data":    dto.ToDataServiceResponse(*servicesData),
	})
}

func (h *ServicesHandler) FindAll(c *gin.Context) {
	userID, vehicleID, exists := h.GetID(c)

	if !exists {
		return
	}

	servicesData, err := h.dataService.GetDataServices(userID, vehicleID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := dto.ToDataServiceResponses(servicesData)

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
	})
}
