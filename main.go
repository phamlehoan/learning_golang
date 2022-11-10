package main

import (
	"hello/book"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, world!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Route("/api/v1/books", func(route fiber.Router) {
		route.Get("/", book.GetBooks)
	})
}

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	setupRoutes(app)

	app.Listen(":3000")
}
