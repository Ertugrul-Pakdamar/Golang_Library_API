package utils

import "github.com/gofiber/fiber/v2"

func SendSuccess(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SendError(c *fiber.Ctx, code int, message string, detail string) error {
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error": fiber.Map{
			"code":    code,
			"details": detail,
		},
	})
}
