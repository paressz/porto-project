package Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"path/filepath"
	"porto-project/api/presenter"
	"porto-project/pkg/model"
	"porto-project/pkg/projects"
	"porto-project/pkg/util/file"
)

func AddProject(s projects.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := new(model.Project)
		name, description := c.FormValue("name"), c.FormValue("description")
		project.Name, project.Description = name, description

		newId, err := gonanoid.New()
		if err != nil {
			log.Debugf("Error generating ID: %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed when generating ID",
				Error:   err.Error(),
			})
		}
		project.Id = "project_" + newId


		img, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Key: 'image' is not found or empty",
				Error:   err.Error(),
			})
		}
		if !file.IsImage(img) {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed to upload image",
				Error:   "Invalid MIME type or extension",
			})
		}
		imgPath, err := file.SaveImage(c, img, project.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed to save image",
				Error:   err.Error(),
			})
		}
		err = file.CompressImage(imgPath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status:  "Failed",
				Message: "Failed when compressing image",
				Error:   err.Error(),
			})
		}
		
		project.ImageUrl = "/api/projects/images/" + filepath.Base(imgPath)
		result, err := s.CreateProject(project)
		return c.Status(fiber.StatusCreated).JSON(presenter.ProjectSuccessResponse{
			Status:  "Success",
			Message: "Project Created",
			Project: result,
		})
	}
}
