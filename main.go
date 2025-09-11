package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"main/auth"
	"main/database"
	"main/examples"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// MongoDB bağlantısı
	err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("MongoDB bağlantısı başlıyor:", err)
	}

	// Program sonlandırma yönlendirmesi
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Drop user database for test
	examples.Example(database.GetContext())

	// Fiber web kurulumu
	fib := fiber.New()

	// Logger middleware: Tüm HTTP isteklerini logla
	fib.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// Fiber yönlendirmesi
	fib.Post("/api/user/register", auth.UserRegister)
	fib.Post("/api/user/login", auth.UserLogin)
	fib.Delete("/api/user/delete", utils.JWTProtected(), auth.UserDelete)

	// Sunucuyu başlat
	go func() {
		if err := fib.Listen("localhost:3000"); err != nil {
			log.Printf("Sunucu durdu: %v\n", err)
		}
	}()

	log.Println("Sunucu başlatıldı: http://localhost:3000")

	// Program sonlandırma işlemleri
	<-c
	log.Println("Program sonlandırılıyor")

	err = database.DisconnectFromMongoDB()
	if err != nil {
		log.Fatal("MongoDB bağlantısı kesildi:", err)
	}

	log.Println("Sunucu Durdu")
}
