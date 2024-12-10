package main

import (
	"sme-backend/database"
	"sme-backend/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// cors 	
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,                    
		AllowMethods:     "GET,POST,PUT,DELETE", 
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, api-key",
	}))
	// routes
	routes.AuthRoute(app)
	// routes.InvestorsRoute(app)
	// routes.UsersRoute(app)



	app.Listen(":8080")
}