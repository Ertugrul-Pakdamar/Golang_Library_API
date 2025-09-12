package services

import (
	"context"
	"main/database"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(ctx context.Context, username string, password string) error {
	count, err := database.GetUsersCollection().CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	role := 1
	if count == 0 {
		role = 0
	}

	user := models.User{
		ID:         primitive.NewObjectID(),
		Username:   username,
		Password:   password,
		Role:       role,
		BooksTaken: []primitive.ObjectID{},
	}

	_, err = database.GetUsersCollection().InsertOne(ctx, user)
	return err
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	filter := bson.M{"username": username}
	var user models.User
	err := database.GetUsersCollection().FindOne(ctx, filter).Decode(&user)
	return user, err
}

func GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	filter := bson.M{"_id": id}
	var user models.User
	err := database.GetUsersCollection().FindOne(ctx, filter).Decode(&user)
	return user, err
}

func DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	user, err := GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	for _, bookID := range user.BooksTaken {
		bookFilter := bson.M{"_id": bookID}
		bookUpdate := bson.M{
			"$inc": bson.M{"borrowed": -1},
		}
		_, err = database.GetBooksCollection().UpdateOne(ctx, bookFilter, bookUpdate)
		if err != nil {
			return err
		}
	}

	userFilter := bson.M{"_id": id}
	_, err = database.GetUsersCollection().DeleteOne(ctx, userFilter)
	return err
}

func SetUserAsAdmin(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"role": 0},
	}
	_, err := database.GetUsersCollection().UpdateOne(ctx, filter, update)
	return err
}

func AddBookToUser(ctx context.Context, userID primitive.ObjectID, bookID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$push": bson.M{"books_taken": bookID},
	}
	_, err := database.GetUsersCollection().UpdateOne(ctx, filter, update)
	return err
}

func RemoveBookFromUser(ctx context.Context, userID primitive.ObjectID, bookID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$pull": bson.M{"books_taken": bookID},
	}
	_, err := database.GetUsersCollection().UpdateOne(ctx, filter, update)
	return err
}
