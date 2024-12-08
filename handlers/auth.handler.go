package handlers

import (
	"sme-backend/database"
	"sme-backend/models/entities"
	"sme-backend/models/requests"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Test (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func Register (ctx *fiber.Ctx) error {
	user := new(requests.Register)

	// Parse the request body
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//validate the request
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//check existing username
	var existingUser entities.User
	result := database.DB.Raw("SELECT 1 FROM users WHERE username = ? LIMIT 1", user.Username).Scan(&existingUser)

	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username already exists",
		})
	}

	//create user
	newUser := entities.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Phone: user.Phone,
		Username: user.Username,
		Password: user.Password,
		Role: user.Role,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errCreateUser.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": newUser,
	})
}