package handlers

import (
	"net/http"
	"strconv"

	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"

	"github.com/gin-gonic/gin"
)

// GetCandidates mengambil semua data kandidat
// Endpoint: GET /api/admin/candidates (untuk Admin) atau GET /api/candidates (untuk Publik)
func GetCandidates(c *gin.Context) {
	var candidates []models.Candidate

	if err := database.DB.Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil kandidat"})
		return
	}

	c.JSON(http.StatusOK, candidates)
}

// CreateCandidate menambahkan kandidat baru
// Endpoint: POST /api/admin/candidates
func CreateCandidate(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	candidate := models.Candidate{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := database.DB.Create(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat kandidat"})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

// DeleteCandidate menghapus kandidat berdasarkan ID
// Endpoint: DELETE /api/admin/candidates/:id
func DeleteCandidate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := database.DB.Delete(&models.Candidate{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kandidat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kandidat berhasil dihapus"})
}

// UpdateCandidate memperbarui data kandidat
// Endpoint: PUT /api/admin/candidates/:id
func UpdateCandidate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var candidate models.Candidate
	if err := database.DB.First(&candidate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kandidat tidak ditemukan"})
		return
	}

	candidate.Name = input.Name
	candidate.Description = input.Description

	if err := database.DB.Save(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update kandidat"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// GetResults menghitung total suara untuk setiap kandidats
func GetResults(c *gin.Context) {
	type Result struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		VoteCount int64  `json:"vote_count"`
	}

	var results []Result

	// Melakukan Join antara tabel candidates dan votes untuk menghitung suara [cite: 22]
	err := database.DB.Model(&models.Candidate{}).
		Select("candidates.id, candidates.name, count(votes.id) as vote_count").
		Joins("left join votes on votes.candidate_id = candidates.id").
		Group("candidates.id").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung hasil"})
		return
	}

	c.JSON(http.StatusOK, results)
}
