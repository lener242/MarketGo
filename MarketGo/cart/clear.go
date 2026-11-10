package cart

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClearCart(c *gin.Context) {
	uid := c.GetInt("userID")

	if err := models.ClearCart(database.DB, uid); err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	handlers.SuccessJSON(c, gin.H{"message": "cleared"})
}
