package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marieesss/API-ManageBudgetApp/database"
	"github.com/marieesss/API-ManageBudgetApp/handlers"
)

func main() {
	database.ConnectDb()
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World  👋!")
	})
	app.Post("/", handlers.CreateUser)
	// Start the server on port 3000
	app.Listen(":3000")
}
