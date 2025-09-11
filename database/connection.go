package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
	Books    *mongo.Collection
	Users    *mongo.Collection
	ctx      context.Context
)

func ConnectToMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	Client = client
	Database = client.Database("golang_api_db")
	Books = Database.Collection("books")
	Users = Database.Collection("users")

	log.Println("Successfully connected to MongoDB!")
	return nil
}

func DisconnectFromMongoDB() error {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := Client.Disconnect(ctx)
		if err != nil {
			return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
		}

		log.Println("Disconnected from MongoDB!")
	}
	return nil
}

func GetDatabase() *mongo.Database {
	return Database
}

func GetBooksCollection() *mongo.Collection {
	return Books
}

func GetUsersCollection() *mongo.Collection {
	return Users
}

func GetContext() context.Context {
	return ctx
}
