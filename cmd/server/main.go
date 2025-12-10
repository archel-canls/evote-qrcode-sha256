package main

import (
	"log"
	"os"

	"evote/internal/config"
	"evote/internal/database"
	"evote/internal/models"
	"evote/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// 1. Load environment variables
	config.LoadEnv()

	// 2. Connect to PostgreSQL
	database.Connect()

	// 3. Auto-migrate DB schema
	err := database.DB.AutoMigrate(
		&models.Voter{},
		&models.Candidate{},
		&models.Vote{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// 4. Setup Gin HTTP server
	r := gin.Default()
	r.Use(cors.Default())

	// 5. Serve static frontend
	r.Static("/static", "./frontend")
	r.LoadHTMLGlob("frontend/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// 6. Register all API routes
	routes.RegisterRoutes(r)

	// 7. Run server on port from .env or fallback to 8080
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 E-Voting server running on port", port)
	r.Run(":" + port)
}
