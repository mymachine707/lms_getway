package models

// LoginModul ...
type LoginModul struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse ...
type TokenResponse struct {
	Token string `json:"token"`
}
