package main

import (
	"PMAB-/api/app"
	"PMAB-/api/database"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := database.OpenDb(false)
	server := app.CrateApp(db)

	server.Listen(":3001")
	defer db.Close()
}
