package handlers

import (
	"net/http"
	"time"

	"evote-qrcode-sha256/internal/crypto"
	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"

	"github.com/gin-gonic/gin"
)

// GET /api/admin/voters
func GetVoters(c *gin.Context) {
	var voters []models.Voter

	if err := database.DB.Find(&voters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pemilih"})
		return
	}

	c.JSON(http.StatusOK, voters)
}

// POST /api/admin/voters
func CreateVoter(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// Generate QR Token (SHA-256)
	raw := input.Email + time.Now().String()
	token := crypto.GenerateHash(raw)

	voter := models.Voter{
		Name:     input.Name,
		Email:    input.Email,
		QRToken:  token,
		HasVoted: false,
	}

	if err := database.DB.Create(&voter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pemilih"})
		return
	}

	c.JSON(http.StatusCreated, voter)
}
