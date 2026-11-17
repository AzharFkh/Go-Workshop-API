package handlers

import (
	"net/http"

	"go_bengkel/internal/dto"
	"go_bengkel/internal/services"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(vehicleService services.VehicleService) *VehicleHandler {
	return &VehicleHandler{vehicleService: vehicleService}
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
//	@Failure		400		{object}	dto.ErrorResponse	"Bad Request - Invalid JSON"
//	@Failure		401		{object}	dto.ErrorResponse	"Unauthorized - User not found in context"
//	@Failure		409		{object}	dto.ErrorResponse	"Conflict - Vehicle creation failed"
//	@Router			/vehicles [post]
func (h *VehicleHandler) Create(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var req dto.VehicleRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := h.vehicleService.CreateVehicle(req, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "vehicle added",
		"vehicle": dto.ToVehicleResponse(*vehicle),
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
//	@Failure		401	{object}	dto.ErrorResponse	"Unauthorized - User not found in context"
//	@Failure		500	{object}	dto.ErrorResponse	"Internal Server Error - Failed to fetch vehicles"
//	@Router			/vehicles [get]
func (h *VehicleHandler) FindAll(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	vehicles, err := h.vehicleService.GetVehicle(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vehicles": dto.ToVehicleResponses(vehicles),
	})
}
