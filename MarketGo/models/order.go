package models

import (
	"database/sql"
	"time"
)

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	Total     int         `json:"total"`
	CreatedAt time.Time   `json:"created_at"`
	Items     []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID       int    `json:"id"`
	OrderID  int    `json:"order_id"`
	DishID   int    `json:"dish_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Subtotal int    `json:"subtotal"`
}

func CreateOrder(db *sql.DB, userID, total int, items []OrderItem) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var orderID int
	err = tx.QueryRow(
		`INSERT INTO orders (user_id, total, created_at) VALUES ($1, $2, NOW()) RETURNING id`,
		userID, total,
	).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	for _, item := range items {
		_, err = tx.Exec(
			`INSERT INTO order_items (order_id, dish_id, name, quantity, price, subtotal)
			 VALUES ($1, $2, $3, $4, $5, $6)`,
			orderID, item.DishID, item.Name, item.Quantity, item.Price, item.Subtotal,
		)
		if err != nil {
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return orderID, nil
}

func GetOrdersByUser(db *sql.DB, userID int) ([]Order, error) {
	rows, err := db.Query(
		`SELECT id, user_id, total, created_at FROM orders WHERE user_id=$1 ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Total, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	for i, o := range orders {
		items, err := getOrderItems(db, o.ID)
		if err != nil {
			return nil, err
		}
		orders[i].Items = items
	}

	return orders, nil
}

func getOrderItems(db *sql.DB, orderID int) ([]OrderItem, error) {
	rows, err := db.Query(
		`SELECT id, order_id, dish_id, name, quantity, price, subtotal FROM order_items WHERE order_id=$1`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []OrderItem
	for rows.Next() {
		var it OrderItem
		if err := rows.Scan(&it.ID, &it.OrderID, &it.DishID, &it.Name, &it.Quantity, &it.Price, &it.Subtotal); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}
