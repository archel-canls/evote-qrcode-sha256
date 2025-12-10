package controllers

import (
	"fmt"

	"evote/internal/crypto"
	"evote/internal/database"
	"evote/internal/models"
	"evote/internal/utils"

	"github.com/gin-gonic/gin"
)

// Submit a vote
func SubmitVote(c *gin.Context) {
	var req struct {
		NIM         string `json:"nim"`
		CandidateID uint   `json:"candidate_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "Invalid request")
		return
	}

	// Create hash: NIM + CandidateID
	hashInput := fmt.Sprintf("%s-%d", req.NIM, req.CandidateID)
	hash := crypto.GenerateHash(hashInput)

	vote := models.Vote{
		VoterNIM:    req.NIM,
		CandidateID: req.CandidateID,
		Hash:        hash,
	}

	database.DB.Create(&vote)

	utils.Success(c, "Vote submitted", vote)
}
