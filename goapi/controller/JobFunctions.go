package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/models"
	"gorm.io/gorm"
)

func CreateJob(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		job := new(models.Job)

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
			"message": "Job created",
			"data":    job,
		})
	}
}

func GetJob(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		job := new(models.Job)
		err := db.Where("job_id = ?", id).First(job).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "Job not found",
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

func GetJobs(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jobs := new([]models.Job)

		err := db.Find(jobs).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve jobs",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"data":    jobs,
			"message": "Retrieved jobs",
		})
	}
}
