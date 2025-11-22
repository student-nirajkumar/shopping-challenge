package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/student-nirajkumar/shopping-challenge/backend/internal/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Read Authorization header
		auth := c.GetHeader("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid auth header"})
			c.Abort()
			return
		}

		// Extract token
		token := strings.TrimPrefix(auth, "Bearer ")

		// Find user with this token
		var user models.User
		if err := models.DB.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", user)

		c.Next()
	}
}
