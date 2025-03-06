package controller

import (
	"github.com/ericNKS/gommerce/app/models"
	"github.com/ericNKS/gommerce/app/repository"
	"github.com/ericNKS/gommerce/app/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New() // Inicializa o validador

type UserController struct{}

func (uc *UserController) Store(ctx *fiber.Ctx) error {
	userCreate := &models.UserCreate{}

	if err := ctx.BodyParser(userCreate); err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var customMessages = map[string]string{
		"Name.required":     "Name is required and should have at least 4 characters",
		"Name.min":          "Name must have at least 4 characters",
		"Email.required":    "A valid email is required",
		"Email.email":       "Please enter a valid email address",
		"Password.required": "Password is required and should have at least 8 characters",
		"Password.min":      "Password must have at least 8 characters",
	}

	if err := validate.Struct(userCreate); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field() + "." + err.Tag()
			if msg, exists := customMessages[fieldName]; exists {
				validationErrors = append(validationErrors, msg)
			} else {
				validationErrors = append(validationErrors, err.Error()) // fallback padrão
			}
		}
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"error": validationErrors})
	}

	ur, err := repository.UserRepository()
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"error": "error in repository"})
	}

	userCreateService := services.UserCreateService(ur)

	user := userCreate.ConvertToUser()
	errs := userCreateService.Execute(user)
	if errs != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"error": err})
	}

	return ctx.JSON(user.Response())
}
