package routes

import (
	"MarketGo/cart"
	"MarketGo/handlers"
	"MarketGo/orders"
	"MarketGo/profile"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", handlers.RegisterHandler)
		api.POST("/login", handlers.LoginHandler)
		api.GET("/menu", handlers.MenuHandler)
	}

	protected := r.Group("/api")
	protected.Use(handlers.AuthMiddleware)
	{
		protected.GET("/profile", profile.ViewProfile)
		protected.PUT("/profile", profile.UpdateProfile)

		protected.GET("/orders", orders.GetOrders)

		c := protected.Group("/cart")
		{
			c.GET("", cart.ViewCart)
			c.POST("/add", cart.AddToCart)
			c.PUT("/update", cart.UpdateQuantity)
			c.DELETE("/remove", cart.RemoveFromCart)
			c.DELETE("/clear", cart.ClearCart)
			c.POST("/checkout", cart.Checkout)
		}
	}

	return r
}
