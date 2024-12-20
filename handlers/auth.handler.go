package handlers

import (
	"sme-backend/database"
	"sme-backend/functions"
	"sme-backend/models/entities"
	"sme-backend/models/requests"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Test (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func TestInvestor (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Inverstor ok",
	})
}

func TestUkm (ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ukm ok",
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

func Login(ctx *fiber.Ctx) error {
	user := new(requests.Login)

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

	var exists bool
	result := database.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? LIMIT 1)", user.Username).Scan(&exists)
	
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	
	if !exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Check if password is correct - password is hashed
	var hashedPassword string
	result = database.DB.Raw("SELECT password FROM users WHERE username = ?", user.Username).Scan(&hashedPassword)
	
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid password",
		})
	}
	

	// Generate JWT token
	token, errGenerateToken := functions.GenerateToken(user.Username)

	if errGenerateToken != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	// Get user from database
	var userData entities.User
	result = database.DB.Raw("SELECT * FROM users WHERE username = ?", user.Username).Scan(&userData)
	
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"token": token,
		"data": 
		fiber.Map{
		"first_name": userData.FirstName, 
		"last_name": userData.LastName, 
		"username": userData.Username,
		"role": userData.Role,},
	})
}
