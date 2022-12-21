package models

import "time"

// Book ..
type Book struct {
	Id         string     `json:"id"`
	Name       string     `json:"name" binding:"required" example:"War and Peace"`
	AuthorID   string     `json:"authorID" binding:"required" example:"uuid"`
	CategoryID string     `json:"categoryID" binding:"required" example:"uuid"`
	LocationID string     `json:"locationID" binding:"required" example:"uuid"`
	ISBN       string     `json:"ISBN" binding:"required" example:"4564231"`
	Quantity   string     `json:"quantity" binding:"required" example:"10"`
	Status     string     `json:"status" binding:"required" example:"enabled"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

// CreateBookModul ..
type CreateBookModul struct {
	Name       string `json:"name" binding:"required" example:"War and Peace"`
	AuthorID   string `json:"authorID" binding:"required" example:"uuid"`
	CategoryID string `json:"categoryID" binding:"required" example:"uuid"`
	LocationID string `json:"locationID" binding:"required" example:"uuid"`
	ISBN       int32  `json:"ISBN" binding:"required" example:"4564231"`
	Quantity   int32  `json:"quantity" binding:"required" example:"10"`
}

// UpdateBookModul ..
type UpdateBookModul struct {
	Id         string `json:"id"`
	Name       string `json:"name" binding:"required" example:"War and Peace"`
	LocationID string `json:"locationID" binding:"required" example:"uuid"`
	Quantity   int32  `json:"quantity" binding:"required" example:"10"`
}
