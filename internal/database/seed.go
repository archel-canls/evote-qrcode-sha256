package database

import (
	"log"

	"evote-qrcode-sha256/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var count int64
	DB.Model(&models.Admin{}).Count(&count)

	if count > 0 {
		log.Println("ℹ️ Admin already exists")
		return
	}

	password := "admin123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("❌ Failed to hash password")
	}

	admin := models.Admin{
		Username: "admin",
		Password: string(hash),
	}

	DB.Create(&admin)
	log.Println("✅ Default admin created (username: admin | password: admin123)")
}
