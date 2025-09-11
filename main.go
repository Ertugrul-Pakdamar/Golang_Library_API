package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"main/database"
	"main/examples"
)

func main() {
	err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx := database.GetContext()

	examples.Example(ctx)

	<-c
	log.Println("Shutting down gracefully...")
	err = database.DisconnectFromMongoDB()
	if err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}
	log.Println("Server stopped!")
}
