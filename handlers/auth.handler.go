package handlers

import (
	"sme-backend/database"
	"sme-backend/models/entities"
	"sme-backend/models/requests"
	"sme-backend/functions"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Test (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func Register(ctx *fiber.Ctx) error {
    user := new(requests.Register)

    // Parse the request body
    if err := ctx.BodyParser(user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    // Validate the request
    validate := validator.New()
    if err := validate.Struct(user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    // Check if username already exists
	var exists bool
	result := database.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? LIMIT 1)", user.Username).Scan(&exists)
	
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	
	if exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username already exists",
		})
	}
	

    // Encrypt password
    hashedPassword, err := functions.HashPassword(user.Password)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }
    user.Password = hashedPassword

    // Generate UUID for new user
    newUserUUID := uuid.New()

    // Create user
    newUser := entities.User{
        IdUser:   newUserUUID,
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

	// Generate JWT token
	token, errGenerateToken := functions.GenerateToken(user.Username)

	if errGenerateToken != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "success",
        "data": newUser,
		"token": token,
    })
}
