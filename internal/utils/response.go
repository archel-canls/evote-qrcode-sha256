package utils

import "github.com/gin-gonic/gin"

// Success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// Error response
func Error(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"success": false,
		"message": message,
	})
}
