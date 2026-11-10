package profile

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewProfile(c *gin.Context) {
	uid := c.GetInt("userID")

	user, err := models.GetUserByID(database.DB, uid)
	if err != nil {
		handlers.ErrorJSON(c, http.StatusNotFound, "user not found")
		return
	}

	username := ""
	if user.Username.Valid {
		username = user.Username.String
	}

	handlers.SuccessJSON(c, gin.H{
		"name":     user.Name,
		"username": username,
		"phone":    user.Phone,
	})
}
