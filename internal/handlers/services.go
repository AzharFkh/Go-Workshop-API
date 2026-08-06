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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid vehicle_id",
		})
		return 0, 0, false
	}
	return userIDRaw.(uint), uint(vehicleID), true
}

// Create godoc
//
//	@Summary		Create service record
//	@Description	Menambahkan data service untuk kendaraan milik user yang sedang login.
//	@Tags			Services
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			vehicle_id	path 	int true	"Vehicle ID"
//	@Param			request body dto.DataServiceRegister true "Service Data"
//	@Success		201 {object} dto.DataServiceResponse
//	@Failure        400 {object} dto.ErrorResponse "Bad Request - Invalid JSON payload"
//	@Failure        401 {object} dto.ErrorResponse "Unauthorized - User ID not found in context"
//	@Failure        500 {object} dto.ErrorResponse "Internal Server Error - Failed to create service record"
//	@Router			/vehicles/{vehicle_id}/dataservices [post]
func (h *ServicesHandler) Create(c *gin.Context) {
	userID, vehicleID, exists := h.GetID(c)

	if !exists {
		return
	}
	var req dto.DataServiceRegister

	if err := c.ShouldBindJSON(&req); err != nil {
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

// FindAll godoc
//
//	@Summary		Get all service records
//	@Description	Mengambil seluruh riwayat service kendaraan milik user yang sedang login.
//	@Tags			Services
//	@Security		BearerAuth
//	@Produce		json
//	@Param			vehicle_id	path	int	true	"Vehicle ID"
//	@Success		200			{object}	dto.DataServiceResponse
//	@Failure        401         {object}    dto.ErrorResponse       "Unauthorized - User ID not found in context"
//	@Failure        404         {object}    dto.ErrorResponse       "Not Found - Vehicle not found or unauthorized"
//	@Failure        500         {object}    dto.ErrorResponse       "Internal Server Error - Failed to fetch service records"
//	@Router			/vehicles/{vehicle_id}/dataservices [get]
func (h *ServicesHandler) FindAll(c *gin.Context) {
	userID, vehicleID, exists := h.GetID(c)

	if !exists {
		return
	}

	servicesData, err := h.dataService.GetDataServices(userID, vehicleID)

	if err != nil {
		// error ketika vehicleID tidak dimiliki oleh userID
		if err.Error() == "vehicle not found or unauthorized" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "vehicle not found or unauthorized",
			})
			return
		}

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
