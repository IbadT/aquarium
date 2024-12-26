package main

import (
	"aquarium/database"
	"aquarium/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()
	database.Connect()

	app.Use(logger.New())

	routes.Setup(app)

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":8080"))
}
