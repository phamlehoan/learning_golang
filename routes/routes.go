package routes

import (
	"hello/controllers/books_controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Route("/api/v1/books", func(route fiber.Router) {
		route.Get("/", books_controller.Index)
		route.Get("/:id", books_controller.Show)
		route.Post("/", books_controller.Create)
		route.Put("/:id", books_controller.Update)
		route.Delete("/:id", books_controller.Delete)
	})
}
