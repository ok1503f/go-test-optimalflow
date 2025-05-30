package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ok1503f/config"
	"github.com/ok1503f/handler"
	"github.com/ok1503f/repository"
	"github.com/ok1503f/routes"
	"github.com/ok1503f/service"
)

func main() {

	db := config.InitDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	routes.RegisterRoutes(app, &userHandler)

	app.Listen(":3000")
}
