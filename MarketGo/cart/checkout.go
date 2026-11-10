package cart

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context) {
	uid := c.GetInt("userID")

	items, err := models.GetCartItems(database.DB, uid)
	if err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	if len(items) == 0 {
		handlers.ErrorJSON(c, http.StatusBadRequest, "корзина пуста")
		return
	}

	total := 0
	var orderItems []models.OrderItem
	for _, it := range items {
		total += it.Subtotal
		orderItems = append(orderItems, models.OrderItem{
			DishID:   it.DishID,
			Name:     it.Name,
			Quantity: it.Quantity,
			Price:    it.Price,
			Subtotal: it.Subtotal,
		})
	}

	orderID, err := models.CreateOrder(database.DB, uid, total, orderItems)
	if err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "failed to save order")
		return
	}

	if err := models.ClearCart(database.DB, uid); err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	handlers.SuccessJSON(c, gin.H{
		"message":  "Заказ оформлен!",
		"order_id": orderID,
		"items":    orderItems,
		"total":    total,
	})
}
