package routes

import (
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"porto-project/database"
	"porto-project/models"
)
func GetProjects(c *fiber.Ctx) error{
	var projects []models.Project
	database.Db.Instance.Find(&projects)
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"projects": projects,
	})
}

func AddProject(c *fiber.Ctx) error {
	project := new(models.Project)
	err := c.BodyParser(&project)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	generatedId, err := gonanoid.New(16)
	if err != nil {
		log.Println("Failed to generate ID")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	project.Id = generatedId
	err = database.Db.Instance.Create(&project).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	type ProjectResponse struct {
    Status  string  `json:"status"`
    Message string  `json:"message"`
    Project models.Project `json:"project"`
}

	return c.Status(fiber.StatusCreated).JSON(ProjectResponse{
		Status: "Success",
		Message: "Project Created",
		Project: *project,
	})
}