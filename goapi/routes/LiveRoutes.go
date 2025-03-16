package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/controller"
	"gorm.io/gorm"
)

func LiveRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/live/:id", controller.GetLive(db))
	api.Post("/live", controller.CreateLive(db))
	api.Get("/live", controller.GetLives(db))
}
