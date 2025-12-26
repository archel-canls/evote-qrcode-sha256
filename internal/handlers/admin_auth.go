package handlers

import (
	"net/http"
	"os"
	"time"

	"evote-qrcode-sha256/internal/database"
	"evote-qrcode-sha256/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest adalah payload untuk login admin
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminLogin menangani autentikasi admin
func AdminLogin(c *gin.Context) {
	var req LoginRequest

	// 1️⃣ Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request tidak valid"})
		return
	}

	// 2️⃣ Cari admin di database
	var admin models.Admin
	if err := database.DB.Where("username = ?", req.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user tidak ditemukan"})
		return
	}

	// 3️⃣ Compare password dengan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password salah"})
		return
	}

	// 4️⃣ Generate JWT
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "defaultsecret" // fallback sementara, sebaiknya wajib di env
	}

	claims := jwt.MapClaims{
		"user": admin.Username,
		"exp":  time.Now().Add(24 * time.Hour).Unix(), // token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal generate token"})
		return
	}

	// 5️⃣ Response
	c.JSON(http.StatusOK, gin.H{
		"message": "login berhasil",
		"token":   tokenString,
	})
}
