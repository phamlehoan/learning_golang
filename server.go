package main

import (
	"hello/models"
	"hello/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()
	defer models.CloseDatabaseConnection(models.DB)

	app := fiber.New()

	app.Static("/", "./public")

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
