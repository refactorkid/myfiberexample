package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		log.Println("Incoming request:", c.Path(), c.Method(), c.OriginalURL(), c.IP())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	userRoutes := app.Group("/users")
	userRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Users list",
		})
	})

	app.Listen(":3000")
}