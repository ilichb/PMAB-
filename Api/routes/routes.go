package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App, db *sql.DB) {

	users := app.Group("/api/v1/test")
	testRoutes(users)

}
