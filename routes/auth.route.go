package routes

import (
	"sme-backend/handlers"
	"github.com/gofiber/fiber/v2"
)
func AuthRoute (app *fiber.App) {
	app.Get("/test", handlers.Test)
	app.Post("/register", handlers.Register)
}