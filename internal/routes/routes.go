package routes

import (
	"evote/internal/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API endpoints
func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	// Voter
	api.POST("/voter/register", controllers.RegisterVoter)

	// Vote
	api.POST("/vote", controllers.SubmitVote)

	// Candidate
	api.POST("/candidate", controllers.AddCandidate)
	api.GET("/candidate", controllers.GetCandidates)
}
