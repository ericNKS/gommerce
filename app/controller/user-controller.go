package controller

import (
	"github.com/ericNKS/gommerce/app/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (uc *UserController) Store(ctx *fiber.Ctx) error {
	user := models.UserCreate{}

	if err := ctx.BodyParser(&user); err != nil {
		ctx.Status(403)
		return err
	}

	return ctx.JSON(models.UserResponse{})
}
