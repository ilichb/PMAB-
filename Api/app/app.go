package app

import (
	"PMAB-/api/routes"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func CrateApp(db *sql.DB) *fiber.App {
	app := fiber.New()

	routes.MainRoute(app, db)

	return app
}
