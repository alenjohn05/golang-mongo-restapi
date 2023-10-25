package main

import (
	"golang-fiber-jwt/configs"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// Database.
	configs.ConnectDB()

	app.Listen(":" + configs.Port())
}
