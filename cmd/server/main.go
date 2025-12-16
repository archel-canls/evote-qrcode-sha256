package main

import (
	"log"

	"evote-qrcode-sha256/internal/config"
	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"
	"evote-qrcode-sha256/internal/routes"
)

func main() {
	// =============================
	// 1. Load Environment Variable
	// =============================
	config.LoadEnv()

	// =============================
	// 2. Connect Database
	// =============================
	database.Connect()

	// =============================
	// 3. Auto Migration
	// =============================
	if err := database.DB.AutoMigrate(
		&models.Admin{},
		&models.Candidate{},
		&models.Voter{},
		&models.Vote{},
	); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}
	log.Println("✅ Database migrated")

	// =============================
	// 4. Setup Router
	// =============================
	r := routes.SetupRouter()

	// =============================
	// 5. Run Server
	// =============================
	port := config.GetEnv("PORT", "8080")
	log.Printf("🚀 Server running on http://localhost:%s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
