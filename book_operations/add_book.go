package book_operations

import (
	"context"
	"main/database"
	"main/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBookToMongoDB(ctx context.Context, title string, author string, isTaken bool) error {
	book := models.Book{
		ID:      primitive.NewObjectID(),
		Title:   title,
		Author:  author,
		IsTaken: isTaken,
	}

	_, err := database.GetBooksCollection().InsertOne(ctx, book)
	return err
}
