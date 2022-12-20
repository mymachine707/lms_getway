package models

import "time"

// Product ...
type Product struct {
	ID           string     `json:"id"`
	Category_id  string     `json:"category_id" binding:"required"`
	Product_name string     `json:"product_name" binding:"required"`
	Description  string     `json:"description" binding:"required"`
	Price        int32      `json:"price" binding:"required"`
	Created_at   time.Time  `json:"created_at"`
	Updated_at   *time.Time `json:"updated_at"`
}

// CreateProductModul ...
type CreateProductModul struct {
	Category_id  string `json:"category_id" binding:"required"`
	Product_name string `json:"product_name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Price        int32  `json:"price" binding:"required"`
}

// UpdateProductModul ...
type UpdateProductModul struct {
	ID          string `json:"id"`
	Description string `json:"description" binding:"required"`
	Price       int32  `json:"price" binding:"required"`
}
