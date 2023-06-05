package main

import (
	"server/configs"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)
	routes.MessageRoute(app)

	app.Listen(":6000")
}
