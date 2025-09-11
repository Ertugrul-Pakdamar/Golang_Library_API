package book_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBookFromMongoDB(ctx context.Context, id primitive.ObjectID) error {
	book := bson.M{"_id": id}

	_, err := database.GetBooksCollection().DeleteOne(ctx, book)
	return err
}
