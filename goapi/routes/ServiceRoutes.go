package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/controller"
	"gorm.io/gorm"
)

func ServiceRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/service", controller.CreateService(db))
	api.Get("/service/:id", controller.GetService(db))
	api.Get("/service", controller.GetServices(db))
}
