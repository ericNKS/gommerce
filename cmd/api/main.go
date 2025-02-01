package main

import (
	"github.com/ericNKS/gommerce/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	app := fiber.New()

	routes.Exec(app)

	app.Listen(":3000")
}
