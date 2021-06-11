package router

import (
	"github.com/Leon4rdoMonteiro/gofiber.api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/v1")

	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)

}
