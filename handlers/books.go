package handlers

import (
	"main/database"
	"main/services"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	var request struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, 400, "Invalid request format", "Request body must be valid JSON")
	}

	if request.Title == "" || request.Author == "" {
		return utils.SendError(c, 422, "Missing required fields", "Title and author are required")
	}

	err := services.CreateBook(database.GetContext(), request.Title, request.Author)
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to add book to library")
	}

	return utils.SendSuccess(c, 201, "Book added to library successfully", nil)
}

func GetAllBooks(c *fiber.Ctx) error {
	books, err := services.GetAllBooks(database.GetContext())
	if err != nil {
		return utils.SendError(c, 500, "Internal server error", "Failed to retrieve library catalog")
	}

	return utils.SendSuccess(c, 200, "Library catalog retrieved successfully", books)
}
