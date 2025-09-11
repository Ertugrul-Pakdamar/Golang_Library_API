package user_operations

import (
	"context"
	"main/database"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByName(username string, ctx context.Context) (models.User, error) {
	user := bson.M{"username": username}

	var result models.User

	err := database.GetUsersCollection().FindOne(ctx, user).Decode(&result)
	return result, err
}
