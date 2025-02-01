package routes

import "github.com/gofiber/fiber/v2"

func Exec(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello teste 3")
	})
}
