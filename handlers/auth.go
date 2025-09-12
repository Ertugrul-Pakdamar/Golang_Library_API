package handlers

import (
	"main/database"
	"main/middleware"
	"main/models"
	"main/services"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, 400, "Invalid request format", "Request body must be valid JSON")
	}

	if user.Username == "" || user.Password == "" {
		return utils.SendError(c, 422, "Missing required fields", "Username and password are required")
	}

	if !utils.IsPasswordValid(user.Password) {
		return utils.SendError(c, 422, "Password validation failed", "Password must be at least 8 characters with uppercase, lowercase and numbers")
	}

	_, err := services.GetUserByUsername(database.GetContext(), user.Username)
	if err == nil {
		return utils.SendError(c, 409, "User already exists", "A user with this username already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to process password")
	}

	err = services.CreateUser(database.GetContext(), user.Username, hashedPassword)
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to create user account")
	}

	token, err := middleware.GenerateJWT(user.Username)
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to generate authentication token")
	}

	return utils.SendSuccess(c, "User registered successfully", fiber.Map{
		"token": token,
	})
}

func UserLogin(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, 400, "Invalid request format", "Request body must be valid JSON")
	}

	user_db, err := services.GetUserByUsername(database.GetContext(), user.Username)
	if err != nil {
		return utils.SendError(c, 401, "Authentication failed", "Invalid username or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password)) != nil {
		return utils.SendError(c, 401, "Authentication failed", "Invalid username or password")
	}

	token, err := middleware.GenerateJWT(user.Username)
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to generate authentication token")
	}

	return utils.SendSuccess(c, "Login successful", fiber.Map{
		"token": token,
	})
}

func UserDelete(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	user, err := services.GetUserByUsername(database.GetContext(), username)
	if err != nil {
		return utils.SendError(c, 404, "User not found", "The specified user does not exist")
	}

	err = services.DeleteUser(database.GetContext(), user.ID)
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to delete user account")
	}

	return utils.SendSuccess(c, "User account deleted successfully", nil)
}
