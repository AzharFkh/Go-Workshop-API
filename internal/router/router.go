package router

import (
	"go_bengkel/internal/handlers"
	"go_bengkel/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	dataServiceHandler *handlers.ServicesHandler,
	vehicleHandler *handlers.VehicleHandler,
) *gin.Engine {
	r := gin.Default()

	public := r.Group("api/users")
	{
		public.POST("/", userHandler.Create)
		public.POST("/login", userHandler.Login)
	}

	userRoute := r.Group("api/users")
	userRoute.Use(middleware.AuthMiddleware())
	{
		userRoute.GET("/", userHandler.FindAll)
		userRoute.GET("/:id", userHandler.FindByID)
		userRoute.DELETE("/:id", userHandler.Delete)
	}

	vehicleRoute := r.Group("api/vehicle")
	vehicleRoute.Use(middleware.AuthMiddleware())
	{
		vehicleRoute.POST("/", vehicleHandler.Create)
		vehicleRoute.GET("/", vehicleHandler.FindAll)

		vehicleRoute.POST("/:vehicle_id/dataservice", dataServiceHandler.Create)
		vehicleRoute.GET("/:vehicle_id/dataservice", dataServiceHandler.FindAll)

	}
	// PR: tambahin validasi pada data service
	// PR: perbaiki service handler

	return r 
}

// refactor route here
