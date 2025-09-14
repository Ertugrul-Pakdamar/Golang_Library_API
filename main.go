package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"main/database"
	"main/handlers"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// MongoDB connection
	err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// Drop Database
	// database.GetUsersCollection().Drop(database.GetContext())
	// database.GetBooksCollection().Drop(database.GetContext())

	// Program termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Fiber setup
	fib := fiber.New()
	fib.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// API routes
	fib.Post("/api/user/register", handlers.UserRegister)
	fib.Post("/api/user/login", handlers.UserLogin)
	fib.Delete("/api/user/delete", middleware.JWTProtected(), handlers.UserDelete)
	fib.Get("/api/user/info", middleware.JWTProtected(), handlers.GetUserInfo)
	fib.Post("/api/book/add", middleware.JWTProtectedAdmin(), handlers.AddBook)
	fib.Get("/api/book/list", middleware.JWTProtected(), handlers.GetAllBooks)
	fib.Post("/api/book/borrow", middleware.JWTProtected(), handlers.BorrowBook)
	fib.Post("/api/book/return", middleware.JWTProtected(), handlers.ReturnBook)

	// Start server
	go func() {
		if err := fib.Listen("localhost:3000"); err != nil {
			log.Printf("Server stopped: %v\n", err)
		}
	}()

	log.Println("Server started: http://localhost:3000")

	// Program termination cleanup
	<-c
	log.Println("Shutting down...")

	err = database.DisconnectFromMongoDB()
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	log.Println("Server stopped")
}
