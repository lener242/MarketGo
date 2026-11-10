package cart

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewCart(c *gin.Context) {
	uid := c.GetInt("userID")

	items, err := models.GetCartItems(database.DB, uid)
	if err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	total := 0
	type CartItem struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
		Price    int    `json:"price"`
		Subtotal int    `json:"subtotal"`
	}

	var result []CartItem
	for _, it := range items {
		result = append(result, CartItem{
			Name:     it.Name,
			Quantity: it.Quantity,
			Price:    it.Price,
			Subtotal: it.Subtotal,
		})
		total += it.Subtotal
	}

	handlers.SuccessJSON(c, gin.H{
		"items": result,
		"total": total,
	})
}
