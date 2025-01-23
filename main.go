package main

import (
	"context"
	"log"
	"rendipilot/simple-golang/database"
	"rendipilot/simple-golang/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber/v2"
	// "github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Load environment variables
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Initialize database connection ONCE
	_, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	defer database.GetDB().Close(context.Background())

	config := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"}, // Allow all methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	app.Use(config)

	// Define endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello Rendy")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Test")
	})

	app.Post("/adduser", handlers.AddUser)

	app.Get("/users", handlers.GetUsers)

	// Start the server
	app.Listen(":3000")
}
