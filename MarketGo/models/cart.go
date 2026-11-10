package models

import (
	"database/sql"
	"time"
)

type CartItem struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	DishID    int       `json:"dish_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func AddToCart(db *sql.DB, userID, dishID, quantity int) error {
	query := `
	INSERT INTO cart (user_id, dish_id, quantity)
	VALUES ($1,$2,$3)
	ON CONFLICT (user_id, dish_id) DO UPDATE SET quantity = cart.quantity + EXCLUDED.quantity
	`
	_, err := db.Exec(query, userID, dishID, quantity)
	return err
}

func UpdateCartQuantity(db *sql.DB, userID, dishID, quantity int) error {
	_, err := db.Exec("UPDATE cart SET quantity=$1 WHERE user_id=$2 AND dish_id=$3", quantity, userID, dishID)
	return err
}

func RemoveFromCart(db *sql.DB, userID, dishID int) error {
	_, err := db.Exec("DELETE FROM cart WHERE user_id=$1 AND dish_id=$2", userID, dishID)
	return err
}

func ClearCart(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM cart WHERE user_id=$1", userID)
	return err
}

func GetCartItems(db *sql.DB, userID int) ([]struct {
	DishID   int
	Name     string
	Quantity int
	Price    int
	Subtotal int
}, error) {
	rows, err := db.Query(`
	SELECT c.dish_id, d.name, c.quantity, d.price, (c.quantity * d.price) AS subtotal
	FROM cart c
	JOIN dishes d ON d.id = c.dish_id
	WHERE c.user_id = $1
	ORDER BY c.id
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type Out struct {
		DishID   int
		Name     string
		Quantity int
		Price    int
		Subtotal int
	}

	var out []Out
	for rows.Next() {
		var it Out
		if err := rows.Scan(&it.DishID, &it.Name, &it.Quantity, &it.Price, &it.Subtotal); err != nil {
			return nil, err
		}
		out = append(out, it)
	}

	res := make([]struct {
		DishID   int
		Name     string
		Quantity int
		Price    int
		Subtotal int
	}, len(out))
	for i := range out {
		res[i] = out[i]
	}
	return res, nil
}
