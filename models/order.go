package models

import "time"

// Order ..
type Order struct {
	ID         string     `json:"id"`
	Product_id string     `json:"product_id" binding:"required" example:"111"`
	Client_id  string     `json:"client_id" binding:"required" example:"111"`
	Count      int32      `json:"count" binding:"required" example:"2"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

// CreateOrderModul ..
type CreateOrderModul struct {
	// Client_id  string       `json:"client_id" binding:"required" example:"111"`
	OrderItems []*OrderItem `json:"orderItems" binding:"required"`
}

type OrderItem struct {
	Product_id string `json:"product_id" binding:"required" example:"111"`
	Quantity   int32  `json:"quantity" binding:"required" example:"111"`
}
