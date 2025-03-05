package Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"porto-project/api/presenter"
	"porto-project/pkg/projects"
)
func GetProjects(s projects.Service) fiber.Handler{
	return func(c *fiber.Ctx) error {
		lastIId := c.QueryInt("last_int_id", 0)
		projectList, pageCount, err := s.GetAllProjects(lastIId)
		if len(projectList) < 1 || projectList == nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				"Failed",
				"Failed fetching projects",
				"No projects with int_id > last_int_id",
			})
		}
		lastIndex := len(projectList)-1
		lastIntId := projectList[lastIndex].IntId
		if err != nil {
			log.Debugf("Failed to fetch projects: %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				"Failed",
				"Unable to fetch projects, try again later",
				err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ProjectsSuccessResponse{
			"Success",
			"Projects fetched",
			lastIntId,
			pageCount,
			projectList,
		})
	}
}

func GetProjectById(s projects.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				"Failed",
				"Invalid Id is empty",
				"",
			})
		}
		project, err := s.GetProjectById(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				"Failed",
				"Unable to fetch project with Id: " + id,
				err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ProjectSuccessResponse{
			"Success",
			"Fetched project with Id: " + project.Id,
			project,
		})
	}
}