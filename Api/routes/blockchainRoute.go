package routes

import (
	"PMAB-/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func blockchainRoutes(api fiber.Router) {

	api.Get("/blockchain", handlers.GetBlockchain)

	api.Post("/blockchain", handlers.PostBlockchain)

}
