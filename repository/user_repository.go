package repository

import (
	"database/sql"
	"fmt"

	"github.com/ok1503f/models"
)

type UserRepository interface {
	CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error)
	FindAll() ([]models.UserResponse, error)
	FindByID(id int) (*models.UserResponse, error)
	FindByEmail(email string) (*models.UserResponse, error)
	UpDateBalance(id int, balance float64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO users (name, email, password, balance) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.Password, user.Balance,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Balance:  user.Balance,
	}, nil
}

func (r *userRepository) FindAll() ([]models.UserResponse, error) {
	rows, error := r.db.Query("SELECT id, name, email, balance FROM users")

	if error != nil {
		return nil, error
	}

	defer rows.Close()

	var users []models.UserResponse

	for rows.Next() {
		var u models.UserResponse
		if error := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Balance); error != nil {
			return nil, error
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) FindByID(id int) (*models.UserResponse, error) {
	query := "SELECT id, name, email, balance FROM users WHERE id = $1"
	var user models.UserResponse

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Balance,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.UserResponse, error) {
	query := "SELECT id, name, email, password, balance FROM users WHERE email = $1"
	var user models.UserResponse

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Balance,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpDateBalance(userId int, newBalance float64) error {
	result, err := r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userId)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", userId)
	}

	return nil
}
