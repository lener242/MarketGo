package models

import "database/sql"

type Dish struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func GetDishByID(db *sql.DB, id int) (*Dish, error) {
	var d Dish
	err := db.QueryRow("SELECT id, name, category, price FROM dishes WHERE id=$1", id).
		Scan(&d.ID, &d.Name, &d.Category, &d.Price)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func GetAllDishes(db *sql.DB) ([]Dish, error) {
	rows, err := db.Query("SELECT id, name, category, price FROM dishes ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Dish
	for rows.Next() {
		var d Dish
		if err := rows.Scan(&d.ID, &d.Name, &d.Category, &d.Price); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

func GetDishesByCategory(db *sql.DB, category string) ([]Dish, error) {
	rows, err := db.Query("SELECT id, name, category, price FROM dishes WHERE category=$1 ORDER BY id", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Dish
	for rows.Next() {
		var d Dish
		if err := rows.Scan(&d.ID, &d.Name, &d.Category, &d.Price); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}
