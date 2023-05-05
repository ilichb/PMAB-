package handlers

import "github.com/gofiber/fiber/v2"

/* Get test */

func GetBlockchain(c *fiber.Ctx) error {
	response, err := services.GetBlockchain()
	if err != nil {
		return c.SendString("Error in get")
	}
	return c.JSON(response)
}

func PostBlockchain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Esto es post Blockchain",
	})
}
