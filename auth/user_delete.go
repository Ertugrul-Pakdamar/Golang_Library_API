package auth

import (
	"main/database"
	"main/user_operations"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func UserDelete(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	// Kullanıcıyı username ile bul
	user, err := user_operations.GetUserByName(username, database.GetContext())
	if err != nil {
		return utils.SendError(c, 404, "Kullanıcı bulunamadı", err.Error())
	}

	err = user_operations.DeleteUserFromMongoDB(database.GetContext(), user.ID)
	if err != nil {
		return utils.SendError(c, 500, "Kullanıcı silinemedi", err.Error())
	}

	return utils.SendSuccess(c, "Hesap silindi", nil)
}
