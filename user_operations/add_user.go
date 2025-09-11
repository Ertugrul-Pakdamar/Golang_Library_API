package user_operations

import (
	"context"
	"main/database"
	"main/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserToMongoDB(ctx context.Context, username string, password string) error {
	user := models.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: password,
		Role:     1,
	}

	_, err := database.GetUsersCollection().InsertOne(ctx, user)
	return err
}
