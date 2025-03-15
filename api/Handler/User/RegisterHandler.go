package User

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/log"
    "porto-project/api/presenter"
    "porto-project/pkg/User"
    "porto-project/pkg/model"
    "porto-project/pkg/util/auth"
)

func RegisterUser(s User.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(model.User)
		err := c.BodyParser(&user)
		if err != nil {
			log.Debugf("Failed registering user: %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status: "Failed",
				Message: "Unable to register",
				Error: err.Error(),
			})
		}
		hash := auth.HashPassword(user.Password)
		user.Password = hash
        err = s.RegisterUser(user)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(presenter.FailedResponse{
				Status: "Failed",
				Message: "Unable to register",
				Error: err.Error(),
			})
        }
		return c.Status(fiber.StatusCreated).JSON(presenter.UserSuccessResponse{
			Status: "Success",
			Message: "Register Success",
			User: presenter.UserResponse{
				Email: user.Email,
				Username: user.Username,
				Name: user.Name,
			},
		})
	}
}