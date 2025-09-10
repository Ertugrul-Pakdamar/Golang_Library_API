package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"main/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
	Books    *mongo.Collection
	Users    *mongo.Collection
)

// ConnectToMongoDB establishes connection to MongoDB
func ConnectToMongoDB() error {
	// Load configuration
	cfg := config.LoadConfig()

	// Set client options
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Set global variables
	Client = client
	Database = client.Database(cfg.DatabaseName)     // Database name
	Books = Database.Collection(cfg.BooksCollection) // Books collection
	Users = Database.Collection(cfg.UsersCollection) // Users collection

	log.Println("Successfully connected to MongoDB!")
	return nil
}

// DisconnectFromMongoDB closes the MongoDB connection
func DisconnectFromMongoDB() error {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := Client.Disconnect(ctx)
		if err != nil {
			return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
		}

		log.Println("Disconnected from MongoDB!")
	}
	return nil
}

// GetDatabase returns the database instance
func GetDatabase() *mongo.Database {
	return Database
}

// GetBooksCollection returns the books collection
func GetBooksCollection() *mongo.Collection {
	return Books
}

// GetUsersCollection returns the users collection
func GetUsersCollection() *mongo.Collection {
	return Users
}
