package router

import (
	"go_bengkel/internal/handlers"
	"go_bengkel/internal/middleware"

	_ "go_bengkel/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	dataServiceHandler *handlers.ServicesHandler,
	vehicleHandler *handlers.VehicleHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
    }))

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
		protected.PATCH("/users/me/password", userHandler.ChangePassword)	// change password
		protected.POST("/vehicles", vehicleHandler.Create)	// create vehicle 
		protected.GET("/vehicles", vehicleHandler.FindAll)	// get all vehicle

		protected.POST("/vehicles/:vehicle_id/dataservices", dataServiceHandler.Create)	// create vehicle data service
		protected.GET("/vehicles/:vehicle_id/dataservices", dataServiceHandler.FindAll)	// find all vehicle data service
	}

	adminRole := protected.Group("")
	adminRole.Use(middleware.RequiredRole("admin"))
	{
		adminRole.GET("/users", userHandler.FindAll)
		adminRole.GET("/users/:id", userHandler.FindByID)	// get user by id
		adminRole.DELETE("/users/:id", userHandler.Delete)
	}

	return r 
}
