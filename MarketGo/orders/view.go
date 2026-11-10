package orders

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	uid := c.GetInt("userID")

	orderList, err := models.GetOrdersByUser(database.DB, uid)
	if err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	if orderList == nil {
		orderList = []models.Order{}
	}

	handlers.SuccessJSON(c, gin.H{"orders": orderList})
}
