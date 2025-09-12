package handlers

import (
	"main/database"
	"main/services"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func BorrowBook(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var request struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, 400, "Invalid request format", "Request body must be valid JSON")
	}

	if request.Title == "" {
		return utils.SendError(c, 422, "Missing required field", "Book title is required")
	}

	user, err := services.GetUserByUsername(database.GetContext(), username)
	if err != nil {
		return utils.SendError(c, 404, "User not found", "The specified user does not exist")
	}

	book, err := services.GetBookByTitle(database.GetContext(), request.Title)
	if err != nil {
		return utils.SendError(c, 404, "Book not found", "The specified book is not available in the library")
	}

	err = services.BorrowBook(database.GetContext(), user.ID, book.ID)
	if err != nil {
		if err.Error() == "Borrowing limit exceeded" {
			return utils.SendError(c, 403, "Borrowing limit exceeded", "You have reached the maximum borrowing limit (2 books)")
		}
		if err.Error() == "Book unavailable" {
			return utils.SendError(c, 409, "Book unavailable", "This book is currently not available for borrowing")
		}
		if err.Error() == "Already borrowed" {
			return utils.SendError(c, 409, "Already borrowed", "You have already borrowed this book")
		}
		return utils.SendError(c, 500, "Internal server error", "Failed to process borrowing request")
	}

	return utils.SendSuccess(c, "Book borrowed successfully", nil)
}

func ReturnBook(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var request struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, 400, "Invalid request format", "Request body must be valid JSON")
	}

	if request.Title == "" {
		return utils.SendError(c, 422, "Missing required field", "Book title is required")
	}

	user, err := services.GetUserByUsername(database.GetContext(), username)
	if err != nil {
		return utils.SendError(c, 404, "User not found", "The specified user does not exist")
	}

	book, err := services.GetBookByTitle(database.GetContext(), request.Title)
	if err != nil {
		return utils.SendError(c, 404, "Book not found", "The specified book is not available in the library")
	}

	err = services.ReturnBook(database.GetContext(), user.ID, book.ID)
	if err != nil {
		if err.Error() == "Not borrowed" {
			return utils.SendError(c, 409, "Not borrowed", "You have not borrowed this book")
		}
		return utils.SendError(c, 500, "Internal server error", "Failed to process return request")
	}

	return utils.SendSuccess(c, "Book returned successfully", nil)
}
