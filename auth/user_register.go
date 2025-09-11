package auth

import (
	"main/database"
	"main/models"
	"main/user_operations"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRegister(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, 400, "Geçersiz istek", err.Error())
	}

	if user.Username == "" || user.Password == "" {
		return utils.SendError(c, 422, "Kullanıcı adı ve şifre zorunludur", "")
	}

	if !utils.IsPasswordValid(user.Password) {
		return utils.SendError(c, 422, "Şifre 8 karakterden uzun büyük, küçük harf ve sayı içermelidir", "")
	}

	_, err := user_operations.GetUserByName(user.Username, database.GetContext())
	if err == nil {
		return utils.SendError(c, 409, "Kullanıcı zaten var", "")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.SendError(c, 500, "Şifre hashlenemedi", err.Error())
	}

	// Bu satırlar gereksiz ve yanlış collection'ı kullanıyor, silindi

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return utils.SendError(c, 500, "Token oluşturulamadı", err.Error())
	}

	user_operations.AddUserToMongoDB(database.GetContext(), user.Username, hashedPassword)
	return utils.SendSuccess(c, "Kayıt başarılı", fiber.Map{
		"token": token,
	})
}
