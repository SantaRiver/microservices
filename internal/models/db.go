package models

import (
	"database/sql"
)

type DBPayment struct {
	ID         int64        `db:"id"`
	UUID       string       `db:"uuid"`
	OrderID    int64        `db:"order_id"`
	TotalPrice TotalPrice   `db:"total_price"`
	Status     int64        `db:"status"`
	CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type DBOrder struct {
	ID         int64          `db:"id"`
	UUID       string         `db:"uuid"`
	UserID     int64          `db:"user_id"`
	BasketID   int64          `db:"basket_id"`
	Status     int64          `db:"status"`
	TotalPrice TotalPrice     `db:"total_price"`
	Error      sql.NullString `db:"error"`
	CreatedAt  sql.NullTime   `db:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at"`
}
type TotalPrice struct {
	Price    float64 `db:"total_price"`
	Currency int     `db:"currency"`
}
type DBBasket struct {
	ID         int64        `db:"id"`
	UserID     int64        `db:"user_id"`
	Products   string       `db:"products"`
	Count      string       `db:"count"`
	TotalPrice float64      `db:"total_price"`
	CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type DBBasketProduct struct {
	ID    int64 `db:"id"`
	Count int64 `db:"quantity"`
}

type DBProduct struct {
	ID          int64        `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	Price       float32      `db:"price"`
	Image       string       `db:"image"`
	Category    string       `db:"category"`
	Subcategory string       `db:"subcategory"`
	Brand       string       `db:"brand"`
	Model       string       `db:"model"`
	Color       string       `db:"color"`
	Size        string       `db:"size"`
	Material    string       `db:"material"`
	Country     string       `db:"country"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type DBUser struct {
	ID         int64         `db:"id"`
	Name       string        `db:"name"`
	Surname    string        `db:"surname"`
	Patronymic string        `db:"patronymic"`
	Type       sql.NullInt16 `db:"type"`
	Birthday   sql.NullTime  `db:"birthday"`
	CreatedAt  sql.NullTime  `db:"created_at"`
	UpdatedAt  sql.NullTime  `db:"updated_at"`
}
