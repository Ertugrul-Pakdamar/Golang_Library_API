package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"main/book_operations"
	"main/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx := database.GetContext()

	err = book_operations.DeleteBooksFromMongoDB(ctx, primitive.ObjectID{0})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	<-c
	log.Println("Shutting down gracefully...")
	err = database.DisconnectFromMongoDB()
	if err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}
	log.Println("Server stopped!")
}
