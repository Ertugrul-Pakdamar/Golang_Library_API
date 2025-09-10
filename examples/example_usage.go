package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"main/database"
	"main/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Connect to MongoDB
	err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer database.DisconnectFromMongoDB()

	// Get collections
	booksCollection := database.GetBooksCollection()
	usersCollection := database.GetUsersCollection()

	// Example: Insert a book
	book := models.Book{
		ID:      primitive.NewObjectID(),
		Title:   "The Go Programming Language",
		Author:  "Alan Donovan",
		IsTaken: false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = booksCollection.InsertOne(ctx, book)
	if err != nil {
		log.Fatal("Failed to insert book:", err)
	}
	fmt.Println("Book inserted successfully!")

	// Example: Insert a user
	user := models.User{
		ID:         primitive.NewObjectID(),
		Username:   "john_doe",
		Password:   "hashed_password_here",
		BooksTaken: []primitive.ObjectID{},
	}

	_, err = usersCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal("Failed to insert user:", err)
	}
	fmt.Println("User inserted successfully!")

	fmt.Println("Database setup completed successfully!")
}
