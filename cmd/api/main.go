package main

import (
	"go_bengkel/internal/config"
	"go_bengkel/internal/handlers"
	"go_bengkel/internal/repository"
	"go_bengkel/internal/router"
	"go_bengkel/internal/services"
	"log"
	_ "go_bengkel/docs"
)


// @title           Go Bengkel API
// @version         1.0
// @description     Backend API untuk aplikasi bengkel.
// @description     
// @description     Cara menggunakan autentikasi pada Swagger:
// @description     1. Login menggunakan endpoint Authentication untuk mendapatkan token.
// @description     2. Klik tombol "Authorize" di atas.
// @description     3. Masukkan token dengan format: Bearer {spasi} {token}.

// @host            localhost:8600
// @BasePath        /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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

	dataServiceRepo := repository.NewServicesRepository(db)
	dataServiceService := services.NewDataService(dataServiceRepo, vehicleRepo)
	dataServiceHandler := handlers.NewServicesHandler(dataServiceService)

	r := router.SetupRouter(userHandler, dataServiceHandler, vehicleHandler)
	
	log.Println("Server running at port 8600... ")
	
	if err := r.Run(":8600"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}	
}
