package models

import "time"

// Client ..
type Client struct {
	Id          string     `json:"id"`
	Firstname   string     `json:"firstname" binding:"required" example:"John"`
	Lastname    string     `json:"lastname" binding:"required" example:"Doe"`
	Username    string     `json:"username" binding:"required" example:"Jdoe002"`
	PhoneNumber string     `json:"phoneNumber" binding:"required" example:"998797456321"`
	Address     string     `json:"address" binding:"required" example:"wall street"`
	Type        string     `json:"type" binding:"required" example:"client"`
	Password    string     `json:"password" binding:"required" example:"client"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// CreateClientModul ..
type CreateClientModul struct {
	Firstname   string `json:"firstname" binding:"required" example:"John"`
	Lastname    string `json:"lastname" binding:"required" example:"Doe"`
	Username    string `json:"username" binding:"required" example:"Jdoe002"`
	PhoneNumber string `json:"phoneNumber" binding:"required" example:"998797456321"`
	Address     string `json:"address" binding:"required" example:"wall street"`
	Type        string `json:"type" binding:"required" example:"client"`
	Password    string `json:"password" binding:"required" example:"client"`
}

// UpdateClientModul ..
type UpdateClientModul struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phoneNumber" binding:"required" example:"998797456321"`
	Address     string `json:"address" binding:"required" example:"wall street"`
	Password    string `json:"password" binding:"required" example:"client"`
}
