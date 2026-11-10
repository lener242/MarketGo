package handlers

import "github.com/gin-gonic/gin"

func SuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "ok", "data": data})
}

func ErrorJSON(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"status": "error", "message": msg})
}
