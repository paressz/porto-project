package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"porto-project/api/routes"
	"porto-project/pkg/projects"
)

func main() {
	app := fiber.New()
	projectRepo := projects.NewRepository()
	projectService := projects.NewService(projectRepo)
	routes.SetupRoutes(app, projectService)
	log.Fatal(app.Listen(":8080"))
}
