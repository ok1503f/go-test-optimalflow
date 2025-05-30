package models

type UserResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type CreateUserRequest struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required"`
	Balance  float64 `json:"balance" default:"100"`
}

type LoginRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type TransferRequest struct {
	FromID int     `json:"from_id" validate:"required"`
	ToID   int     `json:"to_id" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}
