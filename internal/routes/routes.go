package routes

import (
	"net/http"

	"evote-qrcode-sha256/internal/handlers"
	"evote-qrcode-sha256/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// =============================
	// Static / Frontend
	// =============================
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// =============================
	// API Group
	// =============================
	api := r.Group("/api") // ✅ INI YANG KAMU LUPA
	{
		// ---------- Admin Login ----------
		api.POST("/admin/login", handlers.AdminLogin)

		// ---------- Protected Admin ----------
		admin := api.Group("/admin")
		admin.Use(middleware.AdminJWT())
		{
			// Candidates
			admin.GET("/candidates", handlers.GetCandidates)
			admin.POST("/candidates", handlers.CreateCandidate)
			admin.DELETE("/candidates/:id", handlers.DeleteCandidate)

			// Voters
			admin.GET("/voters", handlers.GetVoters)
			admin.POST("/voters", handlers.CreateVoter)
		}
	}

	return r
}
