package book_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBookByName(name string, ctx context.Context) (primitive.ObjectID, error) {
	book := bson.M{"title": name}

	var result struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err := database.GetBooksCollection().FindOne(ctx, book).Decode(&result)
	return result.ID, err
}
