package user_operations

import (
	"context"
	"main/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByName(username string, ctx context.Context) (primitive.ObjectID, error) {
	user := bson.M{"username": username}

	var result struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err := database.GetUsersCollection().FindOne(ctx, user).Decode(&result)
	return result.ID, err
}
