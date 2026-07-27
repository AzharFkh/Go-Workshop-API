package router

import (
	"go_bengkel/internal/handlers"
	"go_bengkel/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_bengkel/docs"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	dataServiceHandler *handlers.ServicesHandler,
	vehicleHandler *handlers.VehicleHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	vehicleRoute := r.Group("api/vehicles")
	vehicleRoute.Use(middleware.AuthMiddleware())
	{
		vehicleRoute.POST("/", vehicleHandler.Create)
		vehicleRoute.GET("/", vehicleHandler.FindAll)

		vehicleRoute.POST("/:vehicle_id/dataservices", dataServiceHandler.Create)
		vehicleRoute.GET("/:vehicle_id/dataservices", dataServiceHandler.FindAll)

	}

	return r 
}
