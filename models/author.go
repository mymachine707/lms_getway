package models

import "time"

// Author ..
type Author struct {
	ID        string     `json:"id"`
	Name      string     `json:"name" binding:"required" example:"Mark Twains"`
	Status    string     `json:"status" binding:"required" example:"enabled"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// CreateAuthorModul ..
type CreateAuthorModul struct {
	Name string `json:"name" binding:"required" example:"Mark Twains"`
}

// UpdateAuthorModul ..
type UpdateAuthorModul struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required" example:"Mark Twains"`
}
