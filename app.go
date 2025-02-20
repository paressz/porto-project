package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"porto-project/database"
	"porto-project/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/projects", routes.GetProjects)
	app.Post("/projects", routes.AddProject)
	app.Use(func (c *fiber.Ctx) error {
		return c.Status(403).SendString("Unauthorized")
	})
}