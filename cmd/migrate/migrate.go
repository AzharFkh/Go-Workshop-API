package main

import (
	"go_bengkel/internal/config"
	"go_bengkel/internal/models"
	"log"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()

	err := config.DB.AutoMigrate(
		&models.User{}, 
		&models.DataService{}, 
		&models.Vehicle{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success")
}