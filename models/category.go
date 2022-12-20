package models

import "time"

// Category ..
type Category struct {
	ID            string     `json:"id"`
	Category_name string     `json:"Category_name" binding:"required" example:"Lavash"`
	Description   string     `json:"Description" binding:"required" example:"Qaynoq va mazali Lavash"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

// CreateCategoryModul ..
type CreateCategoryModul struct {
	Category_name string `json:"Category_name" binding:"required" example:"Lavash"`
	Description   string `json:"Description" binding:"required" example:"Qaynoq va mazali Lavash"`
}

// UpdateCategoryModul ..
type UpdateCategoryModul struct {
	ID            string `json:"id" binding:"required"`
	Category_name string `json:"Category_name" binding:"required" example:"Lavash"`
	Description   string `json:"Description" binding:"required" example:"Qaynoq va mazali Lavash"`
}
