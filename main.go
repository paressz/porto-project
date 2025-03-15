package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"porto-project/api/routes"
	"porto-project/pkg/User"
	"porto-project/pkg/projects"
)
func setupRoutes(router fiber.Router) {
	projectRepo := projects.NewRepository()
	projectService := projects.NewService(projectRepo)
	userRepo := User.NewRepository()
	userService := User.NewService(userRepo)
	routes.SetupLoginRoutes(router, userService)
	routes.SetupProjectsRoutes(router, projectService)
}
func main() {
	app := fiber.New()
	api := app.Group("/api")
	setupRoutes(api)
	app.Use(func (c *fiber.Ctx) error {
		return c.Status(403).SendString("Unauthorized")
	})
	log.Fatal(app.Listen(":8080"))
}
