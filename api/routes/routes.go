package routes

import (
	"github.com/gofiber/fiber/v2"
	projectHandler "porto-project/api/Handler/Projects"
	userHandler "porto-project/api/Handler/User"
	"porto-project/api/middleware"
	"porto-project/pkg/User"
	"porto-project/pkg/projects"
)

func SetupProjectsRoutes(router fiber.Router, service projects.Service) {
	//url: /api/
	router.Get("/projects", projectHandler.GetProjects(service))
	router.Get("/projects/:id", projectHandler.GetProjectById(service))
	router.Static("/projects/images/", "./uploads/projects/")
	router.Use(middleware.JWTProtected())
	router.Post("/projects", projectHandler.AddProject(service))
	router.Delete("/projects/:id", projectHandler.DeleteProject(service))
	router.Put("/projects/:id", projectHandler.EditProject(service))
}

func SetupLoginRoutes(router fiber.Router, service User.Service) {
	router.Get("/login", userHandler.Login(service))
	router.Get("/register", userHandler.RegisterUser(service))
}