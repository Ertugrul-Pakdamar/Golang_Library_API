package handlers

import (
	"main/database"
	"main/services"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	user, err := services.GetUserByUsername(database.GetContext(), username)
	if err != nil {
		return utils.SendError(c, 404, "User not found", "The specified user does not exist")
	}

	user.Password = ""

	return utils.SendSuccess(c, 200, "User information retrieved successfully", user)
}
