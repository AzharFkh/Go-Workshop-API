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

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", userHandler.Login)
		auth.POST("/register", userHandler.Create)
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.PATCH("/users/me/password", userHandler.ChangePassword)
		protected.GET("/users/:id", userHandler.FindByID)
		protected.POST("/vehicles", vehicleHandler.Create)
		protected.GET("/vehicles", vehicleHandler.FindAll)

		protected.POST("/vehicles/:vehicle_id/dataservices", dataServiceHandler.Create)
		protected.GET("/vehicles/:vehicle_id/dataservices", dataServiceHandler.FindAll)
	}

	adminRole := protected.Group("")
	adminRole.Use(middleware.RequiredRole("admin"))
	{
		adminRole.GET("/users", userHandler.FindAll)
		adminRole.DELETE("/users/:id", userHandler.Delete)

	}

	return r 
}
