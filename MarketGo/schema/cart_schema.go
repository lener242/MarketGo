package schema

type CartItemSchema struct {
	DishID   int    `json:"dish_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Subtotal int    `json:"subtotal"`
}
