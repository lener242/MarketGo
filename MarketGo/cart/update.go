package cart

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateReq struct {
	DishID   int `json:"dish_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

func UpdateQuantity(c *gin.Context) {
	uid := c.GetInt("userID")

	var r updateReq
	if err := c.ShouldBindJSON(&r); err != nil {
		handlers.ErrorJSON(c, http.StatusBadRequest, "invalid body")
		return
	}

	if r.Quantity <= 0 {
		if err := models.RemoveFromCart(database.DB, uid, r.DishID); err != nil {
			handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
			return
		}
		handlers.SuccessJSON(c, gin.H{"message": "removed"})
		return
	}

	if err := models.UpdateCartQuantity(database.DB, uid, r.DishID, r.Quantity); err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	handlers.SuccessJSON(c, gin.H{"message": "updated", "quantity": r.Quantity})
}
