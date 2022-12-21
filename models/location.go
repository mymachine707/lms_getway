package models

import "time"

// Location ..
type Location struct {
	ID        string     `json:"id"`
	Name      string     `json:"name" binding:"required" example:"A1"`
	Status    string     `json:"status" binding:"required" example:"enabled"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// CreateLocationModul ..
type CreateLocationModul struct {
	Name string `json:"name" binding:"required" example:"A1"`
}

// UpdateLocationModul ..
type UpdateLocationModul struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required" example:"A1"`
}
