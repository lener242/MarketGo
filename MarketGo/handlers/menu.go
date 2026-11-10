package handlers

import (
	"MarketGo/database"
	"MarketGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MenuHandler(c *gin.Context) {
	category := c.Query("category")
	var (
		dishes []models.Dish
		err    error
	)
	if category == "" {
		dishes, err = models.GetAllDishes(database.DB)
	} else {
		dishes, err = models.GetDishesByCategory(database.DB, category)
	}
	if err != nil {
		ErrorJSON(c, http.StatusInternalServerError, "db error: "+err.Error())
		return
	}
	SuccessJSON(c, gin.H{
		"categories": []string{"Хлебобулочные изделия", "Напитки", "Мясные изделия"},
		"dishes":     dishes,
	})
}
