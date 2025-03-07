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
			return c.Status(fiber.StatusOK).JSON(presenter.ProjectsSuccessResponse{
				Status:    "Success",
				Message:   "Projects fetched",
				PageCount: pageCount,
				Project:   projectList,
			})
		}
		lastIndex := len(projectList)-1
		lastIntId := projectList[lastIndex].IntId
		if err != nil {
			log.Debugf("Failed to fetch projects: %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Unable to fetch projects, try again later",
				Error:   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ProjectsSuccessResponse{
			Status:    "Success",
			Message:   "Projects fetched",
			LastIntId: lastIntId,
			PageCount: pageCount,
			Project:   projectList,
		})
	}
}

func GetProjectById(s projects.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Invalid Id is empty",
			})
		}
		project, err := s.GetProjectById(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Unable to fetch project with Id: " + id,
				Error:   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ProjectSuccessResponse{
			Status:  "Success",
			Message: "Fetched project with Id: " + project.Id,
			Project: project,
		})
	}
}