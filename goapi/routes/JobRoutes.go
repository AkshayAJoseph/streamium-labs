package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/controller"
	"gorm.io/gorm"
)

func JobRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/job", controller.CreateJob(db))
	api.Get("/job/:id", controller.GetJob(db))
	api.Get("/job", controller.GetJobs(db))
}
