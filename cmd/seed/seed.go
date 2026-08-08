package seed

import (
	"go_bengkel/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func helperHashed(password string) (string, error){
    
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil{
        return "", err
    }
    return string(hashed), nil
}

func FeedData(db *gorm.DB) {
	adminRole := models.Role{Name: "admin"}
	userRole := models.Role{Name: "user"}
	db.FirstOrCreate(&adminRole, models.Role{Name: "admin"})
	db.FirstOrCreate(&userRole, models.Role{Name: "user"})

    password, _ := helperHashed("1234")

	adminUser := models.User{
        Name:     "budi_admin",
        Email:    "budi@email.com",
        Password: password,
        RoleID:   adminRole.ID, // as admin
    }
	db.FirstOrCreate(&adminUser, models.User{Name: "budi_admin"})

	normalUser := models.User{
        Name:     "andi_user",
        Email:    "andi@email.com",
        Password: password,
        RoleID:   userRole.ID, // as user
    }
    db.FirstOrCreate(&normalUser, models.User{Name: "andi_user"})

}
