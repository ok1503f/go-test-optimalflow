package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ok1503f/handler"
	"github.com/ok1503f/middleware"
)

func RegisterRoutes(app *fiber.App, h *handler.UserHandler) {
	api := app.Group("/api")
	{
		api.Post("/login", h.Login)
		api.Use(middleware.JWTProtected())
		users := api.Group("/users")
		{
			users.Post("/", h.CreateUser)
			users.Get("/", h.GetAllUsers)
			users.Get("/:id", h.GetUserByID)
			users.Post("/transfer", h.TransferBalance)
		}
	}
}
