package handlers

import "github.com/gofiber/fiber/v2"

/* Get test */

func GetBlockchain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Esto es get Blockchain",
	})
}

func PostBlockchain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Esto es post Blockchain",
	})
}
