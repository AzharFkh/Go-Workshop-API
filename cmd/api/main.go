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

	r := gin.Default()

	public := r.Group("/users")
	{
		public.POST("/", userHandler.Create)
		public.POST("/login", userHandler.Login)
	}

	protected := r.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", userHandler.FindAll)
		protected.GET("/:id", userHandler.FindByID)
		protected.DELETE("/:id", userHandler.Delete)
	}

	r.Run(":8600")
}
