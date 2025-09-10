package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"main/database"
)

func main() {
	// Connect to MongoDB
	err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal
	<-c
	log.Println("Shutting down gracefully...")

	// Disconnect from MongoDB
	err = database.DisconnectFromMongoDB()
	if err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}

	log.Println("Server stopped!")
}
