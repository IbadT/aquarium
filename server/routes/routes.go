package routes

import (
	"aquarium/handlers"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/users", handlers.GetUsers)

	app.Post("/user", handlers.CreateUser)
}
