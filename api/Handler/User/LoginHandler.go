package User

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"porto-project/api/presenter"
	"porto-project/pkg/User"
	"porto-project/pkg/util/auth"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func Login(s User.Service) fiber.Handler  {
	return func(c *fiber.Ctx) error {
		req := new(LoginRequest)
		log.Debug("123" + req.Email + req.Password)
		err := c.BodyParser(&req)
		if err != nil {
			log.Debugf("Failed logging in: %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status: "Failed",
				Message: "Failed logging in",
				Error: err.Error(),
			})
		}
		user, err := s.GetUser(req.Email, req.Password)
		if err != nil {
			log.Debugf("Failed logging in: %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status: "Failed",
				Message: "Failed logging in",
				Error: "Wrong email or password",
			})
		}
		token,err := auth.GenerateToken(user.Email)
		if err != nil {
			log.Debugf("Failed logging in: %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(presenter.FailedResponse{
				Status: "Failed",
				Message: "Failed logging in",
				Error: err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	}
}