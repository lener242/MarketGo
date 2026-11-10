package models

import (
	"database/sql"
	"strings"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Username  sql.NullString `json:"username"`
	Phone     string         `json:"phone"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
}

func CreateUser(db *sql.DB, name, phone, password string) error {
	query := `INSERT INTO users (name, phone, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, name, phone, password, time.Now())
	return err
}

func GetUserByPhone(db *sql.DB, phone string) (*User, error) {
	var user User
	trimmedPhone := strings.TrimSpace(phone)

	query := `SELECT id, name, username, phone, password, created_at 
	          FROM users 
	          WHERE TRIM(phone) = $1`

	err := db.QueryRow(query, trimmedPhone).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Phone,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	query := `SELECT id, name, username, phone, created_at FROM users WHERE id=$1`
	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Phone,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *sql.DB, id int, username, phone string) error {
	query := `UPDATE users SET username=$1, phone=$2 WHERE id=$3`
	_, err := db.Exec(query, username, phone, id)
	return err
}
