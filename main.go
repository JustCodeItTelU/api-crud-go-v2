package main

import (
	"JCI-Go-API/database"
	"JCI-Go-API/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	err := app.Listen("3000")
	if err != nil {
		return
	}
}
