package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ok1503f/models"
	"github.com/ok1503f/service"
	"github.com/ok1503f/utils"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := h.service.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user"})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := h.service.Authenticate(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, error := utils.GenerateJWT(user.ID)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful", "token": token})
}

func (h *UserHandler) TransferBalance(c *fiber.Ctx) error {
	var transfer models.TransferRequest
	if err := c.BodyParser(&transfer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.TransferBalance(transfer.FromID, transfer.ToID, transfer.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to transfer balance"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Balance transferred successfully"})
}
