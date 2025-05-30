package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ok1503f/handler"
)

func RegisterRoutes(app *fiber.App, h *handler.UserHandler) {
	api := app.Group("/api")
	{
		users := api.Group("/users")
		{
			users.Post("/", h.CreateUser)
			users.Get("/", h.GetAllUsers)
			users.Get("/:id", h.GetUserByID)
			users.Post("/login", h.Login)
		}
	}
}
