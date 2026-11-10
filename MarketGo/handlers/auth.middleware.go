package handlers

import (
	"MarketGo/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		c.Abort()
		return
	}

	parts := strings.SplitN(auth, " ", 2)
	tokenStr := auth
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		tokenStr = parts[1]
	}

	userID, err := utils.ParseJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("userID", userID)

	c.Next()
}
