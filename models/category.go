package models

import "time"

// Category ..
type Category struct {
	ID        string     `json:"id"`
	Title     string     `json:"title" binding:"required" example:"Novel"`
	Status    string     `json:"status" binding:"required" example:"enabled"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// CreateCategoryModul ..
type CreateCategoryModul struct {
	Title string `json:"title" binding:"required" example:"Novel"`
}

// UpdateCategoryModul ..
type UpdateCategoryModul struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required" example:"Novel"`
}
