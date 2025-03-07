package Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"porto-project/api/presenter"
	"porto-project/pkg/projects"
)

func DeleteProject(s projects.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectId := c.Params("id")
		err := s.DeleteProject(projectId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Unable to delete project",
				Error:   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "Success",
			"message": "Project deleted",
		})
	}
}
