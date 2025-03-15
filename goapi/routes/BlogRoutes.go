package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/controller"
	"gorm.io/gorm"
)

func BlogRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/blog", controller.CreateBlog(db))
	api.Get("/blog/:id", controller.GetBlog(db))
	api.Get("/blogs", controller.GetBlogs(db))
}
