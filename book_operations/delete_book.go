package book_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBooksFromMongoDB(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := database.GetBooksCollection().DeleteOne(ctx, filter)
	return err
}
