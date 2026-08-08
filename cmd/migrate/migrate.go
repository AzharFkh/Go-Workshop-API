package main

import (
	"go_bengkel/cmd/seed"
	"go_bengkel/internal/config"
	"go_bengkel/internal/models"
	"log"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()

	err := config.DB.AutoMigrate(
		&models.Role{},
		&models.User{}, 
		&models.Vehicle{},
		&models.DataService{}, 
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	seed.FeedData(config.DB)

	log.Println("Migration and data sedding success")
}