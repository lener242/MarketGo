package cart

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type removeReq struct {
	DishID int `json:"dish_id" binding:"required"`
}

func RemoveFromCart(c *gin.Context) {
	uid := c.GetInt("userID")

	var r removeReq
	if err := c.ShouldBindJSON(&r); err != nil {
		handlers.ErrorJSON(c, http.StatusBadRequest, "invalid body")
		return
	}

	if err := models.RemoveFromCart(database.DB, uid, r.DishID); err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	handlers.SuccessJSON(c, gin.H{"message": "removed"})
}
