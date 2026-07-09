package main

import (
	"go_bengkel/internal/config"
	"go_bengkel/internal/handlers"
	"go_bengkel/internal/middleware"
	"go_bengkel/internal/repository"
	"go_bengkel/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	config.ConnectDB()

	db := config.DB

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleService := services.NewVehicleService(vehicleRepo)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	r := gin.Default()

	public := r.Group("/users")
	{
		public.POST("/", userHandler.Create)
		public.POST("/login", userHandler.Login)
	}

	userRoute := r.Group("/users")
	userRoute.Use(middleware.AuthMiddleware())
	{
		userRoute.GET("/", userHandler.FindAll)
		userRoute.GET("/:id", userHandler.FindByID)
		userRoute.DELETE("/:id", userHandler.Delete)
	}

	vehicleRoute := r.Group("/vehicle")
	vehicleRoute.Use(middleware.AuthMiddleware())
	{
		vehicleRoute.POST("/", vehicleHandler.Create)
		vehicleRoute.GET("/", vehicleHandler.FindAll)
	}

	r.Run(":8600")
}
