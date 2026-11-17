package seed

import (
	"log"

	"go_bengkel/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func SetupAndSeedDatabase(db *gorm.DB) {
	roles := []models.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "user"},
	}

	for _, role := range roles {
		if err := db.Where(models.Role{ID: role.ID}).FirstOrCreate(&role).Error; err != nil {
			log.Printf("failed to seed role %s: %v", role.Name, err)
		}
	}

	result := db.Model(&models.User{}).
		Where("role_id IS NULL OR role_id = ?", 0).
		Update("role_id", 2)
	if result.Error != nil {
		log.Printf("failed to update role for old users: %v", result.Error)
	} else if result.RowsAffected > 0 {
		log.Printf("updated role for %d user(s)", result.RowsAffected)
	}

	password, err := hashPassword("1234")
	if err != nil {
		log.Fatalf("failed to hash seed password: %v", err)
	}

	defaultUsers := []models.User{
		{
			Name:     "budi_admin",
			Email:    "budi@email.com",
			Password: password,
			RoleID:   1,
		},
		{
			Name:     "andi_user",
			Email:    "andi@email.com",
			Password: password,
			RoleID:   2,
		},
	}

	for _, user := range defaultUsers {
		if err := db.Where(models.User{Email: user.Email}).FirstOrCreate(&user).Error; err != nil {
			log.Printf("failed to seed user %s: %v", user.Email, err)
		}
	}
}
