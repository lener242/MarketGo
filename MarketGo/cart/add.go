package cart

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addReq struct {
	DishID   int `json:"dish_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

func AddToCart(c *gin.Context) {
	uid := c.GetInt("userID")

	var req addReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handlers.ErrorJSON(c, http.StatusBadRequest, "invalid body")
		return
	}

	if req.Quantity <= 0 {
		handlers.ErrorJSON(c, http.StatusBadRequest, "quantity must be greater than 0")
		return
	}

	if _, err := models.GetDishByID(database.DB, req.DishID); err != nil {
		handlers.ErrorJSON(c, http.StatusNotFound, "dish not found")
		return
	}

	if err := models.AddToCart(database.DB, uid, req.DishID, req.Quantity); err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error: "+err.Error())
		return
	}

	handlers.SuccessJSON(c, gin.H{"message": "added"})
}
