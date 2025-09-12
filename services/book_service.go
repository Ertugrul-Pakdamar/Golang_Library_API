package services

import (
	"context"
	"main/database"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(ctx context.Context, title string, author string) error {
	book, err := GetBookByTitle(ctx, title)
	if err != nil {
		newBook := models.Book{
			ID:       primitive.NewObjectID(),
			Title:    title,
			Author:   author,
			Count:    1,
			Borrowed: 0,
		}
		_, err = database.GetBooksCollection().InsertOne(ctx, newBook)
	} else {
		filter := bson.M{"_id": book.ID}
		update := bson.M{
			"$inc": bson.M{"count": 1},
		}
		_, err = database.GetBooksCollection().UpdateOne(ctx, filter, update)
	}
	return err
}

func GetBookByTitle(ctx context.Context, title string) (models.Book, error) {
	filter := bson.M{"title": title}
	var book models.Book
	err := database.GetBooksCollection().FindOne(ctx, filter).Decode(&book)
	return book, err
}

func GetBookByID(ctx context.Context, id primitive.ObjectID) (models.Book, error) {
	filter := bson.M{"_id": id}
	var book models.Book
	err := database.GetBooksCollection().FindOne(ctx, filter).Decode(&book)
	return book, err
}

func GetAllBooks(ctx context.Context) ([]models.Book, error) {
	cursor, err := database.GetBooksCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []models.Book
	if err = cursor.All(ctx, &books); err != nil {
		return nil, err
	}

	return books, nil
}

func DeleteBook(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := database.GetBooksCollection().DeleteOne(ctx, filter)
	return err
}

func UpdateBookBorrowedCount(ctx context.Context, bookID primitive.ObjectID, increment int) error {
	filter := bson.M{"_id": bookID}
	update := bson.M{
		"$inc": bson.M{"borrowed": increment},
	}
	_, err := database.GetBooksCollection().UpdateOne(ctx, filter, update)
	return err
}
