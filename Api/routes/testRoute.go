package routes

import (
	"PMAB-/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func testRoutes(test fiber.Router) {

	/* test get */
	test.Get("/", handlers.FindTest)

}
