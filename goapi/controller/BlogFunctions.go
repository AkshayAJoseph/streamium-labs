package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/streamium-labs/goapi/models"
	"gorm.io/gorm"
)

func CreateBlog(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		blog := new(models.Blog)

		err := c.BodyParser(blog)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse Body",
				"error":   err.Error(),
			})
		}

		err = db.Create(blog).Error
		if err != nil {

			return c.Status(501).JSON(fiber.Map{
				"message": "Could not create blog",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Blog created",
			"data":    blog,
		})
	}
}

func GetBlog(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		blog := new(models.Blog)
		err := db.Where("blog_id = ?", id).First(blog).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "Blog not found",
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve blog",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Retrieved blog",
			"data":    blog,
		})
	}
}
