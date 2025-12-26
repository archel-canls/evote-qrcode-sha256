package routes

import (
	"evote-qrcode-sha256/internal/handlers"
	"evote-qrcode-sha256/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("frontend/*.html")

	// --- HALAMAN HTML (PUBLIC) ---
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	r.GET("/admin/login", func(c *gin.Context) { c.HTML(http.StatusOK, "admin_login.html", nil) })
	r.GET("/admin", func(c *gin.Context) { c.HTML(http.StatusOK, "admin.html", nil) })
	r.GET("/vote", func(c *gin.Context) { c.HTML(http.StatusOK, "vote.html", nil) })
	r.GET("/voter/create", func(c *gin.Context) { c.HTML(http.StatusOK, "create_voter.html", nil) })

	api := r.Group("/api")
	{
		api.POST("/admin/login", handlers.AdminLogin)
		api.POST("/vote", handlers.SubmitVote)
		api.GET("/candidates", handlers.GetCandidates)

		// ✅ PERBAIKAN: Rute pendaftaran publik (Harus sama dengan di HTML)
		api.POST("/voters/register", handlers.CreateVoter)

		// --- ADMIN (PROTECTED) ---
		admin := api.Group("/admin")
		admin.Use(middleware.AdminJWT())
		{
			admin.GET("/candidates", handlers.GetCandidates)
			admin.POST("/candidates", handlers.CreateCandidate)
			admin.DELETE("/candidates/:id", handlers.DeleteCandidate)
			admin.GET("/voters", handlers.GetVoters)
			admin.GET("/results", handlers.GetResults)
		}
	}
	return r
}
