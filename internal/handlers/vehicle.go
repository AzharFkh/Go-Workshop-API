package handlers

import (
	"go_bengkel/internal/dto"
	"go_bengkel/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(vehicleService services.VehicleService) *VehicleHandler {
	return &VehicleHandler{vehicleService}
}

// Create godoc
//
//	@Summary		Create vehicle
//	@Description	Menambahkan kendaraan baru milik user yang sedang login.
//	@Tags			Vehicles
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.VehicleRegister	true	"Vehicle Data"
//	@Success		201		{object}	dto.VehicleResponse
//	@Router			/vehicles [post]
func (h *VehicleHandler) Create(c *gin.Context) {

	userIDRaw, exist := c.Get("userID")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID := userIDRaw.(uint)

	var req dto.VehicleRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	vehicle, err := h.vehicleService.CreateVehicle(req, userID)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	// optional response
	response := dto.ToVehicleResponse(*vehicle)

	c.JSON(http.StatusCreated, gin.H{
		"message": "vehicle added",
		"vehicle": response,
	})
}

// FindAll godoc
//
//	@Summary		Get all vehicles
//	@Description	Mengambil seluruh kendaraan milik user yang sedang login.
//	@Tags			Vehicles
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	dto.VehicleResponse
//	@Router			/vehicles [get]
func (h *VehicleHandler) FindAll(c *gin.Context) {

	userIDRaw, exist := c.Get("userID")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID := userIDRaw.(uint)

	vehicles, err := h.vehicleService.GetVehicle(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	responses := dto.ToVehicleResponses(vehicles)

	c.JSON(http.StatusOK, gin.H{
		"vehicles": responses,
	})
}
