package user_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUserFromMongoDB(ctx context.Context, id primitive.ObjectID) error {
	user := bson.M{"_id": id}

	_, err := database.GetUsersCollection().DeleteOne(ctx, user)
	return err
}
