package main

import (
	"sme-backend/database"
	"sme-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "sme-backend/models/entities"
)

func main() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		panic("Failed to load ENV")
	}

	// database initialization
	database.DatabaseInit()

	// migrations
	// database.DB.AutoMigrate(&entities.User{})

	app := fiber.New()

	// routes

	routes.AuthRoute(app)
	// routes.InvestorsRoute(app)
	// routes.UsersRoute(app)



	app.Listen(":8080")
}