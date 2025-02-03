package routes

import (
	"github.com/ericNKS/gommerce/app/repository"
	"github.com/gofiber/fiber/v2"
)

func Exec(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		_, err := repository.UserRepository()
		if err != nil {
			return err
		}
		return ctx.SendString("Conexao feita")
	})
}
