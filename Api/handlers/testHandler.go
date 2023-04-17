package handlers

import "github.com/gofiber/fiber/v2"

/* Get test */

func FindTest(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Esto es un test",
	})
}
