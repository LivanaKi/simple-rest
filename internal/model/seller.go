package model

import (
	"errors"
	"time"
)

// Seller (Продавець)
type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Product (Товар)
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerID    int     `json:"seller_id"`
}

// Customer (Покупець)
type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Order (Замовлення)
type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}

// OrderProduct (Товари в замовленні)
type OrderProduct struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (s *Seller) Validation() error {
	if s.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}
