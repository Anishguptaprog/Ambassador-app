package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8000")
}
