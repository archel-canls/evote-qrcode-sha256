package controllers

import (
	"evote/internal/crypto"
	"evote/internal/database"
	"evote/internal/models"
	"evote/internal/utils"

	"github.com/gin-gonic/gin"
)

// Register voter & generate QR Code
func RegisterVoter(c *gin.Context) {
	var req struct {
		NIM string `json:"nim"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "Invalid request")
		return
	}

	// Generate QR for voter
	qr, err := crypto.GenerateQR(req.NIM)
	if err != nil {
		utils.Error(c, "QR Generation failed")
		return
	}

	voter := models.Voter{
		NIM:     req.NIM,
		QRImage: qr,
	}

	database.DB.Create(&voter)

	utils.Success(c, "Voter registered", voter)
}
