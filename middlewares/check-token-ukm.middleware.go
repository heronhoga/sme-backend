package middlewares

import (
	"sme-backend/database"
	"sme-backend/models/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)
func CheckUkm (ctx *fiber.Ctx) error {
	// get token from header
	token := ctx.Get("token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// parse token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// get username from claims
	username, ok := claims["username"].(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// get user from database
	var user entities.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// check role
	if user.Role != "ukm" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}