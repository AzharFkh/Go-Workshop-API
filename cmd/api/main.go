package main

import (
	"log"

	_ "go_bengkel/docs"
	"go_bengkel/internal/config"
	"go_bengkel/internal/handlers"
	"go_bengkel/internal/repository"
	"go_bengkel/internal/router"
	"go_bengkel/internal/services"
)

// @title			Go Bengkel API
// @version			1.0
// @description		Backend API untuk aplikasi bengkel.
// @description
// @description		Cara menggunakan autentikasi pada Swagger:
// @description		1. Login menggunakan endpoint Authentication untuk mendapatkan token.
// @description		2. Klik tombol "Authorize" di atas.
// @description		3. Masukkan token dengan format: Bearer {spasi} {token}.
//
// @host			localhost:8600
// @BasePath		/api
// @schemes			http
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	vehicleRepo := repository.NewVehicleRepository(db)
	dataServiceRepo := repository.NewServicesRepository(db)

	userService := services.NewUserService(userRepo, roleRepo)
	vehicleService := services.NewVehicleService(vehicleRepo)
	dataService := services.NewDataService(dataServiceRepo, vehicleRepo)

	userHandler := handlers.NewUserHandler(userService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)
	dataServiceHandler := handlers.NewServicesHandler(dataService)

	r := router.SetupRouter(userHandler, dataServiceHandler, vehicleHandler)

	log.Println("Server running at port 8600...")
	if err := r.Run(":8600"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
