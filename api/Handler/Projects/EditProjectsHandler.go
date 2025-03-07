package Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"path/filepath"
	"porto-project/api/presenter"
	"porto-project/pkg/model"
	"porto-project/pkg/projects"
	"porto-project/pkg/util"
)

func EditProject(s projects.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := new(model.Project)
		id := c.Params("id")
		project.Id = id
		project.Name, project.Description = c.FormValue("name"), c.FormValue("description")

		img, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Key: 'image' is not found or empty",
				Error:   err.Error(),
			})
		}
		if !util.IsImage(img) {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed to upload image",
				Error:   "Invalid MIME type or extension",
			})
		}
		imgPath, err := util.SaveImage(c, img, project.Id)
		if err != nil {
			log.Debugf("")
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed to save image",
				Error:   err.Error(),
			})
		}
		err = util.CompressImage(imgPath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed when compressing image",
				Error:   err.Error(),
			})
		}
		
		project.ImageUrl = "/api/projects/images/" + filepath.Base(imgPath)
		updatedId, err := s.EditProject(project)
		if err != nil {
			log.Debug("Failed updating project in database: " + err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed updating project",
				Error:   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ProjectSuccessResponse{
			Status:  "Success",
			Message: "Updated project " + updatedId,
			Project: project,
		})
	}
}
