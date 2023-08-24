package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // new
	"github.com/huybne/api-todo-fiber/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})

	})
	api := app.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})
	routes.TodoRoute(api.Group("/todos"))
}
func main() {
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
