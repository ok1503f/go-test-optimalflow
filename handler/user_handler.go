package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ok1503f/models"
	"github.com/ok1503f/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.CreateUserRequest
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	_, err := h.service.CreateUser(&user)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var user models.LoginRequest
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	_, err := h.service.Authenticate(user.Email, user.Password)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).SendString("Login successful")
}
