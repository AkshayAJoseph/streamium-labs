package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/streamium-labs/goapi/database"
	"github.com/streamium-labs/goapi/models"
	"github.com/streamium-labs/goapi/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "Content-Type,Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Failed connection to database")
	}

	models.MigrateUser(db)
	models.MigrateBlog(db)
	models.MigrateService(db)

	routes.UserRoutes(db, app)
	routes.BlogRoutes(db, app)
	routes.ServiceRoutes(db, app)

	app.Listen(":8080")
}
