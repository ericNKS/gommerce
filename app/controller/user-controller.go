package controller

import (
	"github.com/ericNKS/gommerce/app/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New() // Inicializa o validador

type UserController struct{}

func (uc *UserController) Store(ctx *fiber.Ctx) error {
	user := models.UserCreate{}

	if err := ctx.BodyParser(&user); err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(user); err != nil {
		// Se houver erro de validação, retornamos um erro
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"error": validationErrors})
	}

	return ctx.JSON(user)
}
