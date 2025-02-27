package Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"porto-project/api/presenter"
	"porto-project/pkg/projects"
)

func AddProject(s projects.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := new(projects.Project)
		err := c.BodyParser(&project)
		if err != nil {
			log.Debugf("Failed to parse body: %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				`Failed`,
				`Invalid request body`,
				err.Error(),
			})
		}
		newId, err := gonanoid.New()
		if err != nil {
			log.Debugf("Error generating ID: %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				"Failed",
				"Failed when generating ID",
				err.Error(),
			})
		}
		project.Id = newId
		res, err := s.CreateProject(project)
		return c.Status(fiber.StatusCreated).JSON(presenter.ProjectSuccessResponse{
			"Success",
			"Project Created",
			res,
		})
	}
}