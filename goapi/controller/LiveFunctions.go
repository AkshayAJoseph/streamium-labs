package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/models"
	"gorm.io/gorm"
)

func CreateLive(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		job := new(models.Live)

		err := c.BodyParser(job)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse Body",
				"error":   err.Error(),
			})
		}

		err = db.Create(job).Error
		if err != nil {

			return c.Status(501).JSON(fiber.Map{
				"message": "Could not create job",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Live created",
			"data":    job,
		})
	}
}

func GetLive(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		job := new(models.Live)
		err := db.Where("job_id = ?", id).First(job).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "Live not found",
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve job",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Retrieved job",
			"data":    job,
		})
	}
}

func GetLives(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		live := new([]models.Live)

		err := db.Find(live).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve live",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"data":    live,
			"message": "Retrieved live",
		})
	}
}
