package auth

import (
	"main/database"
	"main/models"
	"main/user_operations"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, 400, "Geçersiz istek", err.Error())
	}

	user_db, err := user_operations.GetUserByName(user.Username, database.GetContext())
	if err != nil {
		return utils.SendError(c, 401, "Kullanıcı bulunamadı", "")
	}

	if bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password)) != nil {
		return utils.SendError(c, 401, "Şifre hatalı", "")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return utils.SendError(c, 500, "Token oluşturulamadı", err.Error())
	}

	return utils.SendSuccess(c, "Giriş başarılı", fiber.Map{
		"token": token,
	})
}
