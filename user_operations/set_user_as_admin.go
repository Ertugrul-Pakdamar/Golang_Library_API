package user_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetUserAsAdminByID(ctx context.Context, id primitive.ObjectID) error {
	user := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"role": "admin"},
	}
	_, err := database.GetBooksCollection().UpdateOne(ctx, user, update)
	return err
}
