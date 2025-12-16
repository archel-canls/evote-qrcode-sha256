package handlers

import (
	"net/http"
	"strconv"

	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"

	"github.com/gin-gonic/gin"
)

// GET /api/admin/candidates
func GetCandidates(c *gin.Context) {
	var candidates []models.Candidate

	if err := database.DB.Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil kandidat"})
		return
	}

	c.JSON(http.StatusOK, candidates)
}

// POST /api/admin/candidates
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

// DELETE /api/admin/candidates/:id
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

	c.JSON(http.StatusOK, gin.H{"message": "Kandidat dihapus"})
}
