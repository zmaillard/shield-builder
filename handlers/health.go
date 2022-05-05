package handlers

import "github.com/gofiber/fiber/v2"

func HealthHandler(c *fiber.Ctx) error {
	c.Status(200)
	return nil
}
