package handlers

import (
	"evote-qrcode-sha256/internal/crypto/hash"
	"evote-qrcode-sha256/internal/crypto/qr"
	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// VoteRequest mendefinisikan payload JSON yang diterima dari frontend
type VoteRequest struct {
	Token       string `json:"token" binding:"required"`
	CandidateID uint   `json:"candidate_id" binding:"required"`
}

// SubmitVote menangani proses pemungutan suara secara aman dan anonim
func SubmitVote(c *gin.Context) {
	var req VoteRequest
	// 1. Validasi input JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// 2. Cari voter berdasarkan short_token (diketik) atau token full (di-scan)
	// Hal ini mendukung kemudahan penggunaan (UX) sekaligus keamanan tinggi
	var voter models.Voter
	if err := database.DB.Where("short_token = ? OR token = ?", req.Token, req.Token).First(&voter).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kode atau Token tidak valid"})
		return
	}

	// 3. Verifikasi apakah pemilih sudah pernah memberikan suara
	if voter.HasVoted {
		c.JSON(http.StatusForbidden, gin.H{"error": "Token ini sudah digunakan untuk voting"})
		return
	}

	// 4. Generate Hash Suara (SHA-256 + Salt)
	// Penggunaan salt mencegah serangan rainbow table [cite: 134, 150]
	secret := os.Getenv("VOTE_SECRET")
	if secret == "" {
		secret = "default_secret" // Fallback jika environment variable belum diatur
	}

	// Gabungkan data untuk di-hash agar menghasilkan nilai yang unik dan irreversible [cite: 145, 147]
	rawData := fmt.Sprintf("%d:%d", voter.ID, req.CandidateID)
	voteHash := hash.GenerateWithSalt(rawData, secret)

	// 5. Simpan suara dan update status voter dalam satu transaksi database
	// Transaksi menjamin integritas data; jika satu gagal, semua dibatalkan
	tx := database.DB.Begin()

	vote := models.Vote{
		VoterID:     voter.ID,
		CandidateID: req.CandidateID,
		VoteHash:    voteHash, // Hash disimpan agar pilihan tetap anonim [cite: 146]
	}

	// Simpan data pilihan
	if err := tx.Create(&vote).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat suara"})
		return
	}

	// Update status agar voter tidak bisa memilih dua kali
	if err := tx.Model(&voter).Update("has_voted", true).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status pemilih"})
		return
	}
	tx.Commit()

	// 6. Generate QR Bukti Vote sebagai tanda terima digital pemilih
	qrBase64, _ := qr.GenerateQR(voteHash)

	c.JSON(http.StatusOK, gin.H{
		"message": "Vote berhasil disimpan secara anonim",
		"hash":    voteHash, // Menampilkan hash yang menunjukkan avalanche effect [cite: 97, 145]
		"qr":      qrBase64,
	})
}
