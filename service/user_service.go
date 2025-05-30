package service

import (
	"errors"

	"github.com/ok1503f/models"
	"github.com/ok1503f/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error)
	GetUserByID(id int) (*models.UserResponse, error)
	GetAllUsers() ([]models.UserResponse, error)
	Authenticate(email, password string) (*models.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user *models.CreateUserRequest) (*models.UserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	if user.Balance <= 0 {
		user.Balance = 100
	}
	user.Password = string(hashedPassword)
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(id int) (*models.UserResponse, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) GetAllUsers() ([]models.UserResponse, error) {
	return s.userRepo.FindAll()
}

func (s *userService) Authenticate(email, password string) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
