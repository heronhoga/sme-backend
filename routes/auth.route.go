package routes

import (
	"sme-backend/handlers"
	"github.com/gofiber/fiber/v2"
	"sme-backend/middlewares"
)
func AuthRoute (app *fiber.App) {
	app.Get("/test", middlewares.CheckKey,  handlers.Test)
	app.Post("/register", middlewares.CheckKey, handlers.Register)
	app.Post("/login", middlewares.CheckKey, handlers.Login)

	//check token investor
	app.Get("/test-investor", middlewares.CheckKey, middlewares.CheckInvestor,  handlers.TestInvestor)
	app.Get("/test-ukm", middlewares.CheckKey, middlewares.CheckUkm,  handlers.TestUkm)
}