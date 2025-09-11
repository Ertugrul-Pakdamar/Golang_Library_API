package book_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MarkBookTakenFromMongoDB(ctx context.Context, id primitive.ObjectID) error {
	book := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"is_taken": true},
	}
	_, err := database.GetBooksCollection().UpdateOne(ctx, book, update)
	return err
}
