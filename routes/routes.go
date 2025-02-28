package routes

import (
	"github.com/ericNKS/gommerce/app/controller"
	"github.com/gofiber/fiber/v2"
)

func Exec(app *fiber.App) {
	userController := controller.UserController{}

	app.Post("/user", userController.Store)
}
