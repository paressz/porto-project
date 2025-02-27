package routes

import (
	"github.com/gofiber/fiber/v2"
	Handler "porto-project/api/Handler/Projects"
	"porto-project/pkg/projects"
)

func SetupRoutes(app *fiber.App, service projects.Service) {
	app.Get("/projects", Handler.GetProjects(service))
	app.Post("/projects", Handler.AddProject(service))
	app.Get("/projects/:id", Handler.GetProjectById(service))
	app.Use(func (c *fiber.Ctx) error {
		return c.Status(403).SendString("Unauthorized")
	})
}