package handlers

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"evote-qrcode-sha256/internal/crypto/hash"
	"evote-qrcode-sha256/internal/crypto/qr"
	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"

	"github.com/gin-gonic/gin"
)

// Fungsi pembantu membuat kode pendek 6 karakter
func generateShortCode(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%X", b)[:n]
}

func CreateVoter(c *gin.Context) {
	var input models.Voter
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	// 1. Cek apakah email sudah terdaftar di database
	var existingVoter models.Voter
	err := database.DB.Where("email = ?", input.Email).First(&existingVoter).Error

	if err == nil {
		// EMAIL DITEMUKAN: Ambil QR dari token lama dan kirimkan kembali
		qrBase64, _ := qr.GenerateQR(existingVoter.Token) //

		c.JSON(http.StatusOK, gin.H{
			"message": "Email sudah pernah digunakan! Berikut adalah token Anda:",
			"voter": gin.H{
				"name":        existingVoter.Name,
				"short_token": existingVoter.ShortToken,
				"token":       existingVoter.Token,
			},
			"qr": qrBase64,
		})
		return
	}

	// 2. EMAIL BARU: Lanjutkan proses pendaftaran normal
	input.Token = hash.GenerateWithSalt(input.Email, "VOTER_SECRET") //
	input.ShortToken = generateShortCode(6)
	input.HasVoted = false

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menambahkan voter"})
		return
	}

	qrBase64, _ := qr.GenerateQR(input.Token)

	c.JSON(http.StatusOK, gin.H{
		"message": "voter berhasil ditambahkan",
		"voter": gin.H{
			"name":        input.Name,
			"short_token": input.ShortToken,
			"token":       input.Token,
		},
		"qr": qrBase64,
	})
}

func GetVoters(c *gin.Context) {
	var voters []models.Voter
	if err := database.DB.Find(&voters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil voters"})
		return
	}
	c.JSON(http.StatusOK, voters)
}
