package models

// Rental ..
type Rental struct {
	ID                 string `json:"id"`
	BookID             string `json:"book_id" binding:"required" `
	BookName           string `json:"book_name" binding:"required" `
	UserID             string `json:"user_id" binding:"required" `
	RentalDateTime     string `json:"rental_date_time" binding:"required" `
	ExpectedReturnDate string `json:"expected_return_date" binding:"required" `
	ReturnDate         string `json:"return_date" binding:"required" `
	BookFines          string `json:"book_fines" binding:"required" `
	RentalStatus       string `json:"rental_status" binding:"required" `
}

// CreateRentalModul ..
type CreateRentalModul struct {
	BookID             string `json:"book_id" binding:"required" `
	BookName           string `json:"book_name" binding:"required" `
	UserID             string `json:"user_id" binding:"required" `
	ExpectedReturnDate string `json:"expected_return_date" binding:"required" `
}

// UpdateRentalModul ..
type UpdateRentalModul struct {
	ID           string `json:"id"`
	BookID       string `json:"book_id" binding:"required" `
	BookName     string `json:"book_name" binding:"required" `
	UserID       string `json:"user_id" binding:"required" `
	BookFines    string `json:"book_fines" binding:"required" `
	RentalStatus string `json:"rental_status" binding:"required" `
}
