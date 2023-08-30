package main

import (
	"github.com/gofiber/fiber/v2"
	"todoist/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.GetAllTodos)
	app.Post("/", handlers.CreateTask)
	app.Patch("/:id", handlers.CompletedTask)
	app.Delete("/:id", handlers.DeleteTask)

	app.Listen(":8080")
}
