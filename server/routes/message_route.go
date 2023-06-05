package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func MessageRoute(app *fiber.App) {
	app.Post("/message", controllers.SendMessage)
}
