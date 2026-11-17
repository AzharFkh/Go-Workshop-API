package main

import (
	"log"

	"go_bengkel/cmd/seed"
	"go_bengkel/internal/config"
	"go_bengkel/internal/models"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Vehicle{},
		&models.DataService{},
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	seed.SetupAndSeedDatabase(db)

	log.Println("Migration and data seeding success")
}
