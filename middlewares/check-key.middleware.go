package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func CheckKey(ctx *fiber.Ctx) error {
	errLoad := godotenv.Load()
	if errLoad != nil {
		panic("Failed to load ENV")
	}

	apiKey := ctx.Get("api-key")

	if apiKey != os.Getenv("API_KEY") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}
