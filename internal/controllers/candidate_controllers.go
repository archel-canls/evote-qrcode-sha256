package controllers

import (
	"evote/internal/database"
	"evote/internal/models"
	"evote/internal/utils"

	"github.com/gin-gonic/gin"
)

// Register a new candidate
func AddCandidate(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "Invalid request")
		return
	}

	candidate := models.Candidate{Name: req.Name}
	database.DB.Create(&candidate)

	utils.Success(c, "Candidate added", candidate)
}

// Get all candidates
func GetCandidates(c *gin.Context) {
	var candidates []models.Candidate
	database.DB.Find(&candidates)

	utils.Success(c, "Candidate list", candidates)
}
