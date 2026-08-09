package seed

import (
	"go_bengkel/internal/models"

	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func helperHashed(password string) (string, error) {

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
		db.Where(models.Role{ID: role.ID}).FirstOrCreate(&role)
	}

	result := db.Model(&models.User{}).
		Where("role_id IS NULL OR role_id = ?", 0).
		Update("role_id", 2) // set default as role_id=2 (user)

	if result.Error != nil {
		fmt.Println("failed to update role for old user:", result.Error)
	} else if result.RowsAffected > 0 {
		fmt.Printf("adding %d user role for user '\n", result.RowsAffected)
	}

	password, _ := helperHashed("1234")

	defaultUsers := []models.User{
		{
			Name:     "budi_admin",
			Email:    "budi@email.com",
			Password: password,
			RoleID:   1, // ID Role Admin
		},
		{
			Name:     "andi_user",
			Email:    "andi@email.com",
			Password: password,
			RoleID:   2, // ID Role User
		},
	}

	for _, user := range defaultUsers {
		db.Where(models.User{Email: user.Email}).FirstOrCreate(&user) // FirstOrCreate using email
	}
}
