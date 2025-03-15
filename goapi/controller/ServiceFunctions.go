package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/models"
	"gorm.io/gorm"
)

func CreateService(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		service := new(models.Service)

		err := c.BodyParser(service)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse Body",
				"error":   err.Error(),
			})
		}

		err = db.Create(service).Error
		if err != nil {

			return c.Status(501).JSON(fiber.Map{
				"message": "Could not create service",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Blog created",
			"data":    service,
		})
	}
}

func GetService(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		service := new(models.Service)
		err := db.Where("service_id = ?", id).First(service).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "Service not found",
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve service",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Retrieved service",
			"data":    service,
		})
	}
}
