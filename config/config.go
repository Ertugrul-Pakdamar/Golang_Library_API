package config

import "os"

// Config holds all configuration for our application
type Config struct {
	MongoURI        string
	DatabaseName    string
	BooksCollection string
	UsersCollection string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		MongoURI:        getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		DatabaseName:    getEnv("DB_NAME", "library_db"),
		BooksCollection: getEnv("BOOKS_COLLECTION", "books"),
		UsersCollection: getEnv("USERS_COLLECTION", "users"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
